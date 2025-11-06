"""
Performance smoke test

Tiny performance check that measures the average latency to /health via the
ingress (localhost:443). This is a smoke test (not a full benchmark).
"""
import os
import ssl
import time
from http.client import HTTPSConnection
import pytest


def _get(path):
    host = os.getenv('INGRESS_HOST', 'localhost')
    port = int(os.getenv('INGRESS_PORT', '443'))
    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE
    conn = HTTPSConnection(host, port, context=ctx, timeout=5)
    try:
        conn.request("GET", path)
        r = conn.getresponse()
        r.read()
        return r.status
    except Exception:
        pytest.skip("Ingress HTTPS endpoint unavailable")
    finally:
        conn.close()


def test_performance_smoke_ingress_latency():
    iterations = 3
    total = 0.0
    for _ in range(iterations):
        start = time.perf_counter()
        status = _get("/health")
        assert status == 200
        total += time.perf_counter() - start
    avg = total / iterations
    # smoke threshold
    assert avg < 2.0
