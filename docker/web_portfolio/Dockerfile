# Gunakan image golang versi terbaru sebagai base image
FROM golang:alpine as BuildStage

# Install some missing dependencies
RUN apk update && \
    apk add --no-cache git ca-certificates tzdata gcc libc-dev openssh-client && \
    update-ca-certificates

# Mengatur direktori kerja di dalam container
WORKDIR /go/src/github.com/salehrashid/portfolio-saleh/cmd/web_portfolio

# Menjalankan aplikasi ketika container dijalankan
CMD go run main.go