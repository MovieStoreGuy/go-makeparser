package parser

import (
	"bufio"
	"context"
	"errors"
	"io"

	"github.com/MovieStoreGuy/go-makeparse/pkg/internal/check"
)

type Parser interface {
	// Parser accepts an input to read and then will write the response
	ParseWithContext(ctx context.Context, in io.Reader) (*Graph, error)
}

type parser struct{}

var _ Parser = (*parser)(nil)

func (p *parser) ParseWithContext(ctx context.Context, in io.Reader) (*Graph, error) {
	if check.IsNil(ctx, in) {
		return nil, errors.New("invalid input parameters")
	}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Context has yet to be cancelled
		}
		// Need to read a line from the file
	}
	return &Graph{}, nil
}
