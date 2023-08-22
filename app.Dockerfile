FROM golang:1.19

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY urls.txt ./
COPY .env ./

CMD ["go","run","main.go"]