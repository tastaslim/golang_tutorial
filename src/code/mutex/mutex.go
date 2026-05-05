package main

import (
	"fmt"
	"sync"
)

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	counter := 0
// 	var wg sync.WaitGroup

// 	for range 1000 {
// 		wg.Go(func() {
// 			counter++ // ❌ RACE ==> Result of this could be different on running different times because multiple g routines of trying to increment same variable at the same time
// leading to race condition(if mutex is not used). This can be resolved by taking locks, hence in go world , mutex
// 		})

// 	}
// 	wg.Wait()
// 	fmt.Println(counter)
// }

type Post struct {
	views int
	mut   sync.Mutex
}

func (p *Post) increment() {
	p.mut.Lock()
	defer p.mut.Unlock() // guaranteed to run even on panic
	p.views += 1
}

func (p *Post) getViews() int {
	return p.views
}

func main() {
	myPost := Post{views: 0}
	waitGroup := sync.WaitGroup{}
	for range 100 {
		waitGroup.Go(func() {
			myPost.increment()
		})
	}
	waitGroup.Wait()
	fmt.Println(myPost.getViews()) // Always consistent and 100
}
