#!/bin/sh
set -e

# Only run destructive DB init when explicitly enabled
if [ "${INIT_DB_ON_START}" = "true" ]; then
  echo "INIT_DB_ON_START=true -> Running destructive DB init (database.init_db_pg)..."
  if python -m database.init_db_pg; then
    echo "DB init completed."
  else
    echo "DB init failed or DB not reachable; continuing to start the app."
  fi
else
  echo "Skipping DB init (INIT_DB_ON_START not true)."
fi

PORT_TO_BIND="${PORT:-5000}"
echo "Starting app with gunicorn on port ${PORT_TO_BIND}..."
exec gunicorn -w 2 -k gthread -t 0 -b 0.0.0.0:${PORT_TO_BIND} app:app
