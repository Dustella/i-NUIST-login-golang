# golang NUIST 校园网登录

## 特点

原生跨平台，无需任何运行时即可运行，鲁棒性较好，支持2021之后的校园网登录。  
这个项目的主要目标是在 linux bsd 等不常见环境上进行校园网登录，尤其为了解决 服务器、软路由 的自动登录与校内DDNS问题  
todo: 作为服务监测联网情况，检测到掉线就重拨  

## 功能

- [X] 软路由支持
- [X] 健壮的校园网登录
- [ ] 注册服务
- [ ] 自动安装配置脚本
- [ ] 阿里云DDNS
- [ ] 腾讯云DDNS
- [X] Cloudflare DDNS
- [ ] 多拨支持

## 使用说明

### 一键安装脚本

todo

### 下载运行

#### 提供参数运行

去releases中找到自己机器对应系统、指令集的程序( MacOS 是 Darwin )，然后在终端中执行：

```bash
./login --username=用户名 --password=密码 --isp=运营商
```

例如下面

```bash
./login --username=18812340000 --password=114514 --isp=ChinaNet
```

login是你从releases下得到文件的名字

isp可选参数为：  

运营商|isp内容
---|---
校园网|NUIST
中国移动|CMCC
中国电信|ChinaNet
中国联通|UNION

#### 提供环境变量

你可以在系统环境变量中提供配置。这样，不需要给程序参数也能正常工作。



### DDNS

#### Cloudflare DDNS

目前本工具中 CF 的 DDNS 功能，测试下来非常稳，缺点是，因为是国外DNS，解析记录更新扩散到国内需要一段时间（大约十分钟到半小时不等）

为了使用Cloudflare DDNS，最好你原来就是Cloudflare用户。

你需要提供三个参数

参数名 | 说明
----|----
cfzoneid| CloudFlare ZoneID
cfAPIToken | CloudFlare API令牌
cfhostname | 你要DDNS的域名

前两个参数你可以参考 https://lighti.me/5560.html ，拿到这两个参数就可以

你可以通过运行参数或者环境变量的方式提供给程序这三个参数，参数名和环境变量名都分别为cfzoneid，cfapitoken和cfhostname

如果完整地通过运行参数或者环境变量提供了这三个参数，会自动执行DDNS

## 自己编译

确保你有golang环境，然后

```bash
go build
```

这样，你可以得到一份你的平台下的代码

你也可以在项目文件夹下

```bash
./build.sh
```

在windows下你可以使用 git bash

这样，你会得到全部平台的代码。
