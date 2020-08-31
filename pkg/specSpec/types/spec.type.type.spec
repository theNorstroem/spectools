{
  "name": "type",
  "type": "Type",
  "description": "Defines a type in the furo spec",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "name": {
      "type": "string",
      "description": "Name of the type",
      "__proto": {
        "number": 1
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "name",
        "hint": "Internal name, informative value",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [],
          "flags": null
        }
      },
      "constraints": null
    },
    "type": {
      "type": "string",
      "description": "the type ",
      "__proto": {
        "number": 2
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "type",
        "hint": "The typename, without package. i.e. Person",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [],
          "flags": null
        }
      },
      "constraints": null
    },
    "description": {
      "type": "string",
      "description": "the type description",
      "__proto": {
        "number": 3
      },
      "__ui": {
        "component": "furo-data-textarea-input",
        "flags": [
          "full"
        ],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "description",
        "hint": "Describe what this type is for",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [],
          "flags": null
        }
      },
      "constraints": null
    },
    "__proto": {
      "type": "spec.Typeproto",
      "description": "information for the proto generator, should be removed for the client spec",
      "__proto": {
        "number": 4
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "proto",
        "hint": "",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [],
          "flags": null
        }
      },
      "constraints": null
    },
    "fields": {
      "type": "map<string,spec.Field>",
      "description": "fields of a type",
      "__proto": {
        "number": 5
      },
      "__ui": {
        "component": "",
        "flags": [
          "full"
        ],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "fields",
        "hint": "",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [],
          "flags": null
        }
      },
      "constraints": null
    }
  }
}