
# CPU使用率
top命令
```shell
top - 19:40:58 up 2 days,  6:27,  1 user,  load average: 0.00, 0.03, 0.05
Tasks:  85 total,   1 running,  84 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.3 us,  0.5 sy,  0.0 ni, 99.2 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  3880260 total,  2286856 free,   709400 used,   884004 buff/cache
KiB Swap:        0 total,        0 free,        0 used.  2942692 avail Mem

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
 1236 root      10 -10  128004  13400  10524 S   1.7  0.3  58:54.37 AliYunDun
 1168 root      20   0   42300   4320   2832 S   0.3  0.1   2:28.86 AliYunDunUpdate
    1 root      20   0   43540   3876   2596 S   0.0  0.1   0:01.29 systemd
    2 root      20   0       0      0      0 S   0.0  0.0   0:00.00 kthreadd
    4 root       0 -20       0      0      0 S   0.0  0.0   0:00.00 kworker/0:0
```
top 显示系统总体的cpu和内存使用情况，以及各个进程的资源使用情况。
- us user 用户态CPU时间。
- nice ni 代表低优先级用户态 CPU 时间，也就是进程的 nice 值被调整为 1-19 之间时的 CPU 时间。这里注意，nice 可取值范围是 -20 到 19，数值越大，优先级反而越低。
- sys 内核态CPU时间

每个进程 %CPU 表示进程的CPU使用率，它是用户态和内核态CPU的使用率总和。

可以使用 `perf top`

```shell
Samples: 8K of event 'cpu-clock', 4000 Hz, Event count (approx.): 206539206 lost: 0/0 drop: 0/0
Overhead  Shared     Object        Symbol
  17.62%  [kernel]                 [k] _raw_spin_unlock_irqrestore
  13.33%  [kernel]                 [k] finish_task_switch
   4.78%  [kernel]                 [k] tick_nohz_idle_enter
   4.58%  [kernel]                 [k] tick_nohz_idle_exit
   4.15%  [kernel]                 [k] run_timer_softirq
   2.83%  [kernel]                 [k] __do_softirq
```

第一行包含三个数据，分别是采样数（Samples）、事件类型（event）和事件总数量（Event count）。

- 第一列 Overhead ，是该符号的性能事件在所有采样中的比例，用百分比来表示。
- 第二列 Shared ，是该函数或指令所在的动态共享对象（Dynamic Shared Object），如内核、进程名、动态链接库名、内核模块名等。
- 第三列 Object ，是动态共享对象的类型。比如 [.] 表示用户空间的可执行程序、或者动态链接库，而 [k] 则表示内核空间。
- 最后一列 Symbol 是符号名，也就是函数名。当函数名未知时，用十六进制的地址来表示。

