# notebook store

Microservices architecture for a laptop store backend with authentication, roles (admin/client), and inter-service communication through NATS.

## tools:
- NATS 
- PostgreSQL
- Redis

## setup:
```shell
cd authentication
docker compose build
docker compose up

cd store
docker compose build
docker compose up
```