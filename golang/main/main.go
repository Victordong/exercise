package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type favContextKey string

func main() {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	r, err := http.Get(url)
	if err != nil && r != nil {
		if r.StatusCode == http.StatusOK {
			body, _ := ioutil.ReadAll(r.Body)
			subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
			wg.Add(1)
			go showResp(subCtx, wg)
		}
		r.Body.Close()
	}
	fmt.Println(err)
	select {
	case <-ctx.Done():
		fmt.Printf("stop getting url:%s\n", url)
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("printing ", ctx.Value(favContextKey("resp")))
	select {
	case <-ctx.Done():
		fmt.Println("stop showing resp")
	}
}
