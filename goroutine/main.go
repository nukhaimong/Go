package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup for goroutine synchronization

var wg sync.WaitGroup

func main() {

	start := time.Now()

	// wg.Add(1)
	// go uploadFile()
	wg.Go(uploadFile)
	// wg.Add(1)
	// go saveToDB()
	wg.Go(saveToDB)
	// wg.Add(1)
	// go sendEmail()
	wg.Go(sendEmail)

	wg.Wait()

	fmt.Println("all tasks completed")
	fmt.Println("Time Taken:", time.Since(start))
}

func uploadFile() {
	//defer wg.Done()
	fmt.Println("Uploading File...")
	time.Sleep(3 * time.Second)
	fmt.Println("File Uploaded")
}

func saveToDB() {
	//defer wg.Done()
	fmt.Println("Saving File to DB...")
	time.Sleep(1 * time.Second)
	fmt.Println("File Saved to DB")
}
func sendEmail() {
	//defer wg.Done()
	fmt.Println("Sending Email...")
	time.Sleep(1 * time.Second)
	fmt.Println("Email Sent")
}
