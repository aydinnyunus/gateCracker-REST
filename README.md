# GateCracker REST API

A RESTful API for gateCracker tool that provides default password information for Smart Locks in Turkey.

This API is required to run the [gateCracker](https://github.com/aydinnyunus/gateCracker) tool. The gateCracker application connects to this REST API running on `localhost:8080` to retrieve lock model information and default passwords.

## Installation & Run

```bash
# Clone the repository
git clone https://github.com/aydinnyunus/gateCracker-REST.git

# Navigate to the directory
cd gateCracker-REST

# Install dependencies
go mod tidy

# Run the API server
go run main.go
```

The API will be available at `http://localhost:8080`

## Structure

```
├── go.mod
├── go.sum  
├── simple.go
└── main.go
```

## API Endpoints

### GET /smartLocks
Get all smart locks

### GET /smartLocks/:id
Get a specific smart lock by ID

Example: `GET http://localhost:8080/smartLocks/1`

## Usage with gateCracker

1. Start this REST API server: `go run main.go`
2. Make sure the API is running on `http://localhost:8080`
3. Run the [gateCracker](https://github.com/aydinnyunus/gateCracker) tool which will connect to this API

## Blog Post

- Medium: https://sockpuppets.medium.com/bypassing-door-passwords-4004b8d7995
- GitHub Pages: https://aydinnyunus.github.io/2022/01/07/bypassing-door-passwords/

## Bug / Feature Request 👨‍💻

If you find a bug or would like to request a new feature, please open an issue [here](https://github.com/aydinnyunus/gateCracker-REST/issues/new).

## Connect with me! 🌐

[<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/linkedin.png" title="LinkedIn">](https://linkedin.com/in/yunus-ayd%C4%B1n-b9b01a18a/)       [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/github.png" title="Github">](https://github.com/aydinnyunus)     [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/instagram-new.png" title="Instagram">](https://instagram.com/aydinyunus_/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/twitter.png" title="Twitter">](https://twitter.com/aydinnyunuss)
