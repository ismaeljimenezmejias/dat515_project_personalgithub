# Build stage: not needed in production but needed for installing dependencies
FROM python:3.9-slim AS builder

# Install build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY backend/requirements.txt .

# Install dependencies in builder
RUN pip install --no-cache-dir -r requirements.txt

# Runtime stage
FROM python:3.9-slim

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy installed packages from builder
COPY --from=builder /usr/local/lib/python3.9/site-packages /usr/local/lib/python3.9/site-packages

COPY --from=builder /usr/local/bin /usr/local/bin

# Copy application code
COPY backend/ .

# Fix line endings and make executable
RUN sed -i 's/\r$//' entrypoint.sh && \
    chmod +x ./entrypoint.sh

# Create non-root user
RUN useradd -r -s /bin/false appuser

# Change ownership
RUN chown -R appuser:appuser /app

USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=10s --timeout=3s --retries=10 \
    CMD python -c "import urllib.request; urllib.request.urlopen('http://127.0.0.1:5000/health', timeout=2)" || exit 1

# Run the app
ENTRYPOINT ["./entrypoint.sh"]
