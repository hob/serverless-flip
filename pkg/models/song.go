package models

type Song struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Artist   string   `json:"artist"`
	Key      string   `json:"key"`
	Duration int      `json:"duration"`
	Timing   Timing   `json:"timing"`
	Phrases  []Phrase `json:"phrases"`
}

type Timing struct {
	Upper int `json:"upper"`
	Lower int `json:"lower"`
}

type Phrase struct {
	Bars   float32 `json:"bars"`
	Note   string  `json:"note"`
	Repeat int     `json:"repeat"`
	Lyric  string  `json:"lyric"`
}
