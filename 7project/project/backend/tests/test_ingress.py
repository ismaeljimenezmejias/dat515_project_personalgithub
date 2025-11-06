"""
DEPRECATED/DUPLICATE: ingress E2E tests

This file duplicates the tests in `test_e2e_ingress.py`. It is kept for
backwards compatibility but marked as deprecated; prefer `test_e2e_ingress.py`.
"""

# Skip this module by default — prefer the canonical `test_e2e_ingress.py` file.
import pytest
pytest.skip("Deprecated duplicate of test_e2e_ingress.py — skip this file", allow_module_level=True)

import ssl
import time
from http.client import HTTPSConnection
import pytest


def _get(path):
    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE

    conn = HTTPSConnection("localhost", 443, context=ctx, timeout=5)
    try:
        conn.request("GET", path)
        r = conn.getresponse()
        body = r.read().decode("utf-8")
        headers = dict(r.getheaders())
        return r.status, body, headers
    except Exception:
        pytest.skip("Ingress HTTPS endpoint unavailable")
    finally:
        conn.close()


def test_ingress_health():
    status, body, headers = _get("/health")
    assert status == 200
    assert '"status"' in body
    assert headers.get("Strict-Transport-Security")


def test_ingress_index():
    status, body, _ = _get("/")
    assert status == 200
    assert "Campus Bikes Marketplace" in body


def test_ingress_latency():
    total = 0.0
    for i in range(3):
        start = time.perf_counter()
        status, _, _ = _get("/health")
        assert status == 200
        total += time.perf_counter() - start
    avg = total / 3
    assert avg < 2.0
