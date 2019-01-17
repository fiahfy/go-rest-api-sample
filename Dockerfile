FROM golang

RUN go get github.com/pilu/fresh

CMD ["fresh"]
