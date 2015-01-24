#!/bin/bash
set -euo pipefail # Unofficial bash strict mode
IFS=$'\n\t'

# This script preps the virtual machine for Geomys operations

# Install dependencies
sudo apt-get -y install git gccgo-go tmux

# Kill all running geomys daemons
kill $(ps aux | grep '[g]eo' | awk '{print $2}')
# http://stackoverflow.com/questions/3510673/find-and-kill-a-process-in-one-line-using-bash-and-regex

# Remove existing geomys directory in case of a re-provision
rm -rf /home/vagrant/geomys
mkdir /home/vagrant/geomys
cp -r /vagrant/* /home/vagrant/geomys


# Configure Go
export GOPATH=/home/vagrant/Code/GO
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOPATH/bin

grep -q -F 'GOPATH' /home/vagrant/.bashrc ||  sudo bash -c "echo -e 'export GOPATH=/home/vagrant/Code/GO\nexport GOBIN=$GOPATH/bin' >> /home/vagrant/.bashrc"
grep -q -F 'GOPATH' /root/.bashrc ||  sudo bash -c "echo -e 'export GOPATH=/home/vagrant/Code/GO\nexport GOBIN=$GOPATH/bin' >> /root/.bashrc"

# Install GO dependencies
go get github.com/vishvananda/netlink
go get github.com/olekukonko/tablewriter

cd /home/vagrant/geomys/geomys/georipd
go build georipd.go
go install georipd.go

cd /home/vagrant/geomys/tools/geo-show
go build geo-show.go
go install geo-show.go


# tmux new \; split-window -h