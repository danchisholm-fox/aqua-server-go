###########################################################
# DAN 1.11.25: attempting to containerize my Go server
#    and finally correct VS Code project - new year ❄️ ⛄️⛄️⛄️⛄️⛄️
# reminder to COMMENT MULTIPLE LINES, do CMD + /
#
# TUTORIAL VIDEO: i dont remember (the one where i learned the 'tag' command)
#  nuebooooa
#
# 5.5.25 HOW TO BUILD AND RUN AQUA on AWS infra (ie in the cloud) 
# 🌹 INFRA PREP ⚠️ before doing TF apply, make sure this resource is commented out (else it'll try to build and fail over and over since there's no ECR container to start)
#    this line -->    # resource "aws_ecs_service" "aqua-ecs-service" 
# 🌹 INFRA     🟢🟢>>terraform apply -var-file variables.tfvars -auto-approve
# 🌹 CONTAINER 🟢🟢>> docker build -t aqua-server-1 ../aqua-server-go
#               ⚠️ something boke on the previous step - i think you have to run it from that folder
#                   ok so the go.mod file got messed up, i restored it which was just like 2-3 lines of code
# 🌹 RUN      >> docker run -d -p 127.0.0.1:8072:8072 aqua-server-1   # only if you want to run this locally
#                 fyi, To Run a container in the above, 8072 is the host's port, so 8072 is the container's. running the v5 image of aqua
# 🌹 TAG       🟢🟢>> docker tag aqua-server-1:latest 535002851677.dkr.ecr.us-east-1.amazonaws.com/aqua/aqua-ecr-repo-name:latest
#                  TODO 1.15.24: but how do i get that ECR URI?  i think i have to go to the ECR console and copy it from there
#                   which i wont have access to until i create the ECR repo... so how do i do this with automation?
#                   which youtube video or TF doc did i read that showed me how to do this?
# 🌹 LOGIN ECR 🟢🟢>> aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 535002851677.dkr.ecr.us-east-1.amazonaws.com
# 🌹 PUSH      🟢🟢>> docker push 535002851677.dkr.ecr.us-east-1.amazonaws.com/aqua/aqua-ecr-repo-name:latest
# 🌹 5.5.25 LAUNCH ECS:  uncomment this resource (resource "aws_ecs_service" "aqua-ecs-service") in the main.tf in the other project, then re-run "terraform apply"
# CONNECT: to connect from a client, get the custom domain  in ECS and navigate in Chrome to the url, port 8072
#    ugh, something in the container is crashing.  look at ChatGPT on 5.5.25: https://chatgpt.com/c/68181c91-915c-8008-b16e-e4cb03ab9344
#     i stopped at 7pm and just found out it was crashing when i run the container locally - looks like it is tagged wrong. missing the ":latest"
#
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

