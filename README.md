# memo
这是一个简单的备忘录API服务器。

### 功能：
1. 连接MySQL数据库，一张memo表，一张user表
2. 提供注册 登录和刷新token接口，token使用JWT
3. 路由中间件包含打印接口请求和验证JWT
4. 备忘录的增删改查接口尝试RESTful API规范
5. 提供更新用户信息的接口
6. 文件上传的简单尝试
7. 所有接口返回JSON格式规范化（包含code,msg,time,data）
