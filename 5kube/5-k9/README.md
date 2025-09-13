# Lab 5: Talos Kubernetes Cluster Setup

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Part 5 - K9](#part-5---k9)
- [Next Steps](#next-steps)

## Part 5 - K9

- **Duration:** ~30 minutes
- **Prerequisites:** Completed Part 4.

For troubleshooting, inspecting, and otherwise operating the cluster, using a tool like `k9s` is highly recommended.
It will give you a view like this:

```console
Context: admin@dat515-groupX [RW]                 <0> all   ____  __ ________
Cluster: dat515-groupX                            <1> defau|    |/  /   __   \______
User:    admin@dat515-groupX                               |       /\____    /  ___/
K9s Rev: v0.50.6                                           |    \   \  /    /\___  \
K8s Rev: v1.33.1                                           |____|\__ \/____//____  /
CPU:     n/a                                                        \/           \/
MEM:     n/a
┌───────────────────────────────── pods(default)[6] ─────────────────────────────────┐
│ NAME↑                       PF READY STATUS   RESTARTS IP           NODE      AGE  │
│ mysql-86895c9c5f-jxpr4      ●  1/1   Running         0 10.244.2.9   worker-2  69m  │
│ mysql-86895c9c5f-m6bxn      ●  1/1   Running         0 10.244.0.4   master    69m  │
│ mysql-86895c9c5f-sjn2b      ●  1/1   Running         0 10.244.1.12  worker-1  69m  │
│ wordpress-7bbc564f76-29vxp  ●  1/1   Running         0 10.244.2.10  worker-2  69m  │
│ wordpress-7bbc564f76-xgmhc  ●  1/1   Running         0 10.244.0.5   master    69m  │
│ wordpress-7bbc564f76-xgsbk  ●  1/1   Running         0 10.244.1.13  worker-1  69m  │
│                                                                                    │
└────────────────────────────────────────────────────────────────────────────────────┘
  <pod>
```

k9s specifically, runs on a single binary in your terminal.

Installing can be done with:

```console
curl -sS https://webi.sh/k9s | sh
```

Then be sure to copy-paste the command the installer spits out at you.

Do note that you need to specify the kubeconfig k9s is using. If you have no default KUBECONFIG set, you are free to do so, or pass the kubeconfig directly to k9s:

```console
k9s --kubeconfig=/path/to/kubeconfig
```

Operate k9s as you would VIM.

Please consult the docs for usage: [K9s documentation](https://k9scli.io/topics/commands/)

## Next Steps

Once you understand K9 usage, proceed to [Task 6: Convert Docker to Kubernetes](../6-docker-to-kube/) to conver to your docker lab to kubernetes lab.
