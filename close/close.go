package close

import (
	"io"

	"github.com/go-errors/errors"
)

// 用于匿名实现Closer接口
type AnonymousCloser struct {
	AnonymousClose func() error
}

func (c *AnonymousCloser) Close() error {
	if c.AnonymousClose != nil {
		return c.AnonymousClose()
	}
	return errors.New("close is nil")
}

func New(close func() error) io.Closer {
	return &AnonymousCloser{
		AnonymousClose: close,
	}
}
