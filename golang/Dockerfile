# TODO : image - version 
FROM golang AS base 

WORKDIR /app 

COPY . .

RUN go mod tidy 

FROM base AS dev 

WORKDIR /app

RUN go mod tidy 

CMD ["go", "run", "./main.go"]

FROM base as prod 

RUN go build -o ./build/main ./main.go 

CMD [ "./build/main" ]


