package cmd

type LinkDTO struct {
	LongURL     string `json:"longURL"`
	Description string `json:"description"`
	Expiration  string `json:"expiration"`
	ShortURL    string `json:"shortURL"`
}