package utils

import (
	"math/rand"

	"github.com/alighamgosar/rock_paper_scissors_golang/initializers"
	"github.com/alighamgosar/rock_paper_scissors_golang/models"
	"gorm.io/gorm"
)

// GenerateRandomID generates a random string of the specified length.
func GenerateRandomID(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// CheckRandIDExists checks if the given random ID exists in the database.
func CheckRandIDExists(randID string) bool {
	var games []models.Game
	if err := initializers.DB.Where("random_id = ?", randID).First(&games).Error; err != nil {
		return false
	}
	return true
}

// GenerateAndCheckUniqueRandID generates a random ID and checks if it exists in the database.
func GenerateAndCheckUniqueRandID(db *gorm.DB) string {
	for {
		randID := GenerateRandomID(40)
		if !CheckRandIDExists(randID) {
			return randID
		}
	}
}
