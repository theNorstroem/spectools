---
title: Introduction
type: docs
bookToc: false
description: フロー Spectools, the utility to work with Furo FIDLs 
---

# フロー Furo Spectools
Furo spectools is the main utility to work with furo specs ([FIDLs](https://fidl.furo.pro) ).

Furo spectools contains helpful generators, converters, sanitizer for the furo specs.


{{< columns >}}
## Stay In Sync

The different specification formats can be used as a source or a sink or both of them, but never use sources as source of the specs.

{{< mermaid >}}
graph LR
µSpec --> Spec
Spec --> Proto
Proto --> µSpec
Spec --> µSpec
{{< /mermaid >}}


<--->

## Multiple Sources Of Truth
You have the choice from multiple definition formats. This can be *.proto, one of the FIDLs (µSpec, spec).

{{< /columns >}}


## Add your own extensions
