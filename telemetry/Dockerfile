FROM golang:1.13.15 AS build

WORKDIR /app
COPY . /app/
RUN go build .

FROM scratch AS export
COPY --from=build /app/gaiad_telemetry ./content/gaiad_telemetry
