package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestWatchdog(t *testing.T) {
	wd := NewWD(2*time.Second, func() {
		fmt.Println("lol")
	})

	for i := 0; i < 10; i++ {
		wd.Reset()
		time.Sleep(time.Second)
	}
	time.Sleep(3 * time.Second)
}
