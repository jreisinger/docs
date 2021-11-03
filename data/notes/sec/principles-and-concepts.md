CIA triad - security goals

* Confidentiality - seeks to prevent unauthorized read access to data
* Integrity - seeks to prevent unauthorized write access to data
  * data integrity (modification of data in DB)
  * system integrity (malicious SW to open "back door" to OS) 
* Availability - ensures that information is available when needed

Least privilege

* concerns people and automated tools 
* should be able to access only what they need to do their job, and no more
* deny by default

Defense in depth

* acknowledgement that almost any security control can fail
* multiple layers of overlapping security controls

Think about what you need to protect (assets: VMs, containers, DBs) and who is most likely to cause problems (threat actors: criminals, hacktivists, script kiddies, inside attackers, state actors).

Figure out what needs to talk to what in your application. You should first secure places where line crosses a trust boundary:

![image](https://user-images.githubusercontent.com/1047259/138698724-4a6ecae8-fe54-4d45-b7a8-3b35dfab50e1.png)

Understand what areas you need to secure - this depends on the cloud model you are using and whether you are a consumer or provider:

![image](https://user-images.githubusercontent.com/1047259/138699080-24091008-c78f-48c1-bcc9-e9ac6afd0f8d.png)

Risk is something bad that could happen. Its level is based on its likelihood to happen and its impact. Know your risks (have at least a spreadsheet) and how you approach them:

* avoid the risk - turn off the system, benefits < risk
* mitigate the risk - apply some security measures
* transfer the risk - pay someone else to manage the risk (e.g. insurance)
* accept the risk - benefits > risk

Source: Practical Cloud Security (2019)
