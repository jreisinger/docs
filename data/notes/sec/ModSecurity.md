*2019-10-01*

# ModSecurity

* a WAF engine (library, module) for Apache, Nginx, IIS

ModSecurity 3.0 has a new modular architecture, i.e. it's composed of:

1. [libmodsecurity](https://github.com/SpiderLabs/ModSecurity) (https://www.modsecurity.org) - core component containing the functionality and couple of rules ([docs](https://github.com/SpiderLabs/ModSecurity/wiki))
2. a connector that links libmodsecurity to the web server it is running with - [NGINX](https://github.com/SpiderLabs/ModSecurity-nginx), Apache HTTP Server, and IIS

Files

* `modsecurity.conf` - [ModSecurity directives](https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29#Configuration_Directives) v2.x (like `SecRuleEngine`, `SecRequestBodyAccess`, `SecAuditEngine`)
* `nginx.conf` - ModSecurity-Nginx [connector directives](https://github.com/SpiderLabs/ModSecurity-nginx#usage) (like `modsecurity on`, `modsecurity_rules_file /etc/nginx/modsecurity/nginx-modsecurity.conf`)

Tips

* audit logs are great for visibility but bad for performance - you can disable them via `SecAuditEngine off` (you still have the Nginx error logs)
* you should not inspect static content (like images) for performance reasons

More

* https://github.com/SpiderLabs/ModSecurity/wiki
* https://www.nginx.com/blog/compiling-and-installing-modsecurity-for-open-source-nginx/

## `SecRule` ModSecurity directive

```
SecRule VARIABLES   "OPERATOR"                "TRANSFORMATIONS,ACTIONS"
# E.g.
SecRule REQUEST_URI "@streq /index.php" "id:1,phase:1,t:lowercase,deny"
```

* VARIABLES - where to look (targets)
* OPERATOR - when to trigger a match
* TRANSFORMATIONS - how to normalize VARIABLES data
* ACTIONS - what to do when rule matches

More

* https://www.modsecurity.org/CRS/Documentation/making.html

# OWASP ModSecurity Core Rule Set (CRS)

* definitions of the malicious patterns (signatures, blacklist rules)
* documentation: https://github.com/SpiderLabs/OWASP-CRS-Documentation => https://www.modsecurity.org/CRS/Documentation/

Files

* `crs/setup.conf` - config file
* `crs/rules` - directory with rules (you should modify only `*EXCLUSION-RULES*`)

Tips

* should be used for all ModSecurity deployments
* to [tune](https://www.oreilly.com/ideas/how-to-tune-your-waf-installation-to-reduce-false-positives
) set a high anomaly threshold and progressively lower it

Paranoia levels (FP = false positive - a WAF blocking a valid request):

1. (default) basic security, minimal amount of false positives (FPs)
2. elevated security level, more rules, fair amount of FPs
3. online banking level security, specialized rules, more FPs
4. nuclear powerplant level security, insane rules, lots of FPs

You can configure rules via:

* excludes in `RESPONSE-999-EXCLUSION-RULES-AFTER-CRS.conf` or `REQUEST-900-EXCLUSION-RULES-BEFORE-CRS.conf`
* [whitelists](https://www.modsecurity.org/CRS/Documentation/exceptions.html#exceptions-versus-whitelist) (complicated, performance impact)
* paranoia levels in `crs/setup.conf`
* other configuration options in `crs/setup.conf` (like DOS protection, IP reputation, protocol enforcement)
* "includes" in `nginx-modsecurity.conf` (`Include /etc/nginx/modsecurity/crs/rules/*.conf`) - not supported by docs AFAIK

More

* https://coreruleset.org
* https://coreruleset.org/20171214/practical-ftw-testing-the-core-rule-set-or-any-other-waf/
* https://github.com/SpiderLabs/owasp-modsecurity-crs

# Attacks for testing WAF

```
curl 'https://$FQDN/?exec=/bin/bash'           # Remove Code Execution (RCE)
curl "https://$FQDN/?id=1'%20or%20'1'%20=%20'" # SQL Injection (SQLi)
curl 'https://$FQDN/?page=/etc/passwd'         # Local File Inclusion (LFI)
curl 'https://$FQDN/?<script>'                 # Cross Site Scripting (XSS)
```
