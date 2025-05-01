FROM golang:latest

WORKDIR /app
#Copy everything to the docker directory
COPY statistics/ statistics/

COPY database/ database/

WORKDIR /app/database
RUN go mod tidy


WORKDIR /app/statistics
RUN go mod tidy

#create the executable (main)
RUN go build -o main_statistics .

RUN chmod +x main_statistics

CMD [ "./main_statistics" ]
