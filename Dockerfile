FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

COPY main.go .

EXPOSE 80

CMD [ "go", "run", "main.go" ]