package models

type Movie struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Overview     string  `json:"overview"`
	Release_Date string  `json:"release_date"`
	Genre_Ids    []int   `json:"genre_ids"`
	Vote_Average float64 `json:"vote_average"`
	Poster_Path  string  `json:"poster_path"`
	Media_Type   string  `json:"media_type"`
}

type Data struct {
	Page          int     `json:"page"`
	Results       []Movie `json:"results"`
	Total_Pages   int     `json:"total_pages"`
	Total_Results int     `json:"total_results"`
}
