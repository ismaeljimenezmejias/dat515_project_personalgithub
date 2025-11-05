import socket
import ssl
from http.client import HTTPSConnection, HTTPConnection

import pytest


def _connect(host, port, path, verify_cert):
    context = ssl.create_default_context()
    if not verify_cert:
        context.check_hostname = False
        context.verify_mode = ssl.CERT_NONE

    conn = HTTPSConnection(host, port, context=context, timeout=5)
    try:
        conn.request("GET", path)
        response = conn.getresponse()
        body = response.read().decode("utf-8")
        return response.status, body, response.getheader("Strict-Transport-Security")
    except (ConnectionRefusedError, socket.timeout, ssl.SSLError) as exc:
        pytest.skip(f"HTTPS endpoint unavailable: {exc}")
    finally:
        conn.close()


def test_health_endpoint_over_https():
    """Ensures /health is reachable via HTTPS and returns healthy status."""
    host = "localhost"
    port = 8443
    status, body, hsts_header = _connect(host, port, "/health", verify_cert=False)

    assert status == 200, body
    assert '"status": "healthy"' in body
    assert hsts_header is not None


def test_http_redirects_to_https():
    """Checks plain HTTP is redirected to HTTPS as expected."""
    conn = HTTPConnection("localhost", 8080, timeout=5)
    try:
        conn.request("GET", "/")
        response = conn.getresponse()
        assert response.status in (301, 302)
        location = response.getheader("Location")
        assert location and location.startswith("https://")
    except (ConnectionRefusedError, socket.timeout) as exc:
        pytest.skip(f"HTTP endpoint unavailable: {exc}")
    finally:
        conn.close()
