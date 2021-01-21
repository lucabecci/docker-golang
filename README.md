# Test of Docker with Golang.

> Simple example of api rest in docker without db.

## 1. Pre-requeriments 🛠

You will need this requeriments for good rendiment:

- Go > 10.X
- Docker

## 2. Use the project

### How to run th app with docker

1. In your terminal, navigate to the main directory.
2. Run `docker build . -t go-docker` to create image.
3. Run `docker run -p 3000:3000 go-docker` to run the image.

### How to run th app without docker

1. In your terminal, navigate to the main directory.
2. Run `go build -o main ./cmd/main.go` to create the build.
3. Run `go run ./` to run the build.


## 3. Author 🙎🏻‍♂️

***Luca Becci -** code and documentation*

- [github](https://github.com/lucabecci)
- [twitter](https://twitter.com/lucabecci)
- [linkedin](https://www.linkedin.com/in/luca-becci-b8044b198/)