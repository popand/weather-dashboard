package main

import (
	"html/template"
	"log"
	"net/http"

	"weather-dashboard/weather"

	"go.temporal.io/sdk/client"
)

var weatherTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Weather Dashboard</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .weather-card { 
            border: 1px solid #ccc; 
            padding: 20px; 
            border-radius: 8px;
            margin-top: 20px;
        }
        .form-group { margin-bottom: 15px; }
        input[type="text"] { padding: 8px; width: 200px; }
        button { padding: 8px 16px; background-color: #007bff; color: white; border: none; border-radius: 4px; }
        .ai-message {
            background-color: #e3f2fd;
            border-left: 4px solid #2196f3;
            padding: 15px;
            margin-top: 15px;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <h1>Weather Dashboard</h1>
    <form method="POST">
        <div class="form-group">
            <label for="city">Enter City:</label>
            <input type="text" id="city" name="city" required>
        </div>
        <button type="submit">Get Weather</button>
    </form>
    {{if .Weather}}
    <div class="weather-card">
        <h2>Current Weather</h2>
		<p>Conditions: {{.Weather.City}}</p>
        <p>Temperature: {{.Weather.Temperature}}°C</p>
        <p>Conditions: {{.Weather.Conditions}}</p>
		
        {{if .Weather.AICommentary}}
        <div class="ai-message">
            <p>{{.Weather.AICommentary}}</p>
        </div>
        {{end}}
    </div>
    {{end}}
</body>
</html>
`

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "viaduct.proxy.rlwy.net:46280",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	tmpl := template.Must(template.New("weather").Parse(weatherTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			city := r.FormValue("city")
			if city == "" {
				http.Error(w, "City is required", http.StatusBadRequest)
				return
			}

			options := client.StartWorkflowOptions{
				TaskQueue: "weather-task-queue",
			}

			weatherInfo, err := c.ExecuteWorkflow(r.Context(), options, weather.WeatherWorkflow, city)
			if err != nil {
				http.Error(w, "Failed to execute workflow: "+err.Error(), http.StatusInternalServerError)
				return
			}

			var result weather.WeatherResult
			if err := weatherInfo.Get(r.Context(), &result); err != nil {
				http.Error(w, "Failed to get workflow result: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(w, struct{ Weather *weather.WeatherResult }{Weather: &result}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		tmpl.Execute(w, struct{ Weather *weather.WeatherResult }{Weather: nil})
	})

	log.Println("Server starting on :8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
