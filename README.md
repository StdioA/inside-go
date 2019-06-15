# Inside-Go

![](https://img.shields.io/badge/subject-a_micro_blog_of_my_own-brightgreen.svg?style=flat)
[![](https://img.shields.io/badge/made%20with-%E2%9D%A4-ff69b4.svg)](https://www.stdioa.com/)

[Inside](https://github.com/StdioA/inside) 的姊妹项目，后端使用 gin 重写，用来练习 Golang 后端开发。  
一个迷你博客，使用管理员用户创建及管理 post, 用户可在 post 下面评论，提供归档页面，数据导入导出功能。

## TODO
- [ ] 实现一个账号系统，通过命令添加用户，用户名密码使用 bcrypt 加密存在数据库中，通过 HTTP Basic Auth 登录
- [ ] 数据库配置注入
- [x] 使用 [multitemplate](https://github.com/gin-contrib/multitemplate/) 复用模板文件

前端使用 [Vue.js](http://vuejs.org/) 构建。
