package service

import (
	"WeatherSubscriptionAPI/internal/repository"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// SubscriptionService handles subscribe/confirm/unsubscribe logic.
type SubscriptionService struct {
	Repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}

// Subscribe generates a token and stores a pending subscription.
func (s *SubscriptionService) Subscribe(email, city string) (string, error) {
	// Generate a 16-byte random token (hex-encoded).
	tokenBytes := make([]byte, 16)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)
	// Store in DB.
	if err := s.Repo.Create(email, city, token); err != nil {
		return "", err
	}
	// Simulate sending a confirmation email by printing URLs.
	fmt.Printf("Confirmation link: http://localhost:8080/confirm/%s\n", token)
	fmt.Printf("Unsubscribe link: http://localhost:8080/unsubscribe/%s\n", token)
	return token, nil
}

// Confirm activates the subscription for the given token.
func (s *SubscriptionService) Confirm(token string) (bool, error) {
	return s.Repo.Confirm(token)
}

// Unsubscribe disables the subscription for the given token.
func (s *SubscriptionService) Unsubscribe(token string) (bool, error) {
	return s.Repo.Unsubscribe(token)
}
