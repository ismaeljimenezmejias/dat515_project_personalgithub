"""
Unit test: migrations helper

This file contains tiny unit tests for migrations._is_postgres_connection.
Type: unit test (pure function behavior). No network or real DB required.
"""

import sys
from pathlib import Path

# ensure backend package is importable
sys.path.append(str(Path(__file__).resolve().parents[1]))

# provide light stubs for mysql and psycopg2 so importing db/migrations works in tests
import types, sys
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

from database.migrations import _is_postgres_connection


class DummyPG:
    # emulate a psycopg2 connection class module name
    __class__ = type('cls', (), {'__module__': 'psycopg2.extensions'})


class DummyMySQL:
    __class__ = type('cls', (), {'__module__': 'mysql.connector'})


def test_is_postgres_true():
    conn = DummyPG()
    assert _is_postgres_connection(conn) is True


def test_is_postgres_false():
    conn = DummyMySQL()
    assert _is_postgres_connection(conn) is False
