FROM golang as build

COPY app.go .

ENV GOOS=linux
ENV GOARCH=amd64 
ENV CGO_ENABLED=0

RUN go build -a -ldflags '-extldflags "-static"' -o /app
RUN useradd -u 1000 scratchuser

FROM scratch
COPY --from=build /app /app
COPY --from=build /etc/passwd /etc/passwd

CMD ["/app"]