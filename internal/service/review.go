package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type ReviewService struct {
	repo        *repository.ReviewRepository
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
}

func NewReviewService(repo *repository.ReviewRepository, productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository) *ReviewService {
	return &ReviewService{
		repo:        repo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (s *ReviewService) CreateReview(userID uint, productID uint, rating int, comment string) error {
	// Check if product exists
	_, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	// Check if user has purchased the product
	orders, err := s.orderRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	hasPurchased := false
	for _, order := range orders {
		if order.Status == "delivered" {
			for _, item := range order.Items {
				if item.ProductID == productID {
					hasPurchased = true
					break
				}
			}
		}
		if hasPurchased {
			break
		}
	}

	if !hasPurchased {
		return errors.New("you must purchase the product before reviewing")
	}

	// Check if user has already reviewed the product
	_, err = s.repo.FindByUserAndProduct(userID, productID)
	if err == nil {
		return errors.New("you have already reviewed this product")
	}

	// Validate rating
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// Create review
	review := &models.Review{
		UserID:    userID,
		ProductID: productID,
		Rating:    rating,
		Comment:   comment,
	}

	return s.repo.Create(review)
}

func (s *ReviewService) GetReview(id uint) (*models.Review, error) {
	return s.repo.FindByID(id)
}

func (s *ReviewService) GetProductReviews(productID uint) ([]models.Review, error) {
	return s.repo.FindByProductID(productID)
}

func (s *ReviewService) GetUserReviews(userID uint) ([]models.Review, error) {
	return s.repo.FindByUserID(userID)
}

func (s *ReviewService) UpdateReview(userID uint, reviewID uint, rating int, comment string) error {
	// Get review
	review, err := s.repo.FindByID(reviewID)
	if err != nil {
		return errors.New("review not found")
	}

	// Check ownership
	if review.UserID != userID {
		return errors.New("unauthorized")
	}

	// Validate rating
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// Update review
	review.Rating = rating
	review.Comment = comment

	return s.repo.Update(review)
}

func (s *ReviewService) DeleteReview(userID uint, reviewID uint) error {
	// Get review
	review, err := s.repo.FindByID(reviewID)
	if err != nil {
		return errors.New("review not found")
	}

	// Check ownership
	if review.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(reviewID)
}
