FROM golang:1.20-alpine AS builder

WORKDIR /app/

COPY ["go.mod", "go.sum", "./"]

RUN go mod download
COPY . .
RUN go build -o ./bin/app .

FROM alpine AS runner

ENV GOOS linux
ENV CGO_ENABLED 0

COPY --from=builder /app/bin/app /usr/bin/app
USER nobody:nobody
CMD [ "/usr/bin/app" ]
