package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
)
func TopArticles(limit int) []string {

	resp, err := http.Get(fmt.Sprintf(baseURL, 1))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var first APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&first); err != nil {
		return nil
	}

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

	all := make([]Article, 0)
	for data := range ch {
		all = append(all, data...)
	}

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
			continue
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

	result := make([]string, 0, limit)
	for i := 0; i < limit && i < len(processed); i++ {
		result = append(result, processed[i].Name)
	}
	return result
}