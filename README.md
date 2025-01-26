# Weather Dashboard

A Go-based weather dashboard that uses Temporal workflows to fetch weather data from OpenWeatherMap API. The application demonstrates the use of Temporal for reliable API interactions and provides a simple web interface for users to check weather conditions.

## Features

- Real-time weather data fetching using OpenWeatherMap API
- Temporal workflow implementation for reliable API calls
- Automatic retries on API failures
- Clean web interface
- Unit tests for both activities and workflows

## Prerequisites

- Go 1.21 or later
- Temporal server (using remote instance at viaduct.proxy.rlwy.net:46280)
- OpenWeatherMap API key (provided in the code)

## Project Structure 

```
weather-dashboard/
├── main.go              # Web server and main application
├── worker/
│   └── main.go         # Temporal worker implementation
├── weather/
│   ├── activity.go     # Weather API activity
│   ├── workflow.go     # Temporal workflow definition
│   ├── activity_test.go # Activity tests
│   └── workflow_test.go # Workflow tests
└── README.md
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd weather-dashboard
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

1. Start the Temporal worker:
```bash
go run worker/main.go
```

2. In a separate terminal, start the web server:
```bash
go run main.go
```

3. Access the dashboard at `http://localhost:8088`

## Running Tests

Run all tests:
```bash
go test ./...
```

## API Integration

The application uses OpenWeatherMap API with the following endpoint:
```
https://api.openweathermap.org/data/2.5/weather?q={city}&appid={apiKey}&units=metric
```

## Temporal Workflow Details

- Task Queue: `weather-task-queue`
- Activity Timeout: 10 seconds
- Maximum Retry Attempts: 3

## Web Interface

The dashboard provides:
- City input form
- Weather display card showing:
  - Temperature (°C)
  - Feels Like temperature
  - Humidity (%)
  - Wind Speed (m/s)

## Error Handling

The application includes comprehensive error handling:
- API call failures
- Invalid city names
- Workflow execution errors
- Template rendering errors

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - feel free to use this code for your own projects. 