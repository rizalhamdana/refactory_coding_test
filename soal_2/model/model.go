package model

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Profile  Profile   `json:"profile"`
	Articles []Article `json:"articles:"`
}

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}

type Profile struct {
	FullName string   `json:"full_name"`
	BirthDay string   `json:"birthday"`
	Phones   []string `json:"phones"`
}
