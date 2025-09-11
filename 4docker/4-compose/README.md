# Lab 4: Getting Started with Docker: Deploying a Basic Web App

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Learning Objectives](#learning-objectives)
- [Task 1 - Docker Compose Installation and Basics](#task-1---docker-compose-installation-and-basics)
- [Task 2 - Advanced Compose Features](#task-2---advanced-compose-features)
- [Task 3 - Environment Management](#task-3---environment-management)
- [Task 4 - Monitoring and Debugging](#task-4---monitoring-and-debugging)
- [Task 5 - Production Considerations](#task-5---production-considerations)
- [Task 6 - Cleanup and Review](#task-6---cleanup-and-review)
  - [6.1 System Cleanup](#61-system-cleanup)
- [Troubleshooting](#troubleshooting)
  - [Port Conflicts](#port-conflicts)
  - [Environment Switching Issues](#environment-switching-issues)
  - [Containers Not Starting](#containers-not-starting)
  - [Database Connection Issues](#database-connection-issues)
- [Verification Checklist](#verification-checklist)
- [Next Steps](#next-steps)
- [Quick Reference](#quick-reference)
  - [Essential Compose Commands](#essential-compose-commands)

## Learning Objectives

- **Duration:** ~75 minutes
- **Prerequisites:** Completed Part 3 (Dockerfile and image building)

By the end of this lab, you will be able to:

- Design and implement multi-container applications using Docker Compose
- Configure service dependencies and networking
- Manage persistent storage with volumes
- Implement environment-specific configurations

## Task 1 - Docker Compose Installation and Basics

1. **Check if Docker Compose Already Installed**

   Running with Docker Desktop on macOS or Windows, Docker Compose is included by default.

   ```console
   docker compose version
   ```

2. **Install Docker Compose on Ubuntu**

   ```console
   # For Ubuntu 24.04 (Compose V2)
   sudo apt update
   sudo apt install docker-compose-v2

   # Verify installation
   docker compose version
   ```

   Note: Modern Docker installations include Compose V2.
   Use `docker compose` (no hyphen) instead of `docker-compose`.

3. **Examine the Task 1 Files**

   ```console
   cd task1
   ls -la
   cat docker-compose.yml
   cat html/index.html
   ```

4. **Create Your First Stack**

   ```console
   mkdir simple-stack
   cd simple-stack
   cp -r ../html ../docker-compose.yml .
   ```

5. **Run Your First Stack**

   ```console
   # Start services in the background
   docker compose up -d

   # Check running services
   docker compose ps

   # View logs
   docker compose logs

   # Test the web service
   curl localhost:8080
   ```

6. **Explore the Stack**

   ```console
   # Execute commands in running containers
   docker compose exec web nginx -v
   docker compose exec database mysql --version

   # Check the database
   docker compose exec database mysql -u webuser -pwebpass webapp -e "SELECT 1 as test;"
   ```

7. **Stop the Stack**

   ```console
   # Stop and remove all containers
   docker compose down
   # Remove volumes (optional)
   docker compose down -v
   ```

## Task 2 - Advanced Compose Features

1. **Examine the Task 2 Files**

   ```console
   cd ../../task2
   ls -la
   cat app.py
   cat requirements.txt
   cat Dockerfile
   cat docker-compose.yml
   cat init.sql
   cat nginx.conf
   ```

2. **Create Advanced Application**

   Create the application structure and copy the necessary files:

   ```console
   mkdir webapp-stack
   cd webapp-stack

   # Copy all the provided files
   cp ../app.py ../requirements.txt ../Dockerfile ../docker-compose.yml ../init.sql ../nginx.conf .
   ```

3. **Complete the TODO Items for Task 2**

   Review the Dockerfile and docker-compose.yml and complete the TODO items in the `docker-compose.yml`.

4. **Run and Test the Advanced Stack**

   ```console
   # Build and start all services
   docker compose up -d --build

   # Watch the logs
   docker compose logs -f

   # In another terminal, check service status
   docker compose ps
   ```

   **Note:** The application is configured to use port 5001 (instead of the more common 5000) to avoid conflicts with macOS AirPlay Receiver, which uses port 5000 by default. This ensures the stack works out-of-the-box on all platforms.

5. **Test the Application**

   ```console
   # Test main endpoint
   curl localhost:8080/

   # Test health endpoint
   curl localhost:8080/health

   # Test counter (Redis)
   curl localhost:8080/counter
   curl localhost:8080/counter  # Should increment

   # Test database
   curl localhost:8080/data

   # Add data
   curl -X POST localhost:8080/data \
     -H "Content-Type: application/json" \
     -d '{"message": "Hello from Docker Compose!"}'

   # Check data again
   curl localhost:8080/data
   ```

   **Note:** All requests go through nginx on port 8080. You can also access the app directly on `localhost:5001`.

6. **Stop the Stack**

   ```console
   # Stop and remove all containers
   docker compose down

   # Remove volumes (optional)
   docker compose down -v
   ```

## Task 3 - Environment Management

1. **Examine the Task 3 Files**

   ```console
   cd ../../task3
   cat .env.dev
   cat .env.prod
   cat docker-compose.dev.yml
   cat docker-compose.prod.yml
   ```

2. **Create Environment-Specific Setup**

   Create the environment-specific structure and copy the necessary files:

   ```console
   mkdir environment-demo
   cd environment-demo

   # Copy the base compose file from task2 (required for overrides to work)
   cp ../../task2/docker-compose.yml .
   cp ../../task2/app.py ../../task2/requirements.txt ../../task2/Dockerfile ../../task2/init.sql ../../task2/nginx.conf .

   # Copy environment-specific files from task3
   cp ../../task3/.env.* .
   cp ../../task3/docker-compose.*.yml .
   ```

3. **Complete the TODO Items for Task 2**

   Review the Dockerfile and docker-compose.yml and complete the TODO items in the `docker-compose.yml`.

4. **Test Different Environments**

   ```console
   # Start with development configuration
   docker compose -f docker-compose.yml -f docker-compose.dev.yml --env-file .env.dev up -d --build

   # Test development environment
   curl localhost:5001/health  # Direct app access
   curl localhost:8080/health  # Through nginx

   # View development logs
   docker compose logs app

   # Stop and start production-like environment
   docker compose down  # IMPORTANT: Stop dev environment first!
   docker compose -f docker-compose.yml -f docker-compose.prod.yml --env-file .env.prod up -d --build

   # Check production configuration
   docker compose ps

   # Test production environment (no direct app access, only through nginx)
   curl localhost:80/health    # Production nginx runs on port 80

   # Test load balancing across replicas
   curl localhost:80/counter  # Should increment
   curl localhost:80/counter  # Should increment again (may be served by different replica)

   # Check which containers are running (should see 2 app containers)
   docker compose ps --filter "service=app"
   ```

   **Production Scaling:** The production configuration runs 2 replicas of the app service behind the nginx load balancer, demonstrating horizontal scaling. The base configuration has no direct app port access - the development override adds port 5001 for debugging, while production keeps apps internal. You can see multiple app containers with `docker compose ps`.

   **Port 80 Access:** If you get a "permission denied" error on port 80 (common on macOS/Linux), you can either:

   - Run with `sudo` (not recommended for development)
   - Change the production port to 8081 in `docker-compose.prod.yml` and test with `curl localhost:8081/health`

## Task 4 - Monitoring and Debugging

1. **Examine Task 4 Files - Monitoring and Debug Configuration**

   ```console
   # Navigate to task4 directory
   cd ../../task4

   # Examine all the provided files
   ls -la

   # Look at the monitoring compose file
   cat docker-compose.monitoring.yml

   # Examine the Prometheus configuration
   cat prometheus.yml

   # Look at the debug compose file
   cat docker-compose.debug.yml
   ```

2. **Create Monitoring Setup**

   Create the application structure and copy the necessary files:

   ```console
   mkdir monitoring
   cd monitoring

   # Copy the monitoring files
   cp ../docker-compose.debug.yml ../docker-compose.monitoring.yml ../prometheus.yml .
   ```

3. **Complete the TODO Items for Task 2**

   Review the Dockerfile and docker-compose.yml and complete the TODO items in the `docker-compose.yml`. Copy the necessary files from Task 2:

   ```console
   # Copy base configuration from Task 2
   cp ../task2/docker-compose.yml ./docker-compose.yml
   cp ../task2/app.py ../task2/requirements.txt ../task2/Dockerfile ../task2/init.sql .
   ```

4. **Start Monitoring Services**

   ```console
   # Start the monitoring stack
   docker compose -f docker-compose.monitoring.yml up -d

   # Check running services
   docker compose -f docker-compose.monitoring.yml ps

   # Access Prometheus (http://localhost:9090)
   # Access Grafana (http://localhost:3000 - admin/admin)
   ```

5. **Start Debug Environment**

   ```console
   # Start debug tools
   docker compose -f docker-compose.debug.yml up -d

   # Access Adminer for database management (http://localhost:8081)
   # Access Redis Commander (http://localhost:8082)
   ```

## Task 5 - Production Considerations

1. **Examine Task 5 Files - Production Configuration**

   ```console
   # Navigate to task5 directory
   cd ../../task5

   # Examine all the provided files
   ls -la

   # Look at the production compose file
   cat docker-compose.prod.yml

   # Examine the secrets directory
   ls -la secrets/
   cat secrets/db_root_password.txt
   cat secrets/db_password.txt
   ```

2. **Create Production Setup**

   Create the application structure and copy the necessary files:

   ```console
   # Create new directory for production
   mkdir production-stack
   cd production-stack

   # Copy the production files
   cp -r ../secrets ../docker-compose.prod.yml .

   # Set proper permissions on secrets
   chmod 600 secrets/*.txt
   ```

3. **Complete the TODO Items for Task 2**

   Review the Dockerfile and docker-compose.yml and complete the TODO items in the `docker-compose.yml`. Copy the necessary files from Task 2:

   ```console
   # Copy base configuration from Task 2
   cp ../task2/docker-compose.yml ./docker-compose.yml
   cp ../task2/app.py ../task2/requirements.txt ../task2/Dockerfile ../task2/init.sql .

   # Create ngnix config folder and copy nginx.conf
   mkdir -p nginx
   cd nginx
   cp ../../task2/nginx.conf .
   ```

4. **Deploy Production Stack**

   ```console
   # Set environment variables for production
   export VERSION=1.0
   export REDIS_PASSWORD=secure_redis_password

   # Start production stack
   docker compose -f docker-compose.prod.yml up -d

   # Check service status and resource usage
   docker compose -f docker-compose.prod.yml ps
   docker stats

   # Test the production deployment
   curl localhost/health
   ```

## Task 6 - Cleanup and Review

### 6.1 System Cleanup

1. **Clean Up All Stacks**

   ```console
   # Go back to webapp-stack and stop it
   cd ../../task2/webapp-stack
   docker compose down -v  # -v removes volumes too

   # Stop any other running stacks
   docker compose -f docker-compose.yml -f docker-compose.dev.yml down -v

   # Clean up system
   docker system prune -f
   docker volume prune -f
   ```

2. **Key Takeaways**

   - Use specific service versions for reproducibility
   - Implement proper health checks for all services
   - Use secrets for sensitive data in production
   - Configure resource limits and restart policies
   - Use networks to isolate services
   - Implement proper logging strategies
   - Use environment-specific override files
   - Always use volumes for persistent data
   - Avoid fixed container names in production to enable scaling
   - Use service discovery through network names, not container names

## Troubleshooting

### Port Conflicts

**Common Issue:** `bind: address already in use` errors

**Causes and Solutions:**

- **macOS AirPlay Receiver:** Uses port 5000 by default (this lab uses 5001 to avoid this)
- **Other services:** Check with `lsof -i :PORT` (macOS/Linux) or `netstat -an | findstr :PORT` (Windows)
- **Previous containers:** Clean up with `docker container prune -f`

**To change ports:** Edit the docker-compose.yml port mapping from `"HOST_PORT:CONTAINER_PORT"`

### Environment Switching Issues

**Problem:** `port is already allocated` when switching between dev and prod environments

**Cause:** Previous environment's containers are still running

**Solution:**

```console
# Always stop current environment before starting another
docker compose down  # Stop current stack
docker compose -f docker-compose.yml -f docker-compose.dev.yml down  # Stop dev specifically
docker compose -f docker-compose.yml -f docker-compose.prod.yml down  # Stop prod specifically

# Clean up any orphaned containers
docker container prune -f

# Then start the desired environment
docker compose -f docker-compose.yml -f docker-compose.prod.yml --env-file .env.prod up -d --build
```

### Containers Not Starting

**Problem:** Some containers fail to start or are in "Created" state

**Solution:**

```console
# Stop all compose stacks
docker compose down

# Clean up any orphaned containers
docker container prune -f

# Restart the stack
docker compose up -d --build
```

### Database Connection Issues

**Problem:** App can't connect to database

**Solution:**

```console
# Check if database is healthy
docker compose ps

# Check database logs
docker compose logs database

# Restart just the database
docker compose restart database
```

## Verification Checklist

Before proceeding to Part 5, ensure you can:

- Create multi-service applications with Docker Compose
- Configure service dependencies and health checks
- Manage environment variables and secrets
- Use volumes for persistent storage
- Implement different configurations for dev/prod
- Debug multi-container applications
- Apply production best practices

## Next Steps

Proceed to [Part 5: Final Project](../5-project/README.md) to apply all concepts in a comprehensive project.

## Quick Reference

### Essential Compose Commands

```console
# Lifecycle
docker compose up -d              # Start services in background
docker compose down               # Stop and remove containers
docker compose down -v            # Also remove volumes

# Monitoring
docker compose ps                 # List services
docker compose logs               # View logs
docker compose logs -f service    # Follow logs for specific service

# Operations
docker compose exec service bash  # Execute command in service
docker compose restart service    # Restart specific service
docker compose pull              # Pull latest images

# Multiple files
docker compose -f compose.yml -f override.yml up
```
