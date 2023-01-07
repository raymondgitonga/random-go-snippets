package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentError(t *testing.T) {
	testCases := []struct {
		id             int
		expectedError  error
		expectedResult string
	}{
		{
			id:             0,
			expectedError:  errors.New("error processing comment with id: 0, error: comment not found"),
			expectedResult: "",
		},
		{
			id:             1,
			expectedError:  nil,
			expectedResult: "comment",
		},
	}

	for _, tc := range testCases {
		t.Run("Test Comments", func(t *testing.T) {
			comment, err := GetPost(tc.id)

			assert.Equal(t, tc.expectedResult, comment)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestPostError(t *testing.T) {
	testCases := []struct {
		id             int
		expectedError  error
		expectedResult string
	}{
		{
			id:             0,
			expectedError:  errors.New("error processing post with id: 0, error: post not found"),
			expectedResult: "",
		},
		{
			id:             1,
			expectedError:  nil,
			expectedResult: "post",
		},
	}

	for _, tc := range testCases {
		t.Run("Test Posts", func(t *testing.T) {
			comment, err := GetComment(tc.id)

			assert.Equal(t, tc.expectedResult, comment)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
