package articles

type Article struct {
	Title       *string `json:"title"`
	StoryTitle  *string `json:"story_title"`
	NumComments *int    `json:"num_comments"`
}

type APIResponse struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []Article  `json:"data"`
}
