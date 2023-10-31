# golang NUIST 校园网登录

## ** V2 版本正在开发中，即将添加多拨支持和注册服务支持 **

着急的小伙伴去 Actions 下载编译好的二进制文件，或者自己编译

但是 CI 版本不保证稳定，如果你想要稳定版本，可以下载 release 中的二进制文件

## 特点

原生跨平台，无需任何运行时即可运行，鲁棒性较好，支持 2021 之后的校园网登录。  
这个项目的主要目标是在 linux bsd 等不常见环境上进行校园网登录，尤其为了解决 服务器、软路由 的自动登录与校内 DDNS 问题  
V2 特性: 支持多拨，支持作为服务监测联网情况，检测到掉线自动重拨

## 功能

- :white_check_mark: 软路由支持
- :white_check_mark: 健壮的校园网登录
- :white_check_mark: 多拨支持
- :white_check_mark: Cloudflare DDNS
- :white_large_square: 注册服务
- :white_large_square: 自动安装配置脚本
- :white_large_square: 阿里云 DDNS
- :white_large_square: 腾讯云 DDNS

## 使用说明

### 一键安装脚本

todo

### 下载运行

#### 提供参数运行

去 releases 中找到自己机器对应系统、指令集的程序( MacOS 是 Darwin )，然后在终端中执行：

```bash
./login --username=用户名 --password=密码 --isp=运营商
```

例如下面

```bash
./login --username=18812340000 --password=114514 --isp=ChinaNet
```

login 是你从 releases 下得到文件的名字

isp 可选参数为：

| 运营商   | isp 内容 |
| -------- | -------- |
| 校园网   | NUIST    |
| 中国移动 | CMCC     |
| 中国电信 | ChinaNet |
| 中国联通 | UNION    |

### DDNS

#### Cloudflare DDNS

目前本工具中 CF 的 DDNS 功能，测试下来非常稳，缺点是，因为是国外 DNS，解析记录更新扩散到国内需要一段时间（大约十分钟到半小时不等）

为了使用 Cloudflare DDNS，最好你原来就是 Cloudflare 用户。

你需要提供三个参数

| 参数名     | 说明                |
| ---------- | ------------------- |
| cfzoneid   | CloudFlare ZoneID   |
| cfAPIToken | CloudFlare API 令牌 |
| cfhostname | 你要 DDNS 的域名    |

前两个参数你可以参考 https://lighti.me/5560.html ，拿到这两个参数就可以

你可以通过运行参数或者环境变量的方式提供给程序这三个参数，参数名和环境变量名都分别为 cfzoneid，cfapitoken 和 cfhostname

如果完整地通过运行参数或者环境变量提供了这三个参数，会自动执行 DDNS

## 自己编译

确保你有 golang 环境，然后

```bash
go build
```

这样，你可以得到一份你的平台下的代码

你也可以在项目文件夹下

```bash
./build.sh
```

在 windows 下你可以使用 git bash

这样，你会得到全部平台的二进制文件
