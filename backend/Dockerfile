FROM golang:alpine3.19

# Enviornment variables
ENV APP_NAME backend
ENV PORT 8080


RUN apk update

# Open system port
EXPOSE ${PORT}

# Working directory
WORKDIR /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}

# Install dependecies from mod file
RUN go mod download

# Build application
RUN go build -o ${APP_NAME} ./cmd/backend

# Run application
CMD ./${APP_NAME}
