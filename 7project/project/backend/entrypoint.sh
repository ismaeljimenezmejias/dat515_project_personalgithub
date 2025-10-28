#!/bin/sh
set -e

# Simple entrypoint: run DB init once (idempotent) then start the app.
echo "Running DB init (idempotent)..."
if python init_db_pg.py; then
  echo "DB init completed."
else
  echo "DB init failed or DB not reachable; continuing to start the app."
fi

echo "Starting app..."
exec python app.py
