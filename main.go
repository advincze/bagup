package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, fi := range fileInfos {
		fmt.Printf("%#v \n", fi)
		if strings.HasPrefix(fi.Name(), ".") {
			fmt.Println("dotfile")
		}
	}
}
