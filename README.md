# NUIST 校园网登录 Golang Implemention

## 特点

原生跨平台，无需任何运行时即可运行，鲁棒性较好，支持 2021 之后的校园网登录。

这个项目的主要目标是在 linux, openwrt 等无浏览器环境上进行校园网登录，尤其为了解决 服务器、软路由 的自动登录与校内 DDNS 问题

（V2 特性）支持多拨，支持作为服务无人值守地监测联网情况，检测到掉线自动重拨，支持多账户登录

## 功能

- :white_check_mark: OpenWrt 支持
- :white_check_mark: 健壮的校园网登录
- :white_check_mark: 多拨支持
- :white_check_mark: Cloudflare DDNS
- :white_large_square: 注册服务
- :white_large_square: 自动安装配置脚本
- :white_large_square: 阿里云 DDNS
- :white_large_square: 腾讯云 DDNS

## 使用说明

一般地，你需要去 releases 中找到自己机器对应系统、指令集的二进制文件( MacOS 是 Darwin )，然后在终端中执行它

<!-- ### 一键安装脚本

todo -->

### 普通登录

```bash
./i-nuist-login-xxx login [userstring]
```

**userstring** 是你的用户名和密码，格式为 `username:password@ISP`

isp 可选参数为：

| 运营商   | isp 内容 |
| -------- | -------- |
| 校园网   | NUIST    |
| 中国移动 | CMCC     |
| 中国电信 | ChinaNet |
| 中国联通 | UNION    |

例

`18812340000:114514@CMCC` 是 账户 18812340000 密码 114514 运营商 中国移动

`18812340000:114514@ChinaNet` 是 账户 18812340000 密码 114514 运营商 中国电信

`18812340000:114514@UNION` 是 账户 18812340000 密码 114514 运营商 中国联通

`18812340000:114514@NUIST` 是 账户 18812340000 密码 114514 运营商 校园网

最后，执行指令例如

```bash
./i-nuist-login-xxx login 18812340000:114514@CMCC
```

### Daomon 模式

```bash
./i-nuist-login-xxx daemon [userstring]
```

### 显示调试信息

```bash
./i-nuist-login-xxx info [userstring]
```

### 可选参数

| 参数          | 作用                                                                 |
| ------------- | -------------------------------------------------------------------- |
| -v --verbose  | 显示调试信息，会显示所有的 http 请求和响应，以及登录过程中的各种信息 |
| -h --help     | 显示帮助信息                                                         |
| -u --userpool | 指定用户池文件                                                       |
| -c --config   | 指定配置文件                                                         |
| -s --syncdial | 强制多拨模式                                                         |

### 用户池文件

用户池文件是一个文本文件，每行一个用户，格式为 `username:password@ISP, `

**请务必注意最后的逗号**，逗号是真正的分隔符，回车和空格都会被忽略

例

```text
18812340000:114514@CMCC,
18812340000:114514@ChinaNet,
18812340000:114514@UNION,
18812340000:114514@NUIST,
```

当提供了用户池文件时，程序会自动进入多拨模式，每个接口随机挑选一个用户进行登录

## 配置文件

没做

## 参与开发

### 编译

确保你有 golang 环境，然后

```bash
go build
```

这样，你可以得到一份你的平台下的代码

你也可以在项目文件夹下

```bash
./build.sh
```

这样，你会得到全部平台的二进制文件

### 运行

```bash
go run login.go [args]
```

### 提交 PR

请确保你的代码符合 golang 的规范，使用 `gofmt` 格式化你的代码
