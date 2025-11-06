package articles

// Article représente un article unique récupéré depuis l'API.
type Article struct {
	Title       *string `json:"title"`
	StoryTitle  *string `json:"story_title"`
	NumComments *int    `json:"num_comments"`
}

// APIResponse représente la réponse complète d'une page de l'API.
type APIResponse struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []Article  `json:"data"`
}
