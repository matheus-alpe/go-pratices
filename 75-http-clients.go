package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("init request to google")
	resp, err := http.Get("https://google.com")
	fmt.Println("end request to google")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
