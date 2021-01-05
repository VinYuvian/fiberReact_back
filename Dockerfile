FROM golang:alpine3.12
LABEL maintainer="vin1711 <vinay.nuthipelly@gmail.com>"
WORKDIR /app
ARG BUILD_NUMBER
COPY /Fiber_"$BUILD_NUMBER" /app/Fiber
EXPOSE 3000
CMD ["./Fiber"]
