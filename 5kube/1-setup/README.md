# Lab 5: Talos Kubernetes Cluster Setup

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Part 1: Setup](#part-1-setup)
- [Learning Goals](#learning-goals)
- [Task 1 - Required login](#task-1---required-login)
  - [Prerequisites, needs to be done only once](#prerequisites-needs-to-be-done-only-once)
- [Task 2 - Deploy Talos VMs on OpenStack](#task-2---deploy-talos-vms-on-openstack)
  - [Launch VMs with Talos Image](#launch-vms-with-talos-image)
  - [Note VM IP Addresses](#note-vm-ip-addresses)
  - [Go to each VM and change the ntp server to ntp.ux.uis.no](#go-to-each-vm-and-change-the-ntp-server-to-ntpuxuisno)
- [Task 3 - Configure Talos Cluster](#task-3---configure-talos-cluster)
  - [Generate Talos Configuration](#generate-talos-configuration)
  - [Adjusting the Disk and Scheduling Settings](#adjusting-the-disk-and-scheduling-settings)
  - [Apply Configuration to Nodes](#apply-configuration-to-nodes)
  - [Bootstrap the Cluster](#bootstrap-the-cluster)
- [Task 4 - Verify Cluster Setup](#task-4---verify-cluster-setup)
  - [Get kubeconfig](#get-kubeconfig)
  - [Check Cluster Status](#check-cluster-status)
- [Task 5 - Set kubeconfig environment variable to access the Cluster](#task-5---set-kubeconfig-environment-variable-to-access-the-cluster)
- [Troubleshooting](#troubleshooting)
  - [Common Issues](#common-issues)
  - [Verification Commands](#verification-commands)
- [Next Steps](#next-steps)

## Part 1: Setup

This task covers setting up your Kubernetes development environment using Talos Linux on OpenStack.

## Learning Goals

- **Duration:** ~60–90 minutes
- **Prerequisites:** Access to UiS campus network, OpenStack credentials, and a working JumpVM account

By completing this task, you will:

- Deploy Talos Linux VMs on OpenStack
- Configure a basic Talos Kubernetes cluster
- Verify cluster connectivity and functionality

## Task 1 - Required login

Login to JumpVM with your username and password.

```console
ssh <github-username>@152.94.171.171
or
ssh <github-username>@sshvm.cs.ux.uis.no
```

### Prerequisites, needs to be done only once

1. Add Homebrew to your `PATH`. For example:

```bash
echo >> ~/.bashrc
echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> ~/.bashrc
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
```

## Task 2 - Deploy Talos VMs on OpenStack

### Launch VMs with Talos Image

1. **Login to OpenStack:**

   **Note that certain parts of this lab require UiS campus network access.**

   Login to [cloud.cs.ux.uis.no/horizon](https://cloud.cs.ux.uis.no/horizon) using the details provided to your group.

2. **Manage Security Group Rules:**

   - Navigate to `Project->Network->Security Groups`
   - Select Manage Rules, and click `+ Add Rule`
   - Select Custome TCP from the Rule dropdown menu and select port 50000
   - In the CIDR field, enter the IP address from where you want to connect
     For example, to grant access to the private network: `172.16.0.0/16`
   - Click Add

3. Navigate to Compute → Instances
4. Click "Launch Instance"
5. Configure the following:
   - **Instance Name**: `talos-master` (for control plane)
   - **Source**: Select Talos Linux image (talos 1.10.4)
   - **Flavor**: `m1.large`
   - **Networks**: Select your project network `dat515-network`
   - **Security Groups**: Default (ensure SSH and custom ports are allowed)
6. Repeat for worker nodes: `talos-worker1`, `talos-worker2`
   - **Flavor**: `m1.medium`

### Note VM IP Addresses

Record the IP addresses of your VMs:

- Control Plane: `<CONTROL_PLANE_IP>`
- Worker 1: `<WORKER1_IP>`
- Worker 2: `<WORKER2_IP>`

### Go to each VM and change the ntp server to ntp.ux.uis.no

1. Go to OpenStack page --> Projects --> Compute --> Instances --> Select your VM --> Console --> Click here to show console

2. Once the Console is open --> Press F3 button to go to Network Config

3. Move to Time Servers add `ntp.ux.uis.no` as the NTP server. --> save (No need to add any other details).
   Then you will not see any time query errors on the Talos console.

4. Repeat the above steps for all talos VMs.

## Task 3 - Configure Talos Cluster

### Generate Talos Configuration

These commands should be run on the JumpVM, where you have logged in.

```bash
# Generate cluster configuration, use endpoint-ip for master not any variable here.
# talosctl gen config dat515-groupX https://172.16.X.Y:6443
talosctl gen config dat515-groupX https://<endpoint-ip>:6443

# This creates several files:
# - controlplane.yaml
# - worker.yaml
# - talosconfig
```

### Adjusting the Disk and Scheduling Settings

1. **Edit `controlplane.yaml` and `worker.yaml`** to ensure the correct install disk (line 191):

   ```yaml
   install:
     disk: /dev/vda
   ```

2. **(Optional) Allow workloads on the control plane** by uncommenting:

   ```yaml
   allowSchedulingOnControlPlanes: true
   ```

   near the bottom of `controlplane.yaml`. This will untaint the control plane nodes.

### Apply Configuration to Nodes

```bash
# Apply control plane configuration
CONTROL_PLANE_IP=172.16.X.Y
talosctl apply-config --insecure --nodes $CONTROL_PLANE_IP --file controlplane.yaml

# Apply worker configuration
WORKER_NODE_IP1=172.16.X.Z
WORKER_NODE_IP2=172.16.X.W
talosctl apply-config --insecure --nodes $WORKER_NODE_IP1 --file worker.yaml
talosctl apply-config --insecure --nodes $WORKER_NODE_IP2 --file worker.yaml

# Set talosconfig context
export TALOSCONFIG=$(pwd)/talosconfig
talosctl config endpoint $CONTROL_PLANE_IP
talosctl config node $CONTROL_PLANE_IP
```

### Bootstrap the Cluster

```bash
# Bootstrap etcd on the control plane, no need to pass control plane IP, as we have already set it.
talosctl bootstrap

# Wait for the cluster to be ready (this may take several minutes)
talosctl health --wait-timeout=10m
```

## Task 4 - Verify Cluster Setup

### Get kubeconfig

```bash
# Retrieve kubeconfig from Talos
talosctl kubeconfig .

# Verify kubectl connectivity
kubectl --kubeconfig=kubeconfig get nodes
```

### Check Cluster Status

```bash
# Check all nodes are ready
kubectl --kubeconfig=kubeconfig get nodes

# Check system pods
kubectl --kubeconfig=kubeconfig get pods -n kube-system

# Check cluster info
kubectl --kubeconfig=kubeconfig cluster-info
```

## Task 5 - Set kubeconfig environment variable to access the Cluster

1. Set the `KUBECONFIG` environment variable to point to your kubeconfig:

    ```bash
    echo 'export KUBECONFIG=~/path/to/your/kubeconfig' >> ~/.bashrc
    ```

    Setting this kubeconfig, while running the `kubectl` commands in later, you do not have to pass kubeconfig file.

## Troubleshooting

### Common Issues

**VMs not accessible:**

- Check security group rules allow necessary ports
- Verify network connectivity
- Ensure Talos image is properly deployed

**Talos configuration fails:**

- Use `--insecure` flag for initial configuration
- Check VM IP addresses are correct
- Ensure VMs have booted completely

**Cluster bootstrap fails:**

- Wait sufficient time for etcd to initialize
- Check control plane VM has adequate resources

### Verification Commands

```bash
# Check Talos node status
talosctl --nodes <NODE_IP> health

# View Talos logs
talosctl --nodes <NODE_IP> logs

# Check cluster health
talosctl health

# Verify Kubernetes components
kubectl --kubeconfig=kubeconfig get componentstatuses
```

## Next Steps

Once your cluster is running, proceed to [Task 2: Basic Operations](../2-basic-ops/) to learn fundamental Kubernetes operations.
