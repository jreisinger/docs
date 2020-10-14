# New Linux Server

Things I do after installing a fresh Linux machine.

## Mandatory

 1. Upgrade all packages. Ex. `pacman -Syu`.

 1. Disable root logging in with password, allow only login with ssh key (so scripts using ssh keys can still login). Set `PermitRootLogin prohibit-password` (or `PermitRootLogin without-password`) in `/etc/ssh/sshd_config` and restart `sshd`. `PermitRootLogin prohibit-password` is default in newest versions of OpenSSH.

 1. Make sure [no unnecessary services](https://metacpan.org/pod/App::Monport) are running, ex.:

        service nfs-common stop
        service portmap stop
        update-rc.d nfs-common remove
        update-rc.d portmap remove

 1. Make sure your system keeps exact time. Ex. `aptitude install ntp`. To check the [time is synchronized](https://wiki.archlinux.org/index.php/Systemd-timesyncd) on Arch Linux: `timedatectl status`.

## Optional

 * Install `fail2ban` to block ssh brute-force crackers. Ex. `aptitude install fail2ban`

 * Setup [firewall](https://github.com/jreisinger/varia/blob/master/iptables.sh) to increase network security. Restart `fail2ban` after you install the firewall rules.
 
 * Load some [personalization files and tools](https://github.com/jreisinger/dotfiles).

 * Setup `sudo` to increase security - add the following to `/etc/sudoers`:

        # User privilege specification
        root    ALL=(ALL) ALL
        jbond   ALL=(ALL) ALL

 * Set your timezone. Ex. to UTC:

        $ sudo cp -p /etc/localtime{,.orig}  # don't bother viewing it, it's a binary file
        $ sudo ln -sf /usr/share/zoneinfo/UTC /etc/localtime
        $ date
        Wed Feb 17 08:35:29 UTC 2010

### Perl stuff

 * Install `cpanminus` to have a nice installer of [CPAN](https://metacpan.org/) modules:
 
        sudo apt-get install cpanminus  # or
        sudo cpan App::cpanminus        # or
        curl -L http://cpanmin.us | perl - --sudo App::cpanminus

 * Install `Module::Starter` used for building Perl distributions:

        aptitude install make
        cpanm Test::More  ## upgrading Test::More
        cpanm Module::Starter

 * Upgrade all CPAN modules ([source](http://stackoverflow.com/questions/3727795/how-do-i-update-all-my-cpan-module-to-their-latest-versions)) - can take some time (just to get a list of outdated modules [packages]: `cpan-outdated [-p]`):

        aptitude install gcc
        cpanm App::cpanoutdated
        cpan-outdated -p | cpanm

## More

* http://plusbryan.com/my-first-5-minutes-on-a-server-or-essential-security-for-linux-servers
