# ufw-cidr-autoblock

This tool automatically creates and applies, in conjunction with ufw, firewall block rules based CIDR lists (GEO-IP block)

Currently only IPv4 CIDR blocks are implemented.

The tool also fetches IP's from GitHub meta endpoint and stores those in a SQLite database.

The SQLite database is also used for storing explicit exclusions/inclusions of CIDR zones that one would like to allow/block.

## Credits

Thanks go out to <http://ipverse.net> who provide address block lists aggregated by country.

## ToDo

* Implement IPv6
