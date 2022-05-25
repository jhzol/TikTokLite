# TikTokLite
  字节跳动极简版抖音项目

###  目录结构
- controller 
  - 视图层，处理前端消息
- log 
  - 日志组件
  - 分为`debug` `info` `error` `fatal`四个等级
  - 分为`debug` `debugw` `debugf`三种输出格式，直接调用`log.Debug` `log.Debugw` `log.Debugf`即可
- logfile
  - 日志文件，info等级以下在info+time.log文件,以上在error+time.log文件，可在initlog函数中修改路径（或后续写入config中）
- proto
  - 前端消息结构体，有protobuf文件自动生成
- repository
  - 数据层，直接对数据库进行操作
  - 各model通过函数`GetDB()`获取数据库
- response
  - 对返回消息进行封装
  - 成功调用success，失败调用fail
- routes
  - 路由层
- service  
  - 逻辑层，执行业务操作
  - 从数据层获取数据，封装后返回试图层
- util
  - 工具函数
- config.yaml
  - 配置文件，目前只配置了mysql，minio对象存储，后续redis、nigix？
- TikTokLite.sql
  - 数据库建表文件  

### 配置说明

**目前需要配置的服务**

- **mysql**

  - 配置好mysql数据库，申请读写权限用户

    ```sql
    //创建用户
    CREATE USER user@'%' IDENTIFIED by '123'
    //授予权限并允许远程访问
    GRANT ALL ON *.* TO 'user'@'%' WITH GRANT OPTION;
    ```

  - 执行TikTokLite.sql脚本生成表结构

- **minio对象存储**（目前部署在本机）

  **下载并启动minio**

  linux

  ```bash
  #下载执行文件并修改权限
  wget https://dl.min.io/server/minio/release/linux-amd64/minio
  chmod +x minio
  #启动minio服务并指定文件存储位置
  ./minio server /data
  ```

  windows

  从以下url下载可执行文件[[minio]](http://dl.minio.org.cn/server/minio/release/windows-amd64/minio.exe)

  使用以下命令在 Windows 主机上运行独立的 MinIO 服务器。 将“D:\”替换为您希望 MinIO 存储数据的驱动器或目录的路径。 您必须将终端或 powershell 目录更改为 `minio.exe` 可执行文件的位置，*或*将该目录的路径添加到系统 `$PATH` 中：

  ```cmd
  minio.exe server D:\
  ```

  启动后可以看到以下输出

  ![](https://github.com/jhzol/test/blob/master/image/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202022-05-17%20154717.png?raw=true)

​	在浏览器中输入主机ip:9000即可打开minio提供的前端管理界面，对外访问地址为`服务ip:端口/bucket名称/文件名称`(需将bucket的access policy设置为public）

- **使用手机访问本机服务**（模拟器应该可以不用配置）

  - 手机与本机处于同一局域网
  - 开放windows端口
    - 打开windows设置->windows安全中心->防火墙和网络保护->高级设置
    - 在入栈规则中新建规则
    - 规则类型选择端口，应用于TCP连接、指定端口为服务端口（gin默认8080，minio9000）
    - 一路下一步即可
    - 可用手机浏览器访问主机ip：端口号查看后端是否有访问请求
    - 如果不行可尝试在windows设置->windows安全中心->防火墙和网络保护中关闭防火墙

  - 如果部署在wsl上还需指定端口转发

    - ifconfig查看wsl的ip （一般被eth0）

    - 在win下用管理员模式打开powershell

    - 输入以下命令，表示将任意源访问本机8080端口的请求转发到wsl的ip的8080端口

      ```powershell
      netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=8080 connectaddress=(wsl的ip) connectport=8080
      ```

    - 可使用`netsh interface portproxy show all`命令查看目前的端口转发