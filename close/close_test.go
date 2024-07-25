package close

import (
	"io"
	"log"
	"testing"
)

// 测试匿名实现Close方法
func TestCloser(t *testing.T) {
	c := New(func() error {
		log.Println("closed")
		return nil
	})
	c.Close()
}

type WriteCloser1 struct {
}

func (c *WriteCloser1) Close() error {
	log.Println("Close WriteCloser1")
	return nil
}

func (c *WriteCloser1) Write(p []byte) (n int, err error) {
	log.Println("Write WriteCloser1")
	return 0, nil
}

type WriteCloser2 struct {
	io.WriteCloser
	AnonymousCloser
}

// 由于 WriteCloser、AnonymousCloser 都有Close，因此重写以指定 AnonymousCloser.Close
func (c *WriteCloser2) Close() error {
	return c.AnonymousCloser.Close()
}

// 测试匿名重写Close方法
func TestOverrideClose2(t *testing.T) {
	c := &WriteCloser2{
		WriteCloser: &WriteCloser1{},
		AnonymousCloser: AnonymousCloser{
			AnonymousClose: func() error {
				log.Println("closed123")
				return nil
			},
		},
	}
	c.Close()
}
