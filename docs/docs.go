// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "team.cynexsc@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/address": {
            "post": {
                "description": "Create a new address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Create a new address",
                "parameters": [
                    {
                        "description": "Address data",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddressRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CreatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AddressResponseDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.BadRequestResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/dto.UnprocessableEntityResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register a new user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegistrationRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CreatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserResponseDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.BadRequestResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/dto.UnprocessableEntityResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddressRequestDTO": {
            "description": "Address data for creating a new address",
            "type": "object",
            "required": [
                "city",
                "neighborhood",
                "number",
                "state",
                "street",
                "zip_code"
            ],
            "properties": {
                "city": {
                    "description": "City is the name of the city.\nIt is a required field with a maximum length of 100 characters.",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Joinville"
                },
                "neighborhood": {
                    "description": "Neighborhood is the name of the neighborhood.\nIt is a required field with a maximum length of 100 characters.",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Centro"
                },
                "number": {
                    "description": "Number is the house or building number.\nIt is a required field with a minimum value of 1.",
                    "type": "integer",
                    "minimum": 1,
                    "example": 72
                },
                "state": {
                    "description": "State is the two-letter state code.\nIt is a required field with a fixed length of 2 characters.",
                    "type": "string",
                    "example": "SC"
                },
                "street": {
                    "description": "Street is the name of the street for the address.\nIt is a required field with a maximum length of 100 characters.",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Rua das Flores"
                },
                "zip_code": {
                    "description": "ZipCode is the postal code.\nIt is a required field with a fixed length of 8 characters.",
                    "type": "string",
                    "example": "89201304"
                }
            }
        },
        "dto.AddressResponseDTO": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "City is the name of the city.",
                    "type": "string",
                    "example": "Joinville"
                },
                "id": {
                    "description": "ID is the unique identifier for the address.",
                    "type": "integer",
                    "example": 1
                },
                "neighborhood": {
                    "description": "Neighborhood is the name of the neighborhood.",
                    "type": "string",
                    "example": "Centro"
                },
                "number": {
                    "description": "Number is the house or building number.",
                    "type": "integer",
                    "example": 72
                },
                "state": {
                    "description": "State is the two-letter state code.",
                    "type": "string",
                    "example": "SC"
                },
                "street": {
                    "description": "Street is the name of the street for the address.",
                    "type": "string",
                    "example": "Rua das Flores"
                },
                "zip_code": {
                    "description": "ZipCode is the postal code.",
                    "type": "string",
                    "example": "89201304"
                }
            }
        },
        "dto.BadRequestResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "null"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "title": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "dto.CreatedResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "integer",
                    "example": 201
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "title": {
                    "type": "string",
                    "example": "Created"
                }
            }
        },
        "dto.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "null"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "title": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "dto.UnprocessableEntityResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "name is required"
                },
                "status": {
                    "type": "integer",
                    "example": 422
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "title": {
                    "type": "string",
                    "example": "Unprocessable Entity"
                }
            }
        },
        "dto.UserRegistrationRequestDTO": {
            "type": "object",
            "required": [
                "address",
                "email",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "address": {
                    "description": "Address contains the user's address details.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.AddressRequestDTO"
                        }
                    ]
                },
                "email": {
                    "description": "Email is the user's email address.",
                    "type": "string",
                    "example": "israel@gmail.com"
                },
                "password": {
                    "description": "Password is the user's password.",
                    "type": "string",
                    "minLength": 6,
                    "example": "123456"
                },
                "role": {
                    "description": "Role defines the user's role in the system.",
                    "type": "string",
                    "enum": [
                        "teacher",
                        "student"
                    ],
                    "example": "student"
                },
                "username": {
                    "description": "Username is the unique identifier for the user.",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "Israel"
                }
            }
        },
        "dto.UserResponseDTO": {
            "type": "object",
            "required": [
                "address"
            ],
            "properties": {
                "address": {
                    "description": "Address contains the user's address details.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.AddressResponseDTO"
                        }
                    ]
                },
                "email": {
                    "description": "Email is the user's email address.",
                    "type": "string",
                    "example": "israel@gmail.com"
                },
                "id": {
                    "description": "ID is the unique identifier for the user.",
                    "type": "integer",
                    "example": 1
                },
                "role": {
                    "description": "Role defines the user's role in the system.",
                    "type": "string",
                    "example": "student"
                },
                "username": {
                    "description": "Username is the unique identifier for the user.",
                    "type": "string",
                    "example": "Israel Schroeder"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cynex Class Service",
	Description:      "Cynex Class Service API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
