package apininja

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"m/lift/config"
	"m/lift/db"
	"net/http"
	"net/url"

	"github.com/jackc/pgx/v5"
)

const baseUrl = "https://api.api-ninjas.com/"

type ExercisesResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Muscle string `json:"muscle"`
	Equipment string `json:"equipment"`
	Difficulty string `json:"difficulty"`
	Instructions string `json:"instructions"`
}

type GetExercisesParams struct {
	Name string
	Type string
	Muscle string
}

func FetchExercises(p GetExercisesParams) ([]ExercisesResponse, error) {
	params := url.Values{}
	if p.Name != "" {
		params.Add("name", p.Name)
	}
	if p.Type != "" {
		params.Add("type", p.Type)
	}
	if p.Muscle != "" {
		params.Add("muscle", p.Muscle)
	}

	urlWithParams := baseUrl + "v1/exercises" + "?" + params.Encode()

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Add headers
	req.Header.Add("X-Api-Key", config.Secrets.NinjaApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing body: ", err)
		return nil, err
	}

	var responseObject []ExercisesResponse
	json.Unmarshal(body, &responseObject)

	return responseObject, nil
}

func GetCachedExercises(ctx context.Context, endpoint string) ([]ExercisesResponse, error) {
	var exercises []ExercisesResponse

	// Query for a value based on a single row.
	if err := db.QueryRow(ctx, "SELECT data FROM api_ninjas_exercises WHERE endpoint = $1",
			endpoint).Scan(&exercises); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("GetCachedExercises %s: uncached", endpoint)
			}
			return nil, fmt.Errorf("GetCachedExercises %s: %v", endpoint, err)
	}

	return exercises, nil
}

func SetCachedExercises(ctx context.Context ,endpoint string, data []ExercisesResponse) error {
  query := `INSERT INTO api_ninjas_exercises (endpoint, data) VALUES (@endpoint, @data)`
  args := pgx.NamedArgs{
		"endpoint": endpoint,
    "data": data,
  }

  _, err := db.Insert(ctx, query, args)
  if err != nil {
    return fmt.Errorf("unable to insert row: %v", err)
  }

  return nil
}
