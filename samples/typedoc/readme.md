## Generate a documentation of your types

[Back to overview](../readme.md)

Generate a simple documentation of your types


### usage
Add the scripts to your spec project and register the command and an according flow in the `.spectools` config.

```yaml

commands: #camelcase not allowed, commands con only be executed from a flow
  gen_type_docs: "./scripts/typedoc.sh"
flows:
  doc:
    - gen_type_docs
``` 

run the flow doc with `spectools run --flow=doc`. Or just run `spectools run` (which will run the default flow) if you add the sequence to your default flow.
This will run the flow `doc` which haves a sequence for `gen_type_docs`  which calls the script `./scripts/typedoc.sh`