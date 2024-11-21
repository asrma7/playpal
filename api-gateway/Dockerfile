# base image : specify the enviorment
FROM golang:1.23.2-alpine AS build
# Install Git
RUN apk add git
# working directory
WORKDIR /app
# copy go.mod and go.sum to working directory
COPY go.mod ./
COPY go.sum ./
# download dependencies
RUN go mod download
# copy rest of the files to working directory
COPY . .
# build the application
RUN go build -o ./out/dist ./cmd

#use a smaller base image for final image
FROM alpine:3.14
#copy executable from build stage
COPY --from=build /app/out/dist /app/dist
#set working directory
WORKDIR /app
#set entrypoint for the application
CMD ["/app/dist"]