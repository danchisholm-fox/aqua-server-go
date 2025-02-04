###########################################################
# DAN 1.11.25: attempting to containerize my Go server
#    and finally correct VS Code project - new year ❄️ ⛄️⛄️⛄️⛄️⛄️
# reminder to COMMENT MULTIPLE LINES, do CMD + /
#
# TUTORIAL VIDEO: i dont remember (the one where i learned the 'tag' command)
#  nuebooooa
#
# HOW TO BUILD AND RUN
# BUILD 🟢🟢>> docker build -t aqua-server-1 ../aqua-server-go
# RUN   🟢🟢>> docker run -d -p 127.0.0.1:8062:8080 aqua-server-5  # but dan do i need this if i'm running it on ECS?  i think this is just for localhost... wasnt it?
#                 fyi, To Run a container in the above, 8062 is the host's port, so 8080 is the container's. running the v5 image of aqua
# TAG   🟢🟢>> docker tag aqua-blue5-container:latest 535002851677.dkr.ecr.us-east-1.amazonaws.com/aqua/aqua-ecr:latest
#                  TODO 1.15.24: but how do i get that ECR URI?  i think i have to go to the ECR console and copy it from there
#                   which i wont have access to until i create the ECR repo... so how do i do this with automation?
#                   which youtube video or TF doc did i read that showed me how to do this?
# PUSH  🟢🟢>> docker push 535002851677.dkr.ecr.us-east-1.amazonaws.com/aqua/aqua-ecr:latest
# uh maybe also 🟢🟢>> aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 535002851677.dkr.ecr.us-east-1.amazonaws.com
#
# OTHER COMMANDS: docker ps, docker images, docker stop <id>, docker rm <id>, docker rmi <id>
#
#
###########################################################

# FROM golang:1.18-alpine AS build
FROM golang:1.23-alpine

# WORKDIR is the directory in the container where the commands will be run (not locally)
WORKDIR /aqua-server

# Copy the Go module files
COPY go.* ./

# Download the dependencies
RUN go mod download

COPY . . 

# Build the application
# RUN go build -o /aqua-server/main
RUN go build -o main main.go

# Expose the port your application will listen on
# DAN:  the 1st port number is the port on the actual host machine, like localhost or ec2 instance
#       and the 2nd port is the port of the container (ie the one in EXPOSE):
#            docker run -d -p 127.0.0.1:8074:8017 aqua-server3
#
# Reminder: this port is the one that the server is listening on, so i have to change it if i change this one below
EXPOSE 8072

# Command to run when the container starts
# CMD ["/aqua-server/main"]
CMD ["./main"]

# CMD ["go", "run", "main.go"]

