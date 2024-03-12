package cors

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPatch,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Content-Type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Control"},
		Debug:              true,
	})
	return c
}

//func getNumber(num int, num2 int) int {
//	total := num + num2
//	return total
//}
