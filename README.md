# Cluster Onboarding

This repository contains an example implementation of OpenShift
project onboarding via GitOps.  It's implemented with
[kustomize](https://kustomize.io), but could be probably also done via
helm (see [here](#tools-used)).

## Table of contents

* [General scheme](#general-scheme)
* [Nomenclature](#nomenclature)
* [Use cases considered](#use-cases-considered)
* [Requirements considered](#requirements-considered)
* [Tools used](#tools-used)
* [Tenants organization](#tenants-organization)

## General scheme

The purpose of onboarding is to provide an automated way of creating
required OpenShift resources like

- namespaces
- quotas
- network policies

for developer teams and application operations teams.

Yes we know this is not DevOps, but corresponds to things we see in
real life. But what exactly is DevOps?

We defined two repositories:

- [Onboarding-base](https://github.com/tosmi-gitops/onboarding-base.git)
- [Onboarding](https://github.com/tosmi-gitops/onboarding-base.git)

[Onboarding-base](https://github.com/tosmi-gitops/onboarding-base.git)
is used to provide manifests reused for all other tenants (see
[Nomenclature](#nomenclature)).

[Onboarding](https://github.com/tosmi-gitops/onboarding-base.git)
contains manifest for the actual onboarding of tenants.

## Nomenclature

- Tenant: A team that requires OpenShift resources in one or more clusters.
- Cluster: An OpenShift cluster managed by a platform team

## Use cases considered

1. As developer I would like to get a namespace in a DEV cluster
2. As application ops I would like to create a namespace in the DEV cluster
3. As application ops I would like to create a namespace in the PROD cluster
4. As application ops I would like to overwrite namespace defaults depending on clusters

## Requirements considered

1. A tenant could have multiple namespaces
2. Not all namespaces are in all clusters
3. Namespaces might have different settings per cluster
   e.g.
    - cicd namespace on cluster dev is allowed to create 100 pods
	- cicd namespace on cluster prod is allowed to create 10 pods

## Tools used

We only require Kustomize for rendering manifests to onboard
tenants. Helm is also possible and there are examples available see
[helper-proj-onboarding](https://github.com/tjungbauer/helm-charts/tree/main/charts/helper-proj-onboarding)
and the customers folder in
[openshift-cluster-bootstrap](https://github.com/tjungbauer/openshift-cluster-bootstrap/tree/main/customers)

## Tenants organization

```
tenants/
├── dinosaurs
│   ├── clusters
│   │   ├── in-cluster
│   │   └── prod
│   └── namespaces
│       ├── cicd
│       └── dev
└── goldfish
    ├── clusters
    │   └── prod
    └── namespaces
        └── cicd
```

Every new tenant that requires onboarding is represented by a folder
in the `tenants` directory.

Within tenants we have a `clusters/` folder for every cluster the
tenant should be able to use. In the corresponding cluster folder
(e.g. `prod`) we include namespaces that the cluster requires from the
`namespaces/` folder.

The following diagram depicts connections between clusters, namespaces
and the onboarding-base repository:

![connections](https://raw.githubusercontent.com/tosmi-gitops/onboarding/main/docs/connections.png)
