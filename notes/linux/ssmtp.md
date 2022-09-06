I've found out that one of the simplest ways how to send emails from scripts running on your workstation is <a href="https://packages.debian.org/wheezy/ssmtp">ssmtp</a> (my ex-colleague showed it to me). It's very easy to install and setup. Basically you just edit one or two lines in <code>ssmtp.conf</code>. However there's a caveat; the ssmtp does not consider <code>aliases</code> when sending email. So the cron was trying to send email to non existent email address like root@mybox.local.domain. To <a href="http://raftaman.net/?p=591">fix</a> this problem you have to do the aliasing in the <code>mail</code> program by adding lines like these into <code>/etc/mail.rc</code>:<br />
<br />
<pre class="brush:plain">alias root root&ltusername@company.com&gt
alias postmaster postmaster&ltusername@company.com&gt
alias username username&ltusername@company.com&gt
</pre><br />
I also put this variable into my crontab<br />
<pre class="brush:plain">MAILTO=username@company.com
</pre>
