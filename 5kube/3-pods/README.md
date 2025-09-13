# Lab 5: Talos Kubernetes Cluster Setup

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Part 3 - Deployment: Deploying a Simple Web Application on Kubernetes](#part-3---deployment-deploying-a-simple-web-application-on-kubernetes)
- [Learning Goals](#learning-goals)
- [Understanding Pods](#understanding-pods)
  - [What is a Pod?](#what-is-a-pod)
  - [Pod Lifecycle](#pod-lifecycle)
- [Task 1 - Pods](#task-1---pods)
- [Next Steps](#next-steps)

## Part 3 - Deployment: Deploying a Simple Web Application on Kubernetes

This task covers Kubernetes Pods, the smallest deployable units in Kubernetes.

## Learning Goals

- **Duration:** ~45 minutes
- **Prerequisites:** Completed Part 2

By completing this task, you will:

- Understand Pod concepts and lifecycle
- Create and manage Pods using YAML manifests
- Work with multi-container Pods
- Implement Pod networking and communication
- Configure resource requests and limits

## Understanding Pods

### What is a Pod?

A Pod is the smallest deployable unit in Kubernetes:

- Contains one or more containers
- Shares network and storage
- Scheduled as a single unit
- Ephemeral by nature

### Pod Lifecycle

Pods go through several phases:

- **Pending**: Pod accepted but not yet scheduled
- **Running**: Pod bound to node and containers created
- **Succeeded**: All containers terminated successfully
- **Failed**: All containers terminated, at least one failed
- **Unknown**: Pod state cannot be determined

## Task 1 - Pods

Usually you don't need to create Pods directly, even singleton Pods, as we did in Part 2.
Instead, you create and manage pods using workload resources like Deployment.

1. Create a Deployment Manifest.

   A Deployment ensures that a specified number of pod replicas are running at any given time.
Your task is to update `webapp-deployment.yaml` to define a simple Deployment for a web application using the nginx image.

   Save the following YAML to path:

   ```bash
   mkdir webapp
   cd webapp
   cp ../webapp-deployment.yaml  .
   ```

2. Create a Service Manifest.

   To make your web application accessible, you need to define a Service that exposes your Pods.  
   Your task is to update `webapp-service.yaml` to include a Service definition that exposes port 80 of your application using a NodePort service at port 30080.  
   This will enable external traffic to reach your application, providing load balancing and service

   Save the following YAML to path:

   ```bash
   cp ../webapp-service.yaml  .
   ```

3. Deploy the Application:

   Apply the Deployment and Service manifests:

    ```bash
    kubectl apply -f webapp-deployment.yaml
    kubectl apply -f webapp-service.yaml
    ```

4. Verify the Deployment:

   Check the status of the Deployment and Service:

    ```bash
    kubectl get deployments
    kubectl get services
    ```

   You should see your webapp-deployment with 2 replicas.
   You may need to give it a few moments to bring both replicas online.

5. Access the Web Application:

   Since we used a `NodePort` service, the web application should be accessible on node's IP at port 30080.
   If you're unsure of your node IPs, you can get them with:

    ```bash
    kubectl get nodes -o wide
    ```

   Then, in a web browser (through an ssh tunnel if you are in UiS cloud) or using `curl`, access the web application:

    ```bash
    curl WorkerNodeIP:30080
    ```

   You should see the default nginx welcome page, indicating that your web application is running.

   In order to access the web application from the internet, you need to do port forwarding using an ssh tunnel.

6. Deleting the Service and Deployment.

   Delete the service and deployment using the `kubectl delete` command.

    ```bash
    kubectl delete service webapp-service
    kubectl delete deployment webapp-deployment
    ```

## Next Steps

Once you understand Pod fundamentals, proceed to [Part 4: Deployments](../4-deployments/) to learn about managing Pod replicas and updates.
