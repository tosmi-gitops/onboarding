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
* [Sharing a common base of configurations](#sharing-a-common-base-of-configurations)
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

1. A tenant has one to many namespaces
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

Kustomize allows use to deploy any Kubernetes manifest in any folder
and also patching of existing manifest if there's the need.

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
(e.g. `dinosaurs/clusters/prod`) we include namespaces that the cluster requires
from the `dinosaurs/namespaces/` folder.

The following diagram depicts connections between clusters, namespaces
and the onboarding-base repository:

![connections](https://raw.githubusercontent.com/tosmi-gitops/onboarding/main/docs/connections.png)

## Sharing a common base of configurations

Namespaces in `<tenants>/<tenant name>/namespaces` a sharing a common
base. The base is in a separate GIT repository
[onboarding-base](https://github.com/tosmi-gitops/onboarding-base.git).

Technically `onboarding-base` could be stored in the same repository
as `onboarding` but we have reasons to use a dedicated repository:

- **Making changes to base explicit**: changes in the common base for
  all namespace is potentially a dangerous operation. It could effect
  **all namespaces** on **all cluster**.

  We think that having this common base in a dedicated repository
  strengthens the awareness for this risk.

- **Allow usage of dedicated tags / releases**: We use tags to
  reference a (hopefully ) frozen state of the `onboarding-base`
  repository. If `onboarding-base` changes upstream the changes will
  not effect downstream configuration. See for example
  [here](https://github.com/tosmi-gitops/onboarding/blob/main/tenants/dinosaurs/namespaces/cicd/kustomization.yaml).

`onboarding-base` is a way to share common configurations, relevant
for all namespaces in all clusters. For example it defines the
following t-shirt shape sizes for projects:

- [small](https://github.com/tosmi-gitops/onboarding-base/tree/main/overlays/small)
- [medium](https://github.com/tosmi-gitops/onboarding-base/tree/main/overlays/medium)
- [large](https://github.com/tosmi-gitops/onboarding-base/tree/main/overlays/large)

Because of using kustomize, there are several ways to overwrite base
settings. For example in the CI/CD namespace for the
[goldfish](https://github.com/tosmi-gitops/onboarding/blob/main/tenants/goldfish/namespaces/cicd/kustomization.yaml)
project we use the medium overlay from `onboarding-base`, but
overwrite the memory limit to 7Gi.

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
