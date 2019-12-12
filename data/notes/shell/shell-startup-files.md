(Up-to-date <a href="https://github.com/jreisinger/blog/blob/master/posts/login-scripts.md">source</a> of this post.)

Bash Startup Files
------------------

When `bash` is started it runs a series of scripts to prepare the environment
for user. These scripts, for example, set the environment variables, create
command aliases, run programs.

There are two main types of a shell instance, interactive and noninteractive. However, noninteractive shells (such as those running shell scripts) usually don't read any startup files. To run a noninteractive shell as a login shell, to force it to run startup files, start the shell with the `-l` or `--login` option.

Interactive shell can be a:

 * login shell - a shell started by the `login` program or a remote login server such as SSH (it's startup files are a place for variables like `PATH`, `PS1` and startup programs like `umask`)
 * non-login shell - additional shell you run after you log in - ex. shell inside an X-based terminal (it's startup files are a place for aliases and functions)

... to find out run `echo $0 | perl -lne 'print /^\-/ ? "" : "non-", "login shell"'`

<table>
  <tr>
    <th></th>
    <th>Login shell</th>
    <th>Non-login shell</th>
  </tr>
  <tr>
    <th>Global config</th>
    <td><code>/etc/profile</code> (, <code>/etc/profile.d/</code>)</td>
    <td><code>/etc/bash.bashrc</code>, <code>/etc/bash/bashrc</code>, <code>/etc/bashrc</code></td>
  </tr>
  <tr>
    <th>User config</th>
    <td><code>~/.bash_profile</code>, <code>~/.bash_login</code>, <code>~/.profile</code> (only the first found is run) </td>
    <td><code>~/.bashrc</code></td>
  </tr>
</table>
 
Creating a symlink between `~/.bashrc` and `~/.bash_profile` will ensure that the same startup scripts run for both login and non-login sessions. Debian's `~/.profile` sources `~/.bashrc`, which has the same effect.

More

* [Prostredie shellu](http://wiki.reisinge.net/Linux1/ShellSkriptovanie/02ProstredieShellu) (in Slovak)
* [my dot-files](https://github.com/jreisinger/dot-files)
* ULSAH p. 189
* How Linux Works, 2nd, ch. 13
