FROM golang:latest

WORKDIR /app
#Copy everything to the docker directory
COPY backend/ backend/
COPY database/ database/

WORKDIR /app/database
RUN go mod tidy

WORKDIR /app/backend
RUN go mod tidy

#create the executable (main)
RUN go build -o main_backend .

#I believe that i dont need this since i already have it on docker compose
#EXPOSE 8080
RUN chmod +x main_backend

CMD [ "./main_backend" ]
