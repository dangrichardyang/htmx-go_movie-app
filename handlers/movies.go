package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dangrichardyang/htmx-go_movie-app/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	var data models.Data
	apiKey := os.Getenv("API_KEY")
	apiUrl := os.Getenv("API_BASE_URL")
	trendingUrl := fmt.Sprintf("%v/trending/movie/day?api_key=%v", apiUrl, apiKey)
	res, err := http.Get(trendingUrl)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		mErr := json.Unmarshal(bodyBytes, &data)
		if mErr != nil {
			return mErr
    	}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.SendString(string(jsonData))
}