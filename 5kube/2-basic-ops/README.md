# Lab 5: Talos Kubernetes Cluster Setup

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Part 2: Basic Operations](#part-2-basic-operations)
- [Learning Goals](#learning-goals)
  - [Task 1 - Simple Pod](#task-1---simple-pod)
- [Debugging and Troubleshooting](#debugging-and-troubleshooting)
  - [Pod Debugging](#pod-debugging)
  - [Resource Issues](#resource-issues)
- [Common kubectl Commands Reference](#common-kubectl-commands-reference)
  - [Resource Management](#resource-management)
  - [Information and Debugging](#information-and-debugging)
- [Next Steps](#next-steps)

## Part 2: Basic Operations

This task covers fundamental Kubernetes operations and commands using kubectl.

## Learning Goals

- **Duration:** ~45 minutes
- **Prerequisites:** Completed Part 1

By completing this task, you will:

- Master essential kubectl commands and operations
- Understand Kubernetes API objects and resources
- Develop cluster inspection and debugging skills

### Task 1 - Simple Pod

**Note: This part is provided for your information only.**
It is not required in your specific setup.
Please review and adjust according to your needs.

Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.
For details about [Pods](https://kubernetes.io/docs/concepts/workloads/pods/).

1. Deploy a simple Pod.

    The downloaded file `simple-pod.yaml` and check contents of the file.

    ```console
    mkdir firstpod
    cd firstpod
    wget https://k8s.io/examples/pods/simple-pod.yaml
    cat simple-pod.yaml
    ```

    Deploy and manage:

    ```bash
    # Apply the manifest
    kubectl apply -f simple-pod.yaml
    ```

    See [Workload resources for managing pods](https://kubernetes.io/docs/concepts/workloads/pods/#workload-resources-for-managing-pods) for more information.

2. Check the status of running pods.

   ```bash
   # Check pod status
   kubectl get pods -n development

   # Get pod details
   kubectl describe pod nginx-pod -n development

   # Check pod logs
   kubectl logs nginx-pod -n development

   # Delete the pod
   kubectl delete -f simple-pod.yaml
   ```

3. Testing the nginx pod.

   To test the pod, you can use the `kubectl port-forward` command to forward traffic from a local port to a port on the pod.
   Here port 8080 on the local machine is forwarded to port 80 on the pod.

   ```bash
   kubectl port-forward pod/nginx 8080:80 &
   ```

   Then, using a tool like curl, access the pod:

   ```bash
   curl http://WorkerNodeIP:8080
   ```

   You should see the default nginx welcome page, indicating that your pod is running and serving traffic.
   We won't bother to expose this pod to the internet, as we will deploy a more complex application in the next part.

4. Deleting the Pod.

   Delete the pod using the `kubectl delete` command. It might take a few moments for the pod to be deleted.

   ```bash
   kubectl delete pod nginx
   ```

## Debugging and Troubleshooting

### Pod Debugging

   ```bash
   # Get pod logs
   kubectl logs <pod-name>
   kubectl logs <pod-name> -c <container-name>
   kubectl logs <pod-name> --previous

   # Execute commands in pod
   kubectl exec -it <pod-name> -- /bin/bash
   kubectl exec <pod-name> -- cat /etc/hostname

   # Port forwarding
   kubectl port-forward <pod-name> 8080:80

   # Copy files to/from pod
   kubectl cp local-file.txt <pod-name>:/tmp/
   kubectl cp <pod-name>:/tmp/file.txt ./local-file.txt
   ```

### Resource Issues

   ```bash
   # Check node resource usage
   kubectl describe node <node-name>

   # Check resource requests and limits
   kubectl describe pod <pod-name>

   # View events for debugging
   kubectl get events --field-selector involvedObject.name=<pod-name>
   ```

## Common kubectl Commands Reference

### Resource Management

   ```bash
   # Get resources
   kubectl get <resource>                    # List resources
   kubectl get <resource> <name>             # Get specific resource
   kubectl get <resource> -o yaml            # Output in YAML
   kubectl get <resource> -o json            # Output in JSON
   kubectl get <resource> -o wide            # Extended output

   # Create/Apply resources
   kubectl create -f <file>                  # Create from file
   kubectl apply -f <file>                   # Apply configuration
   kubectl apply -f <directory>              # Apply all files in directory

   # Delete resources
   kubectl delete <resource> <name>          # Delete by name
   kubectl delete -f <file>                  # Delete from file
   kubectl delete <resource> --all           # Delete all resources

   # Update resources
   kubectl edit <resource> <name>            # Edit resource
   kubectl patch <resource> <name> -p <patch> # Patch resource
   kubectl replace -f <file>                 # Replace resource
   ```

### Information and Debugging

   ```bash
   # Describe resources
   kubectl describe <resource> <name>        # Detailed information
   kubectl logs <pod> [-c container]         # View logs
   kubectl exec -it <pod> -- <command>       # Execute command in pod
   kubectl port-forward <pod> <local>:<remote> # Port forwarding

   # Cluster information
   kubectl cluster-info                      # Cluster details
   kubectl get events                        # Cluster events
   kubectl top nodes                         # Node resource usage
   kubectl top pods                          # Pod resource usage
   ```

## Next Steps

Once you're comfortable with basic operations, proceed to [Part 3: Pods and Workloads](../3-pods/) to dive deeper into launching multiple pods and deploying workloads.
