package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed Hidden.mp3
var audio []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("OST_HMM.mp3", audio, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed file/a.txt
//go:embed file/b.txt
//go:embed file/c.txt
var file embed.FS

func TestMultipleEmbed(t *testing.T) {
	a, _ := file.ReadFile("file/a.txt")
	fmt.Println(string(a))

	b, _ := file.ReadFile("file/b.txt")
	fmt.Println(string(b))

	c, _ := file.ReadFile("file/c.txt")
	fmt.Println(string(c))
}

//go:embed file/*.txt
var path embed.FS

func TestEmbedPath(t *testing.T) {
	dir, _ := path.ReadDir("file")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("file/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}
