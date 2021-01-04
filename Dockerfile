FROM golang:1.14 as builder
LABEL maintainer="vin1711 <vinay.nuthipelly@gmail.com>"
WORKDIR /app
RUN apt update && apt install -y git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/Fiber

FROM golang:alpine3.12
COPY --from=builder /app/Fiber /app/Fiber
EXPOSE 3000
CMD ["./Fiber"]
