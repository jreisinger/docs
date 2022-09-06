## Turning off WiFi at night on MikroTik RouterOS

1) In System -> Scripts add DisableWLAN:

![image](https://user-images.githubusercontent.com/1047259/134005654-9c38b1a5-6554-4e88-97a2-129a2bfca832.png)

2) Do anological steps for EnableWLAN.

3) In System -> Scheduler add:

![image](https://user-images.githubusercontent.com/1047259/134005817-16cb1ad3-d60a-4e8f-a459-0a8de081393d.png)

4) Do anological steps for EnableWLAN.

NOTE: To disable/enable ethernet port 4:: `interface ethernet disable ether4` or `interface ethernet enable ether4`

Source: https://goyoambrosio.com/2017/05/how-to-turn-off-wifi-at-night-in-mikrotik-routeros/
