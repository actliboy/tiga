[TOC]
# user接口文档  
----------
# 用户相关  
----------
## 用户注册-v2(`/api/user/v2/add`)  
**func(*service.UserService, *contexti.Ctx, *user.SignupReq) (*response.TinyRep, error)** `/api/user/v2/add` _(Principal jyb)_  
### 接口记录  
|版本|操作|时间|负责人|日志|  
| :----: | :----: | :----: | :----: | :----: |  
|v1.0.0|创建|2019/12/16|jyb|创建|  
|v1.0.1|变更|2019/12/16|jyb|修改测试|  
### 参数信息  
|字段名称|字段类型|字段描述|校验要求|  
| :----  | :----: | :----: | :----: |  
|****|object|||  
|**.**|object|||  
|.|string|||  
|.|string|||  
|.|string|||  
|.|[]string|||  
|**.**|object|||  
|.|number|||  
|.|number|||  
|.|number|||  
||number|||  
||string|||  
|password|string|密码|为必填字段|  
|name|string|昵称|为必填字段|  
|gender|number|性别|为必填字段|  
|mail|string|邮箱||  
|phone|string|手机号||  
|vCode|string|验证码|为必填字段|  
__请求示例__  
```json  
{}  
```  
### 返回信息  
|字段名称|字段类型|字段描述|  
| :----  | :----: | :----: | 
|****|object||  
|**.**|object||  
|.|string||  
|.|string||  
|.|string||  
|.|[]string||  
|**.**|object||  
|.|number||  
|.|number||  
|.|number||  
||number||  
||string||  
|code|number||  
|message|string||  
__返回示例__  
```json  
{
	"code": 0
}  
```  
## 用户注册-v1(`/api/user/v1/add`)  
**func(*service.UserService, *contexti.Ctx, *response.TinyRep) (*response.TinyRep, error)** `/api/user/v1/add` _(Principal jyb)_  
### 接口记录  
|版本|操作|时间|负责人|日志|  
| :----: | :----: | :----: | :----: | :----: |  
|v1.0.0|创建|2019/12/16|jyb|创建|  
|v1.0.1|变更|2019/12/16|jyb|修改测试|  
### 参数信息  
|字段名称|字段类型|字段描述|校验要求|  
| :----  | :----: | :----: | :----: |  
|****|object|||  
|**.**|object|||  
|.|string|||  
|.|string|||  
|.|string|||  
|.|[]string|||  
|**.**|object|||  
|.|number|||  
|.|number|||  
|.|number|||  
||number|||  
||string|||  
|code|number|||  
|message|string|||  
__请求示例__  
```json  
{
	"code": 0
}  
```  
### 返回信息  
|字段名称|字段类型|字段描述|  
| :----  | :----: | :----: | 
|****|object||  
|**.**|object||  
|.|string||  
|.|string||  
|.|string||  
|.|[]string||  
|**.**|object||  
|.|number||  
|.|number||  
|.|number||  
||number||  
||string||  
|code|number||  
|message|string||  
__返回示例__  
```json  
{
	"code": 0
}  
```  
# 瞬间相关  
----------
