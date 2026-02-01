package docs

import (
	"strings"

	"github.com/swaggo/swag"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "support@kodingworks.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/categories": {
            "get": {
                "description": "Get all categories",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["categories"],
                "summary": "List all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Category"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new category",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["categories"],
                "summary": "Create category",
                "parameters": [
                    {
                        "description": "Category data",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                }
            }
        },
        "/api/categories/{id}": {
            "get": {
                "description": "Get category by ID",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["categories"],
                "summary": "Get category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Update category by ID",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["categories"],
                "summary": "Update category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category data",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete category by ID",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["categories"],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {"type": "string"}
                            }
                        }
                    }
                }
            }
        },
        "/api/produk": {
            "get": {
                "description": "Get all products",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["products"],
                "summary": "List all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["products"],
                "summary": "Create product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                }
            }
        },
        "/api/produk/{id}": {
            "get": {
                "description": "Get product by ID with JOIN category",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["products"],
                "summary": "Get product detail with category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Update product by ID",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["products"],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product by ID",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["products"],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {"type": "string"}
                            }
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health check endpoint",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["health"],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "status": {"type": "string"},
                                "message": {"type": "string"}
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                },
                "harga": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/entity.Category"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
type SwaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfoInstance describes our API
var SwaggerInfoInstance = &SwaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "Kasir API",
	Description: "API Kasir dengan Layered Architecture - Week 2 Challenge",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfoInstance
	t := docTemplate
	t = strings.Replace(t, "{{.Title}}", sInfo.Title, -1)
	t = strings.Replace(t, "{{.Description}}", sInfo.Description, -1)
	t = strings.Replace(t, "{{.Version}}", sInfo.Version, -1)
	t = strings.Replace(t, "{{.Host}}", sInfo.Host, -1)
	t = strings.Replace(t, "{{.BasePath}}", sInfo.BasePath, -1)
	t = strings.Replace(t, "{{ marshal .Schemes }}", `["http"]`, -1)
	t = strings.Replace(t, "{{escape .Description}}", sInfo.Description, -1)
	return t
}

func init() {
	swag.Register(swag.Name, &s{})
}
