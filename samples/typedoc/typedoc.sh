#!/bin/bash
# pipe the spec ast to simple-generator and write the result to dist/typedoc.md
spectools exportAsYaml | simple-generator -t scripts/typedoc.tpl > dist/typedoc.md