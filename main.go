package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed Hidden.mp3
var audio []byte

//go:embed file/*txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("OST_HMM.mp3", audio, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("file")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("file/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
