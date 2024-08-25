FROM golang:1.21 as builder

WORKDIR /app

COPY . .

# Run the make command
CMD ["make", "build-all"]