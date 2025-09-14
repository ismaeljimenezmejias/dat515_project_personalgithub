# Use the official Nginx image as base
FROM nginx:latest

# Set maintainer information
LABEL maintainer="isjim2326@uis.no"
LABEL description="My first custom Docker image"

# TODO : Copy website folder with index.html to Nginx document root
# Nginx serves files from /usr/share/nginx/html
COPY website/ /usr/share/nginx/html/

# TODO : Expose port 80
EXPOSE 80

# Default command (inherited from nginx image)
CMD ["nginx", "-g", "daemon off;"]
