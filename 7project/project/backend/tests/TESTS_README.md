How to run tests (docker-compose vs Kubernetes)

This test suite contains unit, integration and E2E tests. The E2E and performance
tests contact services over the network. To make them work in different
environments the E2E tests read host/port from environment variables.

Environment variables (defaults):
- INGRESS_HOST=localhost
- INGRESS_PORT=443
- BACKEND_HOST=localhost
- BACKEND_PORT=8443
- HTTP_HOST=localhost
- HTTP_PORT=8080

Run with docker-compose (example):
1. Start services with docker-compose exposing the ports above, e.g.:

   docker-compose up -d

2. Run tests pointing to the compose host/ports (if using non-default ports set env vars):

   # If you expose ports differently, export the env vars, otherwise defaults work
   $env:INGRESS_HOST = 'localhost'
   $env:INGRESS_PORT = '443'
   pytest -q

Run with Kubernetes (example using port-forward):
1. Port-forward the ingress or service to your localhost:

   kubectl port-forward svc/ingress-nginx-controller 8443:443

2. Run tests pointing to localhost (defaults will work) or change env vars if needed:

   pytest -q

Notes
- Tests that require external endpoints will be skipped automatically if the
  endpoints are unreachable (they use pytest.skip).
- For CI, set the env vars to the service host/ports used in the pipeline.
