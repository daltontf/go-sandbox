package main

type Resource struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Status string `json:"status"`
}
