# go-week2
第二周作业

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

####答案
应该wrap这个error。
因为，warp一下的话就会有堆栈信息了。否则不变追踪。
原则：使用第三方库，调用处就warp下。

代码：

	if err == sql.ErrNoRows {
		err = errors.WithStack(err)
	}
	
