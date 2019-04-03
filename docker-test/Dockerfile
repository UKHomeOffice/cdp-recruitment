FROM golang as build

COPY app.go .

ENV GOOS=linux
ENV GOARCH=amd64 
ENV CGO_ENABLED=0

RUN go build -a -ldflags '-extldflags "-static"' -o /app

FROM scratch
COPY --from=build /app /app
CMD ["/app"]