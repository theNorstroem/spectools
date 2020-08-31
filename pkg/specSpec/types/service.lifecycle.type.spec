{
  "name": "lifecycle",
  "type": "Lifecycle",
  "description": "Lifecycle for a service",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "deprecated": {
      "type": "bool",
      "description": "Is this version deprecated",
      "__proto": {
        "number": 1
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "deprecated",
        "hint": "",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": null
    },
    "info": {
      "type": "string",
      "description": "Inform about the replacement here, if you have one",
      "__proto": {
        "number": 2
      },
      "__ui": {
        "component": "furo-data-textarea-input",
        "flags": ["full"],
        "no_init": false
      },
      "meta": {
        "label": "Deprecation info",
        "hint": "Inform about the replacement here, if you have one",
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
