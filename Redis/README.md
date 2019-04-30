# NoSQL

​	NoSQL（Not Only SQL）数据库的产生就是为了解决大规模数据集合多重数据种类带来的挑战，尤其是大数据应用难题，包括超大规模数据的存储。

​	1.NoSQL数据库种类繁多，但是一个共同的特点都是去掉关系数据库的关系型特性。数据之间无关系，这样就非常容易扩展。也无形之间，在架构的层面上带来了可扩展的能力。

​	2.NoSQL数据库得益于它的无关系性，数据库的结构简单都具有非常高的读写性能，尤其在大数据量下，同样表现优秀。（一般MySQL使用Query Cache，每次表的更新Cache就失效，是一种大粒度的Cache，在针对web2.0的交互频繁的应用，Cache性能不高。而NoSQL的Cache是记录级的，是一种细粒度的Cache，所以NoSQL在这个层面上来说就要性能高很多了）

​	3.NoSQL无需事先为要存储的数据建立字段，随时可以存储自定义的数据格式。而在关系数据库里，增删字段是一件非常麻烦的事情。

RDBMS

- 高度组织化结构化数据
- 结构化查询语言（SQL）
- 数据和关系都存储在单独的表中。
- 数据操纵语言，数据定义语言
- 严格的一致性
- 基础事务

NoSQL

- 代表着不仅仅是SQL
- 没有声明性查询语言
- 没有预定义的模式
- 键-值对存储，列存储，文档存储，图形数据库
- 最终一致性，而非ACID属性
- 非结构化和不可预知的数据
- CAP定理
- 高性能，高可用性和可伸缩性

## Redis

### 常用命令：

#### 客户端

- **使用Redis客户端**

```shell
127.0.0.1:6379> redis-cli -h host -p port -a password
127.0.0.1:6379> redis-cli --raw  //可解决中文乱码问题
```

- **断开Redis客户端和redis服务端连接**

~~~shell
127.0.0.1:6379> quit
OK
~~~

#### 安全

- **查看密码验证**

~~~shell
127.0.0.1:6379> config get requirepass
1) "requirepass"
2) ""
~~~

- **设置密码验证** [默认情况下 requirepass 参数是空的]

~~~shell
127.0.0.1:6379> config set requirepass "123456"
OK
127.0.0.1:6379> config get requirepass
1) "requirepass"
2) "123456"
127.0.0.1:6379> auth "123456"
OK
~~~

#### 数据备份与恢复

- **创建数据库备份**

~~~shell
127.0.0.1:6379> save 
OK
~~~

- **恢复数据** [需将备份文件 (dump.rdb) 移动到 redis 安装目录并启动服务]

~~~shell
#  redis 目录可以使用 CONFIG 命令
127.0.0.1:6379> config GET dir
1) "dir"
2) "/usr/local/redis/bin"
~~~

- **创建备份文件**

~~~shell
127.0.0.1:6379> bgsave
Background saving started
~~~

#### 发布订阅

- **SUBSCRIBE channel [channel ...]** [订阅给定的一个或多个频道的信息。]

~~~shell
127.0.0.1:6379> subscribe redis1 redis2
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "redis1"
3) (integer) 1
1) "subscribe"
2) "redis2"
3) (integer) 2
~~~

- **UNSUBSCRIBE [channel [channel ...]]** [指退订给定的频道。

~~~shell
# 在不同的客户端中有不同的表现
~~~

- **PSUBSCRIBE pattern [pattern ...]** [订阅一个或多个符合给定模式的频道。

~~~shell
127.0.0.1:6379> psubscribe news.* tweet.*
Reading messages... (press Ctrl-C to quit)
1) "psubscribe"                  # 返回值的类型：显示订阅成功
2) "news.*"                      # 订阅的模式
3) (integer) 1                   # 目前已订阅的模式的数量
1) "psubscribe"
2) "tweet.*"
3) (integer) 2
~~~

- **PUNSUBSCRIBE [pattern [pattern ...]]** [退订所有给定模式的频道。

~~~shell
# 在不同的客户端中有不同的表现
~~~

- **PUBSUB subcommand [argument [argument ...]]** [查看订阅与发布系统状态。

```shell
# 1.
# client-1 订阅 news.it 和 news.sport
# client-2 订阅 news.it 和 news.internet
127.0.0.1:6379> PUBSUB channels news.i*
1) "news.internet"
2) "news.it"
# 2.
# client-1 订阅 news.it 和 news.sport
# client-2 订阅 news.it 和 news.internet
127.0.0.1:6379> pubsub numsub news.it news.internet news.sport news.music
1) "news.it"    # 频道
2) "2"          # 订阅该频道的客户端数量
3) "news.internet"
4) "1"
5) "news.sport"
6) "1"
7) "news.music" # 没有任何订阅者
8) "0"
# 3.
# client-1 订阅 news.* 和 discount.*
# client-2 订阅 tweet.*
127.0.0.1:6379> pubsub numpat
(integer) 3
# client-4 订阅 news.* 
127.0.0.1:6379> pubsub numpat
(integer) 4
```

- **PUBLISH channel message** [将信息发送到指定的频道。]

```shell
# subscribe 订阅模式
127.0.0.1:6379> publish redisChat "Redis is a great caching technique"
(integer) 1
127.0.0.1:6379> publish redisChat "Hello Redis"
(integer) 1
# 订阅者的客户端会显示如下消息
1) "message"
2) "redisChat"
3) "Redis is a great caching technique"
1) "message"
2) "redisChat"
3) "Hello Redis"
# psubscribe 订阅模式
PUBLISH redis1 "This is redis 1"
(integer) 1
127.0.0.1:6379> publish redis2 "This is redis 2"
(integer) 1
# 订阅者的客户端会显示如下消息
1) "pmessage"
2) "redis*"
3) "redis1"
4) "This is redis 1"
1) "pmessage"
2) "redis*"
3) "redis2"
4) "This is redis 2"
```

#### 事务

- **MULTI** [标记一个事务块的开始。]

~~~shell
127.0.0.1:6379> multi            # 标记事务开始
OK
127.0.0.1:6379> incr user_id     # 多条命令按顺序入队
QUEUED
127.0.0.1:6379> incr user_id
QUEUED
127.0.0.1:6379> incr user_id
QUEUED
127.0.0.1:6379> ping
QUEUED
~~~

- **EXEC** [执行所有事务块内的命令。]

~~~shell
127.0.0.1:6379> exec       # 执行
1) (integer) 1
2) (integer) 2
3) (integer) 3
4) PONG
~~~

- **WATCH key [key ...]** [监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。]

~~~shell
# 监视 key ，且事务成功执行
127.0.0.1:6379> watch lock_content lock_times
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379> set lock_content "ABC"
QUEUED
127.0.0.1:6379> incr lock_times
QUEUED
127.0.0.1:6379> exec
1) OK
2) (integer) 1
# 监视 key ，且事务被打断
127.0.0.1:6379> watch lock_content lock_times
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379> set lock_content "ABC"# 这时，另一个客户端修改了lock_timesde的值
QUEUED
127.0.0.1:6379> incr lock_times
QUEUED
127.0.0.1:6379> exec                  # 因为lock_times 被修改，ABC的事务执行失败
(nil)
~~~

- **UNWATCH** [取消 WATCH 命令对所有 key 的监视。]

~~~shell
127.0.0.1:6379> unwatch
OK
~~~

- **DISCARD**  [取消事务，放弃执行事务块内的所有命令。]

~~~shell
127.0.0.1:6379> multi
OK
127.0.0.1:6379> ping
QUEUED
127.0.0.1:6379> set greeting "hello"
QUEUED
127.0.0.1:6379> discard
OK
~~~

#### 主从复制

- **Redis服务初始状态**

~~~shell
# 任何一台redis服务最初状态
127.0.0.1:6379> info replication
# Replication 
role:master
connected_slaves:0
master_replid:e2b725b8b0d44ec6f65577350ca32a94e1e70c72
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
~~~

- **配置文件修改细节**

```bash
daemonize yes
pidfile /var/run/redis6379[此处自定义文件名称].pid
port 6379[此处自定义端口号]
logfile "6379[此处自定义文件名称].log"
dbfilename dump6379[此处自定义文件名称].rdb
```

- **SLAVEOF 主库IP 主库PORT [配从不配主，每次与master断开之后，都需要重新连接，除非配置进redis.conf文件]**

~~~shell
# 6379端口
127.0.0.1:6379> keys *
1) "k1"
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=6382,state=online,offset=682,lag=1
slave1:ip=127.0.0.1,port=6381,state=online,offset=682,lag=1
master_replid:80477b1101688537522a385381e2d6817281514b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:682
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:682
~~~

~~~shell
# 6380端口
127.0.0.1:6380> slaveof 127.0.0.1 6379
OK
127.0.0.1:6380> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:0
master_sync_in_progress:0
slave_repl_offset:336
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:5145a76a87c8ad4a10dc7008172e2111e819d6c1
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:336
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:336
127.0.0.1:6380> get k1
"v1"
127.0.0.1:6380> set k2 v2
(error) READONLY You can’t write against a read only slave.
~~~

~~~shell
# 6381端口
127.0.0.1:6381> slaveof 127.0.0.1 6379
OK
127.0.0.1:6381> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:3
master_sync_in_progress:0
slave_repl_offset:406
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:5145a76a87c8ad4a10dc7008172e2111e819d6c1
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:406
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:15
repl_backlog_histlen:392
127.0.0.1:6381> get k2
"v2"
127.0.0.1:6381> set k2 v2
(error) READONLY You can’t write against a read only slave.
~~~

##### 一主二从

- **主库突然停掉，从库如何处理**

~~~shell
# 6379端口
127.0.0.1:6379> shutdown
not connected> exit
127.0.0.1:6379> set k2 v2
OK
~~~

~~~shell
# 6380端口
127.0.0.1:6380> keys *
1) "k1"
# 6379 shutdown，并执行set k2 v2之后
# 当前配置下 从库会待命，不会上位。
127.0.0.1:6380> get k2
"v2"
~~~

~~~shell
# 6381端口
127.0.0.1:6381> keys *
1) "k1"
# 6379 shutdown，并执行set k2 v2之后
# 当前配置下 从库会待命，不会上位。
127.0.0.1:6381> get k2
"v2"
~~~

- **从库中的某一个突然停掉，主库和其他从库如何处理**

~~~shell
# 6379端口
# 6380 shutdown，并执行set k2 v2之后
127.0.0.1:6379> set k2 v2
OK
~~~

~~~shell
# 6380端口
# 配从不配主，每次与master断开之后，都需要重新连接，除非配置进redis.conf文件
127.0.0.1:6380> shutdown
not connected> exit
127.0.0.1:6380> get k2
(nil)
127.0.0.1:6381> slaveof 127.0.0.1 6380
OK
127.0.0.1:6381> get k2
"v2"
~~~

~~~shell
# 6381端口
# 不受影响
127.0.0.1:6381> get k2
"v2"
~~~

##### 主从接力

~~~shell
# 6379端口
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:1
slave0:ip=127.0.0.1,port=6381,state=online,offset=2335,lag=0
master_replid:80477b1101688537522a385381e2d6817281514b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:2335
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:2335
127.0.0.1:6379> set k7 v7
OK
127.0.0.1:6379> get k7
"v7"
~~~

~~~shell
# 6380端口
127.0.0.1:6380> info replication
# Replication
role:slave													    #依旧是slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:1
master_sync_in_progress:0
slave_repl_offset:2447
slave_priority:100
slave_read_only:1
connected_slaves:1
slave0:ip=127.0.0.1,port=6381,state=online,offset=2447,lag=1	#但是有其自己的slave
master_replid:80477b1101688537522a385381e2d6817281514b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:2447
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1370
repl_backlog_histlen:1078
127.0.0.1:6380> get k7
"v7"
~~~

~~~shell
# 6381端口
127.0.0.1:6381> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6380
master_link_status:up
master_last_io_seconds_ago:5
master_sync_in_progress:0
slave_repl_offset:2489
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:80477b1101688537522a385381e2d6817281514b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:2489
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:2489
127.0.0.1:6381> get k7
"v7"
~~~

##### 反客为主

```shell
# 6379端口
127.0.0.1:6379> shutdown
not connected> exit
127.0.0.1:6379> get k2
(nil)
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:0
master_replid:664dbeb865b022f8f820aff1cd437e70214df3bb
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```

```shell
# 6380端口
# 6379 shutdown之后
127.0.0.1:6380> slaveof no one		#反客为主,使当前数据库停止与其他数据库的同步，转成主库
OK
127.0.0.1:6380> info replication
# Replication
role:master
connected_slaves:1
slave0:ip=127.0.0.1,port=6381,state=online,offset=3787,lag=0
master_replid:00d9675440a3ffb68a439726189051c235742ad5
master_replid2:80477b1101688537522a385381e2d6817281514b
master_repl_offset:3787
second_repl_offset:3788
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1370
repl_backlog_histlen:2418
127.0.0.1:6380> set k2 v2
OK
127.0.0.1:6380> get k2
"v2"
```

```shell
# 6381端口
# 6380反客为主之后，把6381当做主库
127.0.0.1:6381> slaveof 127.0.0.1 6381
OK
127.0.0.1:6381> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6381
master_link_status:up
master_last_io_seconds_ago:3
master_sync_in_progress:0
slave_repl_offset:4301
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:00d9675440a3ffb68a439726189051c235742ad5
master_replid2:80477b1101688537522a385381e2d6817281514b
master_repl_offset:4301
second_repl_offset:3788
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:4301
127.0.0.1:6381> get k2
"v2"
```

##### 哨兵模式

~~~shell
# 启动哨兵模式
$ redis-sentinel sentinel.conf
# sentinel.conf内容
sentinel monitor host6380 127.0.0.1 6380 1
~~~

~~~shell
# 6379端口
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=6380,state=online,offset=42,lag=0
slave1:ip=127.0.0.1,port=6381,state=online,offset=42,lag=1
master_replid:b5d0469346311f5ae635393d696b2dc55231ae66
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:42
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:42
127.0.0.1:6379> keys *
1) "k1"
127.0.0.1:6379> shutdown
not connected> exit
# 重新上线 变成slave
127.0.0.1:6379> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6381
master_link_status:up
master_last_io_seconds_ago:0
master_sync_in_progress:0
slave_repl_offset:40048
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:3419275a1e2a88653941d20aabc462150d300fbd
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:40048
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:39893
repl_backlog_histlen:156
127.0.0.1:6379> get k2
"v2"
~~~

~~~shell
# 6380端口
127.0.0.1:6380> keys *
1) "k1"
# 6379 shutdown之后
127.0.0.1:6380> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6381
master_link_status:up
master_last_io_seconds_ago:2
master_sync_in_progress:0
slave_repl_offset:24512
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:3419275a1e2a88653941d20aabc462150d300fbd
master_replid2:b5d0469346311f5ae635393d696b2dc55231ae66
master_repl_offset:24512
second_repl_offset:22320
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:24512
127.0.0.1:6380> get k2
"v2"
~~~

~~~shell
# 6381端口
127.0.0.1:6381> keys *
1) "k1"
# 6379 shutdown之后，变成master
127.0.0.1:6381> info replication
# Replication
role:master						
connected_slaves:1
slave0:ip=127.0.0.1,port=6380,state=online,offset=24911,lag=0
master_replid:3419275a1e2a88653941d20aabc462150d300fbd
master_replid2:b5d0469346311f5ae635393d696b2dc55231ae66
master_repl_offset:24911
second_repl_offset:22320
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:24911
127.0.0.1:6381> set k2 v2
OK
127.0.0.1:6381> get k2
"v2"
# 6379重新上线后
127.0.0.1:6381> info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=6379,state=online,offset=46572,lag=0
slave1:ip=127.0.0.1,port=6380,state=online,offset=46572,lag=1
master_replid:3419275a1e2a88653941d20aabc462150d300fbd
master_replid2:b5d0469346311f5ae635393d696b2dc55231ae66
master_repl_offset:46572
second_repl_offset:22320
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:46572
~~~

### 数据类型：

#### 键(key)

- **DEL key** [该命令用于在 key 存在时删除 key。]

~~~shell
127.0.0.1:6379> del k1
(integer) 1
~~~

- **DUMP key** [序列化给定 key ，并返回被序列化的值。]

```shell
127.0.0.1:6379> dump k
"\x00\x01v\b\x00(%\"\r\xb0!\x00\xf6"
```

- **EXISTS key** [检查给定 key 是否存在。]

```shell
127.0.0.1:6379> exists k
(integer) 1
```

- **EXPIRE key seconds** [为给定 key 设置过期时间，以秒计。]

```shell
127.0.0.1:6379> expire k 5 #剩5秒
(integer) 1
127.0.0.1:6379> get k
(nil)
```

- **EXPIREAT key timestamp** [EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置过期时间。 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。]

```shell
127.0.0.1:6379>
```

- **PEXPIRE key milliseconds** [设置 key 的过期时间以毫秒计。]

```shell
127.0.0.1:6379> pexpire k 5000 #剩5000毫秒
(integer) 1
127.0.0.1:6379> get k
(nil)
```

- **PEXPIREAT key milliseconds-timestamp** [设置 key 过期时间的时间戳(unix timestamp) 以毫秒计。]

```shell
127.0.0.1:6379>
```

- **KEYS pattern** [查找所有符合给定模式( pattern)的 key 。]

```shell
127.0.0.1:6379> keys *
1) "k2"
2) "k1"
3) "k"
```

- **MOVE key db** [将当前数据库的 key 移动到给定的数据库 db 当中。]

```shell
127.0.0.1:6379> move k 1
(integer) 1
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> get k
"v"
```

- **PERSIST key** [移除 key 的过期时间，key 将持久保持。]

```shell
127.0.0.1:6379> expire k 60
(integer) 1
127.0.0.1:6379> persist k
(integer) 1
```

- **PTTL key** [以毫秒为单位,返回给定 key 的剩余的过期时间。]

```shell
127.0.0.1:6379> expire k 60
(integer) 1
127.0.0.1:6379> pttl k
(integer) 57280 #剩57280毫秒
```

- **TTL key** [以秒为单位，返回给定 key 的剩余的过期时间(TTL, time to live)。]

```shell
127.0.0.1:6379> expire k 60
(integer) 1
127.0.0.1:6379> ttl k
(integer) 58 #剩58秒
```

- **RANDOMKEY** [从当前数据库中随机返回一个 key 。]

```shell
127.0.0.1:6379> randomkey
"k"
```

* **RENAME key newkey** [修改 key 的名称，若newkey存在，则覆盖。]

```shell
127.0.0.1:6379> keys *
1) "k"
127.0.0.1:6379> rename k kk
OK
127.0.0.1:6379> keys *
1) "kk"
```

* **RENAMENX key newkey** [仅当 newkey 不存在时，将 key 改名为 newkey 。]

```shell
127.0.0.1:6379> keys *
1) "kk"
2) "k"
127.0.0.1:6379> renamenx k kk
(integer) 0
127.0.0.1:6379> renamenx k kkk
(integer) 1
127.0.0.1:6379> keys *
1) "kk"
2) "kkk"
```

* **TYPE key** [返回 key 所储存的值的类型。]

```shell
127.0.0.1:6379> type k
string
```

#### 字符串(String)

- **SET key value** [设置指定 key 的值。]

~~~shell
# 对不存在的键进行设置，创造新的key-value键值对
# 对存在的键进行设置，则覆盖原有value值
127.0.0.1:6379> set k v
OK
~~~

- **GET key** [获取指定 key 的值。]

~~~shell
# 对不存在的 key 或字符串类型 key 进行 GET
127.0.0.1:6379> get k
(nil)
127.0.0.1:6379> set k v
OK
127.0.0.1:6379> get k 
"v"
# 对不是字符串类型的 key 进行 GET
127.0.0.1:6379> LPUSH db redis mongodb mysql
(integer) 3
redis> get db
(error) ERR Operation against a key holding the wrong kind of value
~~~

- **GETRANGE key start end** [返回 key 中字符串值的子字符。]

~~~shell
127.0.0.1:6379> set k abcdefghi
OK
127.0.0.1:6379> getrange k 0 5
"abcdef"
127.0.0.1:6379> getrange k 0 -1
"abcdefghi"
~~~

- **GETSET key value** [将给定 key 的值设为 value ，并返回 key 的旧值(old value)。]

~~~shell
# 没有旧值，返回 nil
127.0.0.1:6379> getset k 0123456789
(nil)
127.0.0.1:6379> get k
"0123456789"
# 返回旧值 abcdefghi
127.0.0.1:6379> getset k 0123456789
"abcdefghi"
127.0.0.1:6379> get k
"0123456789"
~~~

- **GETBIT key offset** [对 key 所储存的字符串值，获取指定偏移量上的位(bit)。]

~~~shell
# getbit k 1000 
~~~

- **MGET key1 [key2..]** [获取所有(一个或多个)给定 key 的值。] 

~~~shell
127.0.0.1:6379> mget k1 k2 k3
1) "v1"
2) "v2"
3) "v3"
~~~

- **SETBIT key offset value** [对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。]

~~~shell
# setbit k 1000 1
~~~

- **SETEX key seconds value** [将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。]

~~~shell
127.0.0.1:6379> setex k 60 v
OK
127.0.0.1:6379> ttl k
(integer) 60 #剩60秒
127.0.0.1:6379> get k
"v"
~~~

- **SETNX key value** [只有在 key 不存在时设置 key 的值。]

~~~shell
127.0.0.1:6379> setnx k v
(integer) 0
127.0.0.1:6379> del k
(integer) 1
127.0.0.1:6379> setnx k v
(integer) 1
127.0.0.1:6379> get k
"v"
~~~

- **SETRANGE key offset value** [用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始。]

~~~shell
127.0.0.1:6379> get k
"Hello World"
127.0.0.1:6379> setrange k 6 "Redis"
(integer) 11
127.0.0.1:6379> get k
"Hello Redis"
~~~

- **STRLEN key** [返回 key 所储存的字符串值的长度。] 

~~~shell
127.0.0.1:6379> strlen k
(integer) 9
# 不存在的 key 长度为 0
127.0.0.1:6379> strlen nonexisting
(integer) 0
~~~

- **MSET key value [key value ...]** [同时设置一个或多个 key-value 对。]

~~~shell
127.0.0.1:6379> mset rmdbs "MySQL" nosql "MongoDB" key-value-store "redis" 
OK
127.0.0.1:6379> keys *
1) "MySQL"
2) "MongoDB"
3) "redis"
~~~

- **MSETNX key value [key value ...]** [同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。]

~~~shell
# 对不存在的 key 进行 MSETNX
127.0.0.1:6379> msetnx rmdbs "MySQL" nosql "MongoDB" key-value-store "redis"
(integer) 1
127.0.0.1:6379> mget rmdbs nosql key-value-store
1) "MySQL"
2) "MongoDB"
3) "redis"
# MSET 的给定 key 当中有已存在的 key
redis> msetnx rmdbs "Sqlite" language "python"  # rmdbs 键已经存在，操作失败
(integer) 0
redis> EXISTS language                          # 因为 MSET 是原子性操作，language 没有被设置
(integer) 0
redis> GET rmdbs                                # rmdbs 也没有被修改
"MySQL"
~~~

- **PSETEX key milliseconds value** [这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。]

~~~shell
127.0.0.1:6379> psetex k 60000 v
OK
127.0.0.1:6379> pttl k
(integer) 58784 #剩58784毫秒
127.0.0.1:6379> get k
"v"
~~~

- **INCR key** [将 key 中储存的数字值增一。]

~~~shell
127.0.0.1:6379> set k 99
OK
127.0.0.1:6379> incr k
(integer) 100
~~~

- **INCRBY key increment** [将 key 所储存的值加上给定的增量值（increment） 。]

~~~shell
# key 存在且是数字值
127.0.0.1:6379> set k 50
OK
127.0.0.1:6379> incrby k 20
(integer) 70
127.0.0.1:6379> GET k
"70"
# key 不存在时
127.0.0.1:6379> exists k
(integer) 0
127.0.0.1:6379> incrby k 30
(integer) 30
127.0.0.1:6379> get k
"30"
# key 不是数字值时
127.0.0.1:6379> set k "hello world..."
OK
127.0.0.1:6379> incrby k 200
(error) ERR value is not an integer or out of range
~~~

- **INCRBYFLOAT key increment** [将 key 所储存的值加上给定的浮点增量值（increment） 。]

~~~shell
# 值和增量都不是指数符号
127.0.0.1:6379> set k 9.8
OK
127.0.0.1:6379> incrbyfloat k 0.1
"9.9"
# 值和增量都是指数符号，结果格式会被改成非指数符号
127.0.0.1:6379> set k 314e-2
OK
127.0.0.1:6379> incrbyfloat k 1
"1.14"
# 对整数类型执行
127.0.0.1:6379> set k 3
OK
127.0.0.1:6379> incrbyfloat k 1.1
"4.1"
# 将无用的 0 忽略掉，有需要的话，将浮点变为整数
127.0.0.1:6379> set k 3.0
OK
127.0.0.1:6379> incrbyfloat k 1.000000000000000000000    
"4"
~~~

- **DECR key** [将 key 中储存的数字值减一。]

~~~shell
127.0.0.1:6379> set k 100
OK
127.0.0.1:6379> decr k
(integer) 99
~~~

- **DECRBY key decrementkey** [所储存的值减去给定的减量值（decrement） 。]

~~~shell
# key 存在且是数字值
127.0.0.1:6379> set k 100
OK
127.0.0.1:6379> decrby k 20
(integer) 80
# key 不存在时
127.0.0.1:6379> exists k
(integer) 0
127.0.0.1:6379> decr k 10
(integer) -10
# key 不是数字值时
127.0.0.1:6379> set k "hello world..."
OK
127.0.0.1:6379> decrby k 10
(error) ERR value is not an integer or out of range
~~~

- **APPEND key value** [如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。]

~~~shell
# 对不存在的 key 执行 APPEND
127.0.0.1:6379> exists k             
(integer) 0
127.0.0.1:6379> append k "hello world"       # 等同于 SET k "hello world"
(integer) 11                                 # 字符长度
# 对已存在的字符串进行 APPEND
127.0.0.1:6379> append k "hello world"       # 长度从 5 个字符增加到 12 个字符
(integer) 25
127.0.0.1:6379> get k
"hello redis...hello world"
~~~

#### 哈希(Hash)

- **HDEL key field1 [field2]** [删除一个或多个哈希表字段。]

~~~shell
127.0.0.1:6379> hset test k v
(integer) 1
127.0.0.1:6379> hdel test k
(integer) 1
127.0.0.1:6379> hdel test kk
(integer) 0
~~~

- **HEXISTS key field** [查看哈希表 key 中，指定的字段是否存在。]

~~~shell
127.0.0.1:6379> hset test k v
(integer) 1
127.0.0.1:6379> hexists test k
(integer) 1
127.0.0.1:6379> hexists test kk
(integer) 0
~~~

- **HGET key field** [获取存储在哈希表中指定字段的值。]

~~~shell
# 字段存在
127.0.0.1:6379> hset test k v
(integer) 1
127.0.0.1:6379> hget test k
"v"
# 字段不存在
127.0.0.1:6379> hget test kk
(nil)
~~~

- **HGETALL key** [获取在哈希表中指定 key 的所有字段和值。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hset test k2 v2
(integer) 1
127.0.0.1:6379> hgetall test
1) "k1"
2) "v1"
3) "k2"
4) "v2"
~~~

- **HINCRBY key field increment** [为哈希表 key 中的指定字段的整数值加上增量 increment 。]

~~~shell
127.0.0.1:6379> hset test k 5
(integer) 1
127.0.0.1:6379> hincrby test k 1
(integer) 6
127.0.0.1:6379> hincrby test k -1
(integer) 5
127.0.0.1:6379> hincrby test k -10
(integer) -5
~~~

- **HINCRBYFLOAT key field increment** [为哈希表 key 中的指定字段的浮点数值加上增量 increment 。]

~~~shell
127.0.0.1:6379> hset test k 10.50
(integer) 1
127.0.0.1:6379> hincrbyfloat test k 0.1
"10.6"
127.0.0.1:6379> hincrbyfloat test k -5
"5.6"
127.0.0.1:6379> hset test k 5.0e3
(integer) 0
127.0.0.1:6379> hincrbyfloat test k 2.0e2
"5200"
~~~

- **HKEYS key** [获取所有哈希表中的字段。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hset test k2 v2
(integer) 1
127.0.0.1:6379> hkeys test
1) "k1"
2) "k2"
~~~

- **HLEN key** [获取哈希表中字段的数量。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hset test k2 v2
(integer) 1
127.0.0.1:6379> hlen test
(integer) 2
~~~

- **HMGET key field1 [field2]** [获取所有给定字段的值。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hset test k2 v2
(integer) 1
127.0.0.1:6379> hmget test k1 k2 nofield
1) "v1"
2) "v2"
3) (nil)
~~~

- **HMSET key field1 value1 [field2 value2]** [同时将多个 field-value (域-值)对设置到哈希表 key 中。]

~~~shell
127.0.0.1:6379> hmset test k1 v1 k2 v2
(integer) 1
127.0.0.1:6379> hget test k1
"v1"
127.0.0.1:6379> hget test k2
"v2"
~~~

- **HSET key field value** [将哈希表 key 中的字段 field 的值设为 value 。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hget test k1
"v1"
127.0.0.1:6379> hset test k2 v2       # 设置一个新域
(integer) 1
127.0.0.1:6379> hset test k2 vv2 	  # 覆盖一个旧域
(integer) 0
~~~

- **HSETNX key field value** [只有在字段 field 不存在时，设置哈希表字段的值。]

~~~shell
127.0.0.1:6379> hsetnx test k1 v1
(integer) 1
127.0.0.1:6379> hsetnx test k1 vv1
(integer) 0
127.0.0.1:6379> hget test k1
"v1"
127.0.0.1:6379> hsetnx test k1 v1
(integer) 1
127.0.0.1:6379> hsetnx test k1 v1       # 操作无效， k1 已存在
(integer) 0
~~~

- **HVALS key** [获取哈希表中所有值。]

~~~shell
127.0.0.1:6379> hset test k1 v1
(integer) 1
127.0.0.1:6379> hset test k2 v2
(integer) 1
redis 127.0.0.1:6379> hvals test
1) "v1"
2) "v2"
# 空哈希表/不存在的key
127.0.0.1:6379> EXISTS not_exists
(integer) 0
127.0.0.1:6379> HVALS not_exists
(empty list or set)
~~~

- **HSCAN key cursor MATCH pattern** [迭代哈希表中的键值对。]

~~~

~~~

####列表(List)

- **LINDEX key index** [通过索引获取列表中的元素。]

~~~shell
127.0.0.1:6379> LPUSH test v1
(integer) 1
127.0.0.1:6379> LPUSH test v2
(integer) 2
127.0.0.1:6379> LPUSH test v3
(integer) 3
127.0.0.1:6379> lindex test 1
"v2"
~~~

- **LINSERT key [BEFORE|AFTER] pivot value** [在列表的元素前或者后插入元素。]

~~~shell
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lpush test v3
(integer) 3
127.0.0.1:6379> lrange test 0 3        
1) "v3"
2) "v2"
3) "v1"                 
127.0.0.1:6379> linsert test before v1 vv
(integer) 4
127.0.0.1:6379> linsert test after v1 vvv
(integer) 5
127.0.0.1:6379> lrange test 0 4
1) "v3"
2) "v2"
3) "vv"
4) "v1"
5) "vvvv"
~~~

- **LLEN key** [获取列表长度。]

~~~shell
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lpush test v3
(integer) 3
127.0.0.1:6379> llen test
(integer) 3
~~~

- **LPUSH key value1 [value2]** [在列表头部添加一个或多个值。]

```shell
# 从头向尾推入
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v2"
2) "v1"
```

- **LPUSHX key value** [将一个值插入到已存在的**列表头部**，列表不存在时操作无效。]

~~~shell
# 从头向尾推入
127.0.0.1:6379> lpushx test v1
(integer) 0
127.0.0.1:6379> lpush test v2
(integer) 1
127.0.0.1:6379> lpushx test v1
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v2"
~~~

- **RPUSH key value1 [value2]** [在列表中添加一个或多个值。]

~~~shell
# 从尾向头推入
127.0.0.1:6379> rpush test v1
(integer) 1
127.0.0.1:6379> rpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v2"
~~~

- **RPUSHX key value** [将一个值插入到已存在的**列表**，列表不存在时操作无效。]

~~~shell
# 从尾向头推入
127.0.0.1:6379> rpushx test v1
(integer) 0
127.0.0.1:6379> rpush test v2
(integer) 1
127.0.0.1:6379> rpushx test v1
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v2"
2) "v1"
~~~

- **LPOP key** [移出并获取列表的**第一个元素**。]

~~~shell
# lpush 
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v2"
2) "v1"
127.0.0.1:6379> lpop test
"v2"
127.0.0.1:6379> lpop test
"v1"
127.0.0.1:6379> lpop test
(nil)
# rpush 
127.0.0.1:6379> rpush test v1
(integer) 1
127.0.0.1:6379> rpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v2"
127.0.0.1:6379> lpop test
"v1"
127.0.0.1:6379> lpop test
"v2"
127.0.0.1:6379> lpop test
(nil)
~~~

- **RPOP key** [移除并获取列表**最后一个元素**。]

~~~shell
# lpush 
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v2"
2) "v1"
127.0.0.1:6379> rpop test
"v1"
127.0.0.1:6379> rpop test
"v2"
127.0.0.1:6379> rpop test
(nil)
# rpush 
127.0.0.1:6379> rpush test v1
(integer) 1
127.0.0.1:6379> rpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v2"
127.0.0.1:6379> rpop test
"v2"
127.0.0.1:6379> rpop test
"v1"
127.0.0.1:6379> rpop test
(nil)
~~~

- **BLPOP key1 [key2] timeout** [移出并获取列表的**第一个元素**， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。]

```shell
127.0.0.1:6379> lpush test v1
(integer) 1
# 列表不为空，返回一个含有两个元素的列表，第一个元素是被弹出元素所属的 key ，第二个元素是被弹出元素的值。
127.0.0.1:6379> blpop test 10
1) "test"
2) "v1"
# 列表为空，操作会被阻塞，等待10秒后会返回 nil 。
127.0.0.1:6379> blpop test 10
(nil)
(10.10s)
```

- **BRPOP key1 [key2] timeout** [移出并获取列表的**最后一个元素**， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。]

```shell
127.0.0.1:6379> lpush test v1
(integer) 1
# 列表不为空，返回一个含有两个元素的列表，第一个元素是被弹出元素所属的 key ，第二个元素是被弹出元素的值。
127.0.0.1:6379> brpop test 10
1) "test"
2) "v1"
# 列表为空，操作会被阻塞，等待10秒后会返回 nil 。
127.0.0.1:6379> brpop test 10
(nil)
(10.10s)
```

- **BRPOPLPUSH source destination timeout** [从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。]

```shell
127.0.0.1:6379> lpush test1 v1
(integer) 1
127.0.0.1:6379> lpush test1 v2
(integer) 2
127.0.0.1:6379> lrange test1 0 -1
1) "v2"
2) "v1"
127.0.0.1:6379> brpoplpush test1 test2 10
"v1"
127.0.0.1:6379> brpoplpush test1 test2 10
"v2"
127.0.0.1:6379> brpoplpush test1 test2 10
(nil)
(10.07s)
127.0.0.1:6379> lrange test2 0 -1
1) "v2"
2) "v1"
```

- **RPOPLPUSH source destination** [移除列表的最后一个元素，并将该元素添加到另一个列表并返回。]

~~~shell
127.0.0.1:6379> lpush test1 v1
(integer) 1
127.0.0.1:6379> lpush test1 v2
(integer) 2
127.0.0.1:6379> lrange test1 0 -1
1) "v2"
2) "v1"
127.0.0.1:6379> rpoplpush test1 test2 
"v1"
127.0.0.1:6379> rpoplpush test1 test2 
"v2"
127.0.0.1:6379> rpoplpush test1 test2 
(nil)
127.0.0.1:6379> lrange test2 0 -1
1) "v2"
2) "v1"
~~~

- **LRANGE key start stop** [获取列表指定范围内的元素。]

~~~shell
# lpush 
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v2"
2) "v1"
# rpush 
127.0.0.1:6379> rpush test v1
(integer) 1
127.0.0.1:6379> rpush test v2
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v2"
~~~

- **LREM key count value** [移除列表元素。]

~~~shell
# count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
# count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
# count = 0 : 移除表中所有与 VALUE 相等的值。
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v1
(integer) 2
127.0.0.1:6379> lpush test v2
(integer) 3
127.0.0.1:6379> lpush test v3
(integer) 4
127.0.0.1:6379> lpush test v1
(integer) 5
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v3"
3) "v2"
4) "v1"
5) "v1"
127.0.0.1:6379> lrem test -2 v1        # count = -2 
(integer) 2
127.0.0.1:6379> lrange test 0 -1
1) "v1"
2) "v3"
3) "v2"
~~~

- **LSET key index value** [通过索引设置列表元素的值。]

~~~shell
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lpush test v3
(integer) 3
127.0.0.1:6379> lpush test v4
(integer) 4
127.0.0.1:6379> lset test 2 "vv"
OK
127.0.0.1:6379> lrange test 0 -1
1) "v4"
2) "v3"
3) "vv"
4) "v1"
~~~

- **LTRIM key start stop** [对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。]

~~~shell
127.0.0.1:6379> lpush test v1
(integer) 1
127.0.0.1:6379> lpush test v2
(integer) 2
127.0.0.1:6379> lpush test v3
(integer) 3
127.0.0.1:6379> lpush test v4
(integer) 4
127.0.0.1:6379> ltrim test 1 -1
OK
127.0.0.1:6379> lrange test 0 -1
1) "v3"
2) "v2"
3) "v1"
~~~

#### 集合(Set)

- **SADD key member1 [member2]** [向集合添加一个或多个成员。]

~~~shell
127.0.0.1:6379> sadd k v1 
(integer) 1
127.0.0.1:6379> sadd k v2
(integer) 1
127.0.0.1:6379> sadd k v3
(integer) 1
127.0.0.1:6379> smembers k
1) "v1"
2) "v3"
3) "v2"
~~~

- **SCARD key** [获取集合的成员数。]

~~~shell
27.0.0.1:6379> flushall
OK
127.0.0.1:6379> sadd k v1 
(integer) 1
127.0.0.1:6379> sadd k v2
(integer) 1
127.0.0.1:6379> sadd k v1
(integer) 0
127.0.0.1:6379> scard k
(integer) 2
~~~

- **SDIFF key1 [key2]** [返回给定所有集合的差集。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> sdiff k1 k2
1) "a"
2) "b"
~~~

- **SDIFFSTORE destination key1 [key2]** [返回给定所有集合的差集并存储在 destination 中。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> sdiffstore k k1 k2
(integer) 2
127.0.0.1:6379> smembers k
1) "a"
2) "b"
~~~

- **SINTER key1 [key2]** [返回给定所有集合的交集。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> sinter k1 k2
1) "c"
~~~

- **SINTERSTORE destination key1 [key2]** [给定所有集合的交集并存储在 destination 中。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
7.0.0.1:6379> sinterstore k k1 k2
(integer) 1
127.0.0.1:6379> smembers k
1) "c"
~~~

- **SUNION key1 [key2]** [返回所有给定集合的并集。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> sunion k1 k2
1) "a"
2) "b"
3) "c"
4) "d"
5) "e"
~~~

- **SUNIONSTORE destination key1 [key2]** [返回给定所有集合的并集存储在 destination 集合中。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> sunionstore k k1 k2
(integer) 5
127.0.0.1:6379> smembers k
1) "a"
2) "b"
3) "c"
4) "d"
5) "e"
~~~

- **SISMEMBER key member** [判断 member 元素是否是集合 key 的成员。]

~~~shell
127.0.0.1:6379> sadd k v
(integer) 1
127.0.0.1:6379> sismember k v
(integer) 1
127.0.0.1:6379> sismember k vv
(integer) 0
~~~

- **SMEMBERS key** [返回集合中的所有成员。]

~~~shell
127.0.0.1:6379> sadd k v1 
(integer) 1
127.0.0.1:6379> sadd k v2
(integer) 1
127.0.0.1:6379> sadd k v3
(integer) 1
127.0.0.1:6379> smembers k
1) "v1"
2) "v3"
3) "v2"
~~~

- **SRANDMEMBER key [count]** [返回集合中一个或多个随机数。]

~~~shell
127.0.0.1:6379> sadd k v1
(integer) 1
127.0.0.1:6379> sadd k v2
(integer) 1
127.0.0.1:6379> sadd k v3
(integer) 1
127.0.0.1:6379> srandmember k 
"v2"
127.0.0.1:6379> srandmember k 2
1) "v1"
2) "v2"
~~~

- **SMOVE source destination member** [将 member 元素从 source 集合移动到 destination 集合。]

~~~shell
127.0.0.1:6379> sadd k1 a
(integer) 1
127.0.0.1:6379> sadd k1 b
(integer) 1
127.0.0.1:6379> sadd k1 c
(integer) 1
127.0.0.1:6379> sadd k2 c
(integer) 1
127.0.0.1:6379> sadd k2 d
(integer) 1
127.0.0.1:6379> sadd k2 e
(integer) 1
127.0.0.1:6379> smove k1 k2 b
(integer) 1
127.0.0.1:6379> smembers k1
1) "c"
2) "a"
127.0.0.1:6379> smembers k2
1) "d"
2) "c"
3) "b"
4) "e"
~~~

- **SPOP key [count]** [移除并返回集合中的一个或多个随机元素。]

~~~shell
127.0.0.1:6379> sadd k a
(integer) 1
127.0.0.1:6379> sadd k b
(integer) 1
127.0.0.1:6379> sadd k c
(integer) 1
127.0.0.1:6379> spop k
"a"
127.0.0.1:6379> smembers k
1) "c"
2) "b"
127.0.0.1:6379> sadd k d
(integer) 1
127.0.0.1:6379> sadd k e
(integer) 1
127.0.0.1:6379> spop k 3
1) "d"
2) "e"
3) "b"
127.0.0.1:6379> smembers k
1) "c"
~~~

- **SREM key member1 [member2]** [移除集合中一个或多个成员。]

~~~shell
127.0.0.1:6379> sadd k a
(integer) 1
127.0.0.1:6379> sadd k b
(integer) 1
127.0.0.1:6379> sadd k c
(integer) 1
127.0.0.1:6379> srem k a
(integer) 1
127.0.0.1:6379> srem k d
(integer) 0
127.0.0.1:6379> smembers k
1) "c"
2) "b"
~~~

- **SSCAN key cursor [MATCH pattern] [COUNT count]** [迭代集合中的元素。]

~~~shell
127.0.0.1:6379> sadd k hello
(integer) 1
127.0.0.1:6379> sadd k hi
(integer) 1
127.0.0.1:6379> sadd k honour
(integer) 1
127.0.0.1:6379> sadd k world
(integer) 1
127.0.0.1:6379> sadd k redis
(integer) 1
# 迭代集合k，匹配以k开头元素，每次匹配两个。
127.0.0.1:6379> sscan k 0 match h* count 3
1) "3"
2) 1) "hi"
   2) "hello"
127.0.0.1:6379> sscan k 3 match h* count 3
1) "0"
2) 1) "honour"
~~~

#### 有序集合(sorted set)

- **ZADD key score1 member1 [score2 member2]** [向有序集合添加一个或多个成员，或者更新已存在成员的分数。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 3 b
(integer) 1
127.0.0.1:6379> zadd k 2 c
(integer) 1
127.0.0.1:6379> zadd k 1 d
(integer) 1
127.0.0.1:6379> zrange k 0 -1 withscores
1) "d"
2) "1"
3) "c"
4) "2"
5) "b"
6) "3"
7) "a"
8) "4"
127.0.0.1:6379> zrange k 0 -1
1) "d"
2) "c"
3) "b"
4) "a"
~~~

- **ZCARD key** [获取有序集合的成员数。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 3 b
(integer) 1
127.0.0.1:6379> zadd k 2 c
(integer) 1
127.0.0.1:6379> zadd k 1 d
(integer) 1
127.0.0.1:6379> zcard k
(integer) 4
~~~

- **ZCOUNT key min max** [计算在有序集合中指定区间分数的成员数。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 4 b
(integer) 1
127.0.0.1:6379> zadd k 3 c
(integer) 1
127.0.0.1:6379> zadd k 2 d
(integer) 1
127.0.0.1:6379> zcount k 2 3
(integer) 2
127.0.0.1:6379> zcount k 2 4
(integer) 4
~~~

- **ZINCRBY key increment member** [有序集合中对指定成员的分数加上增量 increment。]

~~~shell
127.0.0.1:6379> zincrby k 5 d
"7"
127.0.0.1:6379> zrange k 0 -1 withscores
1) "c"
2) "3"
3) "a"
4) "4"
5) "b"
6) "4"
7) "d"
8) "7"
~~~

- **ZLEXCOUNT key min max** [在有序集合中计算指定字典区间内成员数量。]

~~~shell
127.0.0.1:6379> zadd k 0 a 0 b 0 c 0 d 0 e
(integer) 5
127.0.0.1:6379> zadd k 0 f 0 g
(integer) 2
127.0.0.1:6379> zlexcount k - +
(integer) 7
127.0.0.1:6379> zlexcount k [b [f
(integer) 5
~~~

- **ZRANGE key start stop [WITHSCORES]** [通过索引区间返回有序集合成指定区间内的成员。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 3 b
(integer) 1
127.0.0.1:6379> zadd k 2 c
(integer) 1
127.0.0.1:6379> zadd k 1 d
(integer) 1
127.0.0.1:6379> zrange k 0 -1 withscores
1) "d"
2) "1"
3) "c"
4) "2"
5) "b"
6) "3"
7) "a"
8) "4"
~~~

* **ZREVRANGE key start stop [WITHSCORES]** [返回有序集中指定区间内的成员，通过索引，分数从高到底。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 3 b
(integer) 1
127.0.0.1:6379> zadd k 2 c
(integer) 1
127.0.0.1:6379> zadd k 1 d
(integer) 1
127.0.0.1:6379> zrange k 0 -1 withscores
1) "a"
2) "4"
3) "b"
4) "3"
5) "c"
6) "2"
7) "d"
8) "1"
~~~

- **ZRANGEBYLEX key min max [LIMIT offset count]** [通过字典区间返回有序集合的成员。]

~~~shell
127.0.0.1:6379> zadd k 0 a 0 b 0 c 0 d 0 e 0 f 0 g
(integer) 7
127.0.0.1:6379> zrangebylex k - [c
1) "a"
2) "b"
3) "c"
127.0.0.1:6379> zrangebylex k - (c
1) "a"
2) "b"
127.0.0.1:6379> zrangebylex k [aaa (g
1) "b"
2) "c"
3) "d"
4) "e"
5) "f"
~~~

- **ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT]** [通过分数返回有序集合指定区间内的成员。]

~~~shell
127.0.0.1:6379> zadd salary 2500 jack                     # 测试数据
(integer) 0
127.0.0.1:6379> zadd salary 5000 tom
(integer) 0
127.0.0.1:6379> zadd salary 12000 peter
(integer) 0
127.0.0.1:6379> zrangebyscore salary -inf +inf           # 显示整个有序集
1) "jack"
2) "tom"
3) "peter"
127.0.0.1:6379> zrangebyscore salary -inf +inf withscores# 显示整个有序集及成员score值
1) "jack"
2) "2500"
3) "tom"
4) "5000"
5) "peter"
6) "12000"
127.0.0.1:6379> zrangebyscore salary -inf 5000 withscores # 显示工资小于等于5000所有成员
1) "jack"
2) "2500"
3) "tom"
4) "5000"
127.0.0.1:6379> zrangebyscore salary (5000 400000    # 显示工资大于5000小于等于400000成员
1) "peter"
~~~

- **ZREVRANGEBYSCORE key max min [WITHSCORES]** [返回有序集中指定分数区间内的成员，分数从高到低排序。]

~~~shell
127.0.0.1:6379> zadd salary 2500 jack
(integer) 1
127.0.0.1:6379> zadd salary 5000 tom
(integer) 1
127.0.0.1:6379> zadd salary 12000 peter
(integer) 1
127.0.0.1:6379> zrevrangebyscore salary +inf -inf           # 显示整个倒序集
1) "peter"
2) "tom"
3) "jack"
127.0.0.1:6379> zrevrangebyscore salary +inf -inf withscores# 显示整个倒序集及成员score值
1) "peter"
2) "12000"
3) "tom"
4) "5000"
5) "jack"
6) "2500"
127.0.0.1:6379> zrevrangebyscore salary +inf 5000 withscores # 显示工资大于等于5000所有成员
1) "peter"
2) "12000"
3) "tom"
4) "5000"
127.0.0.1:6379> zrevrangebyscore salary 10000 2000  # 逆序排列薪水介于10000和2000之间的成员
1) "tom"
2) "jack"
~~~

- **ZRANK key member** [返回有序集合中指定成员的索引。]

~~~shell
127.0.0.1:6379> zrange salary 0 -1 withscores        # 显示所有成员及其 score 值
1) "jack"
2) "2500"
3) "tom"
4) "5000"
5) "peter"
6) "12000"
127.0.0.1:6379> zrank salary tom                     # 显示 tom 的薪水排名，第二
(integer) 1
127.0.0.1:6379> zrank salary jack					 # 显示 jack 的薪水排名，第一
(integer) 0
~~~

* **ZREVRANK key member** [返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序。]

~~~shell
127.0.0.1:6379> zrevrange salary 0 -1 withscores     # 显示所有成员及其 score 值
1) "peter"
2) "12000"
3) "tom"
4) "5000"
5) "jack"
6) "2500"
127.0.0.1:6379> zrevrank salary tom					 # 显示 tom 的薪水倒序排名，第二
(integer) 1
127.0.0.1:6379> zrevrank salary jack 				 # 显示 jack 的薪水倒序排名，第三
(integer) 2
~~~

- **ZSCORE key member** [返回有序集中，成员的分数值。]

~~~shell
127.0.0.1:6379> zscore salary jack
"2500"
~~~

- **ZREM key member [member ...]** [移除有序集合中的一个或多个成员。]

~~~shell
127.0.0.1:6379> zadd k 4 a
(integer) 1
127.0.0.1:6379> zadd k 3 b
(integer) 1
127.0.0.1:6379> zadd k 2 c
(integer) 1
127.0.0.1:6379> zadd k 1 d
(integer) 1
127.0.0.1:6379> zrange k 0 -1 withscores
1) "d"
2) "1"
3) "c"
4) "2"
5) "b"
6) "3"
7) "a"
8) "4"
127.0.0.1:6379> zrem k c        # 移除单个元素
(integer) 1
127.0.0.1:6379> zrange k 0 -1
1) "d"
2) "b"
3) "a"
127.0.0.1:6379> zrem k b d      # 移除多个元素
(integer) 2
127.0.0.1:6379> zrange k 0 -1
1) "a"
127.0.0.1:6379> zrem k x        # 移除不存在元素
(integer) 0
~~~

- **ZREMRANGEBYLEX key min max** [移除有序集合中给定的字典区间的所有成员。]

~~~shell
127.0.0.1:6379> zadd k 0 a 0 b 0 c 0 d 0 e
(integer) 5
127.0.0.1:6379> zadd k 0 friend 0 apple 0 give 0 zip
(integer) 4
127.0.0.1:6379> zrange k 0 -1
1) "a"
2) "apple"
3) "b"
4) "c"
5) "d"
6) "e"
7) "friend"
8) "give"
9) "zip"
127.0.0.1:6379> zremrangebylex k [apple (give
(integer) 6
127.0.0.1:6379> zrange k 0 -1
1) "a"
2) "give"
3) "zip"
~~~

- **ZREMRANGEBYRANK key start stop** [移除有序集合中给定的排名区间的所有成员。]

~~~shell
127.0.0.1:6379> zadd salary 2500 jack
(integer) 1
127.0.0.1:6379> zadd salary 5000 tom
(integer) 1
127.0.0.1:6379> zadd salary 12000 peter
(integer) 1
127.0.0.1:6379> zremrangebyrank salary 0 1       # 移除下标 0 至 1 区间内的成员
(integer) 2
127.0.0.1:6379> zrange salary 0 -1 withscores    # 有序集只剩下一个成员
1) "peter"
2) "12000"
~~~

- **ZREMRANGEBYSCORE key min max** [移除有序集合中给定的分数区间的所有成员。]

~~~shell
127.0.0.1:6379> zadd salary 2500 jack
(integer) 1
127.0.0.1:6379> zadd salary 5000 tom
(integer) 1
127.0.0.1:6379> zadd salary 12000 peter
(integer) 1
127.0.0.1:6379> zremrangebyscore salary 5000 10000
(integer) 1
127.0.0.1:6379> zrange salary 0 -1 withscores
1) "jack"
2) "2500"
3) "peter"
4) "12000"
~~~

- **ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]** [计算给定的一个或多个有序集的并集，并将结果集存储在新的有序集合 key 中。]

~~~shell
# 有序集 k1
127.0.0.1:6379> zadd k1 70 a
(integer) 1
127.0.0.1:6379> zadd k1 70 b
(integer) 1
127.0.0.1:6379> zadd k1 99.5 c
(integer) 1
# 有序集 k2
127.0.0.1:6379> zadd k2 88 b
(integer) 1
127.0.0.1:6379> zadd k2 75 c
(integer) 1
127.0.0.1:6379> zadd k2 99.5 d
(integer) 1
# 并集
127.0.0.1:6379> zunionstore k 2 k1 k2
(integer) 4
# 显示有序集内所有成员及其分数值
127.0.0.1:6379> zunionstore k 2 k1 k2 weights 1 3   # k1中成语乘1 k2中成员乘3 并取并集
(integer) 4
127.0.0.1:6379> zrange k 0 -1 withscores
1) "a"
2) "70"
3) "d"
4) "298.5"
5) "c"
6) "324.5"
7) "b"
8) "334"
~~~

- **ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]** [计算给定的一个或多个有序集的交集，并将结果集存储在新的有序集合 key 中。]

~~~shell
# 有序集 k1
127.0.0.1:6379> zadd k1 70 a
(integer) 1
127.0.0.1:6379> zadd k1 70 b
(integer) 1
127.0.0.1:6379> zadd k1 99.5 c
(integer) 1
# 有序集 k2
127.0.0.1:6379> zadd k2 88 b
(integer) 1
127.0.0.1:6379> zadd k2 75 c
(integer) 1
127.0.0.1:6379> zadd k2 99.5 d
(integer) 1
# 交集
127.0.0.1:6379> zinterstore k 2 k1 k2 weights 2 3   # k1中成语乘2 k2中成员乘3 并取交集
(integer) 2
# 显示有序集内所有成员及其分数值
127.0.0.1:6379> zrange k 0 -1 withscores
1) "b"
2) "404"
3) "c"
4) "424"
~~~

- **ZSCAN key cursor [MATCH pattern] [COUNT count]** [迭代有序集合中的元素（包括元素成员和元素分值）。]

~~~shell
127.0.0.1:6379> zadd k 4 hello
(integer) 1
127.0.0.1:6379> zadd k 3 hi
(integer) 1
127.0.0.1:6379> zadd k 2 honour
(integer) 1
127.0.0.1:6379> zadd k 1 world
(integer) 1
127.0.0.1:6379> zadd k 0 redis
(integer) 1
# 迭代集合k，匹配以k开头元素。
127.0.0.1:6379> zscan k 0 match h* count 2
1) "0"
2) 1) "honour"
   2) "2"
   3) "hi"
   4) "3"
   5) "hello"
   6) "4"
~~~

### 服务器配置

```shell
1. Redis默认不是以守护进程的方式运行，可以通过该配置项修改，使用yes启用守护进程
    daemonize no
2. 当Redis以守护进程方式运行时，Redis默认会把pid写入/var/run/redis.pid文件，可以通过pidfile指定
    pidfile /var/run/redis.pid
3. 指定Redis监听端口，默认端口为6379，作者在自己的一篇博文中解释了为什么选用6379作为默认端口，因为6379在手机按键上MERZ对应的号码，而MERZ取自意大利歌女Alessia Merz的名字
    port 6379
4. 绑定的主机地址
    bind 127.0.0.1
5.当 客户端闲置多长时间后关闭连接，如果指定为0，表示关闭该功能
    timeout 300
6. 指定日志记录级别，Redis总共支持四个级别：debug、verbose、notice、warning，默认为verbose
    loglevel verbose
7. 日志记录方式，默认为标准输出，如果配置Redis为守护进程方式运行，而这里又配置为日志记录方式为标准输出，则日志将会发送给/dev/null
    logfile stdout
8. 设置数据库的数量，默认数据库为0，可以使用SELECT <dbid>命令在连接上指定数据库id
    databases 16
9. 指定在多长时间内，有多少次更新操作，就将数据同步到数据文件，可以多个条件配合
    save <seconds> <changes>
    Redis默认配置文件中提供了三个条件：
    save 900 1
    save 300 10
    save 60 10000
    分别表示900秒（15分钟）内有1个更改，300秒（5分钟）内有10个更改以及60秒内有10000个更改。
10. 指定存储至本地数据库时是否压缩数据，默认为yes，Redis采用LZF压缩，如果为了节省CPU时间，可以关闭该选项，但会导致数据库文件变的巨大
    rdbcompression yes
11. 指定本地数据库文件名，默认值为dump.rdb
    dbfilename dump.rdb
12. 指定本地数据库存放目录
    dir ./
13. 设置当本机为slav服务时，设置master服务的IP地址及端口，在Redis启动时，它会自动从master进行数据同步
    slaveof <masterip> <masterport>
14. 当master服务设置了密码保护时，slav服务连接master的密码
    masterauth <master-password>
15. 设置Redis连接密码，如果配置了连接密码，客户端在连接Redis时需要通过AUTH <password>命令提供密码，默认关闭
    requirepass foobared
16. 设置同一时间最大客户端连接数，默认无限制，Redis可以同时打开的客户端连接数为Redis进程可以打开的最大文件描述符数，如果设置 maxclients 0，表示不作限制。当客户端连接数到达限制时，Redis会关闭新的连接并向客户端返回max number of clients reached错误信息
    maxclients 128
17. 指定Redis最大内存限制，Redis在启动时会把数据加载到内存中，达到最大内存后，Redis会先尝试清除已到期或即将到期的Key，当此方法处理 后，仍然到达最大内存设置，将无法再进行写入操作，但仍然可以进行读取操作。Redis新的vm机制，会把Key存放内存，Value会存放在swap区
    maxmemory <bytes>
18. 指定是否在每次更新操作后进行日志记录，Redis在默认情况下是异步的把数据写入磁盘，如果不开启，可能会在断电时导致一段时间内的数据丢失。因为 redis本身同步数据文件是按上面save条件来同步的，所以有的数据会在一段时间内只存在于内存中。默认为no
    appendonly no
19. 指定更新日志文件名，默认为appendonly.aof
     appendfilename appendonly.aof
20. 指定更新日志条件，共有3个可选值： 
    no：表示等操作系统进行数据缓存同步到磁盘（快） 
    always：表示每次更新操作后手动调用fsync()将数据写到磁盘（慢，安全） 
    everysec：表示每秒同步一次（折中，默认值）
    appendfsync everysec
21. 指定是否启用虚拟内存机制，默认值为no，简单的介绍一下，VM机制将数据分页存放，由Redis将访问量较少的页即冷数据swap到磁盘上，访问多的页面由磁盘自动换出到内存中（在后面的文章我会仔细分析Redis的VM机制）
     vm-enabled no
22. 虚拟内存文件路径，默认值为/tmp/redis.swap，不可多个Redis实例共享
     vm-swap-file /tmp/redis.swap
23. 将所有大于vm-max-memory的数据存入虚拟内存,无论vm-max-memory设置多小,所有索引数据都是内存存储的(Redis的索引数据 就是keys),也就是说,当vm-max-memory设置为0的时候,其实是所有value都存在于磁盘。默认值为0
     vm-max-memory 0
24. Redis swap文件分成了很多的page，一个对象可以保存在多个page上面，但一个page上不能被多个对象共享，vm-page-size是要根据存储的 数据大小来设定的，作者建议如果存储很多小对象，page大小最好设置为32或者64bytes；如果存储很大大对象，则可以使用更大的page，如果不 确定，就使用默认值
     vm-page-size 32
25. 设置swap文件中的page数量，由于页表（一种表示页面空闲或使用的bitmap）是在放在内存中的，，在磁盘上每8个pages将消耗1byte的内存。
     vm-pages 134217728
26. 设置访问swap文件的线程数,最好不要超过机器的核数,如果设置为0,那么所有对swap文件的操作都是串行的，可能会造成比较长时间的延迟。默认值为4
     vm-max-threads 4
27. 设置在向客户端应答时，是否把较小的包合并为一个包发送，默认为开启
    glueoutputbuf yes
28. 指定在超过一定的数量或者最大的元素超过某一临界值时，采用一种特殊的哈希算法
    hash-max-zipmap-entries 64
    hash-max-zipmap-value 512
29. 指定是否激活重置哈希，默认为开启（后面在介绍Redis的哈希算法时具体介绍）
    activerehashing yes
30. 指定包含其它的配置文件，可以在同一主机上多个Redis实例之间使用同一份配置文件，而同时各个实例又拥有自己的特定配置文件
    include /path/to/local.conf
```

### Redis 迁移

~~~shell
# 安装相关软件
yum -y install automake
yum -y install libtool
yum -y install autoconf
yum -y install bzip2
# 安装redis-migrate-tool
unzip redis-migrate-tool-master.zip
cd redis-migrate-tool-master
autoreconf -fvi
./configure
make
# 检查安装是否成功，如下图所示即为正确
src/redis-migrate-tool -h
# 修改配置文件
vim rmt.conf
[source]
type: single
servers :
-192.168.1.11:6379
redis_auth:数据库密码
[target]
type: single
servers:
-192.168.1.11:7379
redis_auth:数据库密码
[common]
listen: 0.0.0.0:8888
threads:8
# 启动迁移程序
src/redis-migrate-tool  -c rmt.conf  -o log -d
# 检测
src/redis-migrate-tool  -c rmt.conf  -o log -C redis_check 10000    #默认抽样1000个key

Check job is running...
Checked keys: 1000
Inconsistent value keys: 0
Inconsistent expire keys : 0
Other check error keys: 0
Checked OK keys: 1000

All keys checked OK!
Check job finished, used 9.755s
~~~





