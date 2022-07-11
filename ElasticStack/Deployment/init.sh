#!/bin/bash
SETFILE="/etc/security/limits.conf"
SYSFILE="/etc/sysctl.conf"
ulimit -n 1000000

ret=`sed -n '/soft.*nofile/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* soft nofile 1000000' $SETFILE
else
        sed -i '/soft.*nofile/c\* soft nofile 1000000' $SETFILE
fi

ret=`sed -n '/hard.*nofile/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* hard nofile 1000000' $SETFILE
else
        sed -i '/hard.*nofile/c\* hard nofile 1000000' $SETFILE
fi

ret=`sed -n '/soft.*nproc/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* soft nproc 1000000' $SETFILE
else
        sed -i '/soft.*nproc/c\* soft nproc 1000000' $SETFILE
fi

ret=`sed -n '/hard.*nproc/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* hard nproc 1000000' $SETFILE
else
        sed -i '/hard.*nproc/c\* hard nproc 1000000' $SETFILE
fi

ret=`sed -n '/soft.*memlock/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* soft memlock unlimited' $SETFILE
else
        sed -i '/soft.*memlock/c\* soft memlock unlimited' $SETFILE
fi

ret=`sed -n '/hard.*memlock/p' $SETFILE`
if [ ! "$ret" ]
then
        sed -i '/End of file/i\* hard memlock unlimited' $SETFILE
else
        sed -i '/hard.*memlock/c\* hard memlock unlimited' $SETFILE
fi

ret=`sed -n '/fs.file-max/p' $SYSFILE`
if [ ! "$ret" ]
then
        echo  'fs.file-max=1000000' >> $SYSFILE
else
        sed -i '/fs.file-max/cfs.file-max=1000000' $SYSFILE
fi

ret=`sed -n '/vm.swappiness/p' $SYSFILE`
if [ ! "$ret" ]
then
        echo  'vm.swappiness=1' >> $SYSFILE
else
        sed -i '/vm.swappiness/cvm.swappiness=1' $SYSFILE
fi

sysctl -p