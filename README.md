### 简介
cwrs_go_server 是一个基于go语言开发的web服务框架。
基于gin框架，使用jwt作为token验证框架、gorm作为数据库操作框架、使用viper作为配置文件框架、zap作为日志框架、swagger作为api文档。

### 开发环境
```text
运行环境：go1.24
数据库：mysql8.0
redis：redis7.4
```

### 项目结构
```text
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── logs
│   ├── error.log
│   └── info.log
├── src
│   ├── cwrs_common            #公共模块
│   ├── cwrs_core              #核心模块 (如：gin jwt gorm viper zap swagger等)
│   ├── cwrs_generate          #代码生成 
│   ├── cwrs_routes            #路由模块
│   ├── cwrs_utils             #工具类
│   └── server                 #服务模块
├── application.yaml           #配置文件
├── go.mod
│   └── go.sum
├── main.go
├── README.md
├── run.sh
```

### swagger 文档
```text
  注解	           作用	             示例
@Summary        接口简短描述         @Summary 获取用户列表
@Description    接口详细描述         @Description 获取所有用户的基本信息，支持分页
@Tags           接口分类标签         @Tags users
@Accept         请求数据格式         @Accept json
@Produce        响应数据格式         @Produce json
@Param          请求参数            @Param id path int true "用户ID"
@Success        成功响应            @Success 200 {object} User
@Failure        失败响应            @Failure 404 {object} User
@Router	        路由信息            @Router /users/{id} [get]
```
文档生成命令
在项目根目录下执行命令更新文档:
```bash
swag init
```

### Apifox 文档（个人推荐 这个看着更舒服）
```text
首先要生成swagger文档，然后才能导入到apifox，导入文档步骤如下：
1.安装apifox
2.打开apifox，创建项目，导入数据，选择定时导入，新建（导入频率设置自己喜欢的方式，数据员格式选择OpenAPI/Swagger，数据源名称填写项目名称，数据源URL填写http://localhost:9090/swagger/doc.json）
3.点击运行，生成文档
```

### vscode 运行
```bash
go run .\main.go
```
##### 个人更推荐使用golang编辑器

### linux 运行
```text
1.直接把代码放在linux服务器上运行，需要先安装go环境，然后执行以下命令：
run.sh


2.打包成二进制文件上传到linux服务器，然后执行二进制文件：

    windows 系统建议使用WSL2 虚拟机，然后执行以下命令：
    
    切换工作目录：
        cd /mnt/对应盘符/项目路径
    
    打包命令：
        go build -o 二进制文件名 main.go

注意：1、2选择一个即可，不要同时执行（个人更喜欢打包成二进制文件）。
```

### 业务细节
```text
1.上传文件 使用oss存储 需要在配置文件中配置oss信息 如需本地上传请自行实现
2.所有业务表需有dept_id、created_user_id字段 权限拦截器中用于数据权限控制。
```

### 开发文档
- 查看开发文档：暂无文档，后续补充。

[//]: # (- <a href="/" target="_blank">cwrs_go_server</a>)

### 演示地址
- <a href="http://47.92.24.24:5173/#/login" target="_blank">http://47.92.24.24:5173/#/login</a>
- 账号：admin  密码：123456
- 账号：user  密码：123456

### 联系作者
- QQ交流群：837701104

### GO后端代码仓库
- gitee地址：<a href="https://gitee.com/open-source-project_7/cwrs_go_server" target="_blank">cwrs_go_server</a>

- github地址：<a href="https://github.com/vunwang/cwrs_go_server" target="_blank">cwrs_go_server</a>

### 前端代码仓库
- gitee地址：<a href="https://gitee.com/open-source-project_7/cwrs_vue3" target="_blank">cwrs_vue3</a>

- github地址：<a href="https://github.com/vunwang/cwrs_vue3" target="_blank">cwrs_vue3</a>

### 免责声明
- 1、用户不得利用CWRS管理系统从事非法行为，用户应当合法合规的使用，发现用户在使用产品时有任何的非法行为，CWRS管理系统不承担用户因非法行为造成的任何法律责任，一切法律责任由用户自行承担，如因用户使用造成第三方损害的，用户应当依法予以赔偿。

- 2、所有与使用CWRS管理系统相关的资源直接风险均由用户承担。

### 🎉 关注项目
- 如果项目对您有帮助，请点右上角 "Star" 收藏项目，您的支持是我创作的动力！