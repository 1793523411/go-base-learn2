## 设置Cookie

net/http中提供了如下SetCookie函数，它在w的头域中添加Set-Cookie头，该HTTP头的值为cookie。

`func SetCookie(w ResponseWriter, cookie *Cookie)`

## 获取Cookie

Request对象拥有两个获取Cookie的方法和一个添加Cookie的方法：

### 获取Cookie的两种方法：

// 解析并返回该请求的Cookie头设置的所有cookie
`func (r *Request) Cookies() []*Cookie`

// 返回请求中名为name的cookie，如果未找到该cookie会返回nil, ErrNoCookie。
`func (r *Request) Cookie(name string) (*Cookie, error)`

添加Cookie的方法：

// AddCookie向请求中添加一个cookie。
`func (r *Request) AddCookie(c *Cookie)`

Cookie弥补了HTTP无状态的不足，让服务器知道来的人是“谁”；但是Cookie以文本的形式保存在本地，自身安全性较差；所以我们就通过Cookie识别不同的用户，对应的在服务端为每个用户保存一个Session数据，该Session数据中能够保存具体的用户数据信息。