# Use specific version for reproducibility
FROM python:3.9-slim-bullseye

# Set metadata
# TODO : Update with your email and repository link
LABEL maintainer="your-email@example.com"
LABEL org.opencontainers.image.source="https://github.com/your-repo"

# Install system dependencies in one layer and clean up
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# TODO : Create group and non-root user (flask), Use RUN to execute the command

# TODO : Set working directory to /app

# TODO : Copy requirements.txt first (for better layer caching)

# TODO : Install Python dependencies copied in the previous step

# TODO : Copy application code and change ownership to flask

# TODO : Switch to non-root user (flask)

# Set environment variables
ENV FLASK_ENV=production
ENV PYTHONUNBUFFERED=1

# TODO : Expose port 5000

# TODO : Add health check
# Check the application for health check endpoint
# consider using different timeouts and intervals

# TODO : Run application
