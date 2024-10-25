package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	urls := []string{
		"http://localhost:8080/users",
		"http://localhost:8080/products",
		"http://localhost:8080/orders",
	}
	result := make(chan string)
	for _, url := range urls {
		go fetchAPI(ctx, url, result)
	}

	for range urls {
		fmt.Println(<-result)
	}

}

func fetchAPI(ctx context.Context, url string, result chan string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- fmt.Sprintf("create request error, url = %s, err = %v", url, err)
		return
	}
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("making request error, url = %s, err = %v", url, err)
		return
	}
	defer response.Body.Close()
	result <- fmt.Sprintf("response content, url = %s, tatus = %d", url, response.StatusCode)

}
