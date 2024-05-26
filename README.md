# URL Shortener Service

## Description

A simple URL shortener service implemented in Go using Domain Driven Design (DDD) principles. The service is in-memory
and is deployed on Kubernetes with 2 replicas. Both replicas share a common Redis in-memory database to ensure
consistent data storage and retrieval across the instances.

## Requirements

- Go 1.22.3
- Docker
- Kubernetes
- Redis

## Running the Service

1. Clone the repository.
    ```sh
    git clone https://github.com/gorkemye/link-shortener
    ```
2. Build the Docker image:
    ```sh
    docker build -t link-shortener:latest .
    ```
3. Apply the Kubernetes deployment and service:
    ```sh
    kubectl apply -f k8s/deployment.yaml
    kubectl apply -f k8s/service.yaml
    kubectl apply -f k8s/redis-deployment.yaml
    kubectl apply -f k8s/redis-service.yaml
    ```
4. Kubernetes Utility (Optional)
* Get pods, services and jobs
    ```sh
    kubectl get pods
    kubectl get services
    kubectl get jobs 
   ```
* Delete the entire work environment
    ```sh
    kubectl delete pod --all
    kubectl delete deployment --all
    kubectl delete svc --all
    kubectl delete job --all

    ```



## API Endpoints

- `POST /api/v1/urls/shorten`: Shortens a long URL.
    - Request body:
        ```json
        {
            "long_url": "https://example.com"
        }
        ```
    - Response body:
        ```json
        {
            "short_url": "shortened_url"
        }
        ```
- `GET /api/v1/urls/{short_url}`: Retrieves the long URL for a given short URL.
    - Query parameter: `short_url`
    - Response body:
        ```json
        {
            "long_url": "https://example.com"
        }
        ```

## Running Tests

This section explains how to run tests on Kubernetes and check the results.

### Apply the Test Job

First, apply the test job by running the following command:

```sh
kubectl apply -f k8s/test-job.yaml
```
### Check Current Jobs
You can check the created and running jobs with the following command:

```sh
kubectl get jobs
```

### Check the Log of the Test Performed
To inspect the results of the tests by checking the logs of the test job, use the following command:
```sh
kubectl logs job/link-shortener-test
```
