# Check Web Application Availability

I wanted to see whether a web application is showing a decent "uptime", so I whipped up a small shell script - `check-web-app.sh`:

    #!/bin/bash
    # Usage: $0 <host> <port>

    HOST=$1
    PORT=$2

    function get_resp {
        # Get the HTTP response code
        RESP=$(echo -e "GET / HTTP/1.0\n\n" | nc $HOST $PORT | head -1)
        echo "$TIME;$PORT_STATUS;$RESP"
    }

    while [ 1 -eq 1 ]; do
        # Get actual timestamp
        TIME=$(date "+%F_%T")

        # Get port status
        netcat -w 5 -z $HOST $PORT
        if [ $? -eq 0 ]; then
            PORT_STATUS='open'
        else
            PORT_STATUS='closed'
        fi

        if [ $PORT_STATUS == 'open' ]; then
            get_resp
        else
            echo "$TIME;$PORT_STATUS;n/a"
        fi

        # sleep for random time from 1 to 10 sec
        sleep $[ ( $RANDOM % 10 )  + 1 ]
    done

It uses the "TCP/IP swiss army knife" called `netcat`. First it checks whether the port is open. If so, it gets the first line of HTTP response header. I run it like this and then graph the output in a spreadsheet application:

    ./check-web-app.sh www.google.com 80 >> google.csv &
