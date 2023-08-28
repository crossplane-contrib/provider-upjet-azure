---
title: Quickstart
weight: 1
---
⚠️ **Warning:** The monolithic Azure provider (`upbound/provider-azure`) has been deprecated in favor of the [Azure provider family](https://marketplace.upbound.io/providers/upbound/provider-family-azure/). You can read more about the provider families in our [blog post](https://blog.upbound.io/new-provider-families) and the official documentation for the provider families is [here](https://docs.upbound.io/providers/provider-families/). We will continue support for the monolithic Azure provider until June 12, 2024. And you can find more information on migrating from the monolithic providers to the provider families [here](https://docs.upbound.io/providers/migration/).

# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the Azure official provider.

To use this official provider, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in Azure via Kubernetes.

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.13.0
```

_Note_: official providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

## Install the official Azure provider

Install the official provider into the Kubernetes cluster with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azure
spec:
  package: xpkg.upbound.io/upbound/provider-azure:<version>
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get provider
NAME             INSTALLED   HEALTHY   PACKAGE                                         AGE
provider-azure   True        True      xpkg.upbound.io/upbound/provider-azure:v0.16.0   58s
```

It may take up to 5 minutes to report `HEALTHY`.

## Create a Kubernetes secret
The provider requires credentials to create and manage Azure resources.

### Install the Azure command-line
Generating an [authentication file](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-file-based-authentication) requires the Azure command-line. Follow the documentation from Microsoft to [Download and install the Azure command-line](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli).

### Create an Azure service principal
Follow the Azure documentation to [find your Subscription ID](https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id) from the Azure Portal.

Log in to the Azure command-line.

```command
az login
```

Using the Azure command-line and provide your Subscription ID create a service principal and authentication file.

```command
az ad sp create-for-rbac --sdk-auth --role Owner --scopes /subscriptions/<Subscription ID> 
```

The command generates a JSON file like this:
```json
{
  "clientId": "5d73973c-1933-4621-9f6a-9642db949768",
  "clientSecret": "24O8Q~db2DFJ123MBpB25hdESvV3Zy8bfeGYGcSd",
  "subscriptionId": "c02e2b27-21ef-48e3-96b9-a91305e9e010",
  "tenantId": "7060afec-1db7-4b6f-a44f-82c9c6d8762a",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/"
}
```

Save this output as `azure-credentials.json`.

### Create a Kubernetes secret with the Azure credentials JSON file
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the Kubernetes cluster.

`kubectl create secret generic azure-secret -n upbound-system --from-file=creds=./azure-credentials.json`

View the secret with `kubectl describe secret`
```shell
$ kubectl describe secret azure-secret -n upbound-system
Name:         azure-secret
Namespace:    upbound-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
creds:  629 bytes
```
## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the Azure credentials to the installed official provider.

```yaml
apiVersion: azure.upbound.io/v1beta1
metadata:
  name: default
kind: ProviderConfig
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: azure-secret
      key: creds
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  azure.upbound.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        creds
      Name:       azure-secret
      Namespace:  upbound-system
    Source:       Secret
```

## Create a managed resource
Create a managed resource to verify the provider is functioning. 

This example creates an Azure resource group.

```yaml
apiVersion: azure.upbound.io/v1beta1
kind: ResourceGroup
metadata:
  name: example-rg
spec:
  forProvider:
    location: "East US"
  providerConfigRef:
    name: default
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get managed` to verify resource group creation.

```shell
$ kubectl get managed
NAME                                        READY   SYNCED   EXTERNAL-NAME   AGE
resourcegroup.azure.upbound.io/example-rg   True    True     example-rg      20m
```

Provider created the resource group when the values `READY` and `SYNCED` are `True`.

_Note:_ commands querying Azure resources may be slow to respond because of Azure API response times.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `spec.providerConfigRef.name` value in the `ResourceGroup` doesn't match the `ProviderConfig` `metadata.name`.

```shell
$ kubectl describe resourcegroup
Name:         example-rg
Namespace:
Labels:       <none>
Annotations:  crossplane.io/external-name: example-rg
API Version:  azure.upbound.io/v1beta1
Kind:         ResourceGroup
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Location:  East US
  Provider Config Ref:
    Name:  default
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2022-07-26T16:54:26Z
    Message:               connect failed: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.azure.upbound.io "default" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age              From                                                  Message
  ----     ------                   ----             ----                                                  -------
  Warning  CannotConnectToProvider  2s (x4 over 7s)  managed/azure.upbound.io/v1beta1, kind=resourcegroup  cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.azure.upbound.io "default" not found
```

The output indicates the `ResourceGroup` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`. 

```shell
$ kubectl get providerconfig
NAME                                        AGE
providerconfig.azure.upbound.io/my-config   89s
```
## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `ResourceGroup` object file. It takes a up to 5 minutes for Kubernetes to delete the resource and complete the command.

Verify removal of the resource group with `kubectl get resourcegroup`

```shell
$ kubectl get resourcegroup
No resources found
```