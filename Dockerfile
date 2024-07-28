# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app/service

COPY service/go.mod ./
COPY service/go.sum ./
RUN go mod download

COPY service/*.json ./
COPY service/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o fittrackservice


# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app/service

COPY --from=build-stage service/fittrackservice service/fittrackservice

# FROM node:18-alpine

# WORKDIR /app/frontend

# COPY frontend/public/ ./
# COPY frontend/src/ ./
# COPY frontend/package.json ./

# RUN npm install 

# EXPOSE 5173
EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/fittrackservice"]

# CMD ["npm", "run", "dev"]