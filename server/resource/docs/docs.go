// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/yafgo/yafgo/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "description": "Api Root",
                "tags": [
                    "API"
                ],
                "summary": "ApiRoot",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/admin/menu": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "后台菜单",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/admin/menu/menus": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "Menu list",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "Menu 新增",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/admin/menu/menus/{id}": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "Menu 查询单条",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "Menu 更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "Menu 删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/admin/system/cfg": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "获取当前生效的配置",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/admin/system/cfg_in_redis": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "获取redis中的配置",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "后台"
                ],
                "summary": "更新redis中的配置",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backend.ReqCfgInRedis"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "description": "Api Index Demo",
                "tags": [
                    "API"
                ],
                "summary": "ApiIndex",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/user/info": {
            "get": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/user/login/username": {
            "post": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "description": "用户名登录",
                "tags": [
                    "Auth"
                ],
                "summary": "用户名登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ReqLoginUsername"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/user/register/username": {
            "post": {
                "security": [
                    {
                        "ApiToken": []
                    }
                ],
                "description": "用户名注册",
                "tags": [
                    "Auth"
                ],
                "summary": "用户名注册",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ReqRegisterUsername"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backend.ReqCfgInRedis": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "model.Menu": {
            "type": "object"
        },
        "service.ReqLoginUsername": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.ReqRegisterUsername": {
            "type": "object",
            "required": [
                "password",
                "username",
                "verify_code"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verify_code": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiToken": {
            "description": "接口请求token, 格式: ` + "`" + `Bearer {token}` + "`" + `",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "未分组接口",
            "name": "API"
        },
        {
            "description": "登录相关接口",
            "name": "Auth"
        },
        {
            "description": "后台管理相关接口",
            "name": "后台"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "YAFGO API",
	Description:      "基于 `Gin` 的 golang 项目模板\n- 本页面可以很方便的调试接口，并不需要再手动复制到 postman 之类的工具中\n- 大部分接口需要登录态，可以手动拿到 `登录token`，点击 `Authorize` 按钮，填入 `Bearer {token}` 并保存即可\n- 接口 url 注意看清楚，要加上 `Base URL` 前缀",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
