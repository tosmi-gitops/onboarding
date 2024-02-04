# dinosaurs-accounts

Namespace for an accounting application. The namespaces should be
deployed on different cluster/environments:

- In `test` we would like to have a medium sized namespace
- In `prod` we would require a large sized namespace

The namespace definition is in `base`. For different sizes we use
overlays, to keep it simple we skip the overlays directory.

(clusters)[../../clusters] then use those overlays accordingly.
