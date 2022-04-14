# golang NUIST 校园网登录

## 特点

原生跨平台，无需任何运行时即可运行，鲁棒性较好，支持2021之后的校园网登录。  
这个项目的主要目标是在 linux bsd 等不常见环境上进行校园网登录，尤其为了解决 服务器、软路由 的自动登录与校内DDNS问题  

## 功能

- [X] 软路由支持
- [X] 健壮的校园网登录
- [ ] 注册crontab/服务
- [ ] 自动安装配置脚本
- [ ] 阿里云DDNS
- [ ] 腾讯云DDNS
- [ ] Cloudflare DDNS
- [ ] 多拨支持 

## 使用说明

### 一键安装脚本

todo

### 下载运行

去releases中找到自己机器对应系统、指令集的程序( MacOS 是 Darwin )，然后在终端中执行：

```bash
./login --username=用户名 --password=密码 --isp=运营商
```

例如下面

```bash
./login --username=18812340000 --password=114514 --isp=ChinaNet
```

isp可选参数为：  

运营商|isp内容
---|---
校园网|NUIST
中国移动|CMCC
中国电信|ChinaNet
中国联通|UNION

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