# plentyone
PlentyOne - Coding Challenge
API Gateway with Enhanced Metrics Logging

There are 3 main entities belongig to this project: a microservice, a client and an api gateway.

The microservice only displays the message "Hi! I'm a microservice!"

The client makes a GET request to the microservice and prints the body of the request.

The api gateway relays requests and logs metrics.

Open the app by using the following commands in plentyone dir:

make build

docker compose up -d