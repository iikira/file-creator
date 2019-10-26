package trunctest_test

import (
	"fmt"
	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
	"github.com/iikira/file-creator/trunctest"
	"os"
	"syscall"
	"testing"
	"time"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procSetFileValidData = kernel32.NewProc("SetFileValidData")
)

func TestTrunc1(t *testing.T) {
	filename := "file.txt"
	func() {
		f, err := os.Create(filename)
		if err != nil {
			t.Fatalf("create error: %s\n", err)
		}
		defer f.Close()

		err = f.Truncate(2128 * converter.MB)
		if err != nil {
			t.Fatalf("trunc error: %s\n", err)
		}

		start := time.Now()
		_, err = f.WriteAt([]byte("111"), 227*converter.MB)
		if err != nil {
			t.Fatalf("write error: %s\n", err)
		}

		now := time.Now()
		fmt.Printf("time elapse: %s\n", now.Sub(start))
	}()

	err := os.Remove(filename)
	if err != nil {
		t.Fatalf("remove error: %s\n", err)
	}
}

func TestTrunc2(t *testing.T) {
	trunctest.TestTrunc2()
}
