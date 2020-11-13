---
title: "Overview"
weight: 10
# bookFlatSection: false
bookToc: false
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---


# Overview


{{< mermaid >}}
graph TD

µSpec(µSpec)
Spec([Spec])
µSpec-- spectools -->Spec
Spec-. spectools .->µSpec
someProto[Proto *1] -- protoc-gen-furo-specs --> µSpec

Spec-- spectools -->Es6Module
Spec-- furoc-gen-apidocs *2 -->Docs
Spec-- furoc-gen-u33e  -->U33E
U33E-- simple-generator  -->web-components
Spec-- furoc-gen-ddl *3  -->ddl[(DDL)]
Spec-- spectools  -->xs[...]
Spec-- spectools -->Proto


Proto-- protoc-gen-grpc-gateway  -->Gateway
Proto-- protoc-gen-openapiv2  -->OpenApi
OpenApi-- swagger  -->xo[...]
Proto-- protoc  -->xp[...]

{{< /mermaid >}}

## *1 usualy not the protos that you generate
## *2 not released yet
## *3 not open :-(
