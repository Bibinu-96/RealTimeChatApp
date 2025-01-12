package components

import "context"

type Component interface {
	Run(ctx context.Context) error
	GetName() string
}
