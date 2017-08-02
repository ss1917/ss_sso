## 登录 
### API访问
- 获取配置  
  - **接口地址：** `/v1/accounts/login` 
  - **请求方式：** `POST` 

### 
**输入参数：**
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| username    | string | 是       | 用户名   |
| password    | string | 是       | 密码     |
| autologin   | string | 是       | 是否记住密码 | 

示例: 
- `/v1/accounts/login` 
- `contentType:"application/json"`
```json
{
  "username": "sss",
  "password": "Ss1234--",
  "autologin": "on"
}
```

## 注销 
### API访问
- 获取配置  
  - **接口地址：** `/v1/accounts/logout` 
  - **请求方式：** `GET` 

**输入参数：**
- 无参数 
- 示例: 
- `/v1/accounts/logout` 

## 判断登录状态 
### API访问
- 获取配置  
  - **接口地址：** `/v1/accounts/sso` 
  - **请求方式：** `POST` 

### 
**输入参数：**
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| auth_key    | string | 是       | token    |

示例: 
- `/v1/accounts/sso` 

``` 
Content-Disposition: form-data; name="auth_key"

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNzcyIsInVzZXJpZCI6MiwicmVnaW9uaWQiOjAsImV4cCI6MTUwMjI0Mjk1NywianRpIjoiMTU2MTg3MTgwNjAiLCJpc3MiOiJzaGVuc2h1byIsIm5iZiI6MTUwMTYzODA1N30.pNGIi3LPODDP6PMgAqaYUPCfNu6WGGcNdDW7OVb078k
```

## 鉴定是否有访问权限 
### API访问
- 获取配置  
  - **接口地址：** `/v1/rbac/verify` 
  - **请求方式：** `POST` 

### 
**输入参数：**
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| uid         | string |    是    | 用户id   | 
| uri         | string |    是    | 请求路径 |
| meth        | string |    是    | 请求方法 |

示例: 
- `/v1/rbac/verify`
``` 
Content-Type: multipart/form-data; 

Content-Disposition: form-data; name="uid"
2

Content-Disposition: form-data; name="meth"
GET

Content-Disposition: form-data; name="url"
/mg/umg/role

``` 

## 把用户权限缓存在redis
### API访问
- 获取配置  
  - **接口地址：** `/v1/rbac/verify` 
  - **请求方式：** `PUT` 

### 
**输入参数：**
| 参数名称    | 类型   | 是否必要 | 备注     |
| --          | --     | --       | --       |
| uid         | string |    是    | 用户id   | 

示例: 
- `/v1/rbac/verify`

```
Content-Type: multipart/form-data; 

Content-Disposition: form-data; name="uid"
2

``` 