# FROM golang:alpine AS build
FROM golang AS build
WORKDIR /src
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# RUN apk add --update --no-cache ca-certificates git
# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
# COPY go.* .
COPY go.mod .
COPY go.sum .
RUN  go mod download
COPY . .
RUN  go build -v -o /usr/local/bin/idle

# FROM alpine AS bin
# ENV GIN_MODE=release
# RUN apk add -U tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# COPY --from=build /usr/local/bin/idle /
# EXPOSE 3000
# CMD ["./idle"] 


FROM scratch as bin
ENV GIN_MODE=release \
    ZONEINFO=/zoneinfo.zip \
    TZ=Asia/Shanghai
# COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/share/
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /usr/local/bin/idle /
EXPOSE 3000
CMD ["./idle"] 