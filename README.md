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

