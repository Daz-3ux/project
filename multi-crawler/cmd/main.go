package main

import (
	"fmt"

	"github.com/Daz-3ux/project/tree/main/multi-crawler/pkg/crawl"
)

func main() {
	config := crawl.Crawler{
		StartURL:       "https://movie.douban.com/top250",
		MaxDepth:       2,
		MaxConcurrency: 5,
	}

	results, err := crawl.Crawl(config)
	fmt.Println(results)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
