package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入内容( 输入 exit 退出)：")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		if s == "exit" {
			break
		}
		counts[s]++
	}
	for line, n := range counts {
		if len(counts) == 0 {
			os.Exit(1)
		}
		fmt.Printf("%d\t%s", n, line)
	}
}
