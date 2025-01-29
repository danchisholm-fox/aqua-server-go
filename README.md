# ###############################################################
#
#   WEBSERVER BLUE - started around sept 2024
#
# NEXT STEPS
# DONE ~12.10 - use TF to build and start ec2 infra
# DONE ~12.26 - now need to be able to connect to ec2 instance, 
# DONE ~1.2 - put TF project on github
# DONE 1.11! - containerize the golang Aqua server
# 
# DANHERE 1.11.25 7pm: using TF, upload that container to ECR (SHOULD FIND OUT WHAT ECR DOES 1ST!!)
#         use https://registry.terraform.io/providers/calxus/docker/latest/docs to build the docker image
#   youtube - meh, it's TF containerizing an Nginx server, not a custom golang - https://www.youtube.com/watch?v=IY9-7XVmfKM
#   youtube2 - meh, not sure https://www.youtube.com/watch?v=5dhLy6kcBWQ&t=9s
#   youtube3 - GOOD ECR! https://www.youtube.com/watch?v=8XnqgiQaIkU
#     also GOOOD ECR! - https://www.youtube.com/watch?v=OBDiaKHK75c
# then start the EC2 instance using the container from ECR
# THAT'S IT - YOUR MVP IS DONE
#
# #######################       
#
#  POSTMVP - then set up a sidecar service to do something like send logs to s3 bucket
# then actually work on the API so it does SOMETHING, anything!
# then GHA to do some test validations
# then maybe Argo to do auto CD pulling from the GH repo
#
# ###############################################################



# OLD BELOW

test - dan chisholm - please ignore - 11sept2024

update 3 - i just did some stuffy stuff

update 4 - just messin with new commits 913pm PT


*** GITHUB INSTRUCTIONS ***

TO JUST PUSH A NEW UPDATE (with just the same files as last time)
1. git commit -am "enter your message here"
2. git push


TO SET UP NEW REPO
1. git init
1a. git remote add origin <URL where remote repo is>
2a. git rm README.md
2a. (only if there's stuff on repo that isnt local): git pull origin main
2b. git config pull.rebase false
3. (maybe?) git pull --ff-only


TO COMMIT NEW CODE
OVERALL: git add then commit then push
1. git add -A (or --all)
2. git commit -m "enter your message for the commit here"
3. git push -u origin main. (i think you only use '-u' if you are changing which branch/repo/env you are using since the last push)


WELL, TO DO NEW VS CODE REPO
1. 
2. i think you go to Mac Menu->Code->File->Open Folder
3. 