# TikTokLite
  字节跳动极简版抖音项目

#  目录结构
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
  - 配置文件，目前只配置了mysql，后续redis、nigix？
- TikTokLite.sql
  - 数据库建表文件  