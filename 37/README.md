Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。


推荐以参数的方式显示传递Context
以Context作为参数的函数方法，应该把Context作为第一个参数。
给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
Context是线程安全的，可以放心的在多个goroutine中传递
