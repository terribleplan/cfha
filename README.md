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

```js
{
  "hosts": [ //an array of hosts to check
    {
      "host": "192.168.1.1", //the IP address of the host
      "type": "http", //the type of check to run, can be either http or https
      "options": {} //optional additional parameters to the check module
    },
    {
      "host": "192.168.1.2",
      "type": "https",
      "options": {
        "hostname": "lb-check-hostname" //this is the only currently supported parameter, and will be sent as the "host" header of an http(s) request
      }
    }
  ],
  "cloudflare": { //cloudflare configuration
    "email": "cfemail@example.com",
    "apiKey": "CF_API_KEY", //get this from your cloudflare profile
    "domain": "example.com", //the domain in cloudflare you will be operating on
    "name": "lb.example.com", //the full dns record you wish to edit
    "ttl": "1" //1 is automatic, otherwise see the cloudflare documentation on TTL
  },
  "interval": 1 //how often (in seconds) to ping the server, note that this is a number, not a string
}
```

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
