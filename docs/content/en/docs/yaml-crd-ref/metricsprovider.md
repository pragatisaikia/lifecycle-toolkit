---
title: KeptnMetricsProvider
description: Define a data provider used for metrics and evaluations
weight: 55
---

`KeptnMetricsProvider` defines an instance of the data provider
(such as Prometheus, Dynatrace, or Datadog)
that is used by the [KeptnMetric](metric.md) resource.
One Keptn application can perform evaluations based on metrics
from more than one data provider
and, beginning with the v1alpha3 API version,
can use more than one instance of each data provider.
To implement this, create a `KeptnMetricsProvider` resource
for each instance of each data provider being used,
then reference the appropriate provider
for each metric definition by its name.

## Yaml Synopsis

```yaml
apiVersion: metrics.keptn.sh/v1alpha3
kind: KeptnMetricsProvider
metadata:
  name: <data-source-instance-name>
  namespace: <namespace>
spec:
  type: prometheus | dynatrace | dql | datadog
  targetServer: "<data-source-url>"
  secretKeyRef:
    name: <secret-name>
    key: <secret-key-that-holds-token>
```

## Fields

* **apiVersion** -- API version being used.
`
* **kind** -- Resource type.
   Must be set to `KeptnMetricsProvider`

* **metadata**
  * **name** -- Unique name of this provider,
    used to reference the provider for the
    [KeptnEvaluationDefinition](evaluationdefinition.md)
    and [KeptnMetric](metric.md) resources.
    Names must comply with the
    [Kubernetes Object Names and IDs](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names)
    specification.

    For example, you might define `dev-prometheus`
    for the Prometheus instance that monitors the development deployment,
    and `qa-prometheus` for the Prometheus instance
    that monitors the Prometheus instance that monitors the QA deployment,
    and `prod-dynatrace` for the Dynatrace instance
    that monitors the production deployment.

  * **namespace** -- Namespace where this provider is used.

* **spec**

  * **type** -- The type of data provider for this instance
  * **targetServer** -- URL of the data provider, enclosed in double quotes
  * **secretKeyRef**
    * **name:** -- Name of the token for this data provider
    * **key:** -- Key for this data provider

## Usage

## Examples

### Example 1: Dynatrace data provider

```yaml
apiVersion: metrics.keptn.sh/v1alpha3
kind: KeptnMetricsProvider
metadata:
  name: dynatrace
  namespace: podtato-kubectl
spec:
  targetServer: "<dynatrace-tenant-url>"
  secretKeyRef:
    name: dt-api-token
    key: DT_TOKEN
```

## Files

API Reference:

* [KeptnEvaluationDefinition](../crd-ref/lifecycle/v1alpha3/_index.md#keptnevaluationdefinition)

## Differences between versions

For the `v1alpha2` API version,
Keptn did not support
using more than one instance of a particular data provider
in the same namespace.
In other words, one namespace could support one instance each
of Prometheus, Dynatrace, and Datadog
but could not support, for example, two instances of Prometheus.

The synopsis in those older API versions
only specified the `metadata.name` field
that identified the data provider (`prometheus`, `dynatrace`, or `dql`):

```yaml
apiVersion: metrics.keptn.sh/v1alpha2
kind: KeptnMetricsProvider
metadata:
  name: prometheus | dynatrace |dql
  namespace: <namespace>
spec:
  targetServer: "<data-provider-url>"
  secretKeyRef:
    name: dt-api-token
    key: DT_TOKEN
```

Also note that, for the v1alpha1 and v1alpha2 API versions,
`KeptnMetricsProvider` only specifies the provider
for the `KeptnMetrics` resource.
Beginning with `v1alpha3` API version,
`KeptnMetricsProvider` is also used to specify the provider
for the `KeptnEvaluationDefinition` resource.

## See also

* [KeptnEvaluationDefinition](evaluationdefinition.md)
* [KeptnMetric](metric.md)
