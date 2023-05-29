package oxford

import (
	"fmt"
	"testing"

	"github.com/feckmore/xpwd/domain"
)

func TestGenerateRandomWord(t *testing.T) {
	testCases := []struct {
		name   string
		minIn  int
		maxIn  int
		minOut int
		maxOut int
		errOut error
	}{
		{"min and max are 0", 0, 0, 1, 16, nil},
		{"min is 0 and max is 1", 0, 1, 0, 0, domain.ErrNotEnoughWords},
		{"min is 1 and max is 1", 1, 1, 0, 0, domain.ErrNotEnoughWords},
		{"min is 1 and max is 4", 1, 4, 1, 4, nil},
		{"min is 1 and max is 16", 1, 16, 1, 16, nil},
		{"min is 1 and max is 17", 1, 17, 0, 0, domain.ErrInvalidWordLength},
		{"min is 2 and max is 16", 2, 16, 2, 16, nil},
		{"min is -1 and max is 16", -1, 16, 1, 16, nil},
		{"min is 0 and max is -1", 0, -1, 0, 0, domain.ErrInvalidWordLength},
		{"min is 10 and max is 3", 10, 3, 0, 0, domain.ErrInvalidWordLength},
		{"min is 10 and max is 3", -9, -1, 0, 0, domain.ErrInvalidWordLength},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			generator, gotErr := New(tc.minIn, tc.maxIn)
			if fmt.Sprint(gotErr) != fmt.Sprint(tc.errOut) {
				t.Fatalf("New(%d, %d) returned error: %v, expected error: %v", tc.minIn, tc.maxIn, gotErr, tc.errOut)
			}

			if generator == nil {
				t.SkipNow()
			}

			word := generator.GenerateRandomWord()
			if len(word) < tc.minOut || len(word) > tc.maxOut {
				t.Errorf("GenerateRandomWord() returned word of length %d, expected length between %d and %d", len(word), tc.minOut, tc.maxOut)
			}
		})
	}
}
