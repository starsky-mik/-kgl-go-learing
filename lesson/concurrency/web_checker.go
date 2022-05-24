package concurrency

type WebsiteChecker func(url string) bool
type WebsiteCheckResult struct {
	url    string
	result bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	resultChanel := make(chan WebsiteCheckResult)

	for _, url := range urls {
		go func(u string) {
			resultChanel <- WebsiteCheckResult{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChanel
		result[r.url] = r.result
	}

	return result
}
