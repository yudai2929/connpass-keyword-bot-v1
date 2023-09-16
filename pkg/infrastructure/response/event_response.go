package response

type EventResponse struct {
	ResultsReturned  int `json:"results_returned"`
	ResultsStart     int `json:"results_start"`
	ResultsAvailable int `json:"results_available"`
	Events           []struct {
		EventID     int    `json:"event_id"`
		Title       string `json:"title"`
		Catch       string `json:"catch"`
		Description string `json:"description"`
		EventURL    string `json:"event_url"`
		HashTag     string `json:"hash_tag"`
		StartedAt   string `json:"started_at"`
		EndedAt     string `json:"ended_at"`
		Limit       int    `json:"limit"`
		EventType   string `json:"event_type"`
		Series      struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"series"`
		Address          string `json:"address"`
		Place            string `json:"place"`
		Lat              string `json:"lat"`
		Lon              string `json:"lon"`
		OwnerID          int    `json:"owner_id"`
		OwnerNickname    string `json:"owner_nickname"`
		OwnerDisplayName string `json:"owner_display_name"`
		Accepted         int    `json:"accepted"`
		Waiting          int    `json:"waiting"`
		UpdatedAt        string `json:"updated_at"`
	} `json:"events"`
}
