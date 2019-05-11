# 个人博客系统设计与实现
基于Go语言的个人博客系统

------

### 采用前后端分离的架构

前端使用Vue+ElementUI框架及组件

后端使用 Gin+Gorm 框架实现 RESTful 风格的微服务



### Feature 

- 前后端分离
- RESTful 的风格（一种软件架构风格、设计风格，而**不是**标准，只是提供了一组设计原则和约束条件。它主要用于客户端和服务器交互类的软件。基于这个风格设计的软件可以更简洁，更有层次，更易于实现缓存等机制。）
- 前端 MVC 三层架构， 减少耦合 （Vue 和 AngularJS的选型）
- 前后端采用 JSON 格式交互
- CORS 跨域请求方案
- JWT 认证 （Redis 和 JWT (JSON Web Tokens) 的选型）
- 中间件的应用
- 路由分组
- 后端 MVC 的架构
- MD5 + salt 加密方式
- 权限的控制
- govendor  项目依赖管理



### Todo

- [ ] 博客首页美化、添加标签、添加热门

- [ ] 评论回复管理

- [ ] 文件上传、下载

- [ ] 系统日志

- [ ] 私信管理

- [ ] 权限控制

- [ ] 后台首页统计功能

- [ ] Ngnix 反向代理和负载均衡

  



### Annotation

1. JWT 是一种用于双方之间传递安全信息的简洁的、URL 安全的表述性声明规范。JWT 作为一个开放的标准（RFC 7519），定义了一种简洁的，自包含的方法用于通信双方之间以 Json 对象的形式安全的传递信息。因为数字签名的存在，这些信息是可信的，JWT 可以使用 HMAC 算法或者是 RSA 的公私秘钥对进行签名。简洁(Compact): 可以通过 URL，POST 参数或者在 HTTP header 发送，因为数据量小，传输速度也很快而且 JWT 的 token 串中的负载中包含了所有用户所需要的信息，避免了多次查询数据库。JWT 很好的解决了分布式系统中会话状态保持的问题。 单点登录（SSO）的应用

2.  CSRF（Cross-site request forgery）跨域请求攻击：浏览器对采用了同源策略，所谓同源是指，域名，协议，端口相同。如果是非同源的那么请求就会被浏览器所限制。针对跨域有很多的解决方案，本系统采用 CORS（Cross-Origin Resource Sharing, 跨源资源共享）



### 接口API文档

#### 用户信息

1. 用户登录：POST 

form 表单提交用户名和密码信息

http://localhost:8081/user/login 



2. 用户查询（根据id）： GET

http://localhost:8081/user/query?id=3 



3. 用户信息更新：PUT

http://localhost:8081/user/edit?id=20 

```json
	{
        "username": "chen",
        "nickname": {
            "String": "chen",
            "Valid": true
        },
        "password": "123456",
        "avatar": {
            "String": "http://localhost:8081/1148527767.jpg",
            "Valid": true
        },
        "sign": {
            "String": "生活不止眼前的苟且，还有诗和远方",
            "Valid": true
        }
    }
```



4. 用户删除：DELETE（软删除）

http://localhost:8081/user/delete?id=19 



5. 新增用户：POST

http://localhost:8081/user/add 

```json
{
    "username": "qiang",
    "nickname": {
        "String": "qiang",
        "Valid": true
    },
    "password": "123456",
    "avatar": {
        "String": "http://localhost:8081/1148527767.jpg",
        "Valid": true
    },
    "phone": {
        "String": "1234658",
        "Valid": true
    },
    "email": {
        "String": "15465656565@qq.com",
        "Valid": true
    },
    "birthday": {
        "String": "2010-02-02",
        "Valid": true
    },
    "education": {
        "String": "本科",
        "Valid": true
    }
}
```





### 遇到过的问题

1. gin 框架 axios 的 delete 请求传参问题（参考源码，delete请求接受两个参数，需要对参数的格式设定）

```javascript
var params = {
	'data':{
		'user_name': localStorage.getItem('ms_username'),
		'friend_name': this.friendName
    }
}
```

