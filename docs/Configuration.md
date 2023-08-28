---
title: Configuration
weight: 2
---
⚠️ **Warning:** The monolithic Azure provider (`upbound/provider-azure`) has been deprecated in favor of the [Azure provider family](https://marketplace.upbound.io/providers/upbound/provider-family-azure/). You can read more about the provider families in our [blog post](https://blog.upbound.io/new-provider-families) and the official documentation for the provider families is [here](https://docs.upbound.io/providers/provider-families/). We will continue support for the monolithic Azure provider until June 12, 2024. And you can find more information on migrating from the monolithic providers to the provider families [here](https://docs.upbound.io/providers/migration/).

# Azure official provider documentation
Upbound supports and maintains the Upbound Azure official provider.

## Install the provider
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

### Install the provider

Install the Upbound official Azure provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azure
spec:
  package: xpkg.upbound.io/upbound/provider-azure:<version>
```

Define the provider version with `spec.package`.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-azure   True        True      xpkg.upbound.io/upbound/provider-azure:v0.16.0  62s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1) to view all available `Provider` options.

## Configure the provider
The Azure provider requires credentials for authentication to Azure Cloud Platform. The Azure provider consumes the credentials from a Kubernetes secret object.

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