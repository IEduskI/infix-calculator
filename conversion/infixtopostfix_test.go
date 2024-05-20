package conversion

import (
	"strings"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {

	tests := []struct {
		name    string
		infix   string
		postfix string
		wantErr error
	}{
		{
			name:    "Test Infix To Postfix Success",
			infix:   "1 + 2.5 / 3 * 4",
			postfix: "12.53/4*+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := InfixToPostfix(tt.infix)
			if err != tt.wantErr {
				t.Errorf("Want: %v , got %v", tt.postfix, result)
			}
			if result.String() != tt.postfix {
				t.Errorf("InfixToPostfix() = %v, want %v", result.String(), tt.postfix)
			}
		})

	}
}

func TestTokenization(t *testing.T) {

	tests := []struct {
		name        string
		infix       string
		infixResult []string
		wantErr     error
	}{
		{
			name:        "Test Infix To Postfix Success",
			infix:       "1 + 2.5 / 3 * 4.5",
			infixResult: []string{"1", "+", "2.5", "/", "3", "*", "4.5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Tokenization(tt.infix)
			if err != tt.wantErr {
				t.Errorf("Want: %v , got %v", tt.infixResult, result)
			}
			if convertToString(result) != convertToString(tt.infixResult) {
				t.Errorf("InfixToPostfix() = %v, want %v", result, tt.infixResult)
			}
		})

	}
}

func convertToString(s []string) string {
	return strings.Join(s, "")
}
