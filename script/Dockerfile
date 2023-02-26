FROM arm32v5/golang:1.20 as builder

RUN apt update -y && apt install -y libsdl2-dev

COPY . /app
WORKDIR /app/cmd/robot
RUN go mod tidy

WORKDIR /app
RUN make compile-ev3

FROM arm32v5/debian:buster-slim

COPY --from=builder /app/build/* /app/build/