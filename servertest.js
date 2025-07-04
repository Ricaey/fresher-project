// 引入 express 框架
const express = require('express');
// 创建 express 应用实例
const app = express();

// 定义根路由，访问 / 时返回 "Hello World!"
app.get('/', (req, res) => {
    res.send('Hello World!');
});

// 设置服务器监听端口
const port = 8989;
// 启动服务器并输出启动信息
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}/`);
});