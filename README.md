### 🚘 Reverse-proxy / LB using GO 

**Demo Project**
This project is for learning purposes only.

**[WIP]** Project is still in the work in progress phase.

`goinreverse` is a simple and configurable reverse-proxy and load balancer built with go.

#### Configuration

Configure the proxy easily to suit your needs through the `config.yaml` file:

```yaml
healthcheck: 5 # in seconds
servers:
  - name: andromeda
    url: http://localhost:8081/
  - name: milkyway
    url: http://localhost:8082/
```

Configure the `.env` file. 

```shell script
mv .exmple.env .env
```

I suggets you prefix each environment variable with `export` so that it can be easily populated in the environment with `source .env` we will be using this command throught the project.


#### Testing the proxy

###### Unit tests:

Unit tests are available for the heartbeat package which does the server ping and the main package. Running the unit tests can be achieved using:

- `go test ./...`

###### End-toEnd test:

The end to end test is rather primitive due to time constraints. However it is really effective to see the loadballancer in action.

**Running the test**

*Prerequisites*

- Docker

*Flow*

- `cd test/e2e`
- `docker-compose up --build`
- `cd ../..`
- `source .env`
- `go run main.go`
- *Notice in the console both servers are getting pinged*
- `docker down {container_id}`
- *Notice that the traffic gets redirected to the other server*

🤩 **Congrats** you are done!

##### Dependencies

1. `gopkg.in/yaml.v2` - used for parsing configuration file.
