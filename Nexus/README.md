# Nexus
## 环境搭建
> 使用Docker搭建Nexus私库环境 
~~~shell
docker run -p 10005:8081 --restart=always --name kenny_nexus -v D:/develop/nexus-data:/nexus-data -d sonatype/nexus3:latest
# 设置端口映射，自动重启，名字，文件路径绑定（此处没使用volume），后台运行。
~~~