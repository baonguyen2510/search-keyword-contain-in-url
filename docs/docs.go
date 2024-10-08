// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/keyword/rank/:word": {
            "get": {
                "description": "Get the current rank of a specified keyword.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetKeywordRank"
                ],
                "summary": "Retrieve keyword rank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Keyword to search",
                        "name": "word",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.GetKeywordRankResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/model.KeywordResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Update rank of a specified keyword.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SyncKeywordRank"
                ],
                "summary": "Update keyword rank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Keyword to update",
                        "name": "word",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.KeywordResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/model.KeywordResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GetKeywordRankResponse": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                },
                "rank": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.KeywordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7003",
	BasePath:         "/v1/search",
	Schemes:          []string{"http"},
	Title:            "Search Keyword Service",
	Description:      "Service for search, update ranking of keyword.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
