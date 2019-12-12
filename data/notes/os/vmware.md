Terminology
-----------

Hosts and clusters

![VMware terms](https://raw.github.com/jreisinger/blog/master/files/vmware_terms.png)

* vSphere Client - thick client for managing vCenter servers
* vCenter Server - container for Datacenters
* Datacenter - container for Clusters and Hosts. Large companies might use
    datacenters to represent organizational units.
* Cluster - group of hosts
* Host - computer that uses virtualization software (ESX or ESXi) to run VMs

Networking
* vSphere Distributed Switch - acts as a single virtual switch across all
    associated hosts

Storage
* Datastore - a logical container that holds VMs' files (like `*.vmdk`, `*.log`,
    `*.nvram`). Can exist on different types of physical storage (local disks,
    iSCSI, Fibre Channel SAN, NFS). Can be VMFS-based or NFS-based. Can be
    created by formatting LUNs or by mounting NFS volumes to an existing host.
* More: https://pubs.vmware.com/vsphere-51/topic/com.vmware.vcli.examples.doc/cli_manage_storage.6.1.html
