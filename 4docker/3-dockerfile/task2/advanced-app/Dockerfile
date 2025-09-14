# Use official Python runtime as base image
FROM python:3.9-slim

# Set metadata
# TODO : Update with your email
LABEL maintainer="isjim2326@uis.no"
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

# Update package list and install curl
RUN apt-get update && \
    apt-get install -y curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# TODO : Set working directory (app)
WORKDIR /app

# TODO : Copy application files (app.py) and change ownership (appuser)
COPY app.py /app/
RUN chown -R appuser:appuser /app

# TODO : Switch to non-root user (appuser)
USER appuser

# TODO : Expose the port 8000
EXPOSE 8000

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/ || exit 1

# TODO : Run the application app.py
CMD ["python", "app.py"]