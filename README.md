# 个人博客系统设计与实现
基于Go语言的个人博客系统

------

### 采用前后端分离的架构

前端使用Vue+ElementUI框架及组件

后端使用 Gin+Gorm 框架实现 RESTful 风格的微服务



![后台首页](./public/src/后台首页.jpg)



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

- [x] 博客首页美化、添加标签（分类）、添加归档、添加热门、评论排序

- [x] 评论回复管理

- [x] 文件上传、下载

- [x] 系统日志

- [x] 权限控制

- [x] 关注管理（单向）好友管理（双向）

- [x] 私信管理

- [x] 后台首页统计功能

- [x] 博客前端搜索自动补全（todo: 热门搜索关键字推荐）

- [x] 系统日志导出CSV

- [x] Redis 缓存搜索数据（todo：调优）

- [x] 保存草稿，采用localStorage进行本地存储

- [x] 好友消息推送功能、WebSocket

- [ ] Ngnix 反向代理和负载均衡、分布式部署

  



### Annotation

1. JWT 是一种用于双方之间传递安全信息的简洁的、URL 安全的表述性声明规范。JWT 作为一个开放的标准（RFC 7519），定义了一种简洁的，自包含的方法用于通信双方之间以 Json 对象的形式安全的传递信息。因为数字签名的存在，这些信息是可信的，JWT 可以使用 HMAC 算法或者是 RSA 的公私秘钥对进行签名。简洁(Compact): 可以通过 URL，POST 参数或者在 HTTP header 发送，因为数据量小，传输速度也很快而且 JWT 的 token 串中的负载中包含了所有用户所需要的信息，避免了多次查询数据库。JWT 很好的解决了分布式系统中会话状态保持的问题。 单点登录（SSO）的应用
2.  CSRF（Cross-site request forgery）跨域请求攻击：浏览器对采用了同源策略，所谓同源是指，域名，协议，端口相同。如果是非同源的那么请求就会被浏览器所限制。针对跨域有很多的解决方案，本系统采用 CORS（Cross-Origin Resource Sharing, 跨源资源共享）
3. 搜索自动补全，当用户输入两个及以上字符时，发起sug请求，根据前缀匹配获取所有满足条件的结果。在启动服务的时候，将搜索条件（博客标题、博客分类、标签、作者）生成前缀树，支持中文字符，保存在内存中，然后在之后的每次sug请求，直接从trie中获取满足条件的结果。前缀树的时间复杂度O(K)， K为前缀树的深度。相比简单的字符串匹配算法，线性搜索的复杂度有更好的搜索性能。





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
        "nickname":"chen",
        "password": "123456",
        "avatar": "/1148527767.jpg",
        "sign":"生活不止眼前的苟且，还有诗和远方",
    }
```



4. 用户删除：DELETE（软删除）

http://localhost:8081/user/delete?id=19 



5. 新增用户：POST

http://localhost:8081/user/add 

```json
{
    "username": "qiang",
    "nickname": "qiang",
    "password": "123456",
    "avatar": "http://localhost:8081/1148527767.jpg",
    "phone":"1234658",
    "email": "15465656565@qq.com",
    "birthday": "2010-02-02",
    "education": "本科",
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
} // https://blog.csdn.net/qq383366204/article/details/80268007
```

2. golang结构体json的时间格式化解决方案 <https://www.jianshu.com/p/03003d5cbdbc>

```go

// 实现它的json序列化方法
func (this Log) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Log
	// 定义一个新的结构体
	tmp := struct {
		AliasCom
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		AliasCom:  (AliasCom)(this),
		CreatedAt: this.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: this.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmp)
}
```

3. 系统日志模块，处理器串联，提示解析参数失败，ShouldBindJSON 报EOF错误。原因是c.Request.Body只能读取一次https://blog.csdn.net/impressionw/article/details/84194783， 日志参数首先从body中取，然后从url中的参数中取。
4. 对搜索关键字的前缀树生成，对于只存在英文的情况，实现时都通过固定每个节点的指针数量来降低实现难度，比如使用一个下标与字符一一映射的数组来存储子节点的指针。但是这种请求对中文以及标点符号就无法使用了。这种结构虽然效率较高，但是浪费空间。所以需要实现一种适合动态增加的结构。如果要在一组关键词中，频繁地查询某些关键词，用 Trie 树会非常高效。构建 Trie 树的过程，需要扫描所有的关键词，时间复杂度是 O(n)（n 表示所有关键词的长度和）。但是一旦构建成功之后，后续的查询操作会非常高效。每次查询时，如果要查询的关键词长度是 k，那我们只需要最多比对 k 个节点，就能完成查询操作。跟原本那组关键词的长度和个数没有任何关系。所以说，构建好 Trie 树后，在其中查找关键词的时间复杂度是 O(k)，k 表示要查找的关键词的长度。

```go
// 前缀树
type Trie struct {
	next  map[rune]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie)
	root.isEnd = false
	return root
}

func (this *Trie) Insert(word string) {
	tmp := this
	index := len([]rune(word)) - 1
	for i, v := range []rune(word) {
		if _, exist := tmp.next[v]; !exist {
			node := new(Trie)
			node.next = make(map[rune]*Trie)
			if i == index {
				node.isEnd = true
			}
			tmp.next[v] = node
		}
		tmp = tmp.next[v]
	}
	tmp.isEnd = true
}

func (this *Trie) Search(word string) bool {
	tmp := this
	for _, v := range word {
		if tmp.next[v] == nil {
			return false
		}
		tmp = tmp.next[v]
	}

	return tmp.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	tmp := this
	for _, v := range prefix {
		if tmp.next[v] == nil {
			return false
		}
		tmp = tmp.next[v]
	}
	return true
}

func (this *Trie) GetStartsWith(prefix string) (result []string) {
	tmp := this
	if !tmp.StartsWith(prefix) {
		return
	}
	for _, v := range prefix {
		tmp = tmp.next[v]
	}
	result = tmp.getKey(prefix)
	return
}

func (this *Trie) getKey(prefix string) (result []string) {
	if this.isEnd {
		result = append(result, prefix)
	}
	for key, val := range this.next {
		result = append(result, val.getKey(fmt.Sprintf("%s%s", prefix, string(key)))...)
	}
	return
}
```

