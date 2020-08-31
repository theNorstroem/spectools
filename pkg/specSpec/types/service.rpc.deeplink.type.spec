{
  "name": "servicedeeplink",
  "type": "Servicedeeplink",
  "description": "URL information for the service",
  "__proto": {
    "package": "spec",
    "targetfile": "spec.proto",
    "imports": [
      "furo/meta.proto"
    ],
    "options": null
  },
  "fields": {
    "description": {
      "type": "string",
      "description": "Describe the query params",
      "__proto": {
        "number": 1
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
    "rel": {
      "type": "string",
      "description": "the relationship",
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
        "label": "rel",
        "hint": "like create, update, custommethod",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": null
    },
    "method": {
      "type": "string",
      "description": "method of curl",
      "__proto": {
        "number": 3
      },
      "__ui": {
        "component": "furo-data-collection-dropdown",
        "flags": [],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "Method",
        "hint": "Http vebs like PUT, PATCH,...",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": [
            {
              "id": "GET",
              "display_name": "GET",
              "selected": true,
              "@type": "furo.Optionitem"
            },
            {
              "id": "PUT",
              "display_name": "PUT",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "POST",
              "display_name": "POST",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "PATCH",
              "display_name": "PATCH",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "DELETE",
              "display_name": "DELETE",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "OPTIONS",
              "display_name": "OPTIONS",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "HEAD",
              "display_name": "HEAD",
              "selected": false,
              "@type": "furo.Optionitem"
            },
            {
              "id": "TRACE",
              "display_name": "TRACE",
              "selected": false,
              "@type": "furo.Optionitem"
            }
          ]
        }
      },
      "constraints": null
    },
    "href": {
      "type": "string",
      "description": "The link pattern, like /api/xxx/{qp}/yyy",
      "__proto": {
        "number": 4
      },
      "__ui": {
        "component": "",
        "flags": [
          "double",
          "condensed"
        ],
        "no_init": false,
        "no_skip": false
      },
      "meta": {
        "label": "href",
        "hint": "The link pattern, like /api/xxx/{qp}/yyy",
        "default": "",
        "readonly": false,
        "repeated": false,
        "options": {
          "list": []
        }
      },
      "constraints": {
        "required": {
          "is": "true",
          "message": null
        }
      }
    }
  }
}
