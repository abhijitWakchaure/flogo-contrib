package rest

var jsonMetadata = `{
  "name": "tibco-rest",
  "version": "0.0.1",
  "description": "Simple REST trigger",
  "settings": [
    {
      "name": "port",
      "type": "integer"
    }
  ],
  "outputs": [
    {
      "name": "params",
      "type": "params"
    },
    {
      "name": "content",
      "type": "object"
    }
  ],
  "endpoint": {
    "settings": [
      {
        "name": "method",
        "type": "string"
      },
      {
        "name": "path",
        "type": "string"
      }
    ]
  }
}`