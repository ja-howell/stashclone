package models

type StashItem struct {
	Name string `json:"item_name"`
	Type string `json:"item_type"`
	ID   int    `json:"id"`
}
