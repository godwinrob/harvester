# Harvester
## A fast modern API for storing your resources

### Running Locally

Prerequisites:
  * Go 1.22.4
  * Node.js 22+ (for UI development)
  * Docker
  * [Tilt](https://tilt.dev/) (optional, for enhanced development experience)

#### Option 1: Using Tilt (Recommended)

```bash
# Start everything with a single command
tilt up

# Open Tilt UI at http://localhost:10350
# Press 's' to open in browser

# Stop (Ctrl+C in terminal, then)
tilt down
```

Tilt provides:
- Live reload on code changes
- Dashboard at http://localhost:10350
- Port forwarding for API (3000, 3010), UI (8080), and PostgreSQL (5432)
- Manual triggers for tests and build checks
- UI dev server option for faster frontend iteration

#### Option 2: Using Docker Compose directly

```bash
# Start the API and UI
docker compose -f infrastructure/docker/compose.yaml up --build

# With watch mode for live reload
docker compose -f infrastructure/docker/compose.yaml watch

# Stop and destroy
docker compose -f infrastructure/docker/compose.yaml down
```

### Accessing the Application

| Service | URL | Description |
|---------|-----|-------------|
| **UI** | http://localhost:8080 | Web interface |
| **API** | http://localhost:3000 | REST API (direct access) |
| **Debug** | http://localhost:3010 | Debug endpoints |
| **Tilt Dashboard** | http://localhost:10350 | Development dashboard |

### UI Development

For faster frontend development without Docker rebuilds:

```bash
# Option 1: Use Tilt's ui-dev resource
tilt up
# Then trigger 'ui-dev' in the Tilt dashboard

# Option 2: Run directly
cd ui
npm install
npm run dev
# UI available at http://localhost:5173 with hot reload
```

The UI is built with:
- React 19 + TypeScript
- Tailwind CSS v4
- TanStack Query & Router
- shadcn/ui components

### API Endpoints

When running locally host is available at `localhost:3000`

Postman collection included in root directory with examples

Users:

| Method | Endpoint           | Description           |
|--------|--------------------|-----------------------|
| POST   | /v1/users          | Create a user         |
| POST   | /v1/users/bulk     | Bulk create users     |
| GET    | /v1/users          | List users            |
| GET    | /v1/users/:id      | Get user by ID        |
| PUT    | /v1/users/:id      | Update user           |
| PUT    | /v1/users/role/:id | Update user role      |
| PUT    | /v1/users/bulk     | Bulk update users     |
| DELETE | /v1/users/:id      | Delete user           |
| DELETE | /v1/users/bulk     | Bulk delete users     |

Galaxies:

| Method | Endpoint                | Description           |
|--------|-------------------------|-----------------------|
| POST   | /v1/galaxies            | Create a galaxy       |
| POST   | /v1/galaxies/bulk       | Bulk create galaxies  |
| GET    | /v1/galaxies            | List galaxies         |
| GET    | /v1/galaxies/:id        | Get galaxy by ID      |
| GET    | /v1/galaxies/name/:name | Get galaxy by name    |
| PUT    | /v1/galaxies/:id        | Update galaxy         |
| PUT    | /v1/galaxies/bulk       | Bulk update galaxies  |
| DELETE | /v1/galaxies/:id        | Delete galaxy         |
| DELETE | /v1/galaxies/bulk       | Bulk delete galaxies  |

Resources:

| Method | Endpoint                 | Description           |
|--------|--------------------------|----------------------|
| POST   | /v1/resources            | Create a resource    |
| POST   | /v1/resources/bulk       | Bulk create resources|
| GET    | /v1/resources            | List resources       |
| GET    | /v1/resources/:id        | Get resource by ID   |
| GET    | /v1/resources/name/:name | Get resource by name |
| PUT    | /v1/resources/:id        | Update resource      |
| PUT    | /v1/resources/bulk       | Bulk update resources|
| DELETE | /v1/resources/:id        | Delete resource      |
| DELETE | /v1/resources/bulk       | Bulk delete resources|

### Bulk Operations

All bulk operations support a maximum of **100 items** per request.

#### Bulk Create Request Example
```json
POST /v1/users/bulk
{
  "items": [
    { "name": "John", "email": "john@example.com", "roles": ["USER"], "password": "secret", "passwordConfirm": "secret" },
    { "name": "Jane", "email": "jane@example.com", "roles": ["ADMIN"], "password": "secret", "passwordConfirm": "secret" }
  ]
}
```

#### Bulk Create Response (201)
```json
{
  "items": [
    { "id": "uuid-1", "name": "John", "email": "john@example.com", ... },
    { "id": "uuid-2", "name": "Jane", "email": "jane@example.com", ... }
  ],
  "created": 2
}
```

#### Bulk Update Request Example
```json
PUT /v1/users/bulk
{
  "items": [
    { "id": "uuid-1", "data": { "name": "John Smith" } },
    { "id": "uuid-2", "data": { "enabled": false } }
  ]
}
```

#### Bulk Delete Request Example
```json
DELETE /v1/users/bulk
{
  "ids": ["uuid-1", "uuid-2"]
}
```

#### Validation Error Response (400)
```json
{
  "code": "failed_precondition",
  "message": "validation failed",
  "errors": [
    { "index": 0, "field": "email", "error": "email is required" },
    { "index": 2, "field": "password", "error": "password must be at least 8 characters" }
  ]
}
```

### Project Structure

```
harvester/
├── api/                    # Go API service
│   ├── cmd/service/        # Main entry points
│   └── domain/http/        # HTTP handlers
├── app/                    # Application layer
├── business/               # Business logic layer
├── foundation/             # Infrastructure utilities
├── infrastructure/
│   └── docker/             # Docker configs
│       ├── Dockerfile.harvester
│       ├── Dockerfile.admin
│       ├── Dockerfile.ui
│       ├── compose.yaml
│       └── nginx.conf
├── ui/                     # React frontend
│   ├── src/
│   │   ├── app/            # App shell & routing
│   │   ├── components/     # UI components
│   │   ├── features/       # Feature modules
│   │   └── lib/            # Utilities & API client
│   └── package.json
└── Tiltfile                # Tilt configuration
```
