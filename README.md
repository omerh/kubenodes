# kubenodes

Get a view of pods according to the spread on the nodes.

This app was made to view pods according to well known labels, and not with AWS SDK to determine if an instance is a spot or on-demand, instead it relays on <karpenter.sh> labels (I might add AWS SDK for mode node information).

Well-known label I am using:

- karpenter.sh/capacity-type
- node.kubernetes.io/instance-type
- kubernetes.io/arch
- topology.kubernetes.io/zone

>Please submit PR for more labels, or to add AWS SDK for more instance information

In order to list pod status according to its pods, use the label `app`

```bash
kubenodes --deployment my-app
```

and make sure to label the app with `app: my-app` in kubernetes manifest

you have the option to use `--compact` flag to see a compact list of pods in a single node row.
