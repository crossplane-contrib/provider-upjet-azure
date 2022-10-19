## Authentication mechanisms supported by provider-azure

`provider-azure` needs to authenticate itself to the Azure API. Like other
Crossplane providers, the credentials to be used during authentication can be
configured by means of `azure.upbound.io/v1beta1/ProviderConfig` resources.
`provider-azure` currently supports the following authentication mechanisms:
- Authentication with long-term service principal credentials
- Authentication using System-Assigned Managed Identities
- Authentication using User-Assigned Managed Identities

The authentication mechanism to be used can be selected by setting the
`spec.credentials.source` field of the `ProviderConfig` to one of the following
values: 
- `Secret`
- `SystemAssignedManagedIdentity`
- `UserAssignedManagedIdentity`

`Secret` is for configuring a set of long-term service principal credentials.
`SystemAssignedManagedIdentity` is for configuring authentication using the 
system-assigned managed identity of the AKS cluster.
And `UserAssignedManagedIdentity` is for configuring authentication using the
user-assigned managed identity of the AKS cluster. As it will be detailed below,
when `UserAssignedManagedIdentity` is used, the `spec.clientID` of the
`ProviderConfig` should also be set.

If no authentication mechanism is specified, the default is to use the
`Secret` authentication mechanism. 

Each authentication mechanism may need further configuration specific to it, and
the configuration details of each mechanism will be detailed below together with some
example `ProviderConfig` manifests. 


### Authentication with long-term Service Principal credentials
This authentication method involves making a set of long-term service principal
credentials available to `provider-azure` by means of a Kubernetes `Secret`, and
thus is discouraged from a security perspective. Details can be found in [[1]]
under the ``Create a Kubernetes secret`` section. The
required configuration is a reference to a Kubernetes secret containing the
long-term credentials. And example `Secret` configuration is as follows:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-azure-creds
  namespace: upbound-system
type: Opaque
data:
  creds: <base64-encoded service principal credentials document>
---
apiVersion: azure.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-azure-creds
      namespace: upbound-system
      key: creds

```
The service principal credentials document is a JSON document
and looks like the following:

```
{
  "clientId": "<client ID of the service principal>",
  "clientSecret": "<client secret of the service principal>",
  "subscriptionId": "<subscription ID containing the service principal>",
  "tenantId": "<tenant ID of the service principal>",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/"
}
```

As mentioned above, this authentication method involves a service principal's
long-term credentials and is considered as less secure when compared to other
configurations.

### Authentication using System-Assigned Managed Identities
Authentication using system-assigned managed identities [[2]] is available
when `provider-azure` is running on an AKS cluster and managed identities
have been enabled for the cluster [[3]].

An example `ProviderConfig` is as follows:

```yaml
apiVersion: azure.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: SystemAssignedManagedIdentity
  subscriptionID: <subscription ID>
  tenantID: <tenant ID>
```

One important point to note here is that the AKS cluster's **kubelet
identity** must have the necessary role assignments [[3]] so that
the cloud API calls made by provider-azure will be authorized.
The set of roles required by provider-azure depends on the resources
being managed by it.

As it can be seen in the example `ProviderConfig` above, no explicit credentials
need to be referred by the `ProviderConfig`. This authentication
method, thus, is considered as more secure when compared to the `Secret` method
discussed above. 


### Authentication using User-Assigned Managed Identities
This method is similar to authentication using system-assigned managed
identities and requires that the AKS cluster is using a user-assigned
managed identity as its kubelet identity [[3]]. An example `ProviderConfig`
is as follows:

```yaml
apiVersion: azure.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: UserAssignedManagedIdentity
  clientID: <clientId of the user-assigned managed identity used as the kubelet identity>
  subscriptionID: <subscription ID>
  tenantID: <tenant ID>
```

Again, this `ProviderConfig` requires no long-term service principle credentials,
and is thus considered as more secure when compared to the `Secret` method.

When using authentication with the user-assigned managed identity, you should ensure
that the managed identity being used as the **kubelet identity** in the AKS cluster has
the required role assignments. The set of roles required by provider-azure depends
on the resources being managed by it.Please refer to [[3]] for details.

[1]: https://marketplace.upbound.io/providers/upbound/provider-azure/v0.17.0

[2]: https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview

[3]: https://learn.microsoft.com/en-us/azure/aks/use-managed-identity

