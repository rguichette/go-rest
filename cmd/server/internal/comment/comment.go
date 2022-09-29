package comment

import (
	"context"
	"fmt"
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string // /article/slug --> /xyz is the slug
	Body   string
	Author string
}

// Service - is the struct on which all our logic will be build on top on
type Service struct {
	Store Store
}

// NewService - returns a new pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

//Q1:
