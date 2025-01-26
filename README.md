# Weather Dashboard with AI Commentary

A modern weather dashboard application that combines real-time weather data with AI-generated insights. Built with Go and Temporal workflow engine, this project demonstrates a production-ready microservices architecture.

## Key Features
- Real-time weather data fetching using OpenWeatherMap API
- AI-powered weather commentary using OpenAI's GPT-3.5
- Durable workflow execution with Temporal
- Clean, responsive web interface
- Fault-tolerant design with graceful degradation

## Getting Started

### Prerequisites
- Go 1.x
- Temporal server running locally or accessible endpoint
- OpenWeatherMap API key
- OpenAI API key

### Configuration
Create a `config/config.yaml` file with your API keys:

```
weather:
  api_key: "your-openweathermap-api-key"

openai:
  api_key: "your-openai-api-key"
```

## Project Structure 

```
├── config/
│   ├── config.yaml              # Configuration file (not in version control)
│   └── config.example.yaml      # Example configuration
├── templates/
│   └── index.html               # Web interface template
├── weather/
│   ├── activity.go              # Weather API integration
│   ├── activity_test.go         # Weather activity tests
│   ├── ai_agent.go              # OpenAI integration
│   ├── ai_agent_test.go         # AI agent tests
│   ├── workflow.go              # Temporal workflow definition
│   └── workflow_test.go         # Workflow tests
├── worker/
│   └── main.go                  # Temporal worker implementation
├── main.go                      # Web server and main application entry
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
└── README.md                    # Project documentation
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/popand/weather-dashboard
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

The test suite includes:
- Unit tests for weather activity
- Unit tests for AI commentary generation
- Integration tests for workflows
- Mock servers for external API dependencies

## Technical Stack
- Backend: Go
- Workflow Engine: Temporal
- APIs: OpenWeatherMap, OpenAI
- Frontend: HTML/CSS with vanilla JavaScript
- Configuration: Viper

## Acknowledgments
- OpenWeatherMap for weather data
- OpenAI for AI capabilities
- Temporal for workflow engine
- The Go community for excellent tools and libraries 

## Testing

Run the test suite: 