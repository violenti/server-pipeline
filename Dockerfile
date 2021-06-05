FROM golang:1.16.4-alpine3.13 AS build

WORKDIR /app

COPY . ./
#RUN go get -u github.com/gorilla/mux  &&  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /server-pipeline
RUN go mod download  
RUN go get -u github.com/gorilla/mux  &&  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /server-pipeline


# final stage
#FROM scratch AS runtime
FROM busybox AS runtime

COPY --from=build /server-pipeline /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 8000

CMD ["./server-pipeline"]
