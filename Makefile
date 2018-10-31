build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/flip cmd/flip.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/createSong cmd/createSong/createSong.go