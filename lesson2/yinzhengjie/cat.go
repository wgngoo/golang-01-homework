package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s = os.Args[i]
		printFile(s)
	}

	fmt.Println("yinzhengjie")
}

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
