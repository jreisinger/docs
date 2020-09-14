What does a socket state `TIME_WAIT`, that can be seen in command like `netstat`, mean?

```
$ netstat -tulpan
<...>
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 172.31.11.100:58550     52.59.212.69:80         TIME_WAIT   -
<...>
```

TCP guarantees the reliability by sending back Acknowledgment (ACK) packets back for a single or a range of packets received from the peer. This goes same for the control signals such as termination request/response. [RFC 793](http://tools.ietf.org/html/rfc793) defines the TIME-WAIT state to be as follows:

> TIME-WAIT - represents waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request.

[source](https://serverfault.com/a/329846/99555)
