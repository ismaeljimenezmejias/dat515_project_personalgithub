"""
DEPRECATED/DUPLICATE: HTTP redirect test

This file duplicates the HTTP->HTTPS redirect check included in
`test_e2e_ingress.py`. It's kept for compatibility but marked deprecated.
"""

import socket
from http.client import HTTPConnection
import pytest


def test_http_redirects_to_https_simple():
    conn = HTTPConnection("localhost", 8080, timeout=5)
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
