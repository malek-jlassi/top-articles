package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
)
// TopArticles returns the names of the top 'limit' articles sorted by comment count.
func TopArticles(limit int) []string {
	// Fetch the first page to get total number of pages
	resp, err := http.Get(fmt.Sprintf(baseURL, 1))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var first APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&first); err != nil {
		return nil
	}

	// Fetch all pages concurrently
	ch := make(chan []Article)
	var wg sync.WaitGroup
	for i := 1; i <= first.TotalPages; i++ {
		wg.Add(1)
		go fetchPage(i, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Aggregate all articles
	all := make([]Article, 0)
	for data := range ch {
		all = append(all, data...)
	}

	// Prepare and sort articles by comment count
	type namedArticle struct {
		Name        string
		NumComments int
	}
	var processed []namedArticle
	for _, a := range all {
		var name string
		if a.Title != nil {
			name = *a.Title
		} else if a.StoryTitle != nil {
			name = *a.StoryTitle
		} else {
			continue // skip if both are nil
		}
		numComments := 0
		if a.NumComments != nil {
			numComments = *a.NumComments
		}
		processed = append(processed, namedArticle{Name: name, NumComments: numComments})
	}
	sort.Slice(processed, func(i, j int) bool {
		return processed[i].NumComments > processed[j].NumComments
	})

	// Return top 'limit' article names
	result := make([]string, 0, limit)
	for i := 0; i < limit && i < len(processed); i++ {
		result = append(result, processed[i].Name)
	}
	return result
}