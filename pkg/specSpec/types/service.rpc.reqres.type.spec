{
  "name": "servicereqres",
  "type": "Servicereqres",
  "description": "Repuest and response types for services, used in service.type",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "request": {
      "type": "string",
      "description": "Define the request type, leave this field empty if not needed",
      "__proto": {
        "number": 1
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "request type",
        "hint": "leave empty if not needed",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": null
    },
    "response": {
      "type": "string",
      "description": "Define the response type, leave this field empty if not needed",
      "__proto": {
        "number": 2
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "response type",
        "hint": "leave empty if not needed",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": null
    }
  }
}
