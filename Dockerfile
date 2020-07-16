FROM registry.gitlab.com/fravega-it/arquitectura/godocker:1.13.3 AS build

WORKDIR /app

COPY main.go ./
#RUN go get -u github.com/gorilla/mux  &&  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /server-pipeline
RUN go mod download  && CGO_ENABLED=0  GOOS=linux go build -a -installsuffix cgo -ldflags '-extldfl'

# final stage
#FROM scratch AS runtime
FROM busybox AS runtime

COPY --from=build /server-pipeline /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 8000

CMD ["./server-pipeline"]
