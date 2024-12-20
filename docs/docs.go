// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/info": {
            "get": {
                "description": "Fetches a song by its group and title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Retrieve a song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The group of the song",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The title of the song",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The requested song",
                        "schema": {
                            "$ref": "#/definitions/types.SongDetail"
                        }
                    },
                    "400": {
                        "description": "Invalid or missing query parameters",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/lyrics/{id}": {
            "get": {
                "description": "Retrieves the lyrics of a song in a paginated format",
                "tags": [
                    "Songs"
                ],
                "summary": "Get lyrics of a song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number (default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of verses per page (default: 5)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated lyrics of the song",
                        "schema": {
                            "$ref": "#/definitions/types.PaginatedVersesResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID or pagination parameters",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch the lyrics",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/music": {
            "get": {
                "description": "Retrieves a paginated list of songs with optional filters",
                "tags": [
                    "Songs"
                ],
                "summary": "List all songs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by group name",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by song title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number (default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of songs per page (default: 10)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of songs",
                        "schema": {
                            "$ref": "#/definitions/types.PaginatedSongsResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch the list of songs",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new song to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Add a new song",
                "parameters": [
                    {
                        "description": "Request to add a song",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddSongRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The added song",
                        "schema": {
                            "$ref": "#/definitions/models.Music"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to add the song to the database",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/music/{id}": {
            "put": {
                "description": "Updates an existing song by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Update a song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated song details",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Music"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The updated song",
                        "schema": {
                            "$ref": "#/definitions/models.Music"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID or payload",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update the song",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a song by its ID",
                "tags": [
                    "Songs"
                ],
                "summary": "Delete a song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deletion success message",
                        "schema": {
                            "$ref": "#/definitions/types.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid song ID",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete the song",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Music": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "types.AddSongRequest": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "types.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "types.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "types.PaginatedSongsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Music"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                },
                "totalSongs": {
                    "type": "integer"
                }
            }
        },
        "types.PaginatedVersesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                },
                "totalVerses": {
                    "type": "integer"
                }
            }
        },
        "types.SongDetail": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Music Library API",
	Description:      "This is an API for managing a music library, including adding, updating, deleting, and fetching songs with lyrics.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
