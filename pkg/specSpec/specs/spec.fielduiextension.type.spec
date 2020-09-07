{
  "name": "fielduiextension",
  "type": "Fielduiextension",
  "description": "ui hints for a field",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "component": {
      "type": "string",
      "description": "component hint for ui-builder",
      "__proto": {
        "number": 1
      },
      "meta": {
        "label": "Component hint",
        "hint": "ui-builder will use this element",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": {
        "required": {
          "is": "",
          "message": ""
      }
    }
    },
    "flags": {
      "type": "string",
      "description": "UI element flags like full, double, hidden,...",
      "__proto": {
        "number": 2
      },
      "meta": {
        "label": "Element flags",
        "hint": "UI element flags like full, double, hidden,...",
        "default": null,
        "readonly": false,
        "repeated": true,
        "options": {
          "list": []
        },
        "typespecific": null
      },
      "constraints": null
    },
    "no_init": {
      "type": "bool",
      "description": "Skip adding this field on ui init",
      "__proto": {
        "number": 3
      },
      "meta": {
        "label": "Skip init",
        "hint": "Skip initialization of this field in ui-builder",
        "default": null,
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        },
        "typespecific": null
      },
      "constraints": null
  },
    "no_skip": {
      "type": "bool",
      "description": "do not skip this field, even it is in the default skip list",
      "__proto": {
        "number": 4
      },
      "meta": {
        "label": "do not skip",
        "hint": "never skip initialization of this field in ui-builder",
        "default": null,
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        },
        "typespecific": null
      },
      "constraints": null
  }
}
}
