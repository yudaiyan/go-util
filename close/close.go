package close

import (
	"io"

	"github.com/go-errors/errors"
)

// 用于匿名实现Closer接口
type closer struct {
	close func() error
}

func (c *closer) Close() error {
	if c.close != nil {
		return c.close()
	}
	return errors.New("close is nil")
}

func New(close func() error) io.Closer {
	return &closer{
		close: close,
	}
}
