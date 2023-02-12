package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	// viper.AutomaticEnv()
	// fmt.Println("config file", viper.GetString("env"))
	// fmt.Println("env", viper.Get("NAME"))
	// fmt.Println("env os", os.Getenv("NAME"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// func init() {
// 	viper.SetConfigName("config") // name of config file (without extension)
// 	viper.AddConfigPath(".")      // optionally look for config in the working directory
// 	err := viper.ReadInConfig()   // Find and read the config file
// 	if err != nil {               // Handle errors reading the config file
// 		panic(fmt.Errorf("fatal error config file: %w", err))
// 	}
// }
