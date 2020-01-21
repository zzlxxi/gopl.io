package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts)
			defer f.Close()
		}
	}

	for line, n := range counts {
		if len(counts) == 0 {
			os.Exit(1)
		}
		fmt.Printf("%d\t%s", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	fmt.Println("请输入内容(输入 exit 退出)：")
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if s == "exit" {
			break
		}
		counts[s]++
	}
}
