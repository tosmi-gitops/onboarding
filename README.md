# Cluster onboarding

This repository conatins an example for OpenShift project onboarding via GitOps.
It's implemented with kustomize, but could be probably also done via helm.

## Use cases considered

1. As developer I would like to get a namespace in a DEV cluster
2. As application ops I would like to create a namespace in the DEV cluster
3. As application ops I would like to create a namespace in the PROD cluster
4. As application ops I would like to overwrite namespace defaults

## Requirements considered

1. A tenant could have multiple namespaces
2. Not all namespaces are in all clusters
3. Namespaces might have different settings per cluster
   e.g.
    - cicd ns on cluster dev is allowed to create 100 pods
	- cicd ns on cluster prod is allowed to create 10 pods

## Directory layout

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
