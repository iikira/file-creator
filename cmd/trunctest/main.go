package main

import (
	"github.com/iikira/file-creator/trunctest"
	"time"
)

func main() {
	trunctest.TestTrunc2()
	time.Sleep(3e9)
}
