package service_test

import (
	"strong_password_recommendation/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLogRepository is a mock implementation of the LogRepository interface
type MockLogRepository struct {
	mock.Mock
}

func (m *MockLogRepository) LogRequestResponse(password string, steps int) error {
	args := m.Called(password, steps)
	return args.Error(0)
}

func TestPasswordService_CalculateSteps(t *testing.T) {
	mockRepo := new(MockLogRepository)
	passwordService := service.NewPasswordService(mockRepo)

	testCases := []struct {
		name          string
		password      string
		expectedSteps int
	}{
		{"Too short, no types", "123", 3},
		{"Minimum length with missing types", "aA1", 3},
		{"Already strong", "aa11Baa", 0},
		{"Repeating characters", "aaa111aaa", 3},
		{"Repeating characters episode 2", "aaa111BBB", 3},
		{"Too long", "12345678901234567890aA1", 3},
		{"Missing uppercase and lowercase", "123456", 2},
		{"Missing uppercase and lowercase on 22 character", "1234567891234567891234", 4},
		{"Missing uppercase and lowercase on 22 character with one uppercase and one lowercase", "12345678912345678912Ab", 2},
		{"Missing uppercase and lowercase on 22 character with two uppercase", "12345678912345678912AB", 3},
		{"Missing uppercase and lowercase on 25 character with repeated digit characters", "1234567891234567891234111", 7},
		{"Missing uppercase and lowercase on 26 character with repeated digit characters", "12345678912345678912341111", 8},
		{"Password contains dot", "12.", 3},
		{"Password contains exclamation", "12!", 3},
		{"Password contains dot and exclamation", "1.!", 3},
		{"Password contains digits and repeated exclamations", "12!!!", 2},
		{"Password contains digits and repeated dots", "12...", 2},
		{"Password contains digits, repeated dots and repeated exclamation", "12...!!!", 2},
		{"Password contains only 6 exclamation", "!!!!!!", 3},
		{"Password contains only 20 exclamation", "!!!!!!!!!!!!!!!!!!!!", 6},
		{"Password contains only 20 same digit", "11111111111111111111", 6},

		//TODO
		// {"Ultimate Repeating characters", "aaa111aaaBBB666CCC666lll", 8},
		// {"Password contains only 40 same digit", "1111111111111111111111111111111111111111", 26},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("LogRequestResponse", tc.password, tc.expectedSteps).Return(nil)

			steps := passwordService.CalculateSteps(tc.password)

			assert.Equal(t, tc.expectedSteps, steps)
			mockRepo.AssertExpectations(t)

			mockRepo.ExpectedCalls = nil 
		})
	}
}
