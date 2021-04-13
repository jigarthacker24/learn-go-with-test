package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	reschan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			reschan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-reschan
		results[r.string] = r.bool
	}

	return results
}
