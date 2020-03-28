
/opt/java/jdk1.8.0_131/bin/java 

#运行模式
-server 														

#堆区内存可被分配的最大上限 
-Xmx512m  				

#堆区内存初始内存分配的大小 
-Xms512m  														

#新生代（Eden + 2*S）与老年代（不包括永久区）的比值
-XX:NewRatio=4 								

#Eden区和Survivor区的比值
-XX:SurvivorRatio=8 											

#持久代空间大小
-XX:PermSize=48m

#持久代空间大小
-XX:MaxPermSize=64m 											

#每个线程的堆栈大小
-Xss256k		

#JDK5.0以后每个线程堆栈大小为1M,以前每个线程堆栈大小为256K.更具应用的线程所需内存大小进行 调整.在相同物理内存下,减小这个值能生成更多的线程.但是操作系统对一个进程内的线程数还是有限制的,不能无限生成,经验值在3000~5000左右
一般小的应用， 如果栈不是很深， 应该是128k够用的 大的应用建议使用256k。这个选项对性能影响比较大，需要严格的测试。（校长）
和threadstacksize选项解释很类似,官方文档似乎没有解释,在论坛中有这样一句话:"”
-Xss is translated in a VM flag named ThreadStackSize”
一般设置这个值就可以了
-XX:ThreadStackSize=128k 										

-XX:-ReduceInitialCardMarks 

#垃圾回收统计信息
-XX:+PrintGCDetails  											

#垃圾回收统计信息
-XX:+PrintGCTimeStamps 		

#垃圾回收统计信息
-XX:+PrintHeapAtGC 

-Xloggc:/home/workspace/jvm-log/open-api-global-quartz-GC.log 

#关闭System.gc() 这个参数需要严格的测试
-XX:+DisableExplicitGC											

#使用CMS内存收集
-XX:+UseConcMarkSweepGC 

-XX:+CMSClassUnloadingEnabled 

#CMS并发过程运行时的线程数
-XX:ParallelCMSThreads=4 

#CMS降低标记停顿
-XX:+CMSParallelRemarkEnabled 									

#在FULL GC的时候， 对年老代的压缩 CMS是不会移动内存的， 因此， 这个非常容易产生碎片， 导致内存不够用， 因此， 内存的压缩这个时候就会被启用。 增加这个参数是个好习惯。
可能会影响性能,但是可以消除碎片
-XX:+UseCMSCompactAtFullCollection 								

#CMS作为垃圾回收使用50％后开始CMS收集
-XX:CMSInitiatingOccupancyFraction=50 							

#CMS并发收集器不对内存空间进行压缩,整理,所以运行一段时间以后会产生"碎片",使得运行效率降低.此值设置运行多少次GC以后对内存空间进行压缩,整理.
-XX:CMSFullGCsBeforeCompaction=2 								

#这个可以压缩指针，起到节约内存占用的新参数
-XX:+UseCompressedOops 

#当堆内存空间溢出时输出堆的内存快照
-XX:+HeapDumpOnOutOfMemoryError 

-XX:HeapDumpPath=/home/workspace/jvm_dump/open-api-global-quartz-heapDump.hprof 

-jar open-api-global-quartz-exec.jar 

--eureka.server=http://localhost:8761/eureka 
--environment=Staging --dataCenter=Cloud
————————————————
版权声明：本文为CSDN博主「Men-DD」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/menxu_work/article/details/99576964