package exercises

import (
	"encoding/json"
	"fmt"
	apininja "m/lift/services"
	"net/http"
)

func GetExercises(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get cached current endpoint data
	endpoint := r.URL.Path + "?" + r.URL.RawQuery
	cachedExercises, err := apininja.GetCachedExercises(ctx, endpoint)
	if err != nil {
		fmt.Println("Error getting cached exercises:", err)
	}

	if (len(cachedExercises) > 0) {
		json.NewEncoder(w).Encode(cachedExercises)
		return
	}

	fmt.Println("No cached response, fetching....")
	
	// if no cached endpoint, call apininja exercises api
	nameParam := r.URL.Query().Get("name")
	typeParam := r.URL.Query().Get("type")
	muscleParam := r.URL.Query().Get("muscle")
	
	resp, err := apininja.FetchExercises(apininja.GetExercisesParams{ 
		Name: nameParam,
		Type: typeParam,
		Muscle: muscleParam,
	})

	if err != nil {
		fmt.Println("Error calling apininja get exercises:", err)
	}

	// cache endpoint request and response for next request
	apininja.SetCachedExercises(ctx, endpoint, resp)

	// Encode the data as JSON and write it to the response
	json.NewEncoder(w).Encode(resp)
}