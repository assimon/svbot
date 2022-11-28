
## svbot

短视频直链解析机器人，支持：抖音/火山/快手/绿洲/皮皮虾/微博/微视/西瓜/最右|的短视频无水印解析

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue" alt="license MIT"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Golang-1.17-red" alt="Go version 1.17"></a>
<a href="https://github.com/tucnak/telebot"><img src="https://img.shields.io/badge/Telebot Framework-v3-lightgrey" alt="telebot v3"></a>
<a href="https://github.com/assimon/svbot/releases"><img src="https://img.shields.io/badge/version-0.0.1-green" alt="version 0.0.1"></a>
</p>


## 项目初衷
如果你也讨厌在分享短视频的时候被平台加上牛皮癣式的水印或小尾巴(例如：`xx分享美好生活`)      
那么，`svbot`将帮助你解决这些烦恼~   

## 使用方式

### 一、自行编译
此安装方式多用于开发者，需电脑上安装`go语言`环境。   
[go语言官网](https://golang.org/)

下载：
```shell
# 下载项目
git clone https://github.com/assimon/svbot && cd svbot
```
编译：
```shell
# 编译
go build -o svbot
# 给予执行权限
chmod +x ./svbot
```
配置：
```shell
cp .config.toml.example config.toml
```
执行：
```shell
# 调试启动
./svbot
# nohup 
nohup ./svbot >> run.log 2>&1 &
```

### 二、下载已经编译好的二进制程序
此方式可以直接使用，用于服务器生产环境。
进入打包好的版本列表，下载程序：[https://github.com/assimon/svbot/releases](https://github.com/assimon/svbot/releases)    
配置：
```shell
cp .config.toml.example config.toml
```
运行：
```shell
# linux
# 调试启动
./svbot

# windows
svbot.exe
```

## 配置
申请`Telegram bot api token` [https://t.me/BotFather](https://t.me/BotFather)     
修改`config.toml`里面的`api_token`参数即可，无需其他配置。       
例如：`api_token="xxxxxx"`     

## 使用预览
### 1.复制短视频链接
<img src="https://cloud-dc1-lg.utf8.hk/api/v3/slave/source/0/dXBsb2Fkcy8yMDIyLzExLzI4L3NKU1lVMUtuX2NvcHlfdXJsLmdpZg/copy_url.gif?sign=uvK1moqH5KrFvTzgu_uOd6dpbaojbDXcbFCuf6NC5sc%3D%3A0" width=500 height=1200 />

### 2.发送给机器人
<img src="https://cloud-dc1-lg.utf8.hk/api/v3/slave/source/0/dXBsb2Fkcy8yMDIyLzExLzI4L1hVRXhlRlNsX3ZpZGVvLmdpZg/video.gif?sign=kzoUonIVXC8nEIaE9u5HJwKh6HXbFRuWs63G9Ef10zU%3D%3A0" width=500 height=1200 />

## 鸣谢
部分平台解析灵感来自：[https://github.com/5ime/video_spider](https://github.com/5ime/video_spider)

## 声明`warning！`
项目仅用于`学习交流`使用，禁止用于任何商业目的。       
由于开源项目的分发不可控，故使用者在使用该项目时出现的任何法律问题由使用者自行承担！       