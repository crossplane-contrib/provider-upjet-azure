---
title: Quickstart
weight: 1
---

# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the Azure official provider-family.

To use Azure official provider-family, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in Azure via Kubernetes.

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.19.1
```

_Note_: official providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.13.2-up.2 installed
```

_Note_: Official provider-families only support crossplane version 1.12.1 or UXP version 1.12.1-up.1 or later.

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                       READY   STATUS    RESTARTS   AGE
crossplane-77ff754998-5j9rm                1/1     Running   0          49s
crossplane-rbac-manager-79b8bdd6d8-nwmb9   1/1     Running   0          49s
```

## Install the official Azure provider-family

Install the official provider-family into the Kubernetes cluster with a Kubernetes configuration file.
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

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
NAME                            INSTALLED   HEALTHY   PACKAGE                                                     AGE
provider-azure-management       True        True      xpkg.upbound.io/upbound/provider-azure-management:v0.36.0   101s
upbound-provider-family-azure   True        True      xpkg.upbound.io/upbound/provider-family-azure:v0.36.0       93s
```

It may take up to 5 minutes to report `HEALTHY`.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed)

## Create a Kubernetes secret
The Azure official provider-family requires credentials to create and manage Azure resources.

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
Create a `ProviderConfig` Kubernetes configuration file to attach the Azure credentials to the installed official `provider-azure-management`.

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
Create a managed resource to verify the `provider-azure-management` is functioning. 

This example creates an Azure ManagementGroup resource.

```yaml
apiVersion: management.azure.upbound.io/v1beta1
kind: ManagementGroup
metadata:
  annotations:
    meta.upbound.io/example-id: management/v1beta1/managementgroup
  labels:
    testing.upbound.io/example-name: example_parent
  name: example-parent
spec:
  forProvider:
    displayName: ParentGroup
  providerConfigRef:
    name: default
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get managed` to verify resource group creation.

```shell
$ kubectl get managed
NAME                                                         READY   SYNCED   EXTERNAL-NAME    AGE
managementgroup.management.azure.upbound.io/example-parent   True    True     example-parent   116s
```

Provider created the ManagementGroup when the values `READY` and `SYNCED` are `True`.

_Note:_ commands querying Azure resources may be slow to respond because of Azure API response times.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `spec.providerConfigRef.name` value in the `ManagementGroup` doesn't match the `ProviderConfig` `metadata.name`.

```shell
$ kubectl describe managementgroup
Name:         example-parent
Namespace:
Labels:       testing.upbound.io/example-name=example_parent
Annotations:  crossplane.io/external-create-succeeded: 2023-09-26T13:45:18Z
              crossplane.io/external-name: example-parent
              meta.upbound.io/example-id: management/v1beta1/managementgroup
              upjet.crossplane.io/provider-meta:
                {"e2bfb730-ecaa-11e6-8f88-34363bc7c4c0":{"create":1800000000000,"delete":1800000000000,"read":300000000000,"update":1800000000000}}
API Version:  management.azure.upbound.io/v1beta1
Kind:         ManagementGroup
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Display Name:                ParentGroup
    Parent Management Group Id:  /providers/Microsoft.Management/managementGroups/b9925bc4-8383-4c37-b9d2-fa456d1bb1c7
  Init Provider:
  Management Policies:
    *
  Provider Config Ref:
    Name:  dev
Status:
  At Provider:
    Display Name:                ParentGroup
    Id:                          /providers/Microsoft.Management/managementGroups/example-parent
    Parent Management Group Id:  /providers/Microsoft.Management/managementGroups/b9925bc4-8383-4c37-b9d2-fa456d1bb1c7
  Conditions:
    Last Transition Time:  2023-09-26T13:45:20Z
    Reason:                Available
    Status:                True
    Type:                  Ready
    Last Transition Time:  2023-09-26T13:52:07Z
    Message:               connect failed: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.azure.upbound.io "dev" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age               From                                                               Message
  ----     ------                   ----              ----                                                               -------
  Normal   CreatedExternalResource  7m3s              managed/management.azure.upbound.io/v1beta1, kind=managementgroup  Successfully requested creation of external resource
  Warning  CannotConnectToProvider  7s (x4 over 14s)  managed/management.azure.upbound.io/v1beta1, kind=managementgroup  cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.azure.upbound.io "dev" not found
```

The output indicates the `ManagementGroup` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `dev`. 

```shell
$ kubectl get providerconfig
NAME                                        AGE
providerconfig.azure.upbound.io/dev   89s
```
## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `ManagementGroup` object file. It takes a up to 5 minutes for Kubernetes to delete the resource and complete the command.

Verify removal of the resource group with `kubectl get managementgroup`

```shell
$ kubectl get managementgroup
No resources found
```