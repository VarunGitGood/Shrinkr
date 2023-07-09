package cmd

type LinkDTO struct {
	LongURL     string `json:"longURL"`
	Description string `json:"description"`
	Expiration  string `json:"expiration"`
	ShortURL    string `json:"shortURL"`
}

type LinkInfo struct {
	Key         string `json:"key"`
	Clicks      int    `json:"clicks,omitempty"`
	Passcode    string `json:"passcode,omitempty"`
	Description string `json:"description"`
	LongURL     string `json:"longURL"`
	Created     string `json:"created"`
	CreatedBy   string `json:"createdBy"`
	Expiration  int    `json:"expiration,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Joined   string `json:"joined"`
	Links    int    `json:"links"`
}
