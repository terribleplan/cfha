CFHA
====
CloudFlare High Availability monitor

CFHA is a simple daemon that will continually monitor servers you specify and
add/remove them from a shared "A" record in cloudflare.

Configuration
=============
Configuring CFHA is done through a JSON file, which should be named config.json,
and placed in whatever directory you will run CFHA from.

Example configuration file:

(the configuration format is currently in flux, the example will be updated when it is stable. for now just look at `core/config` and build a json object to match `Config`)


Limitations
===========
CFHA is limited in a few ways, due both to the author, and due to limitations in cloudflare's API / how DNS works.

1. CFHA can only operate on `A` records since it is invalid to have multiple values for CNAME records.
2. You must specify each host by IP address because CFHA does not resolve names to IP addresses yet. (And we need to have IP addresses for A records). An interesting feature of implementing this is that you could create a pool A record and CFHA would monitor all of those via a single host entry pointing to that record, but it would require a lot more logic.
3. CFHA does not operate on IPv6. This is something that is possible (and maybe even simple) to implement, but has not yet been worked on.
4. CFHA is meant to operate on load balancers, and as such only checks for a 200 exit status on a host. You should add a name-based virtual host to nginx that returns 200 on each monitored server, and set the hostname in the check options.

License
=======
CFHA is made available under an ISC license. Anyone is welcome to make contributions with the understanding that they will be released under that same license.
