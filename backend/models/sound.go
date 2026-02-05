package models

type SongInfo struct {
	Count    int `json:"count"`
	Previous any `json:"previous"`
	Next     any `json:"next"`
	Results  []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"results"`
}

type SingleSong struct {
	Song   string `json:"name"`
	Artist string `json:"username"`
}

// MultipleSongs is simply a slice of SingleSong
type MultipleSongs []SingleSong
