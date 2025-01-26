package main

import (
	"log"

	"weather-dashboard/weather"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "viaduct.proxy.rlwy.net:46280",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	w := worker.New(c, "weather-task-queue", worker.Options{})

	w.RegisterWorkflow(weather.WeatherWorkflow)
	w.RegisterActivity(weather.GetWeatherActivity)
	w.RegisterActivity(weather.GetAICommentaryActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker:", err)
	}
}
