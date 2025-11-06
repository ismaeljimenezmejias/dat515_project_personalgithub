# Kubernetes Deployment (HTTPS)

These instructions assume you already built the Docker image locally and loaded it into your cluster (for Minikube, use `minikube image load project-app:latest`).

1. **Ensure ingress controller**
   - Minikube: `minikube addons enable ingress`
   - Other clusters: install nginx ingress controller or equivalent.

2. **Create TLS secret** (reuse the dev cert):
   ```powershell
   kubectl create secret tls bike-market-tls `
     --cert=../certs/dev.crt `
     --key=../certs/dev.key
   ```
   Adjust the relative paths if you run the command from another directory. Add `-n <namespace>` if not using `default`.

3. **Apply manifests**
   ```powershell
   kubectl apply -f configmap.yaml -f redis-deployment.yaml -f redis-service.yaml `
     -f postgres-deployment.yaml -f postgres-service.yaml -f pvc.yaml `
     -f init-db-job.yaml -f backend-deployment.yaml -f backend-service.yaml `
     -f ingress.yaml
   ```

4. **Test access**
   ```powershell
   curl.exe -vk https://localhost/
   ```
   Browsers will warn about the self-signed cert; trust it locally or use `mkcert` for a trusted dev cert.
