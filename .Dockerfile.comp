FROM golang:1.14
LABEL maintainer="vin1711 <vinay.nuthipelly@gmail.com>"
WORKDIR /app
RUN apt update && apt install -y git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 3000
RUN go build
CMD ["./Fiber"]



