

// A Tour of Go complete!

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}


// Crawl TODO
// DONE:同じurlを二度取ってこない様に変える
// DONE:並列実行した時にも同じurlを二度取ってこない実装に変える
// NOTE:mapを使ってキャッシュを保持できる
// NOTE:しかし、それだけでは並列実行時の安全性は保証されない
func Crawl(url string, depth int, fetcher Fetcher) {

	// NOTE:まずはmapでurlのキャッシュを作ってみる
	urlCache := make(map[string]bool)
	mux := new(sync.Mutex)

	crawl := func(string, int, Fetcher){}
	crawl = func(url string, depth int, fetcher Fetcher) {
		// Mutexで危ないアクセスを制限
		mux.Lock()
		if urlCache[url] {
			mux.Unlock()
			return
		}
		urlCache[url] = true
		mux.Unlock()

		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			crawl(u, depth-1, fetcher)
		}
		return
	}
	// 並列実行
	go crawl(url, depth, fetcher)
	crawl(url, depth, fetcher)
}

func main() {

	Crawl("http://golang.org/", 4, fetcher)

}

/* fake fetcher {{{ */

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fake data
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programing Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
/* fake fetcher }}} */

