{
  "name": "rpc",
  "type": "Rpc",
  "description": "Defines a rpc for a service",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [],
    "options": null
  },
  "fields": {
    "description": {
      "type": "string",
      "description": "the service description",
      "__proto": {
        "number": 1
      },
      "__ui": {
        "component": "furo-data-textarea-input",
        "flags": [
          "full"
        ],
        "no_init": false
      },
      "meta": {
        "label": "description",
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
    "rpc_name": {
      "type": "string",
      "description": "RPC name https://developers.google.com/protocol-buffers/docs/proto3#services",
      "__proto": {
        "number": 2
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "rpc name",
        "hint": "The rpc name",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": null
    },
    "deeplink": {
      "type": "spec.Servicedeeplink",
      "description": "This data is needed for...",
      "__proto": {
        "number": 5
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "deeplink",
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
    "data": {
      "type": "spec.Servicereqres",
      "description": "Request and response types for the service",
      "__proto": {
        "number": 3
      },
      "__ui": {
        "component": "",
        "flags": [],
        "no_init": false
      },
      "meta": {
        "label": "proto",
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
    "query": {
      "type": "map<string,spec.Queryparam>",
      "description": "Query params, it is recomended to use string types",
      "__proto": {
        "number": 4
      },
      "__ui": {
        "component": "",
        "flags": ["full"],
        "no_init": false
      },
      "meta": {
        "label": "Query params",
        "hint": null,
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
