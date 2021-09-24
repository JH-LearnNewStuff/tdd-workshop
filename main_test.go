package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const any = "DO_NOT_CHECK_OVERRIDE_VALUE"

func TestCalculatePrice(t *testing.T) {

	testCases := []struct {
		summary                 string
		books 					[]string
		expectedPrice           float64
		expectedErr             bool
	}{
		{
			summary:                 "no default kyma2 value file",
			books: make([]string, 0),
			expectedPrice: 0,
			expectedErr: false,
		},
		{
			summary:                 "Only single book",
			books: []string{ "1" },
			expectedPrice: 8,
			expectedErr: false,
		},
		{
			summary:                 "Multiple same book",
			books: []string{ "1", "1" },
			expectedPrice: 16,
			expectedErr: false,
		},
		{
			summary:                 "Two different book",
			books: []string{ "1", "2" },
			expectedPrice: 15.2,
			expectedErr: false,
		},
		{
			summary:                 "Another two different book",
			books: []string{ "1", "4" },
			expectedPrice: 15.2,
			expectedErr: false,
		},
		{
			summary:                 "Three books in a row",
			books: []string{ "1", "2", "3" },
			expectedPrice: 21.6,
			expectedErr: false,
		},
		{
			summary:                 "Three books in a rwo plus One",
			books: []string{ "1", "2", "3", "1" },
			expectedPrice: 29.6,
			expectedErr: false,
		},
		{
			summary:                 "a lot",
			books: []string{ "1", "2", "1", "2", "3" },
			expectedPrice: 36.8,
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.summary, func(t *testing.T) {
			t.Parallel()
			actual, err := calculatePrice(tc.books)
			require.NoError(t, err)
			require.Equal(t, tc.expectedPrice, actual)
		})
	}
}