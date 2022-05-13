* machine - a virtual machine
* box - a virtual machine image
* base box - used to start the VM from the clean state

[Install](https://www.vagrantup.com/downloads) Vagrant

* on Ubuntu 20.04.4 LTS (focal) I had to remove the apt package and download the binary to get the latest version 

List boxes available locally

    vagrant box list

Initialize vagrant environment

    mkdir ubuntu-focal64
    cd ubuntu-focal64
    vagrant init ubuntu/ubuntu-focal64
    
* `Vagrantfile` is created, you should commit it to version control

Start vagrant environment

    vagrant up

Check machines status

    vagrant status
    
Check machines SSH configuration

    vagrant ssh-config

Ssh to a machine

    vagrant ssh

Clean up

    # save VM's state; fastest to start again; eats most diskspace (hard disk + saved state of RAM)
    vagrant suspend

    # graceful shutdown; slower to start again, still eats disk space (hard disk)
    vagrant halt

    # power down and remove all of the guest hard disks; even slower to
    # start again (reimport of the base box and reprovisioning)
    vagrant destroy [-f]

Show status of all vagrant environments on the host (independent of the directory you're in)

    vagrant global-status [--prune]
    
Multi-machine Vagrantfile

```
# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/focal64"

  config.vm.define "master" do |machine|
    machine.vm.hostname = "master"
    machine.vm.network "private_network", type: "dhcp"
  end

  config.vm.define "worker" do |machine|
    machine.vm.hostname = "worker"
    machine.vm.network "private_network", type: "dhcp"
  end

  # ping worker.localhost
  config.vm.provision "allow_guest_host_resolution", type: "shell",
    inline: <<-SHELL
      apt-get update
      apt-get install -y avahi-daemon libnss-mdns
    SHELL
end
```

To share a folder from the host on the guest, add following to `Vagrantfile`

    config.vm.synced_folder "../../eset-repos", "/shared/eset-repos",
      owner: "jreisinger", group: "jreisinger"

For using Vagrant in a proxy (corporate :-)) environment run [setup_proxy](https://github.com/jreisinger/dotfiles/blob/master/bin/setup_proxy) or these

    # Note the whitespace to prevent saving the credentials in the bash history.
    # You need something like HISTCONTROL=ignoreboth in ~/.bashrc though.
     export http_proxy='http://user:password@host:port'
     export https_proxy='http://user:password@host:port'
    vagrant plugin install vagrant-proxyconf
    
     export VAGRANT_HTTP_PROXY='http://user:password@host:port'
     export VAGRANT_HTTPS_PROXY='http://user:password@host:port'
    vagrant up

Resources

* https://docs.vagrantup.com
* ~~https://sysadmincasts.com/episodes/4-vagrant~~
