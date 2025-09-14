# Use specific version for reproducibility
FROM python:3.9-slim-bullseye

# Set metadata
# TODO : Update with your email and repository link
LABEL maintainer="isjim2326@uis.no"
LABEL org.opencontainers.image.source="https://github.com/ismaeljimenezmejias"

# Install system dependencies in one layer and clean up
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# TODO : Create group and non-root user (flask), Use RUN to execute the command
RUN groupadd -r flask && useradd -r -g flask flask

# TODO : Set working directory to /app
WORKDIR /app

# TODO : Copy requirements.txt first (for better layer caching)
COPY requirements.txt .

# TODO : Install Python dependencies copied in the previous step
RUN pip install --no-cache-dir -r requirements.txt

# TODO : Copy application code and change ownership to flask
COPY app.py .
RUN chown -R flask:flask /app

# TODO : Switch to non-root user (flask)
USER flask

# Set environment variables
ENV FLASK_ENV=production
ENV PYTHONUNBUFFERED=1

# TODO : Expose port 5000
EXPOSE 5000

# TODO : Add health check
# Check the application for health check endpoint
# consider using different timeouts and intervals
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:5000/health || exit 1


# TODO : Run application
CMD ["python3", "app.py"]