# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
 
  config.vm.box = "trusty64"
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  
  config.vm.define "quagga" do |quagga|
    quagga.vm.host_name = "quagga"
    quagga.vm.network "private_network", ip: "192.168.40.30"
    quagga.vm.provision "ansible" do |ansible|
        ansible.playbook = "scripts/quagga.yml"
    end
  end

  config.vm.define "geomys" do |geomys|
    geomys.vm.host_name = "geomys"
    geomys.vm.network "private_network", ip: "192.168.40.31"
    geomys.vm.provision "shell", path: "scripts/geomysdev.sh"
  end
end