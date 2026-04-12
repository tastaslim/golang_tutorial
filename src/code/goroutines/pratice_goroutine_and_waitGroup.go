package goroutines

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func PrintNumber(num int) {
	fmt.Println(num)
}

func DemoGoroutine() {
	fmt.Println("Demo the Goroutines")
	waitGroup := &sync.WaitGroup{} // var waitGroup sync.WaitGroup
	for number := range 10 {
		waitGroup.Go(func() {
			PrintNumber(number)
		})
	}
	waitGroup.Wait()
}

type Post struct {
	UserID int    `json:"userId"` // This tag tells the JSON decoder to look for "userId" in the JSON and put it in the UserID field. Without this tag, it would look for "UserID". Same applicable for all such fields.
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func makeAPICall(ctx context.Context, i int) (Post, error) {
	apiUrl := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return Post{}, fmt.Errorf("Error occurred while created Http Request with context %w", err)
	}
	response, requestError := http.DefaultClient.Do(req)
	if requestError != nil {
		return Post{}, fmt.Errorf("Error occurred while created Http Request with context %w", requestError)
	}
	defer response.Body.Close()
	var post Post
	decodingError := json.NewDecoder(response.Body).Decode(&post)
	if decodingError != nil {
		return Post{}, fmt.Errorf("Error occurred while reading response body with context %w", decodingError)
	}
	return post, nil
}

func WaitGroupDemoWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	waitGroup := &sync.WaitGroup{}
	allResults := make([]Post, 9)
	for i := 1; i < 10; i++ {
		waitGroup.Go(func() {
			result, err := makeAPICall(ctx, i)
			if err != nil {
				fmt.Printf("Error for index %d: %v\n", i, err)
			}
			allResults[i-1] = result
		})
	}
	waitGroup.Wait()
	for _, post := range allResults {
		jsonData, err := json.Marshal(post) // This converts the Post struct into JSON format. The Marshal function returns a byte slice and an error.
		if err != nil {
			fmt.Printf("Error marshaling post: %v\n", err)
			continue
		}
		fmt.Printf("%s\n", jsonData)
	}
}
