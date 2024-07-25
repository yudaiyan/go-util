package close

import (
	"log"
	"testing"
)

func TestCloser(t *testing.T) {
	c := New(func() error {
		log.Println("closed")
		return nil
	})
	c.Close()
}
