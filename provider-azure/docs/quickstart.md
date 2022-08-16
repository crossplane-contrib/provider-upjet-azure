The Upbound Azure Provider is the officially supported provider for the Microsoft Azure Cloud.

View the Azure Provider Documentation for details and configuration options. 

## Quickstart
This guide walks through the process to install Upbound Universal Crossplane and install the Azure official provider.

To use this official provider, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfiguration`, and create a *managed resource* in Azure via Kubernetes.

## Create an Upbound.io user account
Create an account on [Upbound.io](https://cloud.upbound.io/register). 

<!-- Find detailed instructions in the [account documentation](/getting-started/create-account). -->

## Create an Upbound robot account and robot token
Installing an Official Provider requires an Upbound account and associated _Robot Token_.

To create a robot account and robot token in the Upbound Universal Console:
1. Log in to the [Upbound Universal Console](https://cloud.upbound.io) and select **Create New Organization** from the account menu.
2. Provide a unique **Organization ID** and **Display Name**.
3. Select the organization from the account menu.
4. Select **Admin Console**.
5. Select **Robots** from the left-hand navigation. 
6. Select **Create Robot Account**.
7. Provide a **Name** and optional description.
8. Select **Create Robot**.
9. Select **Create Token**.
10. Provide a **Name** for the token.

The console generates an `Access ID` and `Token` on screen. Save this token. The Console can't print the token again.

<!-- Find detailed instructions in the [Robot account and Robot Token](/upbound-cloud/robot-accounts) documentation.  -->

## Install the Up command-line
Install the [Up command-line](https://cloud.upbound.io/docs/cli/install) to connect to Upbound managed control planes.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
up uxp install
```

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
provider-aws-136444403808-cfd7b6dbb-s9l9n   1/1     Running   0             54m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

## Create a Kubernetes imagePullSecret for Upbound
Official providers require a Kubernetes `imagePullSecret` to download and install. The credentials for the `imagePullSecret` are from an Upbound robot token. 

Using the **robot token** generated earlier create an `imagePullSecret` with the command `kubectl create secret docker-registry package-pull-secret`.

```shell
kubectl create secret docker-registry package-pull-secret \
--namespace=upbound-system \
--docker-server=xpkg.upbound.io \
--docker-username=<robot token access ID> \
--docker-password=<robot token value> 
```

Replace `<robot token access ID>` with the `Access ID` of the robot token and `<robot token value>` with the value of the robot token.

Verify the secret with `kubectl get secrets`
```shell
$ kubectl get secrets -n upbound-system package-pull-secret
NAME                  TYPE                             DATA   AGE
package-pull-secret   kubernetes.io/dockerconfigjson   1      23s
```

## Create a Kubernetes imagePullSecret for Upbound
Official providers require a Kubernetes `imagePullSecret` to download and install. The credentials for the `imagePullSecret` are from an Upbound robot token. 

Using the **robot token** generated earlier create an `imagePullSecret` with the command `kubectl create secret docker-registry package-pull-secret`.

```shell
kubectl create secret docker-registry package-pull-secret --namespace=upbound-system --docker-server=xpkg.upbound.io --docker-username=<robot token access ID> --docker-password=<robot token value>
```

Replace `<robot token access ID>` with the `Access ID` of the robot token and `<robot token value>` with the value of the robot token.

Verify the secret with `kubectl get secrets`
```shell
$ kubectl get secrets -n upbound-system package-pull-secret
NAME                  TYPE                             DATA   AGE
package-pull-secret   kubernetes.io/dockerconfigjson   1      23s
```

## Install the official Azure provider in to the managed control plane
<!-- Use the marketplace button -->

Install the official provider into the managed control plane with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azure
spec:
  package: xpkg.upbound.io/upbound/provider-azure:v0.5.1
  packagePullSecrets:
    - name: package-pull-secret
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get provider
NAME             INSTALLED   HEALTHY   PACKAGE                                         AGE
provider-azure   True        True      xpkg.upbound.io/upbound/provider-azure:v0.5.1   58s
```

It may take up to 5 minutes to report `HEALTHY`.

If the `packagePullSecrets` is incorrect the provider returns a `401 Unauthorized` error. View the status and error with `kubectl describe provider`.

```yaml
$ kubectl describe provider
Name:         provider-azure
API Version:  pkg.crossplane.io/v1
Kind:         Provider
# Output truncated
Events:
  Type     Reason         Age                 From                                 Message
  ----     ------         ----                ----                                 -------
  Warning  UnpackPackage  34s (x7 over 100s)  packages/provider.pkg.crossplane.io  cannot unpack package: failed to fetch package digest from remote: HEAD https://xpkg.upbound.io/v2/upbound/provider-azure/manifests/v0.5.1: unexpected status code 401 Unauthorized (HEAD responses have no body, use GET for details)
```
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
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the managed control plane.

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

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.SecretRef.key` must match the value in the `Data` section of the secret.

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

Upbound created the resource group when the values `READY` and `SYNCED` are `True`.

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