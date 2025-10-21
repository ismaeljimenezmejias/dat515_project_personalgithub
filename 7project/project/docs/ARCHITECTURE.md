# Architecture alignment and migration plan

Last updated: 2025-10-21

## Current stack (this repo)

- Web gateway: Nginx (container `nginx`) listening on host port 8080, proxying everything to Flask (`app:5000`).
- Backend: Flask app (container `app`) serving HTML templates and JSON APIs.
- Database: MySQL 8 (`database`), initialized via `database/init.sql`.
- Cache: Redis 7 (`cache`).
- Orchestration: Docker Compose for local/dev.
- Frontend: Server-rendered Jinja templates in `backend/templates/` (there is a `frontend/` folder not yet wired to Nginx).

See:
- `docker-compose.yaml` — service wiring and networks
- `nginx.conf` — upstream to `app:5000` and cache-control headers
- `backend/` — Flask app, routes, templates

## Target design (from project design doc)

- Web gateway: Web Gateway / Ingress → Nginx serving static frontend and reverse proxy to API.
- Frontend: React app served as static files by Nginx (SPA).
- Backend: REST API (Flask/Node) only — JSON contract between UI and logic.
- Database: PostgreSQL with PVC/PV (Kubernetes persistent volumes) in production.
- Orchestration: Kubernetes (dev → Compose is fine; prod → K8s).

## Alignment: what matches vs differs

Matches
- Nginx as entry point → reverse proxy to backend: YES (already in place).
- Backend in Flask (REST endpoints): YES (implemented for bikes and auth basics).
- Stateless API + client-side UI concept: PARTIAL (we currently render server-side HTML; API exists and can be expanded).

Diffs
- Frontend technology: Design says React SPA; current is Jinja server-rendered.
- Database engine: Design says PostgreSQL; current is MySQL 8.
- Platform: Design targets K8s; current is Docker Compose (appropriate for dev).
- API surface: Some listed endpoints (transactions/rentals, single-bike by id) are not yet implemented.

Conclusion: The current structure makes sense and is a valid subset of the design. It’s production-realistic to run “Nginx → Flask → DB/Redis” today. Moving to an API-first + SPA frontend is a clean, incremental evolution.

## Option C (recommended path): API-first backend + separate frontend served by Nginx

Goal: Keep the Flask API as the sole backend surface and serve a decoupled frontend (initially static, later React) from Nginx — matching the design without breaking today’s app.

Phased steps
1) Nginx routing split by path
   - `/api/` → proxy to Flask (`app:5000`)
   - `/` → serve static files from `./frontend` (mounted into the Nginx container)
   - Optional SPA fallback: `try_files $uri /index.html;`

2) Wire the folder
   - Bind-mount `./frontend` to `/usr/share/nginx/html:ro` in the `nginx` service.
   - Start with a minimal static `index.html` that fetches from `/api/...` (no CORS required, same origin).

3) De-risk during transition
   - Keep current server-rendered pages accessible at a subpath (e.g., `/server/`) or keep the old behavior behind a feature flag while the SPA catches up.
   - Alternatively, start by serving the SPA at `/ui/` first, then flip `/` once complete.

4) Frontend evolution
   - Incrementally port pages (home, bikes listing, publish, profile) to the SPA.
   - Reuse existing REST endpoints; add gaps as needed (e.g., `GET /api/bikes/{id}`).

5) Production alignment
   - Swap MySQL → PostgreSQL when convenient (create migration script, or keep MySQL if it’s sufficient for the course scope).
   - Later, deploy to K8s with an Ingress, PersistentVolume for the DB, and CI/CD.

Minimal config sketch (not applied yet)
- `docker-compose.yaml`: add a read-only volume to `nginx`:
  - `./frontend:/usr/share/nginx/html:ro`
- `nginx.conf`: split locations
  - `location /api/ { proxy_pass http://webapp; ... }`
  - `location / { root /usr/share/nginx/html; try_files $uri /index.html; }`

## API coverage vs design list

Implemented (representative)
- `GET /api/bikes` with filters (type, search, price ranges, owner exclusions)
- `POST /api/bikes` (requires login)
- `PUT /api/bikes/{id}` and `DELETE /api/bikes/{id}` with ownership checks
- Basic auth: `/api/login`, `/api/logout`, `/api/me`

Gaps (to add later if needed)
- `GET /api/bikes/{id}` (single bike by id)
- `POST /api/transactions` (purchase records)
- `POST /api/rentals` (rental records)
- `GET /api/users/{id}` (public profile data via API; server template exists)

## Recommended next steps
- Choose Option C and wire Nginx to serve `./frontend` at `/`, keep `/api/` for Flask.
- Start with a minimal static frontend (or scaffold React) and port pages incrementally.
- Defer DB engine switch unless there’s a concrete need; Compose + MySQL is perfectly fine for the course.
- When stable, containerize the SPA build output and plan a light K8s manifest set.

---

If you want, I can apply the Nginx/Compose changes in a “transition-safe” way (serve SPA under `/ui/` first) and drop a minimal stub into `frontend/` so you can iterate without breaking current pages.
