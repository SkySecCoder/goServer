package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	message := "<h1>Placeholder goServer main page...</h1>"

	file, _ := os.Open("base.html")
	data, _ := ioutil.ReadAll(file)

	message = string(data)
	fmt.Println(message)
}
