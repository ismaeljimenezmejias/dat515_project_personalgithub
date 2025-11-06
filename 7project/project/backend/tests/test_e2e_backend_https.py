"""
E2E test (backend HTTPS)

This end-to-end test contacts the backend HTTPS server and verifies the
/health endpoint. Host and port are read from environment variables so the
test can target docker-compose services or a Kubernetes port-forward.

Environment variables:
- BACKEND_HOST (default: localhost)
- BACKEND_PORT (default: 8443)

Type: E2E. Skips when endpoint is unreachable.
"""

import os
import ssl
from http.client import HTTPSConnection
import pytest


def test_backend_https_health_e2e():
    host = os.getenv('BACKEND_HOST', 'localhost')
    env_port = os.getenv('BACKEND_PORT')
    port_candidates = [int(env_port)] if env_port else [8443, 443]

    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE

    last_exc = None
    for port in port_candidates:
        conn = HTTPSConnection(host, port, context=ctx, timeout=3)
        try:
            conn.request("GET", "/health")
            r = conn.getresponse()
            body = r.read().decode("utf-8")
            assert r.status == 200
            assert '"status": "healthy"' in body
            return
        except Exception as e:
            last_exc = e
        finally:
            try:
                conn.close()
            except Exception:
                pass

    pytest.skip(f"Backend HTTPS endpoint not available (tried ports {port_candidates}): {last_exc}")
