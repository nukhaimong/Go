package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)

	go uploadFile(ch)

	fileUrl := <-ch
	fmt.Println("File URL is:", fileUrl)

}

func uploadFile(c chan string) {
	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("Uploaded File")

	fileUrl := "https://example.com/file.jpg"

	c <- fileUrl
}
