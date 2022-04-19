#!/bin/bash
# 修改网卡名
PI(){
grep lv=centos /etc/default/grub
if [ $? = 0 ];then
	NAT=`sed -ri 's/^(GRUB_CMDLINE_LINUX=\"crashkernel=auto ) .* (rhgb quiet)/\1\2 net.ifnames=0 biosdevname=0/g' /etc/default/grub`
fi
}

# 静态IP
IP(){
	IP=`ip a |grep eth0 | sed -rn "2s/[^0-9]*([0-9.]+).*/\1/p"`
	 cat > /etc/sysconfig/network-scripts/ifcfg-eth0 << EOF
BOOTPROTO=static                                                                    
NAME=eth0
DEVICE=eth0
ONBOOT=yes
IPADDR=$IP
GATEWAY=10.0.0.2
PREFIX=24
DNS1=114.114.114.114
DNS2=8.8.8.8
EOF
	systemctl restart network
}

# 阿里源
YUM(){
	mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
	curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo
	yum clean all
	yum makecache

}

# 常用命令下载
CMD(){

	yum install -y wget vim glances atop ncdu iotop iftop nethogs \
		git lrzsz bash-completion gcc gcc-c++ autoconf automake make \
         zlib zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel libffi libffi-devel xz xz-devel	
}

# 修改主机命令提示
PS(){

cat > /etc/profile.d/env.sh << EOF
PS1="\[\e[1;32m\][\[\e[0m\]\t \[\e[1;33m\]\u\[\e[36m\]@\h\[\e[1;31m\] \W\[\e[1;32m\]]\[\e[0m\]\\\\$"
HISTTIMEFORMAT="%F %T"
HISTCONTROL=ignoreboth
EOF
}

# 调整时间
TIME(){
read -p "please input hostname: " name
hostnamectl set-hostname $name

timedatectl set-timezone Asia/Shanghai
timedatectl set-local-rtc 0
systemctl restart rsyslog

}

# 禁用 swap
SWAP(){
if [ "`egrep "vm.swappiness=1" /etc/sysctl.conf|wc -l`" == "0" ];then
        echo "vm.swappiness=1" >> /etc/sysctl.conf
        sysctl -p
else
        echo "禁止swap交换 设置成功或者之前已经设置过了"
fi
}


# 调整内核参数针对 K8S
K8S_conf(){

	cat > /etc/sysctl.d/kubernetes.conf << EOF
net.bridge.bridge-nf-call-iptables=1    #开启网桥模式
net.bridge.bridge-nf-call-ip6tables=1   #开启网桥模式
net.ipv4.ip_forward=1
net.ipv4.tcp_tw_recycle=0
vm.swappiness=0 # 禁止使用 swap 空间，只有当系统 OOM 时才允许使用它
vm.overcommit_memory=1 # 不检查物理内存是否够用
vm.panic_on_oom=0 # 开启 OOM  
fs.inotify.max_user_instances=8192
fs.inotify.max_user_watches=1048576
fs.file-max=52706963
fs.nr_open=52706963
net.ipv6.conf.all.disable_ipv6=1    #关闭IPV6的协议
net.netfilter.nf_conntrack_max=2310720
EOF

sysctl -p /etc/sysctl.d/kubernetes.conf 
}

# 关闭防火墙selinux
SELinux(){

	systemctl stop firewalld && systemctl disable firewalld
	sed -i 's#SELINUX=enforcing#SELINUX=disabled#g' /etc/selinux/config
	setenforce 0

}


PS