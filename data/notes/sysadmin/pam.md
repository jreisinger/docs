# PAM and NSS
###### linux

## Pluggable Authentication Modules

* flexible, modular means of configuring authentication
* sits between the login programs (text mode `login`, graphical XDM, `su`) and account DB
* `/etc/pam.d` contains configuration files for every program that uses PAM (`/etc/pam.d/login`, `/etc/pam.d/gdm`)
* files beginning with `common`, `system` or `config` are common for several services
* _PAM stack_ -- set of modules called to perform a specific task
 * when a program calls a PAM stack, each of the modules in the stack is executed in sequence
 * each module can return a success or a failure code

Each config file consists of one or more PAM stacks (`auth` stack, `session` stack) -- `/etc/pam.d/login`:

    # management_group control_flag module [options]
    auth       optional   pam_faildelay.so  delay=3000000
    auth [success=ok new_authtok_reqd=ok ignore=ignore user_unknown=bad default=die] pam_securetty.so
    auth       requisite  pam_nologin.so
    session [success=ok ignore=ignore module_unknown=ignore default=bad] pam_selinux.so close
    session       required   pam_env.so readenv=1
    session       required   pam_env.so readenv=1 envfile=/etc/default/locale
    @include common-auth
    auth       optional   pam_group.so
    session    required   pam_limits.so
    session    optional   pam_lastlog.so
    session    optional   pam_motd.so
    session    optional   pam_mail.so standard
    @include common-account
    @include common-session
    @include common-password
    session [success=ok ignore=ignore module_unknown=ignore default=bad] pam_selinux.so open

**management groups**

* account -- account mngt. based on other things than authentication (time of day, available system resources)
* auth -- authentication (username/password)
* password -- password (or other authentication type) updates
* session -- sets up and cleans up user's session

**control flags** -- determine how the success (`0`) or failure (`1`) of the module will affect stack execution

* requisite
* required
* sufficient
* optional

Sample mini-stack:

    auth required   pam_unix.so try_first_pass
    auth sufficient pam_ldap.so try_first_pass

possible outcomes:
<table>
  <tr>
    <th></th>
    <th>pam_ldap.so 0</th>
    <th>pam_ldap.so 1</th>
  </tr>
  <tr>
    <th>pam_unix.so 0</th>
    <td>Stack 0</td>
    <td>Stack 0</td>
  </tr>
  <tr>
    <th>pam_unix.so 1</th>
    <td>Stack 1</td>
    <td>Stack 1</td>
  </tr>
</table>

**common PAM modules**

* `pam_unix.so` -- traditional Unix authentication based on `/etc/password`, `/etc/shadow`
* `pam_listfile.so` -- searches the specified file for rules to allow/deny access
* `pam_cracklib.so` -- checks a password strength
* `pam_limits.so` -- sets login session limits on MEM, CPU (`/etc/security/limits.conf`)

## Name Service Switch

provides system tools with lists of users and groups, maps UID to username, identifies users' home dirs, etc.

`/etc/nsswitch.conf`:

    passwd:         compat
    group:          compat
    shadow:         compat

    hosts:          files dns
    networks:       files

    protocols:      db files
    services:       db files
    ethers:         db files
    rpc:            db files

    netgroup:       nis

the order is important

- - -

Source

* Roderick W. Smith: LPIC-2 (2011)
