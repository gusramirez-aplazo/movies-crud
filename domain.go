package main

import "github.com/google/uuid"

type MovieWithoutId struct {
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Movie struct {
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
	ID       uuid.UUID `json:"id"`
}

type Director struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

type MovieResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Content any    `json:"content"`
	Ok      bool   `json:"ok"`
}
