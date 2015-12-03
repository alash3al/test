package test

import (
	"github.com/alash3al/aof"
	"fmt"
	"os"
)

func run() {
	db, _ := aof.Open(os.Args[1], 0777)
	db.ReverseScan([]byte("\n"), func(d []byte, atEOF bool) bool {
		fmt.Println(string(d))
		return true
	})
}
