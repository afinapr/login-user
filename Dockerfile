# Step 1 build executable binary
FROM golang:alpine as builder

# Install git
COPY . $GOPATH/src/login_user
WORKDIR $GOPATH/src/login_user

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /login_user .


FROM alpine:3.4
RUN apk update -y

EXPOSE 8080

# Copy our static executable
COPY --from=builder /login_user /login_user
ENTRYPOINT ["./login_user"]