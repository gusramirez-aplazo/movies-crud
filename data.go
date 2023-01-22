package main

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
)

var movies []Movie

func getMovies() []Movie {
	return movies
}

func addMovie(movie MovieWithoutId) Movie {
	movieToInsert := Movie{
		ID:       uuid.New(),
		Title:    movie.Title,
		Isbn:     movie.Isbn,
		Director: movie.Director,
	}
	movies = append(movies, movieToInsert)

	return movieToInsert
}

func removeMovie(id uuid.UUID) (Movie, error) {
	var movie Movie
	for idx, item := range movies {
		if item.ID == id {
			movie = item
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}

	if reflect.DeepEqual(movie, Movie{}) {
		return Movie{}, errors.New("id not found")
	}

	return movie, nil
}

func findMovieReference(id uuid.UUID) (*Movie, error) {
	var movie *Movie
	for index, item := range movies {
		if item.ID == id {
			movie = &movies[index]
			break
		}
	}

	if movie != nil && movie.ID == id {
		return movie, nil
	}

	return nil, errors.New("can not find movie")
}

func getMovieById(id uuid.UUID) (Movie, error) {
	var movie Movie
	ref, err := findMovieReference(id)

	if err == nil && ref != nil {
		movie = *ref
		return movie, nil
	}

	return Movie{}, errors.New("id not founded")
}

func updateMovie(id uuid.UUID, data MovieWithoutId) (Movie, error) {
	var updated Movie
	movie, err := findMovieReference(id)

	if err == nil && movie != nil {
		movie.Isbn = data.Isbn
		movie.Title = data.Title
		movie.Director = data.Director
		updated = *movie
		return updated, nil
	}

	return Movie{}, errors.New("id not founded")
}
