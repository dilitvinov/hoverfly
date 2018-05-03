package testdata

var TemplatingDisabled = `{
    "data": {
        "pairs": [
            {
                "response": {
                    "status": 200,
                    "body": "{{ Request.QueryParam.singular }}",
                    "encodedBody": false,
                    "templated" : false
                },
                "request": {
                    "method": {
						"exactMatch": "GET"
                    }
                }
            }
        ],
        "globalActions": {
            "delays": []
        }
    },
    "meta": {
        "schemaVersion": "v3",
        "hoverflyVersion": "v0.10.2",
        "timeExported": "2017-02-23T12:43:48Z"
    }
}`