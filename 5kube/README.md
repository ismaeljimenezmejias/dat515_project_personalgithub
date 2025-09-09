# Lab 5: Talos Kubernetes Cluster Setup

| Lab 5:           | Talos Kubernetes Cluster Setup |
| ---------------- | ------------------------------ |
| Subject:         | DAT515 Cloud Computing         |
| Deadline:        | **September 26, 2025 18:00**   |
| Expected effort: | 5-10 hours                     |
| Grading:         | Pass/fail                      |
| Submission:      | Individually                   |

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Background: What is Talos?](#background-what-is-talos)
- [Learning Goals](#learning-goals)
- [Lab Tasks](#lab-tasks)
  - [1. Setup](#1-setup)
  - [2. Deploy a Simple Application](#2-deploy-a-simple-application)
  - [3. Deploy pods](#3-deploy-pods)
  - [4. Deployments Wordpress + MySQL](#4-deployments-wordpress--mysql)
  - [5. K9](#5-k9)
  - [6. Docker to Kubernets](#6-docker-to-kubernets)
- [Submission](#submission)
- [Quick References](#quick-references)
- [Getting Started](#getting-started)

## Introduction

This lab introduces you to **Kubernetes**. In this lab, you will deploy a Kubernetes cluster using **Talos Linux** on OpenStack VMs and learn to manage containerized applications using Kubernetes.

## Background: What is Talos?

Talos OS is a lightweight, immutable Linux distribution built for running Kubernetes clusters securely and efficiently. Unlike traditional Linux distributions, Talos is managed entirely via an API and does not provide a shell or SSH access. This design reduces the attack surface and operational complexity, making it ideal for production environments.

**Why Talos?**

- Security: No SSH, minimal OS surface.
- Consistency: Immutable infrastructure.
- Automation: API-driven management.

## Learning Goals

By completing this lab, you will:

- Deploy a Talos-based Kubernetes cluster on OpenStack
- Understand core Kubernetes concepts and architecture
- Work with Pods, Deployments, Services, and other Kubernetes resources
- Deploy and manage multi-tier applications in Kubernetes
- Understand Kubernetes networking and storage concepts
- Troubleshoot common Kubernetes issues

## Lab Tasks

The goal of this lab is to understand how to deploy, expose, and manage services on a Kubernetes cluster running Talos. Each task below includes a brief explanation and a link to detailed instructions.

### 1. Setup

This section covers the setup of your Kubernetes environment using Talos Linux.

1. Deploy Talos VMs on OpenStack
2. Configure Talos cluster
3. Verify cluster connectivity

For detailed setup instructions, see [1-setup](1-setup/).

### 2. Deploy a Simple Application

Learn fundamental Kubernetes operations and concepts.

- Master essential kubectl commands
- Cluster inspection and debugging skills

For detailed instructions, see [2-basic-ops](2-basic-ops/).

### 3. Deploy pods

Understand Kubernetes Pods, the smallest deployable units.

- Learn to create and manage Kubernetes Pods
- Hands-on deployment of a simple web application

For detailed instructions, see [3-pods](3-pods/).

### 4. Deployments Wordpress + MySQL

- Learn about Kubernetes Deployments for managing application replicas
- Deploy a complete WordPress and MySQL application

For detailed instructions, see [4-deployments](4-deployments/).

### 5. K9

- Learn to use k9s, a terminal-based UI for Kubernetes cluster management and monitoring
- Troubleshooting, inspection, and operational tasks using an interactive dashboard interface

For detailed instructions, see [5-k9](5-k9/).

### 6. Docker to Kubernets

- Learn to migrate existing Docker applications to Kubernetes manifests
- Understand the transition from container-based development to orchestrated production deployments using Kubernetes resources

For detailed instructions, see [6-docker-to-kube/](6-docker-to-kube/).

## Submission

- Submit your code to the GitHub repository before the deadline.
- Attend the lab session to demonstrate your deployment and obtain approval.
- If you submit before the lab session but cannot get approval on the last date (because of the queue), we will try to approve it in the next lab session without any issues.

## Quick References

Each lab directory contains:

- README.md with complete instructions
- example-files/ directory with templates

## Getting Started

1. Navigate to the specific lab directory
2. Read the README.md file completely
3. Use the provided templates as starting points
4. Follow the step-by-step instructions
5. Verify your work with the provided checklists

Good luck with your Kubernetes journey!
