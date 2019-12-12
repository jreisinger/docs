(Up-to-date [source](https://github.com/jreisinger/blog/blob/master/posts/clone-and-resize-kvm-virtual-machine.md) of this post.)

I needed to upgrade (from Squeeze to Wheezy) some important virtual servers. As
I wanted a minimal impact of the upgrade, I chose this procedure:

1. Create identical copy of the server to upgrade
2. Upgrade the copy
3. Upgrade the server if everything worked ok with the copy

The servers to upgrade were virtual machines (VMs) running on KVM. I also
discovered that some servers needed more space because their disks had filled
up during upgrade. So disk resize was needed. The following steps did the task:

(1) Copy the image (.qcow2) and the configuration (.xml) files to some other
location. The image file should ideally be copied from a snapshot to avoid data
inconsistencies a running machine could create.

(2) Edit the following fields in the copied .xml file accordingly

    name
    uuid
    source dev    # make sure you enter the copied image path!
    mac address
    source bridge # change the VLAN to avoid IP address conflicts

(3) Boot the cloned VM and change the hostname and IP address by editing these files:

    /etc/network/interfaces
    /etc/hostname
    /etc/hosts

(4) Change back the VLAN and shutdown the cloned VM

(5) Increase the disk size

    # convert the qcow image to a plain raw file
    qemu-img convert system.qcow -O raw system.raw
    
    # create a dummy file (filled with zeros) of the size of extra space you want to add to your image (here 1GB)
    dd if=/dev/zero of=zeros.raw bs=1024k count=1024
    
    # add your extra space to your raw system image without fear
    cat system.raw zeros.raw > big.raw
    
    # finally convert back your raw image to a qcow file not to waste space
    qemu-img convert big.raw -O qcow growed-system.qcow

(6) Boot the cloned VM and using cfdisk delete the old small partition and
create a new one with the free space

(7) Increase the filesystem using:

    e2fsck -f
    resize2fs

Make sure the VM's image file (.qcow) has the correct access rights, otherwise
your system might have disk related problems (I was bitten by this and got
helped by my nice colleague).
