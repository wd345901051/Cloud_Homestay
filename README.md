# MyCode

#### 介绍
一个基于go的gin框架，结合gorm，ajax，jQuery写出来的类似于酒店管理系统的网站，是学生时期的一项课设作业。
本项目为开源项目，欢迎大家使用调整更新。
相关的业务信息：
（1）酒店的首页显示：读取数据库中的数据展示到首页。
（2）代码的框架设计：采用MVC模式进行设计
（3）酒店的管理：房间管理与用户管理
（4）网页路由的设置：仿照正常的业务逻辑实现路由跳转
（5）管理员身份、游客身份和用户身份：对数据库中的用户表进行相应的字段设置
（6）网站的稳定性：本地开发模式，暂不考虑
（7）网站的额外业务：由开发小组进行确认分析设计
（8）网站的设计基础：由本小组成员根据小组实际情况进行确认

#### 运行
使用 docker build . -t cloud_homestay 构建镜像并运行