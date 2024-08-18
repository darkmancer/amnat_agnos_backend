package service

import (
	"strong_password_recommendation/internal/core/port"
	"unicode"
)

type PasswordService struct {
	repo port.LogRepository
}

func NewPasswordService(repo port.LogRepository) *PasswordService {
	return &PasswordService{repo: repo}
}

func (s *PasswordService) CalculateSteps(password string) int {
	length := len(password)
	needsLower, needsUpper, needsDigit := true, true, true
	repeatFixes := 0
	steps := 0

	var lastChar1, lastChar2 rune

	for i, char := range password {
		if unicode.IsLower(char) {
			needsLower = false
		} else if unicode.IsUpper(char) {
			needsUpper = false
		} else if unicode.IsDigit(char) {
			needsDigit = false
		}

		if i > 1 && char == lastChar1 && lastChar1 == lastChar2 {
			repeatFixes++
			lastChar1 = 0
		} else {
			lastChar2 = lastChar1
			lastChar1 = char
		}
	}

	missingTypes := 0
	if needsLower {
		missingTypes++
	}
	if needsUpper {
		missingTypes++
	}
	if needsDigit {
		missingTypes++
	}

	if length < 6 {
		steps = max(6-length, max(missingTypes, repeatFixes))
	} else {
		steps += max(missingTypes, repeatFixes)
		if length > 20 {
			steps += length - 20
		}
	}

	s.repo.LogRequestResponse(password, steps)

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
