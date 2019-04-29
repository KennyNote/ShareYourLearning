# Docker Command
## 容器生命周期管理
### 👉run

#### Description

> Run a command in a new container

#### Usage

~~~shell
$ docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
~~~

#### Options

- **-a , --attach = []** : 指定标准输入输出内容类型，可选 **STDIN / STDOUT / STDERR** 三项；
- **--add-host** : 添加host和ip之间的映射
- **-c , --cpu-shares = 0** : 设置容器CPU权重，在CPU共享场景使用；
- **--cpuset** : 绑定容器到指定CPU运行，此参数可以用来容器独占CPU；
- **--cap-add = []** : 添加权限；
- **--cap-drop = []** : 删除权限；
- **--cidfile = ""** : 运行容器后，在指定文件中写入容器PID值，一种典型的监控系统用法；
- **-d , --detach = false** : 后台运行容器，并返回容器ID；
- **--device = []** : 添加主机设备给容器，相当于设备直通；
- **--dns** : 指定容器使用的DNS服务器，默认和宿主一致；
- **--dns-opt** : 设置DNS选项
- **--dns-option** : 设置DNS选项
- **--dns-search** : 指定容器DNS搜索域名，写入到容器的**/etc/resolv.conf**文件，默认和宿主一致；
- **-e , --env** : 设置环境变量；
- **--env-file = []** : 从指定文件读入环境变量；
- **--entrypoint = "" **: 覆盖image的入口点；
- **--expose = []** : 开放一个端口或一组端口；
- **-h , --hostname = ""** : 指定容器的主机名指定容器的hostname；
- **--health-cmd** : 检测容器健康；
- **-i , --interactive = false** : 以交互模式运行容器，通常与 -t 同时使用；
- **--ip** = [] : 设置固定IPv4
- **--ip6** = [] :  设置固定IPv6
- **-l , --label** : 设置容器元数据
- **--link = []** : 指定容器间的关联，使用其他容器的IP、env等信息；
- **--lxc-conf = []** : 指定容器的配置文件，只有在指定**--exec-driver =  lxc**时使用；
- **-m , --memory** :设置容器使用内存最大值；
- **--name** : 为容器指定一个名称；
- **--net** : 指定容器的网络连接类型，支持 **bridge/host/none/container**四种类型；
- **-p , --publish** : 端口映射，格式为：**主机(宿主)端口:容器端口**；
- **-P , --publish-all = false** : 端口随机映射；
- **--privileged = false** : 指定容器是否为特权容器，特权容器拥有所有的capabilities；
- **--restart = ""** : 指定容器停止后的重启策略，待详述
- **--rm = false** : 指定容器停止后自动删除容器(不支持以docker run -d启动的容器)
- **--sig-proxy = true** : 设置由代理接受并处理信号，但是SIGCHLD、SIGSTOP和SIGKILL不能被代理
- **-t , --tty = false** : 为容器重新分配一个伪输入tty终端，通常与 -i 同时使用；
- **-u , --user = ""** : 指定容器的用户
- **-v , --volume = []** : 给容器挂载存储卷，挂载到容器的某个目录
- **--volumes-from = []** : 给容器挂载其他容器上的卷，挂载到容器的某个目录
- **-w , --workdir = ""** :指定容器的工作目录
- 未完待续。。。
### 👉start/stop/restart
### 👉kill
### 👉rm
### 👉pause/unpause
### 👉create
### 👉exec
## 容器操作
### 👉ps
### 👉inspect
### 👉top
### 👉attach
### 👉events
### 👉logs
### 👉wait
### 👉export
### 👉port
## 容器rootfs命令
### 👉commit
### 👉cp
### 👉diff
## 镜像仓库
### 👉login
### 👉pull
### 👉push
### 👉search
## 本地镜像管理
### 👉images
### 👉rmi
### 👉tag
### 👉build
### 👉history
### 👉save
### 👉import
### 👉info|version
### 👉info
### 👉version