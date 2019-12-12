(Up-to-date [source](https://github.com/jreisinger/blog/blob/master/posts/ssh-tunnel.md) of this post.)

Tunneling is the process of packaging and transporting one network connection using another one.

## Forwarding remote port (firewall tunneling via SSH)

We want to allow the tech access the incomp (intranet) host from the outcomp.sk (Internet) host:

![SSH Tunneling](https://raw.github.com/jreisinger/blog/master/files/ssh_tunneling.png)

1) Redirect the port 2222 on outcomp.sk to port 22 on incomp:

        incomp:~$ ssh -R 2222:localhost:22 user@outcomp.sk
        outcomp.sk:~$ while [ 1 ]; do date; sleep 300; done  # to keep the connection open
    
2) Connect to intranet host:

        outcomp.sk:~$ ssh -p 2222 root@localhost
        
Note that you need to have SSH server running on both incomp and outcomp.sk.

## Forwarding local port

We want to connect to a home router's web interface (192.168.1.1:80) to make some configuration changes. The home router is not accessible from Internet. However we can ssh (through port forwarding on the home router) to a host behind the home router (homebox.duckdns.org):

    localhost$ ssh -L 8080:192.168.1.1:80 homebox.duckdns.org

Nov we enter `http://localhost:8080` into the browser on localhost.

---

We want to connect to a remote database running on dbserver but it is configured to allow connections only from localhost (127.0.0.1). We use port 3307 on the client because the default 3306 port is already being used (you are running MySQL server on the client).

    client:~$ ssh -L 3307:localhost:3306 root@dbserver
    client:~$ mysql -u root -p dbname -P 3307

## See also

 * [ssh -R 2222:localhost:22 user@host](http://explainshell.com/explain?cmd=ssh+-R+2222%3Alocalhost%3A22+user%40host)
 * [ssh -L 3307:localhost:3306 user@host](http://explainshell.com/explain?cmd=ssh+-L+3307%3Alocalhost%3A3306+user%40host)
 * SSH, The Secure Shell: The Definitive Guide, 2nd Ed. (2005) - Connecting Through a Gateway Host
