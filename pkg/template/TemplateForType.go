package template

const TypeTPL = `{
  "name": "{{.type}}{{$type := .type}}",
  "type": "{{.type}}",
  "description": "{{.description}}",
  "__proto": {
    "package": "{{.package}}",
    "targetfile": "{{.target}}",
    "imports": [
      {{.protoImports}}
    ],
    "options": {}
  },
{{- $protonum:=1}}{{$first := true}}
  "fields": {
{{- range $key, $field := .fields }}
{{- $protonum = $protonum | add1}}
{{- $f:= $field }}{{if not $first}},{{end}}{{$first = false}}
    "{{$key}}": {
      "type": "{{$f.type}}",
      "description": "{{$f.description}}",
      "__proto": {
        "number": {{$protonum}}
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "label.{{$type | lower}}.{{$key}}",
        "hint": "",
        "default": "{{$f.defaultValue}}",
        "readonly": {{$f.readonly}},
        "repeated": {{$f.repeated}},
        "options": {
          "list": [],
          "flags": {}
        }
      },
      "constraints": {
      {{- if $f.required}}
        "required": {
          "is": "true",
          "message": "{{$key}} is required"
        }{{end}}
      }
    }
{{- end}}
  }
}`
