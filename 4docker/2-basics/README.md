# Lab 4: Getting Started with Docker: Deploying a Basic Web App

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Learning Objectives](#learning-objectives)
- [Task 1 - Container Lifecycle Management](#task-1---container-lifecycle-management)
- [Task 2 - Image Management](#task-2---image-management)
- [Task 3 - Networking and Port Mapping](#task-3---networking-and-port-mapping)
- [Task 4 - Environment Variables and Configuration](#task-4---environment-variables-and-configuration)
- [Task 5 - Volume Basics](#task-5---volume-basics)
- [Task 6 - Debugging and Troubleshooting](#task-6---debugging-and-troubleshooting)
- [Task 7 - Resource Management](#task-7---resource-management)
- [Task 8 - Cleanup and Best Practices](#task-8---cleanup-and-best-practices)
- [Verification Checklist](#verification-checklist)
- [Next Steps](#next-steps)
- [Quick Reference](#quick-reference)
  - [Essential Commands](#essential-commands)

## Learning Objectives

- **Duration:** ~45 minutes
- **Prerequisites:** Completed Part 1 (Docker installed and working)

By the end of this lab part, you will be able to:

- Master essential Docker commands and container lifecycle
- Understand the difference between images and containers
- Work with Docker networking and port mapping
- Debug and troubleshoot common container issues

## Task 1 - Container Lifecycle Management

1. **Working with Pre-built Images - Pull and Run Your First Web Server**

   ```console
   # Pull the Apache HTTP server image
   docker pull httpd

   # Run the web server
   docker run -d -p 80:80 --name my-web-server httpd
   ```

2. **Verify the Container is Running**

   ```console
   # List running containers
   docker ps

   # Check the web server works
   curl localhost
   curl http://localhost:80
   ```

   You should see: `<html><body><h1>It works!</h1></body></html>`

3. **Access from Outside the VM**

   - Open your browser and navigate to `http://<your-floating-ip>`
   - You should see the "It works!" page

4. **Container Management Commands**

   ```console
   # List all containers (running and stopped)
   docker ps -a

   # View container logs
   docker logs my-web-server

   # Execute commands inside a running container
   docker exec -it my-web-server /bin/bash
   # (Type 'exit' to leave the container)

   # Stop the container
   docker stop my-web-server

   # Start the container again
   docker start my-web-server

   # Remove the container (must be stopped first)
   docker stop my-web-server
   docker rm my-web-server
   ```

5. **Exercise: Container Exploration**

   Run an interactive Ubuntu container and explore its file system:

   ```console
   # Run Ubuntu interactively
   docker run -it --name ubuntu-explorer ubuntu:latest /bin/bash

   # Inside the container, try these commands:
   ls -la
   cat /etc/os-release
   ps aux
   whoami

   # Exit the container
   exit

   # Clean up
   docker rm ubuntu-explorer
   ```

## Task 2 - Image Management

1. **Working with Image Commands**

   ```console
   # List local images
   docker images

   # Search for images on Docker Hub
   docker search nginx

   # Pull specific version (1.29 is the latest version as of this writing)
   docker pull nginx:1.29
   # You may also use the `latest` tag for testing, but it's better to use specific versions in production.
   docker pull nginx:latest

   # Inspect an image
   docker image inspect nginx:1.29
   docker image inspect nginx:latest

   # Remove an image
   docker rmi nginx:1.29
   docker rmi nginx:latest
   ```

2. **Understanding Image Layers and Exploring Image History**

   ```console
   # See image layers
   docker history httpd

   # Check image size and usage
   docker system df
   ```

## Task 3 - Networking and Port Mapping

1. **Multiple Port Mappings**

   ```console
   # Run Nginx with custom port mapping
   docker run -d -p 8080:80 --name nginx-8080 nginx

   # Run another instance on different port
   docker run -d -p 8081:80 --name nginx-8081 nginx

   # Test both instances
   curl localhost:8080
   curl localhost:8081
   ```

2. **Network Inspection**

   ```console
   # Inspect container network settings
   docker inspect nginx-8080 | grep IPAddress

   # Check port mappings
   docker port nginx-8080

   # List Docker networks
   docker network ls
   ```

3. **Exercise: Multi-Container Setup - Run Multiple Services**

   ```console
   # Run a simple Python web server
   docker run -d -p 3000:8000 --name python-server python:3.9 \
     python -m http.server 8000

   # Run a Node.js hello world (if available)
   docker run -d -p 3001:3000 --name node-server node:16 \
     sh -c "echo 'console.log(\"Hello World\");' > app.js && node app.js"

   # Check all running containers
   docker ps
   ```

## Task 4 - Environment Variables and Configuration

1. **Using Environment Variables - Run Container with Environment Variables**

   ```console
   # Run MySQL with environment configuration
   docker run -d \
     --name test-mysql \
     -e MYSQL_ROOT_PASSWORD=mypassword \
     -e MYSQL_DATABASE=testdb \
     -p 3306:3306 \
     mysql:8.0

   # Check if MySQL is ready
   docker logs test-mysql
   ```

2. **Connect to MySQL**

   ```console
   # Connect to MySQL inside the container
   docker exec -it test-mysql mysql -uroot -pmypassword testdb

   # Inside MySQL, try:
   SHOW DATABASES;
   EXIT;
   ```

## Task 5 - Volume Basics

1. **Temporary Storage vs Persistent Storage - Demonstrate Data Loss**

   ```console
   # Create a container and add some data
   docker run -it --name temp-container ubuntu:latest

   # Inside the container:
   echo "This data will be lost" > /tmp/myfile.txt
   cat /tmp/myfile.txt
   exit

   # Remove and recreate - data is gone
   docker rm temp-container
   docker run -it --name temp-container ubuntu:latest
   ls /tmp/  # myfile.txt is not there
   exit
   docker rm temp-container
   ```

2. **Using Volumes for Persistence**

   ```console
   # Create a volume
   docker volume create my-volume

   # Use the volume
   docker run -it --name persistent-container \
     -v my-volume:/data ubuntu:latest

   # Inside the container:
   echo "This data will persist" > /data/persistent.txt
   exit

   # Create new container with same volume
   docker rm persistent-container
   docker run -it --name new-container \
     -v my-volume:/data ubuntu:latest

   # Check the data is still there:
   cat /data/persistent.txt
   exit

   # Clean up
   docker rm new-container
   docker volume rm my-volume
   ```

## Task 6 - Debugging and Troubleshooting

1. **Common Debugging Commands - Container Inspection**

   ```console
   # Start my-web-server again if not running
   docker run -d -p 80:80 --name my-web-server httpd
   # Get detailed container information
   docker inspect my-web-server

   # Check resource usage
   docker stats

   # View running processes in container
   docker top my-web-server
   ```

2. **Log Analysis**

   ```console
   # Follow logs in real-time
   docker logs -f my-web-server

   # Get last 10 lines of logs
   docker logs --tail 10 my-web-server

   # Logs with timestamps
   docker logs -t my-web-server
   ```

3. **Exercise: Troubleshoot a Broken Container - Intentionally Break and Fix a Container**

   In this exercise, you will run a container with an invalid image tag to simulate a failure, investigate the issue, and then fix it by running the correct image.

   ```console
   # Run a container that will fail (using a valid image but an invalid command)
   docker run -d --name broken-nginx -p 80:80 nginx nginx-does-not-exist

   # This will fail - investigate why
   docker ps -a
   docker logs broken-nginx

   # Fix it
   docker rm broken-nginx
   docker run -d --name working-nginx -p 80:80 nginx:latest

   # Verify it works
   curl localhost
   ```

## Task 7 - Resource Management

1. **Setting Resource Limits - Run Containers with Resource Limits**

   ```console
   # Limit memory and CPU
   docker run -d --name limited-container \
     --memory="256m" \
     --cpus="0.5" \
     nginx

   # Check resource usage
   docker stats limited-container --no-stream

   # Clean up
   docker stop limited-container
   docker rm limited-container
   ```

## Task 8 - Cleanup and Best Practices

1. **System Cleanup - Clean Up Your System**

   ```console
   # Stop all running containers
   docker stop $(docker ps -q)

   # Remove all stopped containers
   docker container prune

   # Remove unused images
   docker image prune

   # Remove unused volumes
   docker volume prune

   # Remove unused networks
   docker network prune

   # Remove everything unused (be careful!)
   docker system prune -a
   ```

2. **Best Practices Summary - Key Takeaways**

   - Always name your containers (`--name`)
   - Use specific image tags, not `latest` in production
   - Clean up regularly to save disk space
   - Use volumes for persistent data
   - Monitor resource usage
   - Read container logs when debugging

## Verification Checklist

Before proceeding to Part 3, ensure you can:

- Pull and run containers successfully
- List, start, stop, and remove containers
- Map ports and access services
- Use environment variables
- Work with volumes for persistent storage
- Debug container issues using logs and inspection

## Next Steps

Proceed to [Part 3: Dockerfile and Image Building](../3-dockerfile/README.md) to learn how to create custom Docker images.

## Quick Reference

### Essential Commands

```console
# Container management
docker run -d -p HOST_PORT:CONTAINER_PORT --name NAME IMAGE
docker ps / docker ps -a
docker stop/start/restart CONTAINER
docker rm CONTAINER
docker logs CONTAINER

# Image management
docker images
docker pull IMAGE
docker rmi IMAGE

# System information
docker system info
docker system df
docker stats
```
