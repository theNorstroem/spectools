{
  "name": "Queryparam",
  "type": "Queryparam",
  "description": "Defines a queryparam field (for rpc type)",
  "__proto": {
    "package": "spec",
    "imports": [],
    "targetfile": "spec.proto"
  },
  "fields": {
    "type": {
      "description": "the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar",
      "meta": {
        "label": "type",
        "hint": "Use a scalar type or a already defined one"
      },
      "type": "string",
      "__proto": {
        "number": 2
      }
    },
    "description": {
      "description": "the field description",
      "meta": {
        "label": "description"
      },
      "type": "string",
      "__proto": {
        "number": 1
      },
      "__ui": {
        "component": "furo-data-textarea-input",
        "flags": [
          "full"
        ]
      }
    },
    "meta": {
      "description": "meta information for the client, like label, default, repeated, options...",
      "type": "furo.FieldMeta",
      "meta": {
        "label": "meta"
      },
      "__proto": {
        "number": 3
      }
    },
    "constraints": {
      "description": "constraints for a field, like min{}, max{}, step{}. Not used at the moment",
      "type": "map<string,furo.FieldConstraint>",
      "meta": {
        "label": "constraints"
      },
      "__proto": {
        "number": 4
      },
      "__ui": {
        "flags": [
          "full"
        ]
      }
    }
  }
}
