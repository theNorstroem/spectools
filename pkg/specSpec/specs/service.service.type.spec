{
  "name": "service",
  "type": "Service",
  "description": "Defines a service",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "name": {
      "type": "string",
      "description": "Describe the rpcs or so",
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
        "label": "Name",
        "hint": "optional name",
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
    "version": {
      "type": "string",
      "description": "The version number, use semver",
      "__proto": {
        "number": 3
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "Version",
        "hint": "use semver",
        "default": "0.0.1",
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
      "description": "Describe the rpcs or so",
      "__proto": {
        "number": 2
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
        "label": "Description",
        "hint": "Describe the main purpose of this service",
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
    "lifecycle": {
      "type": "spec.Lifecycle",
      "description": "todo replace with lc type",
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
        "label": "",
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
    "__proto": {
      "type": "spec.Typeproto",
      "description": "information for the proto generator, should be removed for the client spec",
      "__proto": {
        "number": 5
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
    "services": {
      "type": "map<string,spec.Rpc>",
      "description": "RPCs for the service",
      "__proto": {
        "number": 6
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": true,
        "no_skip": false
      },
      "meta": {
        "label": "rpc services",
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