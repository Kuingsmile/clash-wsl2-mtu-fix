# clash-wsl2-mtu-fix
修复wsl2镜像网络模式与clash tun模式共同使用时，mtu错误导致的无法联网问题

## 使用方式

1. 从release下载打包的二进制文件，放到自己定义的目录，给以执行权限
2. 下载`clash-wsl2-mtu-fix.service`文件，放入`/etc/systemd/system`文件夹，然后执行`systemctl enable clash-wsl2-mtu-fix`命令
