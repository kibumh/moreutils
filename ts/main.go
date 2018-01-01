package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()
		t := time.Now()
		fmt.Println(t.Format("1-2 03:04:05"), l)
	}
}
