

## DNS

DNS -> Momain Name System 域名系统 -> 名字和ip的一个映射 -> 方便用户更好的访问互联网
1. DNS -> 流量调度到哪个IDC
2. 智能DNS -> L4负载均衡 (lvs/F5/dpvs) -> L7负载均衡（ng的衍生）
3. -> 业务网关 gateway -> 微服务 -> 数据层 -> 物理层

本地 cat /etc/resolv.conf

ping www.baidu.com


dig shxx.sg +trace

; <<>> DiG 9.10.6 <<>> shxx.sg +trace
;; global options: +cmd
.			52600	IN	NS	g.root-servers.net.
.			52600	IN	NS	h.root-servers.net.
.			52600	IN	NS	i.root-servers.net.
.			52600	IN	NS	j.root-servers.net.
.			52600	IN	NS	k.root-servers.net.
.			52600	IN	NS	l.root-servers.net.
.			52600	IN	NS	m.root-servers.net.
.			52600	IN	NS	a.root-servers.net.
.			52600	IN	NS	b.root-servers.net.
.			52600	IN	NS	c.root-servers.net.
.			52600	IN	NS	d.root-servers.net.
.			52600	IN	NS	e.root-servers.net.
.			52600	IN	NS	f.root-servers.net.
;; Received 460 bytes from 10.12.2.4#53(10.12.2.4) in 33 ms

;; Received 38 bytes from 192.33.4.12#53(c.root-servers.net) in 21 ms