---
title: Configuration
weight: 2
---

# Azure official provider-family documentation
Upbound supports and maintains the Upbound Azure official provider-family.

## Install the provider-family-azure
### Prerequisites
#### Upbound Up command-line
The Upbound Up command-line simplifies configuration and management of Upbound
Universal Crossplane (UXP) and interacts with the Upbound Marketplace to manage
users and accounts.

Install `up` with the command:
```shell
curl -sL "https://cli.upbound.io" | sh
```
More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

#### Upbound Universal Crossplane
UXP is the Upbound official enterprise-grade distribution of Crossplane for
self-hosted control planes. Only Upbound Universal Crossplane (UXP) supports
official providers.

Official providers aren't supported with open source Crossplane.

Install UXP into your Kubernetes cluster using the Up command-line.

```shell
up uxp install
```

Find more information in the [Upbound UXP documentation](https://docs.upbound.io/uxp/).

### Install the provider-family-azure

Install the Upbound official Azure provider-family with the following configuration file.
For instance, let's install the `provider-azure-management`

_Note_: The first provider installed of a family also installs an extra provider-family Provider.
The provider-family provider manages the ProviderConfig for all other providers in the same family.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azure-management
spec:
  package: xpkg.upbound.io/upbound/provider-azure-management:<version>
```

Define the `provider-azure-management` version with `spec.package`.

Install the `provider-azure-management` with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                            INSTALLED   HEALTHY   PACKAGE                                                     AGE
provider-azure-management       True        True      xpkg.upbound.io/upbound/provider-azure-management:v0.36.0   101s
upbound-provider-family-azure   True        True      xpkg.upbound.io/upbound/provider-family-azure:v0.36.0       93s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1) to view all available `Provider` options.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed)

## Configure the provider-family-azure
The Azure provider-family requires credentials for authentication to Azure Cloud Platform. The Azure provider-family consumes the credentials from a Kubernetes secret object.

### Install Azure CLI

Follow the documentation from Microsoft to [download and install the Azure
command-line](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli).

Log in to the Azure command-line:

```shell
az login
```

### Create an Azure service principal

Follow the Azure documentation to [find your Subscription
ID](https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id)
from the Azure Portal.

Using the Azure command-line, provide your Subscription ID to create a  service
principal and an authentication file.

```shell
# Replace <Subscription ID> with your subscription ID.
az ad sp create-for-rbac --sdk-auth --role Owner --scopes /subscriptions/<Subscription ID> \
  > azure.json
```

Here is an example key file:
```json
{
  "clientId": "00000000-0000-0000-0000-000000000000",
  "clientSecret": "XXX",
  "subscriptionId": "00000000-0000-0000-0000-000000000000",
  "tenantId": "00000000-0000-0000-0000-000000000000",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/"
}
```

### Generate a Kubernetes secret

Use the JSON file to generate a Kubernetes secret.

`kubectl create secret generic azure-secret --from-file=creds=./azure.json`

### Create a ProviderConfig object
Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: azure.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: azure-secret
      key: creds
```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.