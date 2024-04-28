# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.define "ubuntu18" do |server|
    server.vm.box = "reclusive/bionic64-salt"
    server.vm.box_version = "0.0.2023-04-09-1907"
    server.ssh.private_key_path = ["C:/Users/Alexander/.vagrant.d/insecure_private_key","C:/Users/Alexander/.vagrant.d/20170926_vagrant_private_key"]
  end  

  config.vm.define "ubuntu20" do |server|
    config.vm.box = "reclusive/focal64"
    config.vm.box_version = "0.0.2023-08-03-1835"
    server.ssh.private_key_path = ["C:/Users/Alexander/.vagrant.d/insecure_private_key","C:/Users/Alexander/.vagrant.d/20170926_vagrant_private_key"]

  end  
  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
  # documentation for more information about their specific syntax and use.
  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    echo "vagrant" | su --login vagrant
    wget -q https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
    pwd
    ls -la /usr/local/
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local/ -xzf go1.20.7.linux-amd64.tar.gz
    ls -la /usr/local/
    # There is an issue in making these binaries available on the PATH
    # When updating /etc/environment I cannot download anything
    echo "PATH=$PATH:/usr/local/go/bin" | sudo tee /etc/environment
    source /etc/environment
    echo PATH="$PATH:$(go env GOPATH)/bin" | sudo tee /etc/environment
    source /etc/environment
    go version
    # For the VSCode go extension to work we need gopls
    go install -v golang.org/x/tools/gopls@latest
    # Ginkgo needs to be installed for the ginkgo extension to work
    go install github.com/onsi/ginkgo/v2/ginkgo@v2.12.0
    go install github.com/maxbrunsfeld/counterfeiter/v6
    ls -l /home/vagrant/
    ls -l /home/vagrant/go/bin/
    whoami
  SHELL
end