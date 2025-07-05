# 评论区网页项目

该项目是一个简单的评论区网页，允许用户提交评论并查看其他用户的评论。

## 项目结构

```
fresher-project
├── config.json                    # 数据库配置
├── rear
│   ├── main.go                    # Go 后端入口
│   ├── server.go                  # Go 后端接口实现
│   └── db
│       ├── init.go                # 数据库初始化
│       └── model.go               # 数据库模型
│   └── go.mod                     # Go module 配置
│   └── go.sum                     # Go module 依赖锁定
├── src
│   ├── index.html                 # 前端主页面
│   ├── styles
│   │   └── style.css              # 前端样式
│   ├── scripts
│   │   └── main.js                # 前端逻辑
│   └── assets                     # 静态资源（可为空）
├── .gitignore                     # Git 忽略文件
└── README.md                      # 项目说明文档
```
