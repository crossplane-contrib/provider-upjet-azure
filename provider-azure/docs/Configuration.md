# Configuration

Install official provider on self-hosted control planes using Universal Crossplane (`UXP`).

## Install the provider
Official providers require a Kubernetes `imagePullSecret` to install. 
<!-- vale gitlab.Substitutions = NO --> 
<!-- Details on creating an `imagePullSecret` are available in the [generic provider documentation](/providers/#create-a-kubernetes-imagepullsecret) -->
<!-- vale gitlab.Substitutions = YES --> 

_Note:_ if you already installed an official provider using an `imagePullSecret` a new secret isn't required.

Install the Upbound official Azure provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azure
spec:
  package: xpkg.upbound.io/upbound/provider-azure:latest
  packagePullSecrets:
    - name: package-pull-secret
```

Define the provider version with `spec.package`. This example uses version `latest`.

The `spec.packagePullSecrets.name` value matches the Kubernetes `imagePullSecret`. The secret must be in the same namespace as the Upbound pod.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get provider`.

```shell
$ kubectl get providers
NAME             INSTALLED   HEALTHY   PACKAGE                                         AGE
provider-azure   True        True      xpkg.upbound.io/upbound/provider-azure:latest   11m
```

View the [Provider CRD definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1) to view all available `Provider` options.

## Configure the provider
The Azure provider requires credentials for authentication to Azure Cloud. The
Azure provider consumes the credentials from a Kubernetes secret object.

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

The command generates a JSON file like this:
```json
{
  "clientId": "5d73973c-1933-4621-9f6a-9642db949768",
  "clientSecret": "24O8Q~db2DFJ123MBpB25hdESvV3Zy8cSd",
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

Use the JSON file to generate a Kubernetes secret.

`kubectl -n upbound-system create secret generic azure-secret --from-file=creds=./azure.json`

### Create a ProviderConfig object
Apply the secret in a `ProviderConfig` Kubernetes configuration file.

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

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl -n upbound-system create secret generic <name>` command.
