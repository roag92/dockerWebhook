Docker Webhook
=======

A PoC for provisioning a Docker Image stored in DockerHub installing a Webhool agent to recreate container in every new image.

# Requirements
 
  - Go 1.13
  
## Configure

Create the `.env` file based on the `.env.dist`

```.env
SECRET_KEY=CREATE_A_SECRET_TOKEN
TAG_KEY=CREATE_A_SECRET_TOKEN
HOST_URL="http://localhost"
HOST_PORT=8080
```

Before to continue you should configure the dependencies of your Docker Image. You should create a new .env file and the docker-compose.yml file into the `config` directory.
