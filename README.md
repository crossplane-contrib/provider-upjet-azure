# Upbound Azure Provider

`provider-azure` is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Microsoft Azure](https://azure.microsoft.com/).

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/upbound/official-providers/provider-azure/releases):
```
kubectl crossplane install provider crossplane/provider-azure:v0.1.0
```

Alternatively, you can use declarative installation:
```
kubectl apply -f examples/install.yaml
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/official-providers/provider-azure).

## Contributing

Please refer to the [Adding New Resources](/docs/adding-resources.md) guide.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane/provider-azure/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://upboundio.slack.com)
* Twitter: [@upbound_io](https://twitter.com/upbound_io)

## Governance and Owners

`provider-azure` is owned and maintained by Upbound.

## Code of Conduct

provider-azure adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

# TODO(aru): determine license 
`provider-azure` is under ?.
