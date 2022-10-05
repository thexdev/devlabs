# Docker Basic Load Balancing

This experiment is aimed to demonstrate how to load balancing two services (Rest API) with an isolated environment.

## Tech-Stack

- Docker
- PHP (v8.1)
- Nginx (v1.23.1)

## Up & Running

Here we expect you have install Docker on your machine.

### 1. Clone the repository

```bash
git clone git@github.com:thexdev/devlabs.git
```

### 2. Build all images

Build the `proxy` image.

```bash
cd proxy; docker build -t proxy:0.1.0 .
```

Build the `service` image.

```bash
cd service; docker build -t service:0.1.0 .
```

### 3. Creating custom network

```bash
docker network create docker-basic-lb-network
```

### 4. Running all containers

Running the service containers.

```bash
docker run --name service-1 -e SERVICE_ID=1 -e PORT=8000 --network docker-basic-lb-network -d service:0.1.0
```

```bash
docker run --name service-2 -e SERVICE_ID=2 -e PORT=8001 --network docker-basic-lb-network -d service:0.1.0
```

Running the proxy cointainer.

```bash
docker run --name proxy-1 -p 80:80 --network docker-basic-lb-network -d proxy:0.1.0
```

## Testing

1. Open Postman
2. Hit endpoint http://localhost with GET method
