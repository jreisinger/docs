# KVM Virtual Machines Management

Virtualization seems to be really useful and I've seen it being used in all companies I've worked so far. I've spent more or less time working with WMWare and Xen but recently I've been working mainly with KVM. 

KVM uses the Linux kernel itself for the hypervisor role; memory management and scheduling are handled by the host's kernel, and guest machines are normal Linux processes which you can see and manage through standard commands like `top`, `ps`, and `kill`.

The management tools I've used are <a href="http://virt-manager.org/">virt-manager</a> (GUI) and <code>virsh</code> (command line). Here are the basic commands for managing virtual machines (VMs) under KVM.

<h2>virt-manager</h2>

This tool is fairly self-explanatory. The only thing I'll mention is how to get it running from a Windows machine via putty (see <a href="http://www.math.umn.edu/systems_guide/putty_xwin32.html">here</a> for more details):

 1. Install and start Xming - an X(11) server for Windows (Xming Server:0.0). You don't need to be admin.
 1. Setup putty to forward X connections: "Enable X11 forwarding" => "X display location: localhost:0"

## virsh

List VMs

    virsh -c qemu:///system list --all

Install VM

    virt-install \
                 --name vm01 \
                 --ram 256 \
                 --disk /images/vm/vm01.qcow2,size=10 \
                 --cdrom /images/iso/debian-6.0.6-amd64-netinst.iso \
                 --os-variant debiansqueeze \
                 --connect qemu:///system \
                 --virt-type kvm \
                 --network bridge=br0 \
                 --vnc \
                 --prompt
                 
* double check the first 5 parameters (`--name` to `--os-variant`)
* only `qcow2` disk supports [snapshotting](http://wiki.libvirt.org/page/VM_lifecycle#Taking_a_Snapshot_of_a_guest_domain)

Stop VM

    virsh -c qemu:///system shutdown vm01

.. or
    
    virsh -c qemu:///system destroy vm01
    
Start VM

    virsh -c qemu:///system create /path/to/host.xml

Remove VM

    virsh -c qemu:///system undefine vm01
    
More

* ULSAH 4th, p. 995
* [Guest Management](http://wiki.libvirt.org/page/Main_Page#Guest_Management) (libvirt)
* [Manage your virtual machines](https://help.ubuntu.com/community/KVM/Managing) (Ubuntu)
* [converting dist to qcow2](http://forums.fedoraforum.org/showthread.php?t=260126)
