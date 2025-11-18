package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (string, []string, error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println("error here, ", err)
		return
	}

	fmt.Printf("For the body, %s\n", body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}

}

func main() {
	Crawl("https://www.google.com", 3, fetcher)
}

type fakeFetcher map[string]*fakeResults

type fakeResults struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("Not found %s", url)
}

var fetcher = fakeFetcher{
	"https://www.google.com": &fakeResults{
		"Enter google",
		[]string{"https://www.apple.com", "https://www.nvidia.com"},
	}, "Round 2": &fakeResults{
		"https://www.nvdia.com",
		[]string{"https://www.apple.com", "https://www.google.com"},
	}, "https://www.apple.com": &fakeResults{
		"Enter Apple",
		[]string{"https://www.nvdia.com", "https://www.google.com"},
	},
}
