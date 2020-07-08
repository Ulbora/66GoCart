// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/Ulbora/Six910/blob/master/tos.html",
        "contact": {
            "name": "API Support",
            "url": "http://www.ulboralabs.com/contact/form"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://github.com/Ulbora/Six910/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rs/store/add": {
            "post": {
                "description": "Adds a new store to the system (only for OAuth2 stores)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Add a new store",
                "parameters": [
                    {
                        "description": "store",
                        "name": "store",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/six910-database-interface.Store"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OAuth2 client ID only for OAuth2 stores",
                        "name": "clientId",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID only for OAuth2 stores",
                        "name": "userId",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/managers.ResponseID"
                        }
                    }
                }
            }
        },
        "/rs/store/delete/{storeName}/{localDomain}": {
            "delete": {
                "description": "Delete a store from the system (only for OAuth2 stores)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Delete a store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store name",
                        "name": "storeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "store localDomain",
                        "name": "localDomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "apiKey required for non OAuth2 stores only",
                        "name": "apiKey",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "store name",
                        "name": "storeName",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "store localDomain",
                        "name": "localDomain",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OAuth2 client ID only for OAuth2 stores",
                        "name": "clientId",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID only for OAuth2 stores",
                        "name": "userId",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/managers.Response"
                        }
                    }
                }
            }
        },
        "/rs/store/get/{storeName}/{localDomain}": {
            "get": {
                "description": "Get details of a store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Get details of a store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store name",
                        "name": "storeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "store localDomain",
                        "name": "localDomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "apiKey required for non OAuth2 stores only",
                        "name": "apiKey",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "store name",
                        "name": "storeName",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "store localDomain",
                        "name": "localDomain",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OAuth2 client ID only for OAuth2 stores",
                        "name": "clientId",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID only for OAuth2 stores",
                        "name": "userId",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/six910-database-interface.Store"
                        }
                    }
                }
            }
        },
        "/rs/store/update": {
            "put": {
                "description": "Update store data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Update a store",
                "parameters": [
                    {
                        "description": "store",
                        "name": "store",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/six910-database-interface.Store"
                        }
                    },
                    {
                        "type": "string",
                        "description": "apiKey required for non OAuth2 stores only",
                        "name": "apiKey",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "store name",
                        "name": "storeName",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "store localDomain",
                        "name": "localDomain",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OAuth2 client ID only for OAuth2 stores",
                        "name": "clientId",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "User ID only for OAuth2 stores",
                        "name": "userId",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/managers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "managers.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "managers.ResponseID": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "six910-database-interface.Store": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "entered": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "localDomain": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "oauthClientId": {
                    "type": "integer"
                },
                "oauthSecret": {
                    "type": "string"
                },
                "remoteDomain": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "storeName": {
                    "type": "string"
                },
                "storeSlogan": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:3002",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Six910 API",
	Description: "This is the Six910 (six nine ten) API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
