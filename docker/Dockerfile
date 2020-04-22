FROM golang:latest 
WORKDIR /app

# copy to container
COPY . .

# build the app
RUN go build

EXPOSE 8000

ENTRYPOINT [ "./docker" ]