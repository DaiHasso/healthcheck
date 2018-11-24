FROM golang:alpine as healthcheck_builder

RUN apk update -y && apk add upx

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-s -w" \
-o /healthcheck .

RUN upx --best --ultra-brute /healthcheck

FROM scratch

COPY --from=healthcheck_builder /healthcheck /healthcheck
