# Harvester
## A fast modern API for storing your resources

### Running Locally

Prerequisites:
  * Go 1.24+
  * Node.js 22+ (for UI development)
  * Docker & Docker Compose
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

### UI

The web interface is built with React 19, TypeScript, Tailwind CSS v4, TanStack Query/Router/Table, and shadcn/ui components.

**Pages:**

| Route | Description |
|-------|-------------|
| `/` | Dashboard with navigation cards |
| `/users` | User management table (CRUD, bulk ops) |
| `/galaxies` | Galaxy management table (CRUD, bulk ops) |
| `/resources` | Resource table with group/type filters |
| `/login` | Login page (placeholder) |

**Resource Filtering:**

The resources page supports server-side filtering with two dropdown filters:
- **Resource Group** - filters resources whose type belongs to the selected group (e.g., "Metal", "Organic")
- **Resource Type** - filters to a specific resource type; narrows to types within the selected group when a group is active

Both filters use AND logic when combined.

#### UI Development

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

### API Endpoints

When running locally host is available at `localhost:3000`

Postman collection included in root directory with examples.

All list endpoints support pagination via `page` and `row` query parameters, and sorting via `orderBy`.

#### Users

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

**Query params:** `user_id`, `name`, `email`, `start_created_date`, `end_created_date`
**Order fields:** `user_id`, `name`, `email`, `roles`, `guild`, `date_created`, `enabled`

#### Galaxies

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

**Query params:** `galaxy_id`, `name`, `date_created`
**Order fields:** `galaxy_id`, `name`, `date_created`

#### Resources

| Method | Endpoint                 | Description            |
|--------|--------------------------|------------------------|
| POST   | /v1/resources            | Create a resource      |
| POST   | /v1/resources/bulk       | Bulk create resources  |
| GET    | /v1/resources            | List resources         |
| GET    | /v1/resources/:id        | Get resource by ID     |
| GET    | /v1/resources/name/:name | Get resource by name   |
| PUT    | /v1/resources/:id        | Update resource        |
| PUT    | /v1/resources/bulk       | Bulk update resources  |
| DELETE | /v1/resources/:id        | Delete resource        |
| DELETE | /v1/resources/bulk       | Bulk delete resources  |

**Query params:** `resource_id`, `name`, `resource_type`, `resource_group`, `added_at`
**Order fields:** `resource_id`, `name`, `resource_type`, `verified`, `unavailable_at`, `added_at`, `cr`, `cd`, `dr`, `fl`, `hr`, `ma`, `pe`, `oq`, `sr`, `ut`, `er`

#### Resource Types

| Method | Endpoint                           | Description            |
|--------|------------------------------------|------------------------|
| GET    | /v1/resource-types                 | List resource types    |
| GET    | /v1/resource-types/:resource_type  | Get resource type      |
| POST   | /v1/resource-types                 | Create resource type   |
| POST   | /v1/resource-types/bulk            | Bulk create types      |
| PUT    | /v1/resource-types/:resource_type  | Update resource type   |
| DELETE | /v1/resource-types/:resource_type  | Delete resource type   |

**Query params:** `resourceType`, `resourceTypeName`, `resourceCategory`, `resourceGroup`, `enterable`, `containerType`

#### Resource Groups

| Method | Endpoint                              | Description           |
|--------|---------------------------------------|-----------------------|
| GET    | /v1/resource-groups                   | List resource groups  |
| GET    | /v1/resource-groups/:resource_group   | Get resource group    |

**Query params:** `resourceGroup`, `groupName`, `groupLevel`, `containerType`

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

### Configuration

Environment variables (prefix `HARVESTER_`):

| Variable | Default | Description |
|----------|---------|-------------|
| `HARVESTER_WEB_APIHOST` | `0.0.0.0:3000` | API listen address |
| `HARVESTER_WEB_DEBUGHOST` | `0.0.0.0:3010` | Debug listen address |
| `HARVESTER_WEB_READTIMEOUT` | `5s` | HTTP read timeout |
| `HARVESTER_WEB_WRITETIMEOUT` | `10s` | HTTP write timeout |
| `HARVESTER_WEB_IDLETIMEOUT` | `120s` | HTTP idle timeout |
| `HARVESTER_WEB_SHUTDOWNTIMEOUT` | `20s` | Graceful shutdown timeout |
| `HARVESTER_WEB_CORSALLOWEDORIGINS` | `*` | CORS allowed origins |
| `HARVESTER_DB_HOST` | `postgres` | Database host |
| `HARVESTER_DB_USER` | `postgres` | Database user |
| `HARVESTER_DB_PASSWORD` | `postgres` | Database password |
| `HARVESTER_DB_NAME` | `postgres` | Database name |
| `HARVESTER_DB_DISABLETLS` | `true` | Disable database TLS |
| `HARVESTER_DB_MAXIDLECONNS` | `0` | Max idle DB connections |
| `HARVESTER_DB_MAXOPENCONNS` | `0` | Max open DB connections |
| `HARVESTER_DB_RESET` | `false` | Drop all tables before migration |
| `HARVESTER_SEED_RESOURCES` | `false` | Seed random test resources |

### Project Structure

```
harvester/
├── api/                    # Go API service
│   ├── cmd/service/        # Main entry points
│   │   └── harvester/      # Harvester API server
│   └── domain/http/        # HTTP handlers by domain
│       ├── galaxyapi/
│       ├── resourceapi/
│       ├── resourcegroupapi/
│       ├── resourcetypeapi/
│       └── userapi/
├── app/                    # Application layer (models, filters)
│   └── domain/
│       ├── galaxyapp/
│       ├── resourceapp/
│       ├── resourcegroupapp/
│       ├── resourcetypeapp/
│       └── userapp/
├── business/               # Business logic layer (entities, stores)
│   ├── domain/
│   │   ├── galaxybus/
│   │   ├── resourcebus/
│   │   ├── resourcegroupbus/
│   │   ├── resourcetypebus/
│   │   └── userbus/
│   └── sdk/
│       ├── migrate/        # Database migrations & seeds
│       └── sqldb/          # Database utilities
├── foundation/             # Cross-cutting concerns
│   ├── logger/
│   ├── validate/
│   └── web/                # Web framework
├── infrastructure/
│   └── docker/             # Docker configs
│       ├── Dockerfile.harvester
│       ├── Dockerfile.admin
│       ├── Dockerfile.ui
│       ├── compose.yaml
│       └── nginx.conf
├── ui/                     # React frontend
│   └── src/
│       ├── app/            # App shell & routing
│       ├── components/     # Shared UI components (shadcn/ui)
│       ├── features/       # Feature modules
│       │   ├── galaxies/
│       │   ├── resources/
│       │   ├── resource-types/
│       │   └── users/
│       └── lib/            # API client & utilities
├── Harvester.postman_collection.json
└── Tiltfile
```
