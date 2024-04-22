# Project: Multi-Region Disaster Recovery (DR) for a Microservices Application

## Description 
This project focuses on building a robust and automated DR solution for a microservices application deployed across multiple GCP regions.

## Define The business problems

## Requirement
- Languages: version of Golang `go1.22.2`| version of python ``3.11``
- GCP:
  - ``Google Kubernetes Engine (GKE)`` for containerized application deployment in both primary and DR regions.
  - ``Cloud Storage`` for storing application artifacts and configuration files.
  - Cloud SQL or ``Cloud Spanner`` for managing the application database.
  - ``Cloud Load Balancing`` for distributing traffic across application instances.
  - ``Cloud Monitoring and Cloud Logging`` for centralized application health monitoring and logging.
- Infrastructure Tools:
  - ``GitHub Actions`` for automating the CI/CD pipeline.
  - ``Docker`` ``KB8`` for containerizing and Orchestration the application.
  - ``Terraform`` for infrastructure provisioning and configuration management in both regions.
  - `Vault` for securely storing sensitive application secrets.
  - `ArgoCD` for GitOps-based application deployment and management.
- API Design 
    - `gRPC` for high-performance scenarios

## Project structure Using UML 




## Features 
This project contains the following features:


## Installation 
To clone the project, create a folder ``multi-region-dr-microservices``
```
mkdir multi-region-dr-microservices
cd multi-region-dr-microservices
git clone https://github.com/zacksfF/Multi-Region-Disaster-Recovery-for-a-Microservices-Application.git
code . "for vscode"
```

