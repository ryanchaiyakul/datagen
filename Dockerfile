# Get base image

FROM golang:1.12.0-alpine3.9

# Copy and prepare the files

RUN mkdir /app
COPY . /app
WORKDIR /app

# Generate 'datagen' binary

RUN go build

# Expose port 8080 for HTTP API

EXPOSE 8080

# Call 'datagen'

CMD ["/app/datagen"]
