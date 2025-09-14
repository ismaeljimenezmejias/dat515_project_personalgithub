FROM python:3.9-slim

# Build arguments
ARG APP_VERSION=1.0
ARG BUILD_DATE
ARG VCS_REF

# Labels for metadata
# TODO : Update with your email
LABEL maintainer="isjim2326@uis.no"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.version=$APP_VERSION

# Environment variables from build args
ENV APP_VERSION=$APP_VERSION

# TODO : Create a non-root user and group (appuser)
RUN groupadd -r appuser && useradd -r -g appuser appuser

# TODO : Set the working directory to /app
WORKDIR /app

# TODO : Copy the current directory contents into the container at /app
COPY . /app

# TODO : Switch to the appuser
USER appuser

# TODO : Expose port 8000
EXPOSE 8000

# Set the default command to run when the container starts
CMD ["python", "-m", "http.server", "8000"]
