FROM golang:1.14.3-alpine AS build
WORKDIR /src

ENV USER=titanic
ENV UID=10001 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# Install GCC to run tests properly
RUN apk add build-base

# Start by copying dependencies first.
COPY vendor ./vendor
COPY go.mod go.sum ./

# Copy the code itself.
COPY . .

# Since $GOPATH=/go, install will build the binary under /go/bin
RUN go install ./cmd/server

FROM alpine:3.7

# Copy users
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

# Use an unprivileged user.
USER titanic:titanic

# Copy the binary generated in previous stage
COPY --from=build /go/bin/server /usr/local/bin

# Copy the templates, because they are loaded in runtime. The
# template folder used by the server is defined at
# `endpoint/config.go`.
COPY --from=build /src/html_templates /src/html_templates
ENTRYPOINT ["/usr/local/bin/server"]
EXPOSE 8080
