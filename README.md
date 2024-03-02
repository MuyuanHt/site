## site-collaboration-platform

#### windows 下执行 deploy 中的 shell 脚本

- 如果使用终端无法成功执行 sh 命令，可以尝试打开开发者选项或者使用 git bash 执行
- 在当前目录打开 git bash 并使用 bash 执行

#### 部署时修改文件配置时需要注意的地方

- 在 conf 目录下的 `app_conf.go` 文件中定义了几个常量，修改文件配置的话按需修改

#### 使用到的一些第三方库或实现方法

- 配置信息：viper
- 密码加密：bcrypt
- 自增uid：go-snowflake
- 登录认证：jwt-go (rsa 非对称加密)