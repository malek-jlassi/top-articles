package articles

import "testing"

func TestTopArticlesReturnsResults(t *testing.T) {
    result := TopArticles(3)
    if len(result) != 3 {
        t.Errorf("Expected 3 articles, got %d", len(result))
    }
}