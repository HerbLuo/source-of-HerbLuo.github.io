# MacOS
rust mysql错误

```
library not found for -lmysqlclient
```

# WSL
启用`SSH Server`
```shell
sudo systemctl enable sshd
sudo systemctl start sshd
```

如果失败,可以使用 `sudo /usr/sbin/sshd -d` 查看日志

如果原因是确密钥,执行 `/usr/bin/ssh-keygen -A` 即可

更改`sshd`配置以支持远程访问
```shell
sudo vim /etc/ssh/sshd_config
# 确保以下选项正确
PasswordAuthentication yes
```
```
sudo systemctl restart sshd
```

# 阿里云,腾讯云,华为云等
使用`autossh`需修改`ssh server`的配置
```shell
GatewayPorts yes
```


# Linux
sudo 免密码
```shell
sudo su -
chmod u+w /etc/sudoers
vim /etc/sudoers
# 文件尾添加 START
用户名 ALL=(ALL:ALL) NOPASSWD:ALL
# 文件尾添加 END
chmod u-w /etc/sudoers
```

# Manjaro
1. 排序并增加China源，以及添加`archlinuxcn`

```shell
sudo pacman-mirrors -i -c China -m rank
```
运行完成后会弹出对话框要求选择源，选择23个速度快的

添加 `archlinuxcn`
（可选，建议照做，下方某些步骤的包只存在于`archlinuxcn`当中）
```
sudo vi /etc/pacman.conf
```
在文件末尾添加
```
[archlinuxcn]
SigLevel = Optional TrustedOnly
Server = https://mirrors.ustc.edu.cn/archlinuxcn/$arch
```
导入GPG Key
```
sudo pacman -Syy && sudo pacman -S archlinuxcn-keyring
```

2. 升级系统

```shell
sudo pacman -Syyu
```

3. 安装 `vim`
```
sudo pacman -S vim
```

4. 英文系统 安装中文（搜狗）输入法
```
sudo pacman -S fcitx-im #默认全部安装

sudo pacman -S fcitx-configtool

# 如果不使用搜狗输入法，以下可以省略（上述步骤做完后，自带中文输入法，够用）
sudo pacman -S fcitx-sogoupinyin 

vim ~/.xprofile
# 添加以下内容
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
```
重启，需进入`Fcitx Configuration`添加`Pinyin`,注意复选框的勾去掉

5. 安装`nodejs`, `jdk`, `jetbrains`, `yarn`, `chrome`, `code`

6. 安装 `MariaDb`
```
sudo pacman -S mysql
sudo mysqld --initialize --user=mysql --basedir=/usr --datadir=/var/lib/mysql
sudo systemctl start mysqld # 启动MariaDb
sudo mysqladmin -u root password "123456" -p # 为root、用户添加密码
sudo systemctl enable mysqld # 设置mariaDb开机自启
mysql -uroot -p # 输入设置的的密码，登录数据库
# localhost 可能无法连接，使用127.0.0.1即可
```

7. 链接VPN `fmt1.link.ac.cn`，同步chrome账户

8. 生成 `ssh` 密钥对
```
ssh-keygen -t rsa -b 10240
```

9. `github`, `coding` 添加公钥

10. 安装 `Navicat`

11. 配置 `ssh`
```
vim ~/.ssh/config
```
输入以下内容
```
Host free
User yyy
Hostname 121.40.244.xxx
```
复制公钥
```
ssh-copy-id -i ~/.ssh/id_rsa.pub free
```
12. 安装vmware
```
sudo pacman -S fuse2 gtkmm  pcsclite libcanberra
sudo pacman -S linux-headers
sudo pacman -S ncurses5-compat-libs
sudo pacman -S vmware-workstation
sudo systemctl enable vmware-networks.service  vmware-usbarbitrator.service vmware-hostd.service
sudo systemctl start  vmware-networks.service  vmware-usbarbitrator.service vmware-hostd.service
# 重启4or
sudo modprobe -a vmw_vmci vmmon
```
13. 微信安装
```

```
14. 开机自动挂载硬盘
```
sudo blkid
sudo vim /etc/fstab
#模板 UUID=3b8a9a2b-f9a0-4117-866b-554afca8a469 /run/media/herbluo/2t ext4 defaults,noatime 0 0
sudo mount -a
```

# Linux
```shell
nc -uz 192.168.1.30 6379 #某端口是否开放

nmap -sS192.168.1.0/24 #局域网主机及端口
```

# CentOS

查看`centos`版本 

`rpm -q centos-release`

安装mysql 

1. 官网选择`MySQL Yum Repository`
2. 根据`centos`版本选择对应的`Red Hat Enterprise Linux`版本。
3. 点入链接后复制对应的地址，使用`wget`下载
4. 
```
rpm -ivh mysql[tab]
yum install -y mysql-community-server
systemctl start mysqld.service
systemctl status mysqld.service

cat /var/log/mysqld.log #查看默认密码
mysql -uroot -p
alter user 'root'@'localhost' identified by 'mi_ma' #可与默认密码相同
```


