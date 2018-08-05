package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type Fib struct {
	n   int
	res int
}

func worker(wg *sync.WaitGroup, num int, jobs <-chan Fib, results chan<- Fib, stopCh <-chan struct{}) {
	defer wg.Done()
	fmt.Printf("worker %d started work \n", num)
	for {
		select {
		case f := <-jobs:
			res := fib(f.n)
			f.res = res
			fmt.Printf("worker %d counted fib(%d) = %d \n", num, f.n, f.res)
			results <- f
		case <-stopCh:
			fmt.Printf("worker %d stoped work \n", num)
			return
		}
	}

}

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	jobNum := flag.Int("job", 20, "number of jobs to run")
	workerNum := flag.Int("worker", 10, "number of workers")

	jobs := make(chan Fib, *jobNum)
	results := make(chan Fib, *jobNum)
	stopCh := make(chan struct{})

	wg := sync.WaitGroup{}
	//start workers
	for i := 0; i < *workerNum; i++ {
		wg.Add(1)
		go worker(&wg, i, jobs, results, stopCh)
	}

	//run jobs
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("Exit jobs goroutine")
				return
			default:
				f := Fib{n: randInt(10, 20)}
				jobs <- f
				time.Sleep(100 * time.Millisecond)
			}

		}
	}()

	go func() {
		for res := range results {
			fmt.Printf("For n=%d fib = %d\n", res.n, res.res)
		}
		fmt.Println("Exit results goroutine")
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt for shutdown
	select {
	case sig := <-sigC:
		fmt.Printf("Shutdown request (signal: %v)\n", sig)
	}

	//finish workers as want to shutdown
	close(stopCh)
	wg.Wait()
	close(results)
	fmt.Println("Exit ...")

}
func fib(n int) int {
	if n <= 2 {
		return 1
	}

	var x = 1
	var y = 1
	var ans = 0
	for i := 2; i < n; i++ {
		ans = x + y
		x = y
		y = ans
	}
	return ans
}
