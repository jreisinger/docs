Least privilege

* people and automated tools should be able to access only what they need to their job, and no more
* deny by default

Defense in depth

* acknowledgement that almost any security control can fail
* multiple layers of overlapping security controls

Think about what you need to protect (assets) and who is most likely to cause problems (threat actors): criminals, hacktivists (script kiddies), inside attackers, state actors.

Figure out what needs to talk to what in your application. You should secure first places where line crosses a trust boundary:

![image](https://user-images.githubusercontent.com/1047259/138698724-4a6ecae8-fe54-4d45-b7a8-3b35dfab50e1.png)

Understand what areas you need to secure - this can depend on the cloud model you are using:

![image](https://user-images.githubusercontent.com/1047259/138699080-24091008-c78f-48c1-bcc9-e9ac6afd0f8d.png)

Risk is something bad that could happen. Its level is based on its likelihood to happen and its impact. Know your risks (have at least a spreadsheet) and how you approach them:

* avoid the risk - turn off the system, benefits are less than the rist
* mitigate the risk - apply some securit measures
* transfer the risk - pay someone else to manage the rist (e.g. insurance)
* accept the risk - benefits are greater than the risk

Source: Practical Cloud Security (2019)
