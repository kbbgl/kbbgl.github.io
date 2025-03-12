---
slug: use-local-terraform-provider
title: How To Use Local Terraform Provider
description: A tutorial how to use a local Terraform provider
authors: [kgal-akl]
tags: [tutorial,how-to,terraform]
---

Let's say we have a repository for our own Terraform provider in `~/dev/kgal-tf-provider` and we want to use it in a `main.tf`. This document explains how to set it up.

## Prerequisites

`go` and `terraform` should be installed.

## Create Terraform Local Plugin Directory

Terraform expects the providers to be found in a specific directory structure.

We need to create this directory structure:

```bash
PROVIDER_VERSION=0.0.1
ARCH=darwin_amd64
PROVIDER_NAME=kgal/tf-provider
mkdir -p ~/.terraform.d/plugins/local/$PROVIDER_NAME/$PROVIDER_VERSION/$ARCH
```

In the example above, we're going to be creating this structure for a provider which will be built in aMac OS with Apple Silicon. You will need to modify that according to your system. To see what architecture you have, run:

```bash
go env GOOS GOARCH
darwin
amd_64
```

Once we have the Terraform local plugin directory structure, we can mode to building the provider and moving it to this directory.

## Build the Provider Locally

```bash
cd ~/dev/kgal-tf-provider
go mod tidy
go build -o ~/.terraform.d/plugins/local/$PROVIDER_NAME/$PROVIDER_VERSION/$ARCH/kgal-tf-provider
```

## Tell Terraform to Use Local Provider


```title=main.tf
terraform {
  required_providers {
    kgal-tf = {
        source = "local/kgal/tf-provider"
        version = "0.0.1"
    }
  }
}

provider "kgal-tf" { 
  # ....
}

resource "some_resource" "test_res" {
  # ...
}
```

## Reinitialize Terraform

```bash
terraform init -upgrade
terraform plan
terraform apply
```

### After Changes

If you need to modify the provider code, you will need to modify the provider version, [rebuild the provider](#build-the-provider-locally) and `terraform init -upgrade`. 
