# superjuice-go

superjuice-go is a simple Go + HTMX app for calculating superjuice recipes.

## Run locally

### Using Air
    
```bash
go get -u github.com/cosmtrek/air
air
```

### Using Go

```bash
go run main.go
```

## Build

```bash
docker build --platform linux/amd64 -t 192.168.254.57:5000/superjuice:latest . && docker push 192.168.254.57:5000/superjuice:latest
docker stop superjuice-go
docker rm superjuice-go
docker pull localhost:5000/superjuice:latest
docker run -d --name superjuice-go -p 8000:8080 localhost:5000/superjuice
```

## License

MIT