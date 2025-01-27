package main

import (
	"log"
	"weather-dashboard/weather"

	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Initialize viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading config:", err)
	}

	c, err := client.Dial(client.Options{
		HostPort: viper.GetString("temporal.host_port"),
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
