"""
LEGACY: Duplicate integration test (kept for compatibility)

This file is a legacy duplicate of `test_integration_db.py`. It performs the
same in-memory DB integration flow (signup -> create bike -> list bikes).
It is kept to avoid breaking older references but may be removed later.
"""

import os
import sqlite3
from http import cookies
import json


class _SqliteCompatCursor:
    def __init__(self, conn):
        self._cur = conn.cursor()

    def execute(self, sql, params=None):
        # replace %s with ? for sqlite
        if params is None:
            params = ()
        q = sql.replace('%s', '?')
        # remove RETURNING id (sqlite may not support) and emulate
        self._last_sql = q
        self._cur.execute(q, params)
        return self

    def fetchone(self):
        return self._cur.fetchone()

    def fetchall(self):
        return self._cur.fetchall()

    @property
    def lastrowid(self):
        return self._cur.lastrowid


class SqliteCompatConn:
    def __init__(self):
        self.conn = sqlite3.connect(':memory:')
        self.conn.row_factory = sqlite3.Row

    def cursor(self):
        return _SqliteCompatCursor(self.conn)

    def commit(self):
        return self.conn.commit()

    def close(self):
        # keep underlying in-memory DB open for the lifetime of the test
        # routes call conn.close(), but we want to reuse the same connection
        return


def _init_schema(conn):
    cur = conn.cursor()
    # Minimal schema for tests
    cur.execute('''CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE, password_hash TEXT, created_at TEXT)''')
    cur.execute('''CREATE TABLE bikes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        sale_price REAL,
        rental_price REAL,
        sale_type TEXT,
        model TEXT,
        description TEXT,
        bike_condition TEXT,
        image_url TEXT,
        owner_id INTEGER,
        created_at TEXT,
        location_name TEXT,
        latitude REAL,
        longitude REAL
    )''')
    conn.commit()


def _get_cookiejar_from_resp(resp):
    # resp is a werkzeug response
    raw = resp.headers.get('Set-Cookie')
    if not raw:
        return {}
    simple = cookies.SimpleCookie()
    simple.load(raw)
    return {k: v.value for k, v in simple.items()}


def test_signup_and_create_bike(monkeypatch):
    # Disable migrations on app import
    os.environ['MIGRATE_ON_START'] = 'false'

    # Patch db.get_connection to return sqlite compat
    import types, sys
    from pathlib import Path
    # ensure backend path is on sys.path so we can import local modules like db and app
    sys.path.append(str(Path(__file__).resolve().parents[1]))
    # inject lightweight stubs so importing db.py won't fail if mysql/psycopg2 are missing
    mysql_mod = types.ModuleType('mysql')
    mysql_connector = types.ModuleType('mysql.connector')
    def _fake_mysql_connect(**kwargs):
        raise RuntimeError('mysql connector stub - should not be used in test')
    mysql_connector.connect = _fake_mysql_connect
    mysql_mod.connector = mysql_connector
    sys.modules['mysql'] = mysql_mod
    sys.modules['mysql.connector'] = mysql_connector
    psycopg2_mod = types.ModuleType('psycopg2')
    def _fake_psycopg2_connect(*args, **kwargs):
        raise RuntimeError('psycopg2 stub - should not be used in test')
    psycopg2_mod.connect = _fake_psycopg2_connect
    sys.modules['psycopg2'] = psycopg2_mod
    # stub redis module if missing
    redis_mod = types.ModuleType('redis')
    class _FakeRedis:
        @staticmethod
        def from_url(url):
            return None
        def __init__(self, *a, **k):
            pass
    redis_mod.from_url = _FakeRedis.from_url
    redis_mod.Redis = _FakeRedis
    sys.modules['redis'] = redis_mod
    # stub flask_session so Session(app) won't error
    flask_sess = types.ModuleType('flask_session')
    class _FakeSession:
        def __init__(self, *a, **k):
            pass
        def __call__(self, app):
            return None
    flask_sess.Session = _FakeSession
    sys.modules['flask_session'] = flask_sess

    import db as dbmod

    # create a single shared in-memory DB for the test run
    _shared_conn = SqliteCompatConn()
    _init_schema(_shared_conn)

    def _fake_conn():
        return _shared_conn

    monkeypatch.setattr(dbmod, 'get_connection', _fake_conn)

    # import app after monkeypatch
    from app import app as flask_app

    client = flask_app.test_client()

    # signup
    resp = client.post('/api/signup', json={'name': 'testuser', 'password': 'pass'})
    assert resp.status_code == 201
    data = resp.get_json()
    assert data.get('ok')

    # create a bike (session is maintained by test_client)
    resp2 = client.post('/api/bikes', json={'title': 'My Bike', 'sale_price': 100.0, 'sale_type': 'venta'})
    # debug: if there's an error, print body for diagnosis
    if resp2.status_code != 201:
        print('CREATE BIKE failed status:', resp2.status_code)
        print(resp2.get_data().decode('utf-8', errors='replace'))
    assert resp2.status_code == 201
    d2 = resp2.get_json()
    assert d2.get('success') or d2.get('ok') or resp2.status_code == 201

    # fetch bikes
    resp3 = client.get('/api/bikes')
    assert resp3.status_code == 200
    j = resp3.get_json()
    assert 'bikes' in j
    assert j['count'] >= 1
