package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const baseURL = "https://jsonmock.hackerrank.com/api/articles?page=%d"

// fetchPage récupère une page donnée d'articles.
func fetchPage(page int, wg *sync.WaitGroup, ch chan<- []Article) {
	defer wg.Done()
	url := fmt.Sprintf(baseURL, page)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erreur lors du fetch page %d: %v\n", page, err)
		return
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Printf("Erreur de décodage page %d: %v\n", page, err)
		return
	}
	ch <- apiResp.Data
}
