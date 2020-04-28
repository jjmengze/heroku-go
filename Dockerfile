From golang:1.13

ADD main.go ./
RUN go build -o main
CMD ./main
