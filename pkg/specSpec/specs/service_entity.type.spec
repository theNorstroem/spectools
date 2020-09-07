{
  "name": "service_entity",
  "type": "SpecServiceEntity",
  "description": "SpecEntity with Spec",
  "__proto": {
    "package": "spec",
    "imports": [
      "furo/meta.proto",
      "furo/link.proto"
    ],
    "targetfile": "spec.proto"
  },
  "fields": {
    "data": {
      "description": "contains a spec.Spec",
      "type": "spec.Service",
      "__proto": {
        "number": 1
      }
    },
    "links": {
      "description": "Hateoas links",
      "type": "furo.Link",
      "meta": {"repeated": true},
      "__proto": {
        "number": 2
      }
    },
    "meta": {
      "description": "Meta for the response",
      "type": "furo.Meta",
      "__proto": {
        "number": 3
      }
    }
  }
}
