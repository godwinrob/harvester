# -*- mode: Python -*-

# Harvester Tiltfile for local development using Docker Compose

# ============================================================================
# Docker Compose Mode - No Kubernetes required!
# Just run: tilt up
# ============================================================================

# Use docker compose for local development
docker_compose('infrastructure/docker/compose.yaml')

# ============================================================================
# Resource Configuration
# ============================================================================

# Configure PostgreSQL resource
dc_resource(
    'postgres',
    labels=['database'],
)

# Configure Admin resource (migrations/seeding)
dc_resource(
    'admin',
    resource_deps=['postgres'],
    labels=['tools'],
)

# Configure Harvester API resource
dc_resource(
    'harvester',
    resource_deps=['admin'],
    labels=['api'],
    links=[
        link('http://localhost:3000/v1/users', 'API: Users'),
        link('http://localhost:3000/v1/galaxies', 'API: Galaxies'),
        link('http://localhost:3000/v1/resources', 'API: Resources'),
    ],
)

# Configure UI resource
dc_resource(
    'ui',
    resource_deps=['harvester'],
    labels=['frontend'],
    links=[
        link('http://localhost:8080', 'Open UI'),
    ],
)

# ============================================================================
# Local Resources (optional convenience commands)
# ============================================================================

# Run tests
local_resource(
    'test',
    cmd='go test ./...',
    auto_init=False,
    trigger_mode=TRIGGER_MODE_MANUAL,
    labels=['tools'],
)

# Build check
local_resource(
    'build-check',
    cmd='go build ./...',
    auto_init=False,
    trigger_mode=TRIGGER_MODE_MANUAL,
    labels=['tools'],
)

# UI dev server (for faster development without Docker rebuild)
local_resource(
    'ui-dev',
    serve_cmd='cd ui && npm run dev',
    auto_init=False,
    trigger_mode=TRIGGER_MODE_MANUAL,
    labels=['frontend'],
    links=[
        link('http://localhost:5173', 'Open UI (Dev)'),
    ],
)
