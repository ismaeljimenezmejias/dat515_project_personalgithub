# Use official Python runtime as base image
FROM python:3.9-slim

# Set metadata
# TODO : Update with your email
LABEL maintainer="your-email@example.com"
LABEL version="2.0"
LABEL description="Advanced Python web application"

# Set environment variables
ENV APP_ENV=production
ENV APP_VERSION=2.0
ENV PORT=8000
ENV PYTHONUNBUFFERED=1

# Create app directory and user for security
RUN groupadd -r appuser && useradd -r -g appuser appuser
RUN mkdir -p /app && chown appuser:appuser /app

# TODO : Set working directory (app)

# TODO : Copy application files (app.py) and change ownership (appuser)

# TODO : Switch to non-root user (appuser)

# TODO : Expose the port 8000

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/ || exit 1

# TODO : Run the application app.py
