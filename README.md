# CricRadio (Containerizing using Docker)
## Listen to the Commentary of Live Cricket Matches.

## Components of the Project

### ✨ MySQL Database

### ✨ Kafka
* Introduction to Kafka - https://www.youtube.com/watch?v=heR3I3Wxgro
* Running Kafka on Docker with Compose - https://www.youtube.com/watch?v=ncTosfaZ5cQ

### ✨ Kafka's REST API

### ✨ Go Microservice (Heart of the Project)
* Get Live Matches from ESPNCricInfo API and update to mySQL database.
* Scrape the Web page of each match to get ball-to-ball commentary.
* Create a unique topic in Kafka for each match and update ball-to-ball commentary to Kafka.


### ✨ Web Microservice (Face of the Project)
* Developed Using React.
* Display List of Live Matches.
* Pull the text Commentary from Kafka and convert it into speech.
* Play the commentary of the Selected Match. 

## Run the project on your local machine

### ✨ Install Docker & Docker Compose
### ✨ Clone the Code into your machine.
### ✨ Start the Docker Daemon
### ✨ change the directory to "CricRadio" & Run the following command
```
docker-compose up
```

### ✨ Go to localhost:3000 to listen to the commentary. It's that simple. If you have any problem, create an issue here.

## Building or Pulling Docker Images for each MicroService to run as a Container

### ✨ GO MicroService
* Write the Dockerfile
* Run the following command to build the Docker Image inside "cricradio-go-svc" directory
```
docker build -t suhasgumma/cricradio-go-svc:latest . 
```

* Create a Repository in Dockerhub & Push the Image.
 ```
 docker push suhasgumma/cricradio-go-svc:latest   
 ```
* Command to pull the image from Dockerhub
```
docker pull suhasgumma/cricradio-go-svc
```

### ✨ Web MicroService
* Write the Dockerfile
* Run the following commands to build the Docker Image inside "cricradio-web-svc" directory
```
docker build -t suhasgumma/cricradio-web-svc:latest . 
```

* Create a Repository in Dockerhub & Push the Image.
 ```
 docker push suhasgumma/cricradio-web-svc:latest   
 ```
* Command to pull the image from Dockerhub
```
docker pull suhasgumma/cricradio-web-svc
```
* Change the base image of node to "alpine" version to reduce the image size.
```
docker build -t suhasgumma/cricradio-web-svc:alpine . 
docker push suhasgumma/cricradio-web-svc:alpine  
```

### ✨ MySQL 
* Use the official Docker Image - https://hub.docker.com/_/mysql
* Set the root password by using environment variable "MYSQL_ROOT_PASSWORD"

### ✨ Kafka 
* Use the Confluent Community Docker Image for Apache Kafka.
* More Info here - https://hub.docker.com/r/confluentinc/cp-kafka/

## Using Volumes to Persist Data
* If the Container stops running for any reason, the data will be lost if not stored in volumes.
* In Our project, SQL databases and Kafka topics are vulnerable to data loss.
* To counter that, we use named volumes to persist data on the local machine. 

## Writing docker-compose.yml
* List all the containers as services.
* Set the required environment variables for each service.
* Map the ports required.
* Define volumes required for the service if needed.
* List all the volumes needed in the volumes section.

## Running Containers


