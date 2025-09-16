FROM python:3.9-slim

# Install system dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    default-mysql-client \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app 
# TODO: Copy requirements.txt and install Python packages
COPY requirements.txt .   
RUN pip install --no-cache-dir -r requirements.txt

# TODO: Copy application to /app
COPY app.py . 

# Create non-root user and switch to it
RUN useradd -r -s /bin/false appuser
USER appuser

# TODO: Expose the port your application runs on
EXPOSE 5000 

# TODO: Health check to ensure the application is running
HEALTHCHECK --interval=10s --timeout=3s --retries=10 \
  CMD python -c "import urllib.request; urllib.request.urlopen('http://127.0.0.1:5000/health', timeout=2)" || exit 1
  
CMD ["python", "app.py"]
