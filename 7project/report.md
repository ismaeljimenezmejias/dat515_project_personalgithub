# Project Report

> **Instructions**:
> This template provides the structure for your project report.
> Replace the placeholder text with your actual content.
> Remove instructions that are not relevant for your project, but leave the headings along with a (NA) label.

## Project Overview

**Project Name**: Bike Student Marketplace

**Group Members**:

- 283596, Ismael Jiménez Mejías, ismaeljimenezmejias

**Brief Description**:
Bike Student Marketplace is a small web application for buying, selling and searching second‑hand bicycles within the university community. It uses a Flask backend with server‑rendered templates (Jinja) for the UI and a simple JSON API for dynamic actions (login, create/edit/delete bikes, search and filters). When running locally the stack uses MySQL and Redis (via Docker Compose); in production on Railway the container connects to a managed PostgreSQL instance. The goal was a reliable app that runs the same codebase in both environments.



## Architecture Overview

### High-Level Architecture


The app is a classic server‑rendered web app with a JSON API. Locally (Docker Compose) Nginx reverse proxies to the Flask container, which renders HTML (Jinja templates) and exposes `/api/*` endpoints. The local database is MySQL and Redis is available for caching/health checks. In production the same Flask image is deployed directly to Railway without Nginx; it connects to Railway's managed PostgreSQL database (and optional Redis service).

```mermaid
graph LR
  Browser[User Browser]

  subgraph Local (Docker Compose)
    Browser -->|HTTP| Nginx[Nginx]
    Nginx -->|HTTP| FlaskLocal[Flask App]
    FlaskLocal -->|SQL| MySQL[(MySQL)]
    FlaskLocal -->|Cache| Redis[(Redis)]
  end

  subgraph Production (Railway)
    Browser -->|HTTPS| FlaskProd[Flask App]
    FlaskProd -->|SQL| Postgres[(Railway PostgreSQL)]
    FlaskProd -->|Cache| RedisManaged[(Railway Redis)]
  end
```


### Components

- Frontend/UI (Flask + Jinja templates): index, about, publish, view bikes, bike detail, login, profile; search and filters rendered server-side.
- Backend (Flask blueprints in `backend/routes`): JSON endpoints for bikes (`/api/bikes`, `/api/bikes/<id>`), auth (`/api/login`, `/api/logout`, `/api/me`), and health (`/health`).
- Database: MySQL locally via Docker Compose; managed PostgreSQL on Railway in production.
- Cache: Redis locally via Docker Compose; optional Redis add-on on Railway.
- Web server: Nginx reverse proxy only in the local Docker composition; production routes directly to Flask.
- Containerization / Local infra: Dockerfile for backend; Docker Compose to run app + MySQL + Redis + Nginx locally.
- Deployment: Flask Docker image deployed to Railway through GitHub; environment variables configured in Railway project (REDIS_URL & DB_URL).


### Technologies Used

- **Backend**: Python 3.9, Flask, Flask Blueprints, Jinja2, redis-py, mysql-connector-python (in local)
- **Database**: MySQL 8 (Dockerized local dev), Railway managed PostgreSQL (production)
- **Cache**: Redis 7 (Docker local), Railway Redis service (production optional)
- **Cloud Services**: Railway (app container hosting, PostgreSQ, Redis), GitHub integration for deployments
- **Container Orchestration**: Docker, Docker Compose
- **Other**: Nginx reverse proxy (local stack)

## Prerequisites

### System Requirements

- **Local development (Docker Compose)**
  - Operating system: Windows 10/11, macOS, or Linux with Docker Desktop / Docker Engine
  - CPU: Hardware virtualization enabled (Docker requirement)
  - Memory: 4 GB minimum, 8 GB recommended so MySQL and Redis containers do not swap during seed loads and when multiple services run simultaneously
  - Storage: ~3 GB free for container images, database volume, and logs
- **Production usage (Railway deployment)**
  - Modern web browser and internet access; all server resources run on Railway’s infrastructure

### Required Software

- Docker Desktop 4.0+ (includes Docker Engine 20.10+ and Docker Compose v2)
- Ensure Docker Desktop is running before using `docker compose`; the CLI alone cannot manage containers without the daemon (could be Docker Desktop in Windows, or a remote daemon running in VM).
- Git and GitHub
- Railway account (for deploying the container and provisioning PostgreSQL/Redis services)

### Dependencies

**Runtime containers**

- Backend base image: `python:3.9-slim`
- Database image: `mysql:8.0` (local development)
- Cache image: `redis:7-alpine` (local development)
- Reverse proxy image: `nginx:alpine` (local development)
- Railway managed PostgreSQL (production)

- TODO: Optimizar la imagen del backend para tamaño/eficiencia (multi-stage, eliminar deps de build, considerar `python:3.10-slim`/`3.11-slim`).
  
  

**Python packages (installed via `backend/requirements.txt`)**

- `flask==2.3.3`
- `mysql-connector-python==8.1.0`
- `redis==4.6.0` (optional in production unless a Redis service is configured)
- `gunicorn` (WSGI server para producción)
- `psycopg2-binary` (driver para PostgreSQL en producción)
- `python-dotenv` (carga de variables de entorno en desarrollo)
- `Flask-Session==0.5.0` (gestión de sesiones en servidor)

## Build Instructions

### 1. Clone the Repository


```bash
git clone https://github.com/dat515-2025/ismaeljimenezmejias-labs

cd .\ismaeljimenezmejias-labs\7project\project\

```

### 2. Install Dependencies

```bash
No manual action needed when building with Docker; the backend Dockerfile copies `backend/requirements.txt` and runs `pip install --no-cache-dir -r requirements.txt` inside the image.
```

### 3. Build the Application

```bash
# Build all the images 
docker compose up --build -d

# Inspect container status (alternatively use Docker Desktop)
docker compose ps

# Build the backend image once when targeting Kubernetes
docker build -t bike-market-backend -f backend/Dockerfile .
```

### 4. Configuration

Local Docker Compose already sets the application credentials (`DB_HOST=database`, `DB_USER=webuser`, `DB_PASSWORD=webpass`, `DB_NAME=webapp`, `REDIS_HOST=cache`). Optional knobs are `SECRET_KEY` (Flask session signing) and `INIT_DB_ON_START` (forces a destructive PostgreSQL init when true); if you do nothing, sensible defaults apply.

All of the following scripts are invoked automatically by Docker or the app entrypoint; there are no extra configuration files to create.

- `backend/db.py`: central database connector; it detects whether `DB_URL` is present. Without it the app connects to the MySQL container using the vars above, otherwise it opens the Railway PostgreSQL URL.
- `backend/redis_client.py`: wraps the Redis client. Locally it points to the `cache` service; in production it reads `REDIS_URL` if you provision Redis on Railway. The module returns `None` when Redis is unavailable so the app still runs.
- `backend/migrations.py`: executes lightweight, idempotent schema adjustments (adds `password_hash`, optional bike location columns). It runs on startup and safely no-ops if the changes already exist.
- `database/init.sql`: bootstrap script mounted into the MySQL container. It creates the tables, seeds three demo users (Alice, Bob, Charlie), sample bikes, and a messages table.
- `backend/init_db_pg.py`: helper used only when deploying to Railway. When `INIT_DB_ON_START=true` it runs the equivalent schema creation against PostgreSQL to ensure prod has the same tables.


## Deployment Instructions

### Local Deployment (docker compose)

```bash
# Start the stack using existing images
docker compose up -d

# Follow backend logs until it reports ready
docker compose logs -f app
```

### Local Deployment (kubernetes)

```bash
# Make sure the backend image exists locally
docker build -t bike-market-backend -f backend/Dockerfile .

# Apply every manifest in k8s/ (Deployments, Services, Job, ConfigMap, PVC)
kubectl apply -f k8s/

# Confirm the cluster is reachable and pods are up
kubectl get nodes

#Check every pod runs correctly
kubectl get pods
```



### Cloud Deployment

```text
Railway workflow (production deploy):

1. Create a Railway project and add the “Deploy from GitHub” service pointing to `ismaeljimenezmejias/dat515_project_personalgithub` (main branch). Railway will auto-build on every push.

2. In the same project provision the managed PostgreSQL database and, if needed, the Redis service. Copy their public URLs.

3. In the backend service → Variables tab set `DB_URL` to the PostgreSQL URL, `REDIS_URL` to the Redis URL (optional), and `SECRET_KEY` to a strong value. Leave `INIT_DB_ON_START` unset unless you need to re-run the schema bootstrap.

4. Commit and `git push` to main. Railway builds the Docker image automatically; monitor the Deployments tab until it turns green. The app is then available at the Railway-provided URL.
```

### Verification

```bash
# Local:
docker compose ps                      # containers should be Up/healthy
curl http://localhost:8080/health       # returns {"status":"ok"}

# Railway:
- Use the Railway Logs tab or open https://railway.com/project/19d78c66-c328-4e09-a486-cbc548433250?environmentId=c0aef4e5-d796-4860-a454-e8779f5683f0/health
```

## Testing Instructions

### Unit Tests

```bash
# Commands to run unit tests
# For example:
# go test ./...
# npm test
```

### Integration Tests

```bash
# Commands to run integration tests
# Any setup required for integration tests
```

### End-to-End Tests

```bash
# Commands to run e2e tests
# How to set up test environment
```

## Usage Examples


### Basic Functionality

- GET /api/bikes
  - Returns JSON list of bikes. Supports query params: `search`, `type=venta|alquiler`, `minSalePrice`, `maxSalePrice`, `minRentalPrice`, `maxRentalPrice`.
- GET /api/bikes/<id>
  - Returns a single bike's details.
- POST /api/login
  - Body (JSON): `{ "name": "alice", "password": "..." }`. Sets server session cookie on success.
- POST /api/logout
  - Clears session.
- GET /api/me
  - Returns auth status and current user (useful to confirm session/cookies).
- POST /api/bikes
  - Create a new bike (requires authentication). The UI uses a form; the JSON API accepts fields such as `title`, `sale_price`, `sale_type`, `description`, `image_url`, `latitude`, `longitude`, etc.
- PUT /api/bikes/<id>
  - Update a bike (authenticated owner only).
- DELETE /api/bikes/<id>
  - Delete a bike (authenticated owner only).
- POST /api/messages
  - Send a message to a bike's owner. Body (JSON): `{ "bike_id": <id>, "receiver_id": <user_id>, "content": "..." }`.
- GET /api/messages/bike/<bike_id>
  - Returns message thread for that bike (authenticated and participant required).
- GET /api/messages/conversations
  - Returns user's message conversations summary.

### Advanced Features

- Authentication
  - POST /api/login
    - Body (JSON): `{ "name": "alice", "password": "..." }`. Sets server session cookie on success.
  - POST /api/logout
    - Clears session.
  - GET /api/me
    - Returns auth status and current user (useful to confirm session/cookies).

- Messaging
  - POST /api/messages
    - Send a message to a bike's owner. Body (JSON): `{ "bike_id": <id>, "receiver_id": <user_id>, "content": "..." }`.
  - GET /api/messages/bike/<bike_id>
    - Returns message thread for that bike (authenticated and participant required).
  - GET /api/messages/conversations
    - Returns user's message conversations summary.

---

## Presentation Video

**YouTube Link**: [Insert your YouTube link here]

**Duration**: [X minutes Y seconds]

**Video Includes**:

- [ ] Project overview and architecture
- [ ] Live demonstration of key features
- [ ] Code walkthrough
- [ ] Build and deployment showcase

### Troubleshooting

Below are the main troubleshooting patterns we hit and the concise mitigations we applied:

- OpenStack deploys blocked progress — mitigation: switch to local Docker Desktop for fast iteration and use Railway for an early cloud preview.
- DB incompatibilities (MySQL ↔ PostgreSQL) — mitigation: add small, idempotent migration checks and test the schema with the PostgreSQL driver used in production.
- Redis/session differences across environments — mitigation: make Redis optional, add health checks and validate session behaviour end‑to‑end in prod.
- Messaging persistence (brief) — during manual testing some messages were not saved and later triggered errors; we corrected the database implementation and validated the message flow.

### Quick debug commands

Main debugging functions I used were:

- Follow application logs
docker compose logs -f app

- Check the app health endpoint
curl http://localhost:8080/health

- Open a shell on the database container (inspect DB state interactively)
docker compose exec database sh

- Ping the local Redis service
docker compose exec cache redis-cli ping

- Re-run the app's lightweight schema ensure step
docker compose exec app python -c "from migrations import ensure_schema; ensure_schema()"

---

## Self-Assessment Table

> Be honest and detailed in your assessments.
> This information is used for individual grading.
> Link to the specific commit on GitHub for each contribution.

| Task/Component                                                      | Assigned To | Status        | Time Spent   | Difficulty | Notes                                         |
| ------------------------------------------------------------------- | ----------- | ------------- | ------------ | ---------- | --------------------------------------------- |
| Project Setup & Repository                                          | Ismael      | ✅ Complete    | 2.5 hours    | Medium     | Repo layout, initial Docker skeleton          |
| [Design Document](https://github.com/dat515-2025/group-name)         | Ismael      | ✅ Complete    | 4.0 hours    | Easy       | First meeting / design                        |
| [Backend API Development](https://github.com/dat515-2025/group-name) | Ismael      | ✅ Complete    | 12.0 hours   | Hard       | Phase 1 + Phase 2 backend work                |
| [Database Setup & Models](https://github.com/dat515-2025/group-name) | Ismael      | ✅ Complete    | 7.0 hours    | Medium     | Persistent DB, migrations, seed data         |
| [Frontend Development](https://github.com/dat515-2025/group-name)    | Ismael      | ✅ Complete    | 5.0 hours    | Medium     | UI completed: listing, detail, publish forms  |
| [Docker Configuration](https://github.com/dat515-2025/group-name)    | Ismael      | ✅ Complete    | 2.5 hours    | Easy       | Dockerfile, docker-compose                    |
| [Cloud Deployment](https://github.com/dat515-2025/group-name)        | Ismael      | ✅ Complete    | 10.0 hours   | Hard       | Railway deployment, env config                |
| [Testing Implementation](https://github.com/dat515-2025/group-name)  | Ismael      | 🔄 In Progress | TBD          | Medium     | Writing unit & integration tests locally      |
| [Local Kubernetes (minikube/k3d)]                                   | Ismael      | 🔄 In Progress | TBD          | Medium     | Setup manifests, local cluster, deploy tests  |
| [Documentation](https://github.com/dat515-2025/group-name)           | Ismael      | ✅ Complete    | 3.5 hours    | Easy       | README, report updates                        |
| [Presentation Video](https://github.com/dat515-2025/group-name)      | Ismael      | ✅ Complete    | 2.0 hours    | Medium     | Demo recording / outline                      |

**Legend**: ✅ Complete | 🔄 In Progress | ⏳ Pending | ❌ Not Started

## Hour Sheet

> Link to the specific commit on GitHub for each contribution.

### Ismael Jiménez Mejías

| Date      | Activity                         | Hours | Description                                                                 |
| --------- | -------------------------------- | -----:| --------------------------------------------------------------------------- |
| 2025-10-15| Design / first meeting           | 4.0   | Design and project understanding                                            |
| 2025-10-17| Initial setup                    | 2.5   | Repo layout, directory structure, Docker skeleton                           |
| 2025-10-22| Backend / functionality (phase 1)| 7.0   | Persistent DB, search & filters, Docker + persistent volumes                |
| 2025-10-29| Backend / functionality (phase 2)| 5.0   | Extended features (messaging, auth/authorization)                           |
| 2025-11-02| Use of the web app, deployment   | 10.0  | Manual QA, deployment to Railway, verification and iteration                |
| 2025-11-03| Documentation                    | 3.5   | README, report updates, design doc edits                                    |
| 2025-11-XX| Testing (TODO)                   | TBD   | Unit/integration tests not implemented yet                                  |
| **Total** |                                  | **32.0** |                                                                             |


---

## Final Reflection

### What We Learned


I learned how to build and deploy a full‑stack Flask app end‑to‑end and how the pieces fit together: container images, a persistent database, session state (Redis), server‑rendered pages and a small JSON API. I improved at debugging across the stack (app logs, Docker builds, env vars) and at shipping features in small, traceable commits.

- Iterate locally with Docker Desktop for fast feedback.
- Deploy early to catch infra-specific issues.
- Pin images/dependencies and keep the Dockerfile lean for reproducibility.


### Challenges Faced


- Stuck on OpenStack early → switched to Docker Desktop and continued locally to unblock progress.  
- Local MySQL vs Railway PostgreSQL incompatibilities → added small migration scripts, driver checks and tested against Railway Postgres.  
- Redis/session inconsistencies between environments → made Redis usage defensive (health checks / optional) and validated session behavior in prod.  
- Image build and env var failures on Railway → iterated Dockerfile, fixed env config, and used deployment logs to target fixes.



### If We Did This Again

Next time I would:
- Add automated tests early.
- Record short troubleshooting notes for each deployment fix.
- Optimize the Dockerfile (multi-stage) from the start and consider pinning base images when reproducibility is required.
- I now have a clearer understanding of the overall workflow and goals, which helps me plan and execute tasks more accurately from the start.


### Individual Growth

#### Ismael Jiménez Mejías

Working on this project greatly improved my end-to-end understanding of web app development and deployment. 

Now I have a clearer mental model of the steps involved (environment setup, containerization, DB provisioning and troubleshooting) so future work will be faster and more focused. I also learned to start deployment early to surface infra issues sooner and to keep short troubleshooting notes for fixes.

I became more confident diagnosing problems across layers (Flask app, Docker, DB, cloud) and learned to change strategy when something blocks progress (switching from OpenStack to Docker Desktop).


---

**Report Completion Date**: 14/11/2025
**Last Updated**: 14/11/2025
