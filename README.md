# load-balancer

## Step 2

Let's extend the load balancer to distribute incoming requests between two servers using the round-robin scheduling algorithm. Additionally, we'll ensure that the load balancer only forwards traffic to healthy servers.

In this modification, we introduced a backendURLs slice to hold the URLs of the backend servers. The currentIdx variable keeps track of the index of the last used backend server. The getNextBackend function implements the round-robin scheduling algorithm.

Ensure Health Checks:

To make the load balancer send traffic only to healthy servers, you need to implement health checks. One simple way is to check the availability of each backend server using a health check endpoint.

In this modification, we added a getNextHealthyBackend function to select the next healthy backend server based on the round-robin algorithm. The isBackendHealthy function performs a simple health check by sending a GET request to a health check endpoint (/health) on each backend server.

Ensure that your backend servers are running and accessible. Then, run your load balancer (lb.go), and you should be able to send requests to http://localhost/ and observe the responses alternating between the servers on port 8080 and 8081. The load balancer will route traffic only to healthy servers.
