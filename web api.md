# Web API的设计与开发规范
## 概念

* API目标群体

```
对外、对内
```

* 设计原则

```
通用并易于理解和使用，考虑终端用户使用体验,减少API请求数
```


* REST API的几个等级：

```
REST LEVEL0：使用HTTP
REST LEVEL1: 引入资源的概念
REST LEVEL2：引入HTTP动词（GET/POST/PUT/DELETE等）
REST LEVEL3：引入HATEOAS概念(Hypermedia as the engine of application state)

LEVEL3中的HATEOAS概念思路是API返回的数据中应该包括下一步行为对应的URI是什么，客户端请求下一步的URI就可以得到进一步的数据。
```

## 设计流程
资源识别、端点（URL）设计、响应数据设计、缓存设计、安全设计、版本控制
市场调研、需求分析、系统分析、系统设计、编码、测试、部署、运维

### 资源识别
* 绘制页面状态图
* 列出功能列表
* 整理功能列表，基于名词、数据为表寻找资源,CRUD


### 端点（URI)设计


* URI短小且容易输入

```
好的例子：http://api.example.com/search
坏的例子：http://api.example.com/service/api/search。

域名已经是api，在URI就没有必要重复一些毫无意义的单词。
```

* URI一眼看懂

```
不要轻易使用缩写：http://api.example.com/sv/u
使用更地道的英语表达：比如搜索接口一般用search而不是find，可以多参照一些国外大厂的API。
获取某个商品信息的URI应该长这样：http://api.example.com/v1/items/12346
```

* URI只有小写字母组成

```
URI应该使用小写，禁止大小写混写。
```

* URI不能反映了服务端的架构
 
```
错误的例子：
http://api.example.com/cgi-bin/get_user.php?user=100
```

* URI规则统一

```
URI中的词汇和结构应该保持统一。

错误的例子：
获取好友信息：http://api.example.com/friends?id=100
发送信息：http://api.example.com/friend/100/messages

一个正确的例子：
获取好友信息：http://api.example.com/friends/100
发送信息：http://api.example.com/friends/100/messages

```

* 资源一般都会有两个端点：集合（表）、单个数据（行）。例如：

```
http://api.example.com/users
http://api.example.com/users/:id
```

* 使用合适的HTTP方法

```
HTTP方法表示"进行怎样的操作"，URI表示"资源"，HTTP和URI一起则表示"对什么资源做什么操作"。

GET：获取资源

获取ID=100的好友信息，
GET http://api.example.com/friends/100

POST：新增资源

添加一位好友，相当于新建一个好友关系：
POST http://api.example.com/friends

PUT：更新已有资源

更新ID=100的好友信息（例如：更新备注信息），
PUT http://api.example.com/friends/100

PATCH：更新部分资源

和PUT类似，只是强调更新资源的部分信息，不常用。

DELETE：删除资源

删除ID=100的好友信息：
DELETE http://api.example.com/friends/100

HEAD：获取资源的元信息

使用_method参数或X-Http_Method_Overide发送http方法
```

* URI里用到的资源采用复数形式

```
因为URI表示资源的集合，所以作者是建议总是使用复数形式。

正确的例子：http://api.example.com/friends/100
错误的例子：http://api.example.com/friend/100

另外，因为REST风格强调URI是资源，所以不应该在URI里出现动词，因为动作是HTTP方法表达的。

错误的例子：http://api.example.com/get_friend?id=100
```

* URI里不能有空格符以及需要编码的字符

```
URL是会被urlencode编码的，所以不要在URI里使用空格（会被编码成+）、UTF-8字符、乱七八糟的符号等。

即不要影响URI的可读性。
```

* 查询参数
http://api.example.com/users?field=a,b,c&a=12&order=x
http://api.example.com/users?a=1&b=c&order=a&aa=desc

```
投影，   select a,b,c  from  xxx where a = 11 order by a desc
过滤
排序
分页
```

* 使用OAuth2.0进行登录

```
Authorization Code
	使用场合：第三方公司使用

Impliit
	使用场合：第三方公司使用

Resoure Owner Password Credentials
	使用场合：公司内部使用客户端使用
	Client_id, Client_secret
		识别应用身份，频率控制，授权范围
		以Base认证的形式经base64编码放入Authorization头部
	端口： 
			/oauth2/token
			获取token：grant_type(password),username,password,scope
			刷新token：grant_type(refresh_token),refresh_token
	
Client Credentials
	使用场合：不需要登录的资源。需传递Client_id, Client_secret
	

token无效返回401

```



### 响应数据设计

* 响应数据格式使用JSON作为默认格式

* 支持通过查询参数来指定数据格式

```
如果服务端支持多种返回数据格式，那么客户端可以指定。

通过get参数：
https://api.example.com/v1/users?format=xml 

通过扩展名：
https://api.example.com/v1/users.xml

通过HTTP头部：
GET /v1/users
Host: api.example.com
Accept: application/xml

建议首先使用HTTP头部，因为更符合HTTP协议规范；其次使用查询参数，避免使用扩展名。
```


* 响应数据的内容可以从客户端指定

```
有时调用方只需要部分信息，比如：用户信息接口只希望获取用户ID，这样可以节约通讯量。

此类接口可以通过类似fields的get参数来指定返回哪些信息：
http://api.example.com/v1/users/12345?fields=name,age，a,c

另外也可以提前准备几种返回值的组合，称为响应群（response group），比如:
http://api.example.com/v1/users/12345?group=basic_info
其中，basic_info表示返回用户的基础信息，例如name和age。
```

* 错误信息

```
错误码放在http header里，而不是放在body里。

出错时，响应信息应该包含2部分：
1，错误码
2，错误原因

建议使用http header返回错误码，每种错误码的含义如下：

1xx：消息
2xx：成功
3xx：重定向
4xx：客户端原因引起的错误
5xx：服务端原因引起的错误

400

错误的例子：
HTTP/1.1 200 OK
{
	"error_code": 500,
	"error_msg": "参数错误",
	"data": {}
}
HTTP返回200，但内容却表达了500失败。

建议用下面这种方式：
HTTTP 响应头错误码：大概错误范围
具体错误码：
1001： 用户名重复
2001：券
3001：活动
HTTP/1.1 500 参数错误
{	
	"errors": [
		{
			"message": "参数fields错误",
			"code": 1001
		},
		{
			"developerMessage": "参数last_id错误",
			"userMessage": "参数last_id错误",
			"code": 2002,
			"info": "http://doc.example.com"
		}
	]
}

详细错误代码
1xxx: 通用错误
2xxx: 券错误
3XXX: 活动错误
...
```

* 出错时不应该返回HTML数据

```
当服务端发生错误时，很多web框架会打印一个html错误信息页面。

对于API来说，当发生错误时也应该返回一个合法的JSON结构，因为客户端假设服务端返回JSON，返回HTML可能导致异常。
```

* 响应数据的结构尽量做到扁平化

```
不要在JSON中增加无意义的多余层级，尽可能扁平化。

一个错误的例子：

{
	"id": 12345,
	"name": "hahaha",
	"profile": {
		"birthday": "0203",
		"gender": "male",
		"language": ["zh", "en"]
	}
}

增加profile并没有带来什么价值，不如扁平化：
{
	"id": 12345,
	"name": "hahaha",
	"birthday": "0203",
	"gender": "male",
	"language": ["zh", "en"]
}

不仅访问起来方便，而且传输的内容也少了。
```
* 响应数据用对象来描述，而不是用数组

```
作者建议JSON返回值总是使用{}作为返回值的最外层，而不要直接返回数组[]。

正确的例子：
{
	"articles": [
		{"id": 1},
		{"id": 2}
		...
	]
}
错误的例子：
[
	{"id": 1},
	{"id", 2}
]

这样做有2个次要的理由：
1，因为从字面看，articles能直接表达数据的含义
2，客户端在处理JSON应答时，可以统一将最外层作为对象去解析，不需要为数组做适配。
```

* 响应数据的名称所选用的单词的意思是否和大部分API相同

```
对非英语母语的人，多模仿大厂使用的常见单词。
```

* 响应数据的名称有没有用尽可能少的单词来描述

```
关于用户注册时间字段，

错误的例子：userRegistraionDataTime
这个单词太长了，很容易打错，也不容易记忆。

正确的例子：registeredAt
```

* 响应数据的名称由多个单词连接而成时，连接方法在整个API里是否一致

```
有几种连接单词的方法：
1，user_id：蛇形法
2，user-id：脊柱法
3，userId：驼峰法

在JSON和Javascript中，都是建议使用驼峰法的，但是保持风格一致是最重要的。
```

* 响应数据的名称不要缩写形式

```
尽量避免奇怪的缩写，比如timezone写成tz。

如果出于数据量大小的考虑而采用缩写，属于特殊情况。
```

* 响应数据的名称的单复数形式要一致

```
只为数组采用复数，比如friends。
其他情况使用单数。

一致的命名风格，API使用者会顺其自然，养成习惯。
```

* 服务器端在维护时返回503状态码

```
当服务器需要停机维护时，按照Google爬虫的建议，应该返回503错误码，并且在header中告知维护的结束时间：

503 Service Temporarily Unavailable
Retry-After: Mon, 2 Dec 2013 03:00:00 GMT

这遵循HTTP1.1规范，客户端需要实现逻辑去识别这个情况，但是至少google爬虫会去理解这些信息。
```

* 返回媒体类型

```
有的HTTP客户端会校验应答中的Content-Type字段，因此服务端如果返回的是JSON，那么就应该返回Content-Type: application/json而不是Content-Type: text/html，这样避免一些严格的客户端出现解析失败。
```

* 支持CORS

```
浏览器有同源策略，禁止跨域Ajax请求。

API可以支持CORS跨域资源共享，比如http://www.example.com请求http://api.example.com的API时应该携带请求的来源：
Origin: http://www.example.com

服务端只允许某些来源的跨域调用，如果Origin合法就在返回中携带：

Access-Control-Allow-Origin: http://www.example.com
或者
Access-Control-Allow-Origin: *

浏览器看到这样的应答，就会把ajax请求正常执行完成，否则会报告ajax调用失败。

对于一些特殊场景，浏览器会采用"事先请求"的方式，先通过一个OPTION方法调用到对应的接口来试探服务端是否返回Access-Control-Allow-Origin，如果没有返回则不发起真正的数据请求。

CORS客户端默认不会传输cookie，我们在发起ajax前设置XHTTPRequest.withCredentials=true，并且服务端必须返回header：Access-Control-Allow-Credentials: true，否则这次ajax调用将报告失败。
```

* 选择一种超媒体类型作为响应(Collection+json、HAL、Siren)

```
集合：
{ "collection" :
  {
    "version" : "1.0",
    "href" : "http://example.org/friends/",
    
    "links" : [
      {"rel" : "feed", "href" : "http://example.org/friends/rss"}
    ],
    
    "items" : [
      {
        "href" : "http://example.org/friends/jdoe",
        "data" : [
          {"name" : "full-name", "value" : "J. Doe", "prompt" : "Full Name"},
          {"name" : "email", "value" : "jdoe@example.org", "prompt" : "Email"}
        ],
        "links" : [
          {"rel" : "blog", "href" : "http://examples.org/blogs/jdoe", "prompt" : "Blog"},
          {"rel" : "avatar", "href" : "http://examples.org/images/jdoe", "prompt" : "Avatar", "render" : "image"}
        ]
      },
      
      {
        "href" : "http://example.org/friends/msmith",
        "data" : [
          {"name" : "full-name", "value" : "M. Smith", "prompt" : "Full Name"},
          {"name" : "email", "value" : "msmith@example.org", "prompt" : "Email"}
        ],
        "links" : [
          {"rel" : "blog", "href" : "http://examples.org/blogs/msmith", "prompt" : "Blog"},
          {"rel" : "avatar", "href" : "http://examples.org/images/msmith", "prompt" : "Avatar", "render" : "image"}
        ]
      },
      
      {
        "href" : "http://example.org/friends/rwilliams",
        "data" : [
          {"name" : "full-name", "value" : "R. Williams", "prompt" : "Full Name"},
          {"name" : "email", "value" : "rwilliams@example.org", "prompt" : "Email"}
        ],
        "links" : [
          {"rel" : "blog", "href" : "http://examples.org/blogs/rwilliams", "prompt" : "Blog"},
          {"rel" : "avatar", "href" : "http://examples.org/images/rwilliams", "prompt" : "Avatar", "render" : "image"}
        ]
      }      
    ],
    
    "queries" : [
      {"rel" : "search", "href" : "http://example.org/friends/search", "prompt" : "Search",
        "data" : [
          {"name" : "search", "value" : ""}
        ]
      }
    ],
    
    "template" : {
      "data" : [
        {"name" : "full-name", "value" : "", "prompt" : "Full Name"},
        {"name" : "email", "value" : "", "prompt" : "Email"},
        {"name" : "blog", "value" : "", "prompt" : "Blog"},
        {"name" : "avatar", "value" : "", "prompt" : "Avatar"}
        
      ]
    }
  } 
}

item: 
{ "collection" :
  {
    "version" : "1.0",
    "href" : "http://example.org/friends/",
    
    "links" : [
      {"rel" : "feed", "href" : "http://example.org/friends/rss"},
      {"rel" : "queries", "href" : "http://example.org/friends/?queries"},
      {"rel" : "template", "href" : "http://example.org/friends/?template"}
    ],
    
    "items" : [
      {
        "href" : "http://example.org/friends/jdoe",
        "data" : [
          {"name" : "full-name", "value" : "J. Doe", "prompt" : "Full Name"},
          {"name" : "email", "value" : "jdoe@example.org", "prompt" : "Email"}
        ],
        "links" : [
          {"rel" : "blog", "href" : "http://examples.org/blogs/jdoe", "prompt" : "Blog"},
          {"rel" : "avatar", "href" : "http://examples.org/images/jdoe", "prompt" : "Avatar", "render" : "image"}
        ]
      }
    ]
  } 
}

Error: 
{ "collection" :
  {
    "version" : "1.0",
    "href" : "http://example.org/friends/",
    
    "error" : {
      "title" : "Server Error",
      "code" : "X1C2",
      "message" : "The server have encountered an error, please wait and try again."
    }
  } 
}
```

## 缓存设计

* 使用Cache-Control、ETag、Last-Modified、Vary等首部以便客户端采用合适的缓存策略

```
缓存模型分2种：
1，过期模型：Expires、Cache-Control
2，验证模型：Last-Modified、ETag

过期模型是指，浏览器在过期之间直接使用本地缓存文件，下面是一个例子：
Expires: Fri, 01 Jan 2016 00:00:00 GMT
Cache-Control: max-age=3600

Cache-Control是HTTP1.1协议出现的，Expires是HTTP1.0，前者优先级更高。
并且HTTP1.1协议也规定，缓存时间不应超过1年，但实际上客户端可能没有遵循这个约束。

验证模型是指，客户端照常发起请求，但在header中携带附加条件，服务器根据附加条件判断若数据没有修改则返回304，客户端直接使用本地缓存即可，否则返回200并携带内容。

下面是个例子，

请求：
GET /v1/users/12345
If-Modified-Since: Tue, 01 Jul 2014 00:00:00 GMT
If-None-Match: "ff39b31e285573ee373af0d492aca581"
应答：
HTTP/1.1 304 Not Modified
Last-Modified: Tue, 01 Jul 2014 00:00:00 GMT
ETag: "ff39b31e285573ee373af0d492aca581"

需要注意ETag分为强验证和弱验证：
强验证是指资源的真实内容完全不能变，弱验证是指逻辑上资源没有改变即可。
```

* 不想缓存的数据有没有添加Cache-Control: no-cache首部信息

```
如果不希望被客户端缓存，可以指定Cache-Control: no-cache。

如果你的API前面存在反向代理缓存，可以额外声明Cache-Control: no-store，这样代理服务器也不会缓存数据了。

客户端可能多次请求同一个API，但是请求的http header不同，导致返回的内容结构不同，比如：

客户端携带 Accept: application/json，则服务端返回的是JSON。
客户端携带 Accept: application/xml，则服务端返回的是XML。

如果反向代理根据URI缓存，则会导致无法根据客户端的要求返回正确格式，此时我们API应该在返回值里携带Vary: Accept，这样缓存服务器会为不同的Accept分别缓存。
```

## 版本控制

* 有没有对API进行版本管理

```
一般API会不断的迭代功能，有时会出现无法向下兼容的情况。

通常老客户端会依旧使用老版本的API，新客户端使用新版本的API，并在合适的时机完全下线掉老版本的API。
```

* API版本的命名有没有遵循语义化版本控制规范

```
作者介绍了语义化版本控制，通常版本号是a.b.c这样的，分别表示主版本号，次版本号，补丁版本号。

1，如果软件API没有变更，只是修复服务端BUG，那么就增加补丁版本号
2，对软件API实施了向下兼容的变更，增加次版本号
3，对软件API实施了不向下兼容的变更时，增加主版本号

```

* 有没有在URI里嵌入主版本编号，并且能够让人一目了然

```
对于Web API来说，作者建议在URI中嵌入主版本号即可，例如：
http://api.example.com/v1/users

整体原则是，尽量保持向下兼容，这样URI不会改变，老用户不需要迁移。

还有一个问题是，如果不带版本号访问应该套用哪个版本的接口？谷歌的做法是使用最老版本，这样就不会影响那些老用户了。
```

* 有没有考虑API终止提供时的相关事项

```
停止API时应该让API返回410错误码，它代表接口不再对外公开。

如果客户端是公司的产品，则可以强制客户端升级，避免停止API导致用户无法使用。
```

* 有没有在文档里明确注明API的最低提供期限

```
错误的例子：

该API 2018-06-01下线，请注意迁移。

正确的例子：

该API将继续维护12个月，请您尽快迁移。

错误的例子把期限说的太死了，而正确的例子则留了余地（比如再维护额外的12个月），使用者的感受会好很多。

```

### 安全设计

* 使用HTTPS来提供API

* 过滤输入，转义输出

* 通过浏览器访问的API有没有使用XSRF token

* API在接收参数时有没有仔细检查非法的参数

* 有没有做到即使请求重复发送，数据也不会多次更新

* 有没有实施访问限速

```
限速是为了保护API服务，避免超过负载。

限速一般是针对每个用户的，限速的单位是多少分钟内最多访问多少次。

从存储上可以采用Redis，key的数量大概是"API的数量 * 用户数量"。

API超出限速应该返回429 Too Many Requests的http code，最好还能给出Retry-After告知多久后可以继续使用。

```

### 编写文档
[api-blueprint](https://www.jianshu.com/p/d39c3553e25a)