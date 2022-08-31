package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"time"
)

func makeRequest(url string) {
	timeout := time.Duration(600 * time.Second)
	client := http.Client{Timeout: timeout}
	_, err := client.Get(url)
	if err != nil {
		log.Print("GET error:", err)
		return	
	}
	return
}

func thread(domain string, lines []string) []string {
	var data[]string
	for i:=0; i < 9; i++  {
		url := domain + lines[i]
		makeRequest(url)
		}
		return data
}

func main() {
	file, err := os.Open("0501/0501.txt")
	if err != nil{
		log.Fatalln(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stack []string
	for scanner.Scan() {
		stack = append(stack, scanner.Text())
	}
	workers := []string{
		"http://164.92.148.204:8080/?url=",
		"http://146.190.28.135:8080/?url=",
		"http://146.190.24.55:8080/?url=",
		"http://167.99.211.138:8080/?url=",
		"http://104.248.193.222:8080/?url=",
		"http://164.92.152.249:8080/?url=",
		"http://159.223.238.112:8080/?url=",
		"http://164.92.156.135:8080/?url=",
		"http://159.223.10.39:8080/?url=",
		"http://164.92.146.36:8080/?url="}
	
	go thread(workers[1], stack[0:7])
	go thread(workers[2], stack[7:14])
	go thread(workers[3], stack[14:21])
	go thread(workers[4], stack[21:28])
	go thread(workers[5], stack[28:35])
	go thread(workers[6], stack[35:42])
	go thread(workers[7], stack[42:49])
	go thread(workers[8], stack[49:56])
	go thread(workers[9], stack[56:63])
	
	thread(workers[0], stack[63:70])
}