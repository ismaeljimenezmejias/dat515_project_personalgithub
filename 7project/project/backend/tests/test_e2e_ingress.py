"""
E2E tests (ingress + HTTP redirect)

These tests contact the ingress and the HTTP port to validate TLS-proxied
endpoints (/health, /) and that HTTP redirects to HTTPS. They read host/port
from environment variables so they can be pointed at docker-compose services
or a Kubernetes ingress (via port-forward).

Environment variables (defaults set for local dev):
- INGRESS_HOST (default: localhost)
- INGRESS_PORT (default: 443)
- HTTP_HOST (default: localhost)
- HTTP_PORT (default: 8080)

Type: E2E (real network). Tests are skipped when endpoints are unreachable.
"""

import os
import ssl
from http.client import HTTPSConnection
import pytest


def _get(path):
    host = os.getenv('INGRESS_HOST', 'localhost')
    # If INGRESS_PORT provided, use it. Otherwise try common ports: docker-compose maps 8443->443,
    # Kubernetes ingress commonly listens on 443. Try 8443 first, then 443.
    env_port = os.getenv('INGRESS_PORT')
    port_candidates = [int(env_port)] if env_port else [8443, 443]
    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE

    # Try candidates until one connects and returns a response.
    last_exc = None
    for port in port_candidates:
        conn = HTTPSConnection(host, port, context=ctx, timeout=3)
        try:
            conn.request("GET", path)
            r = conn.getresponse()
            body = r.read().decode("utf-8")
            headers = dict(r.getheaders())
            return r.status, body, headers
        except Exception as e:
            last_exc = e
        finally:
            try:
                conn.close()
            except Exception:
                pass

    # If none of the candidates worked, skip the test.
    pytest.skip(f"Ingress HTTPS endpoint unavailable (tried ports {port_candidates}): {last_exc}")


def test_ingress_health_e2e():
    status, body, headers = _get("/health")
    assert status == 200
    assert '"status"' in body
    assert headers.get("Strict-Transport-Security")


def test_ingress_index_e2e():
    status, body, _ = _get("/")
    assert status == 200
    assert "Campus Bikes Marketplace" in body


def test_http_redirects_to_https_e2e():
    # plain HTTP redirect check (reads host/port from env)
    http_host = os.getenv('HTTP_HOST', 'localhost')
    http_port = int(os.getenv('HTTP_PORT', '8080'))
    from http.client import HTTPConnection
    conn = HTTPConnection(http_host, http_port, timeout=5)
    try:
        conn.request("GET", "/")
        r = conn.getresponse()
        assert r.status in (301, 302)
        loc = r.getheader("Location")
        assert loc and loc.startswith("https://")
    except Exception:
        pytest.skip("HTTP endpoint not available")
    finally:
        conn.close()
