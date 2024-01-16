package cmd

import (
	"fmt"
	"os"
  "net/http"
  "encoding/json"
	dotenv "github.com/joho/godotenv"
)

type App struct {

UnlockedUrl string `json:"unlocked_url"`

UnlockedKey string `json:"unlocked_key"`

CanvasUrl string `json:"canvas_url"`

CanvasKey string `json:"canvas_key"`

}

var service = http.HandleFunc("/userservice", func(w http.ResponseWriter, r *http.Request) {
  var body, err = ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Println("Error reading body")
  }
  var app App
  app = json.Unmarshal(body, )
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})


func InitApp(unlocked_url string, unlocked_key string, canvas_url string, canvas_key string) *App {
  if os.Stat(".env") == nil {
    err := dotenv.Load()
    if err != nil {
  fmt.Println("Error loading .env file")
    }

  if unlocked_url == "" {
   unlocked_url = os.Getenv("UNLOCKED_URL") 
  }
  if unlocked_key == "" {
    unlocked_key = os.Getenv("UNLOCKED_KEY") 
  }
  if canvas_url == "" {
    canvas_url = os.Getenv("CANVAS_URL") 
  }
  if canvas_key == "" {
    canvas_key = os.Getenv("CANVAS_KEY") 
  }
  var app App = App{
    UnlockedUrl: unlocked_url,
    UnlockedKey: unlocked_key,
    CanvasUrl: canvas_url,
    CanvasKey: canvas_key,
  }
  return &app
}
