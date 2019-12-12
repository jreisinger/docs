# Store Nmap Results to Excel File

**1.** Run `nmap`, e.g. (to select random IPs use
[rand_lines.pl](https://github.com/jreisinger/varia/blob/master/rand_lines.pl)):

    cat ips.txt | grep -v '^#' | nmap -PN -sV -oA scan_results -iL -

  * `-PN` -- scan also "non-pingable" hosts (in case you're scanning many hosts it can take a while)
  * `-oA` -- store output to `scan_results.xml` and two other formats

Sample `ips.txt`:

    # host1
    192.168.1.11
    # host2
    192.168.1.19

**2.** Run [conversion
script](https://github.com/jreisinger/blog/blob/master/code/nmap_xml2csv.pl):

    perl nmap_xml2csv.pl scan_results.xml > scan_results.csv

**3.** Import `scan_results.csv` to Excel (`Data => From Text`) and edit as
needed

Sample output:

![Output](https://raw.github.com/jreisinger/blog/master/files/nmap2xls.jpg)
