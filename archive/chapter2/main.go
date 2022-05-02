package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)

	f, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	// returnする時に実行
	defer f.Close()

	// ファイルに書き込み
	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to write message to file \n: %v", err)
	}

	fmt.Println("Done!")
}
