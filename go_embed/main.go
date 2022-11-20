package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed files/*.json
var data []byte

//go:embed files/*.txt
var fs embed.FS

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	embedSingleFile()
	embedMultipleFiles()
}

func embedSingleFile() {
	user := &User{}

	_ = json.Unmarshal(data, user)

	fmt.Println(user)
}

func embedMultipleFiles() {
	names, _ := fs.ReadFile("files/names.txt")
	fmt.Println(string(names))

	places, _ := fs.ReadFile("files/places.txt")
	fmt.Println(string(places))
}
