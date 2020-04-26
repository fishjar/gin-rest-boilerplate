// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-26 21:47:42.961677443 +0800 CST m=+0.039526103

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login/account": {
            "post": {
                "description": "帐号登录...",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "帐号登录",
                "parameters": [
                    {
                        "description": "登录",
                        "name": "参数",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AuthAccountLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthAccountLoginRes"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户列表...",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserListRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Auth": {
            "type": "object",
            "properties": {
                "authCode": {
                    "description": "鉴权识别码",
                    "type": "string"
                },
                "authName": {
                    "description": "鉴权名称",
                    "type": "string"
                },
                "authType": {
                    "description": "鉴权类型",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "expireTime": {
                    "description": "过期时间",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "string"
                },
                "isEnabled": {
                    "description": "是否启用",
                    "type": "boolean"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "user": {
                    "description": "用户",
                    "type": "object",
                    "$ref": "#/definitions/model.User"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "string"
                },
                "verifyTime": {
                    "description": "认证时间",
                    "type": "string"
                }
            }
        },
        "model.AuthAccountLoginReq": {
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
        "model.AuthAccountLoginRes": {
            "type": "object",
            "required": [
                "accessToken",
                "expiresIn",
                "message",
                "tokenType"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "description": "过期时间（分钟）",
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "tokenType": {
                    "type": "string"
                }
            }
        },
        "model.Group": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "string"
                },
                "leader": {
                    "description": "队长",
                    "type": "object",
                    "$ref": "#/definitions/model.User"
                },
                "leaderId": {
                    "description": "队长ID",
                    "type": "string"
                },
                "name": {
                    "description": "组名称",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "users": {
                    "description": "队员",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "model.Menu": {
            "type": "object",
            "properties": {
                "children": {
                    "description": "子菜单",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "icon": {
                    "description": "菜单图标",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "string"
                },
                "name": {
                    "description": "菜单名称",
                    "type": "string"
                },
                "parent": {
                    "description": "父菜单",
                    "type": "object",
                    "$ref": "#/definitions/model.Menu"
                },
                "parentId": {
                    "description": "父ID",
                    "type": "string"
                },
                "path": {
                    "description": "菜单路径",
                    "type": "string"
                },
                "roles": {
                    "description": "角色",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Role"
                    }
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "model.PaginReq": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                }
            }
        },
        "model.Role": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "string"
                },
                "menus": {
                    "description": "菜单",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                },
                "name": {
                    "description": "角色名称",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "users": {
                    "description": "用户",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "地址",
                    "type": "string"
                },
                "auths": {
                    "description": "帐号",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Auth"
                    }
                },
                "avatar": {
                    "description": "昵称",
                    "type": "string"
                },
                "birthday": {
                    "description": "生日",
                    "type": "string"
                },
                "bloodType": {
                    "description": "血型(ABO)",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "friends": {
                    "description": "友",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "groups": {
                    "description": "组",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Group"
                    }
                },
                "height": {
                    "description": "身高(cm)",
                    "type": "number"
                },
                "homepage": {
                    "description": "个人主页",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "string"
                },
                "intro": {
                    "description": "简介",
                    "type": "string"
                },
                "lives": {
                    "description": "生活轨迹",
                    "type": "string"
                },
                "luckyNumbers": {
                    "description": "幸运数字",
                    "type": "string"
                },
                "mobile": {
                    "description": "手机",
                    "type": "string"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "notice": {
                    "description": "备注",
                    "type": "string"
                },
                "roles": {
                    "description": "角色",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Role"
                    }
                },
                "score": {
                    "description": "积分",
                    "type": "integer"
                },
                "tags": {
                    "description": "标签",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userNo": {
                    "description": "编号",
                    "type": "integer"
                }
            }
        },
        "model.UserListRes": {
            "type": "object",
            "required": [
                "page",
                "rows",
                "size",
                "total"
            ],
            "properties": {
                "page": {
                    "type": "integer"
                },
                "rows": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Host:        "localhost:4000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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
