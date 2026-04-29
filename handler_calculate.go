package main

import (
	"encoding/json"
	"net/http"
)

func handlerCalculate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Weight float64 `json:"Weight"`
		Reps   int     `json:"Reps"`
	}
	type response struct {
		OneRepMax float64 `json:"one_rep_max"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	oneRepMax := Calculate1RM(params.Weight, params.Reps)

	respondWithJSON(w, http.StatusCreated, response{
		OneRepMax: oneRepMax,
	})
}

// Calculate1RM returns the estimated one-rep max based on Brzycki's formula
func Calculate1RM(weight float64, reps int) float64 {
	if reps <= 1 {
		return weight
	}
	return weight / (1.0278 - (0.0278 * float64(reps)))
}
