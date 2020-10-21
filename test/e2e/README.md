### Simple end to end test

This is for the purpose to help in developing the proxy. It runs 2 simple NodeJS application servers, both with a `/healhCheck` endpoin and running on different ports.

- Server 1: `8081`
- Server 2: `8082`

For demo purposes we can stop the running server using `docker stop {container_id}` to see the traffic flow to the second server.

Further we can automate this. However, due to time constraints it's manual for the moment.
