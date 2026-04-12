// From go version 1.25 onwards, the WaitGroup type has been enhanced to support a new method called Go. This method allows you to launch a goroutine
// and automatically manage the wait group counter without needing to call Add and Done separately. Here's how you can use it:
/*package goroutines

import (
	"fmt"
	"sync"
)

func PrintNumber(numberToPrint int, waitGroup *sync.WaitGroup) {
	// defer waitGroup.Done()
	fmt.Println(numberToPrint)
}

func main() {
	var waitGroup sync.WaitGroup
	num := 10
	for number := range num {
		waitGroup.Go(func() {
			PrintNumber(number, &waitGroup)
		})
	}
	waitGroup.Wait()
}
*/

package goroutines

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Demo() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Go(func() {
		req, err := http.NewRequestWithContext(ctx, "GET", "http://httpbin.org/delay/3", nil) // We need to make delay less than the context timeout to see the request complete successfully.
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		_, requestError := http.DefaultClient.Do(req)
		if requestError != nil {
			fmt.Println("Request cancelled:", requestError)
			return
		}
		fmt.Println("Request finished.")
	})

	wg.Wait()
	fmt.Println("Done.")
}
