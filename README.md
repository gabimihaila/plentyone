PlentyOne - Coding Challenge
API Gateway with Enhanced Metrics Logging

There are 3 main entities belongig to this project: a microservice, a client and an api gateway:

 - the microservice only displays the message "Hi! I'm a microservice!";

 - the client makes a GET request to the microservice and prints the body of the request;

 - the api gateway relays requests and logs metrics;

**Client part**

 Client makes a GET request to ```http://192.168.1.206:8084/serv0/home```. For Docker it was needed to use the ip address to make a request, because docker also virtualizes the network.
 For a better working change ```192.168.1.206``` with your ip address.
 Also make the same change in ```api_gateway/config.json``` file.

 There is also a bearer token for authorization purposes. The JWT secret for the token is stored in the ```config.json``` file from the ```apy_gateway``` directory.

 The Dockerfile for this is in the ```plentyone``` directory named as ```Dockerfile_client```.

**Microservice part**

 It displays the message "Hi! I'm a microservice!" on PORT 8081.

 The Dockerfile for this is in the ```plentyone``` directory named as ```Dockerfile_microservice```.

**API Gateway part**

 API Gateway makes possible the connection between the client and the microservice. It work on PORT 8084.

 In the ```config.json``` file is stored information for the project, such as: the API gateway port, the JWT secret and some destinations urls for the microservice (for scalability). Normally this file should remain in the local project, or stored, but not placed on Github, but for the technical challenge I made it "public".

 The ```gateway.log``` file stores all log data for the project, API metrics, errors, latency, and other metrics and information messages.

 The ```config``` directory with the ```config.go``` file helps for loading configuration from the ```config.json``` file.

 In the ```handlers``` directory there is the ```handle.go``` file that has a proxy handler which forwards request to the microservice.

 The ```jwt_auth``` directory with the ```auth.go`` file is a standard HTTP middleware for JWT authentication.

 The ```logger/logger.go``` logs details of different messages, such as errors, panic, infomation messages, warnings.

 In the ```request_manager/utils.go``` there are 2 functions, one for tracking the number of the requests per endpoint and logging the information, and the other one for tracking the latency of the requests.

 The Dockerfile for this is in the ```plentyone``` directory named as ```Dockerfile_api_gateway```.

<br>

The ```docker-compose.yml``` file stores configuration for the client, microservice, api-gateway and a postgres database.

The ```Makefile``` is used for building the client, microservice and api-gateway.

<br>

**Open the app by using the following commands in plentyone dir:**
```
$ make build
```
and

```
$ docker compose up -d
```
