package main

import (
	"github.com/kyma-incubator/reconciler/pkg/reconciler/workspace"
	"github.com/stretchr/testify/require"
	"testing"
)

const any = "DO_NOT_CHECK_OVERRIDE_VALUE"

func TestCalculatePrice(t *testing.T) {

	testCases := []struct {
		summary                 string
		books 					[]string
		expectedPrice           int
		expectedErr             bool
	}{
		{
			summary:                 "no default kyma2 value file",
			books: make([]string, 0),
			expectedPrice: 0,
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