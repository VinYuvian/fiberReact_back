FROM golang:alpine3.12
LABEL maintainer="vin1711 <vinay.nuthipelly@gmail.com>"
WORKDIR /app
COPY /Fiber /app/Fiber
EXPOSE 3000
CMD ["./Fiber"]
