# Use image
FROM --platform=linux/amd64 golang:1-alpine AS build
# Create working directory
WORKDIR /app
# Copy files to working directory
COPY go.mod go.sum ./
# Download packages
RUN go mod download
# Copy rest of files
COPY . .
# Set architecture and build
RUN GOARCH=amd64 GOOS=linux go build -o main .

# Use image
FROM --platform=linux/amd64 alpine:3.20
# Create working directory
WORKDIR /app
# Copy built app
COPY --from=build /app/main /app/main
# Allow connections on 443
EXPOSE 443
# Run executable 
CMD [ "./main" ]

