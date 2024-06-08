# Harvester API
## A fast modern API for storing your resources

### Running Locally

Prerequisites:
  * Go 1.22.4
  * Docker
  * make

```bash
# Start the API
  make compose-up
```

```bash
# Stop and destory the API
  make compose-down
```

### API Endpoints

When running locally host is available at `localhost:3000`

Users:

| Method | Endpoint           |
|--------|--------------------|
| POST   | /v1/users/         |
| GET    | /v1/users/         |
| GET    | /v1/users/:id      |
| PUT    | /v1/users/:id      |
| PUT    | /v1/users/role/:id |
| DELETE | /v1/users/:id      |

Galaxies:

| Method | Endpoint                |
|--------|-------------------------|
| POST   | /v1/galaxies/           |
| GET    | /v1/galaxies/           |
| GET    | /v1/galaxies/:id        |
| GET    | /v1/galaxies/name/:name |
| PUT    | /v1/galaxies/:id        |
| DELETE | /v1/galaxies/:id        |

Resources:

| Method | Endpoint                 |
|--------|--------------------------|
| POST   | /v1/resources/           |
| GET    | /v1/resources/           |
| GET    | /v1/resources/:id        |
| GET    | /v1/resources/name/:name |
| PUT    | /v1/resources/:id        |
| DELETE | /v1/resources/:id        |
