
http://www.52im.net/thread-3287-1-1.html

https://www.bilibili.com/video/BV1qJ411w7du?from=search&seid=4318128653382485453
select
	fd_set 使用数组实现  
		1.fd_size 有限制 1024 bitmap
			fd【i】 = accept()
		2.fdset不可重用，新的fd进来，重新创建
		3.用户态和内核态拷贝产生开销
		4.O(n)时间复杂度的轮询
		成功调用返回结果大于 0，出错返回结果为 -1，超时返回结果为 0
		具有超时时间

poll
	基于结构体存储fd
	struct pollfd{
		int fd;
		short events;
		short revents; //可重用
	}
	解决了select的1,2两点缺点

epoll
	解决select的1，2，3，4
	不需要轮询，时间复杂度为O(1)
	epoll_create  创建一个白板 存放fd_events
	epoll_ctl 用于向内核注册新的描述符或者是改变某个文件描述符的状态。已注册的描述符在内核中会被维护在一棵红黑树上
	epoll_wait 通过回调函数内核会将 I/O 准备好的描述符加入到一个链表中管理，进程调用 epoll_wait() 便可以得到事件完成的描述符


    重排 触发事件的放在最前面+ 返回前几个被触发了事件

	两种触发模式：
		LT:水平触发
			当 epoll_wait() 检测到描述符事件到达时，将此事件通知进程，进程可以不立即处理该事件，下次调用 epoll_wait() 会再次通知进程。是默认的一种模式，并且同时支持 Blocking 和 No-Blocking。
		ET:边缘触发
			和 LT 模式不同的是，通知之后进程必须立即处理事件。
			下次再调用 epoll_wait() 时不会再得到事件到达的通知。很大程度上减少了 epoll 事件被重复触发的次数，
			因此效率要比 LT 模式高。只支持 No-Blocking，以避免由于一个文件句柄的阻塞读/阻塞写操作把处理多个文件描述符的任务饿死。
ssd和机械
哪些数据库对ssd优化