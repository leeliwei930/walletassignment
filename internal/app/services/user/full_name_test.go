package user

import (
	"context"
	"testing"

	"github.com/leeliwei930/walletassignment/ent"
	"github.com/stretchr/testify/assert"
)

func TestGetFullName(t *testing.T) {
	// Create test cases
	testCases := []struct {
		name     string
		user     *ent.User
		expected string
	}{
		{
			name: "normal case",
			user: &ent.User{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: "John Doe",
		},
		{
			name: "empty names",
			user: &ent.User{
				FirstName: "",
				LastName:  "",
			},
			expected: " ",
		},
		{
			name: "first name only",
			user: &ent.User{
				FirstName: "John",
				LastName:  "",
			},
			expected: "John ",
		},
		{
			name: "last name only",
			user: &ent.User{
				FirstName: "",
				LastName:  "Doe",
			},
			expected: " Doe",
		},
	}

	// Create service instance
	service := &userService{}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := service.GetFullName(context.Background(), tc.user)
			assert.Equal(t, tc.expected, result)
		})
	}
}
