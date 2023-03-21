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
* [Additional Folders](#additional-folders)
* [OpenShift Gitops](#openshift-gitops)

## General scheme

The purpose of platform onboarding is to provide an automated way of
creating required OpenShift resources like

- namespaces
- quotas
- network policies

for developer teams and application operations teams.

Yes we know this is not DevOps, but corresponds to things we see in
real life. But what exactly is DevOps anyways?

We defined two repositories:

- [Onboarding-base](https://github.com/tosmi-gitops/onboarding-base.git)
- [Onboarding](https://github.com/tosmi-gitops/onboarding-base.git)

[Onboarding-base](https://github.com/tosmi-gitops/onboarding-base.git)
is used to provide manifests reused for all tenants (see
[Nomenclature](#nomenclature)).

[Onboarding](https://github.com/tosmi-gitops/onboarding-base.git)
contains manifest for the actual onboarding of tenants.

## Nomenclature

- Tenant: A team that requires OpenShift resources in one or more clusters.
- Cluster: An OpenShift cluster managed by a platform team

## Use cases considered

1. As developer I would like to get namespaces in a DEV cluster
2. As application ops I would like to create namespaces in the DEV cluster
3. As application ops I would like to create namespaces in the PROD cluster
4. As application ops I would like to overwrite namespace defaults
   depending on clusters
5. As a platform team I would like to leverage a common base for all
   Kubernetes manifests.


## Requirements considered

1. A tenant as one to many namespaces
2. Not all tenant namespaces are in all clusters
3. Namespaces might have different settings per cluster
   e.g.
    - CI/CD namespace on cluster dev is allowed to create 100 pods
	- CI/CD namespace on cluster prod is allowed to create 10 pods

## Tools used

We required [kustomize](https://kustomize.io) for rendering manifests
to onboard tenants. Helm is also possible and there are examples
available, see
[helper-proj-onboarding](https://github.com/tjungbauer/helm-charts/tree/main/charts/helper-proj-onboarding)
and the customers folder in
[openshift-cluster-bootstrap](https://github.com/tjungbauer/openshift-cluster-bootstrap/tree/main/customers)

## Tenants organization

We leverage the following folder structure for tenants:

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

## Additional Folders

- `onboarding`: Contains ArgoCD ApplicationSets for creating ArgoCD
  Applications from the tenants folder.
- `argocd`: Additional ArgoCD configuration like repositories and clusters.
- `docs`: Diagrams and other documentation artifacts

## OpenShift Gitops

OpenShift GitOps leverages ArgoCD to sync Kubernetes manifests from a
GIT repository to a cluster.

An ArgoCD `Application` defines which repository to sync into which cluster.

ArgoCD uses
[projects](https://argo-cd.readthedocs.io/en/stable/user-guide/projects/)
to organize applications in groups.

Specifically we use
[ApplicationSets](https://argo-cd.readthedocs.io/en/stable/user-guide/application-set/)
to generate required Applications. See
[onboarding/onboarding-applicationset.yaml](onboarding/onboarding-applicationset.yaml)
for more information.
