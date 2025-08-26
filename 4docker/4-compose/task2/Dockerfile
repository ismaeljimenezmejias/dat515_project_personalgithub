FROM python:3.9-slim

# Install system dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    default-mysql-client \
    && rm -rf /var/lib/apt/lists/*

# TODO: Copy requirements.txt and install Python packages

# TODO: Copy application to /app

# Create non-root user and switch to it
RUN useradd -r -s /bin/false appuser
USER appuser

# TODO: Expose the port your application runs on

# TODO: Health check to ensure the application is running

CMD ["python", "app.py"]
