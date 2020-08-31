{
  "name": "fieldproto",
  "type": "Fieldproto",
  "description": "Proto options for a field",
  "__proto": {
    "package": "spec",
    "imports": [],
    "targetfile": "spec.proto"
  },
  "fields": {
    "number": {
      "description": "The field numbers are used to identify your fields in the message binary format, and should not be changed once your message type is in use.",
      "type": "int32",
      "meta": {
        "label": "Proto field number",
        "hint": "must be unique in fields"
      },
      "constraints": {
        "required": {
          "is": "true",
          "message": "a unique number is needed"
        }
      },
      "__proto": {
        "number": 2
      },
      "__ui": {
        "component": "furo-data-number-input"
      }
    },
    "oneof": {
      "description": "Assign field to a protobuf oneof group.",
      "type": "string",
      "meta": {
        "label": "Proto oneof group",
        "hint": ""
      },
      "constraints": {
      },
      "__proto": {
        "number": 3
      },
      "__ui": {
        "component": "furo-data-text-input"
      }
    }
  }
}
