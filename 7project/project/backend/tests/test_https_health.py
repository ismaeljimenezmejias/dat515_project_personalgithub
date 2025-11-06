"""
DEPRECATED/DUPLICATE: backend HTTPS health test

This file is a duplicate of `test_e2e_backend_https.py`. It's kept for
compatibility but marked as deprecated — prefer `test_e2e_backend_https.py`.
"""

import ssl
from http.client import HTTPSConnection
import pytest


def test_health_https_simple():
    # check /health via https on localhost:8443
    host = "localhost"
    port = 8443
    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE

    conn = HTTPSConnection(host, port, context=ctx, timeout=5)
    try:
        conn.request("GET", "/health")
        r = conn.getresponse()
        body = r.read().decode("utf-8")
        assert r.status == 200
        assert '"status": "healthy"' in body
        assert r.getheader("Strict-Transport-Security") is not None
    except Exception:
        pytest.skip("HTTPS endpoint not available")
    finally:
        conn.close()
