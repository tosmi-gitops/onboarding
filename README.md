* Requirements

- A tenant has multiple namespaces

- Namespaces can have different quotas per cluster
  - e.g.
    - cicd ns on cluster dev is allowed to create 100 pods
	- cicd ns on cluster prod is allowed to create 10 pods

- Not every namespace is on every cluster
