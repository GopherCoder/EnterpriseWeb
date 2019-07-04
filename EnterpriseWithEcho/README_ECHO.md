# ECHO 使用文档


## 请求

``` 
GET /v1/api/wish/:wish_id?search='A'
```

- c.Param("wish_id") // 路径参数
- c.QueryParam("search") // 查询参数

``` 
POST /v1/api/wish

{
    "data": {
        "title": "echo"
    }
}
```

- c.Bind(&param) // 请求参数

## 响应

- c.JSON

## 中间件

- e.Use
- 自定义

``` 
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error
	}
}
```

