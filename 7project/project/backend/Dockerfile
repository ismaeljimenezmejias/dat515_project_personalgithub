FROM python:3.9-slim

# Minimal system deps
RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app 
# TODO: Copy requirements.txt and install Python packages
COPY requirements.txt .   
RUN pip install --no-cache-dir -r requirements.txt

# Copy application to /app
COPY . . 

# Normalize line endings (if copied from Windows) and make entrypoint executable
RUN sed -i 's/\r$//' entrypoint.sh && \
  chmod +x ./entrypoint.sh

# Create non-root user and switch to it
RUN useradd -r -s /bin/false appuser
USER appuser

# Expose the port (Cloud Run defaults to 8080)
EXPOSE 8080 

# Optional local healthcheck (Cloud Run ignores Dockerfile healthchecks)
HEALTHCHECK --interval=10s --timeout=3s --retries=10 \
  CMD python -c "import urllib.request; urllib.request.urlopen('http://127.0.0.1:8080/health', timeout=2)" || exit 1

# Use entrypoint script which will initialize DB (idempotent) then start the app
ENTRYPOINT ["./entrypoint.sh"]
