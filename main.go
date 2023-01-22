package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func allMoviesController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(response).Encode(newSuccessResponse(getMovies()))

	if err != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
	}
}

func createMovieController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var movieWithoutId MovieWithoutId

	bodyRequestError := json.NewDecoder(request.Body).Decode(&movieWithoutId)

	if bodyRequestError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
		return
	}

	movie := addMovie(movieWithoutId)

	_ = json.NewEncoder(response).Encode(newSuccessResponse(movie))
}

func movieByIdController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	id, uuidError := uuid.Parse(params["id"])

	if uuidError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("the param id is not valid", 400))
		return
	}

	findedMovie, movieErr := getMovieById(id)

	if movieErr != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse(movieErr.Error(), 400))
		return
	}

	err := json.NewEncoder(response).Encode(newSuccessResponse(findedMovie))

	if err != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
		return
	}
}

func updateMovieByIdController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	id, uuidError := uuid.Parse(params["id"])

	if uuidError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("the param id is not valid", 400))
		return
	}
	var movieWithoutId MovieWithoutId

	bodyRequestError := json.NewDecoder(request.Body).Decode(&movieWithoutId)

	if bodyRequestError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
		return
	}

	updatedMovie, updateError := updateMovie(id, movieWithoutId)

	if updateError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse(updateError.Error(), 400))
		return
	}

	err := json.NewEncoder(response).Encode(newSuccessResponse(updatedMovie))

	if err != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
		return
	}
}

func deleteMovieByIdController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	id, uuidError := uuid.Parse(params["id"])

	if uuidError != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("the param id is not valid", 400))
		return
	}

	removedMovie, movieErr := removeMovie(id)

	if movieErr != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse(movieErr.Error(), 400))
		return
	}

	err := json.NewEncoder(response).Encode(newSuccessResponse(removedMovie))

	if err != nil {
		_ = json.NewEncoder(response).Encode(newErrorResponse("internal server error", 500))
		return
	}
}

func init() {
	addMovie(
		MovieWithoutId{
			Isbn:  "der2393",
			Title: "Test movie",
			Director: &Director{
				ID:        uuid.New(),
				Firstname: "John",
				Lastname:  "Doe",
			},
		},
	)

}

func main() {
	const port = 3000
	router := mux.NewRouter()

	router.HandleFunc("/movies", allMoviesController).Methods("GET")
	router.HandleFunc("/movies/{id}", movieByIdController).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovieByIdController).Methods("DELETE")
	router.HandleFunc("/movies", createMovieController).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovieByIdController).Methods("PUT")

	http.Handle("/", router)

	fmt.Printf("Server listen on port: %v", port)

	address := fmt.Sprintf(":%v", port)

	log.Fatal(http.ListenAndServe(address, router))
}
