package models

type Skill struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	TagId int    `json:"tag_id"`
}
