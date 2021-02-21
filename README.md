# kubet
A simple manifest template generator for Kubernetes.

## Installation

```
$ go get -u github.com/codingconcepts/kubet
```

## Usage

Available commands:
```
deployment  Create a deployment manifest template.
namespace   Create a namespace manifest template.
service     Create a service manifest template.
```

Output to stdout:
```
$ kubet namespace name=MY_NAMESPACE
apiVersion: v1
kind: Namespace
metadata:
  name: MY_NAMESPACE
```

Output to a file:
```
$ kubet namespace name=MY_NAMESPACE > namespace.yaml
```