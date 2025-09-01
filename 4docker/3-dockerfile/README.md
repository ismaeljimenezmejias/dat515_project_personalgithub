# Lab 4: Getting Started with Docker: Deploying a Basic Web App

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Learning Objectives](#learning-objectives)
- [Task 1 - Your First Dockerfile](#task-1---your-first-dockerfile)
- [Task 2 - Dockerfile Instructions Deep Dive](#task-2---dockerfile-instructions-deep-dive)
- [Task 3 - Best Practices and Optimization](#task-3---best-practices-and-optimization)
- [Task 4 - Multi-Stage Builds](#task-4---multi-stage-builds)
- [Task 5 - Building for Production](#task-5---building-for-production)
- [Task 6 - Security Best Practices](#task-6---security-best-practices)
- [Task 7 - Image Registry Operations](#task-7---image-registry-operations)
- [Task 8 - Cleanup and Review](#task-8---cleanup-and-review)
- [Verification Checklist](#verification-checklist)
- [QuickFeed Checks](#quickfeed-checks)
- [Next Steps](#next-steps)

## Learning Objectives

- **Duration:** ~60 minutes
- **Prerequisites:** Completed Part 2 (Docker basics)

By the end of this lab, you will be able to:

- Write efficient Dockerfiles following best practices
- Build custom Docker images with proper layering
- Implement security best practices in image creation
- Use multi-stage builds for optimized images

## Task 1 - Your First Dockerfile

1. **Examine the Task 1 Files**

   First, examine the provided files in the `task1/` directory:

   ```console
   cd task1
   ls -la
   cat index.html
   cat Dockerfile
   ```

2. **Create Project Structure**

   Create the project structure and copy the necessary files:

   ```console
   mkdir my-web-app
   cd my-web-app
   mkdir website
   cp ../index.html website/
   cp ../Dockerfile .
   ```

3. **Complete the TODO Items for Task 1**

   Review the Dockerfile and complete the TODO items in the `Dockerfile`.

4. **Build and View Your Image**

   ```console
   # Build the image
   docker build -t my-web-app:v1.0 .
   # List images to see your new image
   docker images
   ```

5. **Run and Test Your Custom Image**

   ```console
   docker run -d -p 8080:80 --name my-app my-web-app:v1.0
   curl localhost:8080
   # Access from browser: http://your-floating-ip:8080
   ```

## Task 2 - Dockerfile Instructions Deep Dive

1. **Examine the Task 2 Files**

   ```console
   cd ../../task2
   cat app.py
   cat Dockerfile
   ```

2. **Create Advanced Application**

   Create the application structure and copy the necessary files:

   ```console
   mkdir advanced-app
   cd advanced-app
   cp ../app.py .
   cp ../Dockerfile .
   ```

3. **Complete the TODO Items for Task 2**

   Review the Dockerfile and complete the TODO items in the `Dockerfile`.

4. **Build and Run Advanced App**

   ```console
   docker build -t advanced-app:v2.0 .

   # Run with custom environment variables
   docker run -d -p 8001:8000 \
     --name advanced-app \
     -e APP_ENV=staging \
     -e APP_VERSION=2.1 \
     advanced-app:v2.0

   curl localhost:8001

   # Check health status
   docker ps  # Look for "healthy" status
   ```

## Task 3 - Best Practices and Optimization

1. **Examine the Task 3 Files**

   ```console
   cd ../../task3
   cat requirements.txt
   cat app.py
   cat Dockerfile
   cat .dockerignore
   ```

2. **Create Optimized Application**

   Create the application structure and copy the necessary files:

   ```console
   mkdir optimized-app
   cd optimized-app
   cp ../requirements.txt ../app.py ../Dockerfile ../.dockerignore ./
   ```

3. **Complete the TODO Items for Task 3**

   Review the Dockerfile and complete the TODO items in the `Dockerfile`.

4. **Compare and Analyze Image Sizes and Layers**

   ```console
   # Build optimized image
   docker build -t optimized-app:latest .

   # Compare image sizes
   docker images | grep -E "(advanced-app|optimized-app|my-web-app)"

   # Analyze image layers
   docker history optimized-app:latest

   # Run the optimized app
   docker run -d -p 8002:5000 --name optimized-app optimized-app:latest

   # Test it
   curl localhost:8002
   curl localhost:8002/health
   ```

## Task 4 - Multi-Stage Builds

1. **Examine the Task 4 Files**

   ```console
   cd ../../task4
   cat main.go
   cat go.mod
   cat Dockerfile
   ```

2. **Create Multi-Stage Application**

   Create the application structure and copy the necessary files:

   ```console
   mkdir go-app
   cd go-app
   cp ../main.go ../go.mod ../Dockerfile .
   ```

3. **Complete the TODO Items for Task 4**

   Review the Dockerfile and complete the TODO items in the `Dockerfile`.

4. **Build and Compare Multi-Stage Image**

   ```console
   # Build multi-stage image
   docker build -t go-app:multistage .

   # Compare sizes (notice how small the Go app is!)
   docker images | head -10

   # Run the Go application
   docker run -d -p 8003:8080 --name go-app go-app:multistage

   # Test it
   curl localhost:8003
   curl localhost:8003/health
   ```

## Task 5 - Building for Production

1. **Examine the Task 5 Files**

   ```console
   cd ../../task5
   cat Dockerfile
   cat index.html
   ```

2. **Create Production App**

   Create the project structure and copy the necessary files:

   ```console
   mkdir production-app
   cd production-app
   cp ../index.html ../Dockerfile .
   ```

3. **Complete the TODO Items for Task 5**

   Review the Dockerfile and complete the TODO items in the `Dockerfile`.

4. **Build and Inspect Labels**

   ```console
   # Build with build arguments
   docker build \
     --build-arg APP_VERSION=2.0 \
     --build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
     --build-arg VCS_REF=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown") \
     -t production-app:2.0 .

   # Inspect the labels
   docker inspect production-app:2.0 | grep -A 10 '"Labels"'
   ```

## Task 6 - Security Best Practices

1. **Scan Images for Vulnerabilities**

   ```console
   # Scan one of your images
   docker scout quickview optimized-app:latest

   # Get detailed vulnerability report
   docker scout cves optimized-app:latest
   ```

2. **Secure Dockerfile Checklist**

   - Using specific image tags (not 'latest')
   - Running as non-root user
   - Using minimal base images (alpine, slim)
   - Cleaning up package managers cache
   - Using .dockerignore to exclude sensitive files
   - Setting proper file ownership
   - Adding health checks
   - Using multi-stage builds to reduce final image size
   - Scanning for vulnerabilities

## Task 7 - Image Registry Operations

1. **Practice Tagging Images and Versioning**

   ```console
   # Tag images with different versions
   docker tag optimized-app:latest optimized-app:v1.0
   docker tag optimized-app:latest optimized-app:stable

   # List all your images
   docker images | grep optimized-app

   # Practice semantic versioning
   docker tag go-app:multistage go-app:1.0.0
   docker tag go-app:multistage go-app:1.0
   docker tag go-app:multistage go-app:latest
   ```

## Task 8 - Cleanup and Review

1. **Clean Up Your Environment**

   ```console
   # Stop all containers
   docker stop $(docker ps -q) 2>/dev/null || true

   # Remove containers
   docker container prune -f

   # Remove images you don't need
   docker image prune -f

   # Check remaining images
   docker images
   ```

## Verification Checklist

Before proceeding to Part 4, ensure you can:

- Build images with appropriate tags
- Write Dockerfiles with proper instruction ordering for optimal caching
- Use multi-stage builds effectively to reduce image size
- Run containers as non-root users
- Use .dockerignore to exclude unnecessary files
- Use build arguments for configurable builds
- Add health checks and metadata
- Implement security best practices
- Scan images for security vulnerabilities

## QuickFeed Checks

This lab has some **QuickFeed** checks to help you and us to validate your work.
We provide an even simpler check that you can run locally to check that your Dockerfiles are correctly configured.

```sh
cd 4docker/3-dockerfile
go test -v
```

If this test passes on your local machine, then the same test should also pass on QuickFeed.
QuickFeed has a few additional checks for the content of the Dockerfiles.

## Next Steps

Proceed to [Part 4: Docker Compose and Multi-Container Applications](../4-compose/README.md) to learn container orchestration.
