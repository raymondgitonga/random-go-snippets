package errors

import (
	"fmt"
)

type CommentError struct {
	id           int
	errorMessage string
}

func (e CommentError) Error() error {
	return fmt.Errorf("error processing comment with id: %v, error: %s", e.id, e.errorMessage)
}

type PostError struct {
	id           int
	errorMessage string
}

func (p PostError) Error() error {
	return fmt.Errorf("error processing post with id: %v, error: %s", p.id, p.errorMessage)
}

func GetComment(id int) (string, error) {
	if id == 0 {
		return "", CommentError{
			id:           id,
			errorMessage: "comment not found",
		}.Error()
	}
	return "comment", nil
}

func GetPost(id int) (string, error) {
	if id == 0 {
		return "", PostError{
			id:           id,
			errorMessage: "post not found",
		}.Error()
	}
	return "post", nil
}
