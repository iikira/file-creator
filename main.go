package main

import (
	"flag"
	"log"
	"os"
)

var (
	size int64
	path string
)

func init() {
	flag.Int64Var(&size, "s", 256*1024, "size of file")
	flag.StringVar(&path, "p", "file.txt", "file path")
	flag.Parse()
}

func main() {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Truncate(path, size)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("file create succeed, path: %s, size: %d\n", path, size)
	file.Close()
}
