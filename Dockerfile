# Step 1 build executable binary
FROM golang:alpine as builder

# Install git
COPY . $GOPATH/src/login-user
WORKDIR $GOPATH/src/login-user

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /login-user .


FROM alpine:3.4
RUN apk update -y

RUN touch .env


# Copy our static executable
COPY --from=builder /login-user /login-user
ENTRYPOINT ["./login-user"]