"""Database package: connectors, schema init and migrations.

Contains:
  - init_db_pg: destructive PostgreSQL bootstrap (dev/one-off)
  - migrations: idempotent additive changes for both MySQL/PostgreSQL
  - init.sql: legacy MySQL first-run schema (docker compose)
"""
