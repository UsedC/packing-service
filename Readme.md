# Packing Service

Packing Service uses a coin change problem solution using DP to find an optimal amount of packs for items.

## How to run

You can just run the project using your IDE or you can build a binary or a Docker image.

```bash
make build
```

```bash
make docker-build
```

The server will then start. You can also source .env file or specify the ENV by hand. The config also comes with defaults.

## Usage

You can run unit tests for the calculator using

```bash
make test
```
---
After the server is started, you can use it's API.

This method will return the current pack sizes:

```bash
curl -X GET http://localhost:8080/pack-sizes
```
---
You can also specify pack sizes on the fly:
```bash
curl -X POST http://localhost:8080/set-pack-sizes -H "Content-Type: application/json" -d '{"pack_sizes": [250, 500, 1000, 2000, 10000, 5000]}'
```
---
To calculate the packs, use:
```bash
curl -X GET http://localhost:8080/calculate-packs -H "Content-Type: application/json" -d '{"total": 12001}'
```