# Lab 5: Talos Kubernetes Cluster Setup

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Part 4 - Deploying WordPress and MySQL on Kubernetes](#part-4---deploying-wordpress-and-mysql-on-kubernetes)
- [Learning Goals](#learning-goals)
- [Understanding Deployments](#understanding-deployments)
  - [What is a Deployment?](#what-is-a-deployment)
  - [Deployment Architecture](#deployment-architecture)
- [Task - Creating Your Deployment](#task---creating-your-deployment)
- [Next Steps](#next-steps)

## Part 4 - Deploying WordPress and MySQL on Kubernetes

This part covers Kubernetes Deployments for managing application replicas.

## Learning Goals

- **Duration:** ~45 minutes
- **Prerequisites:** Completed Part 3.

By completing this task, you will:

- Understand Deployment concepts and architecture
- Create and manage Deployments
- Implement scaling strategies
- Perform rolling updates and rollbacks
- Work with ReplicaSets

## Understanding Deployments

### What is a Deployment?

A Deployment provides declarative updates for Pods and ReplicaSets:

- Manages desired state for application replicas

### Deployment Architecture

```text
Deployment
└── ReplicaSet
    └── Pod
    └── Pod
    └── Pod
```

## Task - Creating Your Deployment

1. Deploy MySQL.

    You need to define a Deployment for MySQL in `mysql-deployment.yaml`.  
    This Deployment should ensure that a MySQL Pod is always running with the desired configuration.

    Review the `mysql-deployment.yaml` file and complete the TODO items to complete the Task.

    Copy the `mysql-deployment.yam` in to the deployment folder.

    ```bash
    mkdir deployment
    cd deployment
    cp ../mysql-deployment.yaml .
    cat mysql-deployment.yaml
    ```

   Apply the Deployment:

    ```bash
    kubectl apply -f mysql-deployment.yaml
    ```

2. Create a Service for MySQL:

    To enable WordPress to connect to the MySQL database, you need to create a Service for MySQL.
    Review the `mysql-service.yaml` file and complete the TODO items to complete the Task.
    Copy the `mysql-service.yaml` file into your `deployment` folder:

    ```bash
    cp ../mysql-service.yaml .
    ```

    Apply the MySql Service:

    ```bash
    kubectl apply -f mysql-service.yaml
    ```

3. Deploy WordPress.

    Copy the `wordpress-deployment.yaml` in to the deployment folder.
    Review the `wordpress-deployment.yaml` file and complete the TODO items to complete the Task.

    ```bash
    cp ../wordpress-deployment.yaml .
    ```

   Apply the Wordpress Deployment:

    ```bash
    kubectl apply -f wordpress-deployment.yaml
    ```

4. Create a Service for WordPress.

    This step will create a Service that makes your WordPress application accessible to users.
    Copy the `wordpress-service.yaml` in to the deployment folder.
    Review the `wordpress-service.yaml` file and complete the TODO items to complete the Task.

    ```bash
    cp ../wordpress-service.yaml .
    ```

   Apply the Wordpress Service:

    ```bash
    kubectl apply -f wordpress-service.yaml
    ```

5. Access WordPress.

    Since WordPress is exposed using a NodePort service, it will be accessible externally on the node's IP address at a port number in the 30000+ range assigned by Kubernetes.  
To determine which NodePort has been assigned to the WordPress, get svc for wordpress:

    ```bash
    kubectl get svc wordpress
    ```

   Then, in a web browser with the ssh tunnel, access WordPress:

    ```bash
    http://< INTERNAL-IP>:<NODE_PORT>
    ```

## Next Steps

Once you master Deployments, proceed to [Task 5: K9](../5-k9/) to learn about troubleshooting, inspecting, and operating your talos cluster.
