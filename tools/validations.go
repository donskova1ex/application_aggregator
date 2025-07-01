package tools

import (
	"github.com/google/uuid"
	"regexp"
)

func ValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func ValidPhone(phoneNumber string) bool {
	cleanedPhone := regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	pattern := `^(?:\+7|8|7)\s?9\d{2}\s?\d{3}\s?\d{2}\s?\d{2}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(cleanedPhone)
}
