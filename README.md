random_hosts
============

This tool quickly generates random domains and IP addresses in the syntax needed 
for /etc/hosts or C:\Windows\System32\drivers\etc\hosts files.

**Usage:**

    random_hosts -n 10 >> /etc/hosts

**Options:**

    -n [int] The number of IP/domain pairs to generate
    -t [com,net,org] The TLDs to randomly append to the domains generated
    -m [int] The minimum character length of the second-level domain to generate
    -M [int] The maximum character length of the second-level domain to generate

Why?!
-----
During the [great VAC debacle of 2014](http://www.reddit.com/r/Games/comments/1y1uuc/vac_now_reads_all_the_domains_you_have_visited/), 
I decided to do a little traffic analysis to try to confirm or disconfirm the 
report. Since Windows adds values in its hosts file to the DNS cache, I knew 
that I could quickly generate a lot of dummy entries. I hacked this up in Go 
so I could cross-compile it quickly.

Issues
------
A lot, probably. Notably, the IP addresses it generates don't use up the whole 
IP address space. To save time, I just limited each octet to Class A, B, or C 
values.
