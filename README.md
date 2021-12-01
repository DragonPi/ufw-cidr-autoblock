# ufw-cidr-autoblock

This tool automatically creates and applies, in conjunction with ufw, firewall block rules based CIDR lists (GEO-IP block)

Currently only IPv4 CIDR blocks are implemented.

The tool also fetches IP's from GitHub meta endpoint and stores those in a SQLite database.

The SQLite database is also used for storing explicit exclusions/inclusions of CIDR zones that one would like to allow/block.

## ToDo

* ~~Implement version~~ ---> Reports the applications version
* Implement apply ---------> Apply the rules to the firewall
* Implement revert --------> Revert to previous ruleset
* Implement block/unblock -> Block/unblock individual zones/countries on the fly
* Implement reset ---------> Reset all the rules
* Implement report --------> Report about all the blocked countries/zones
* Implement IPv6 ----------> Currently only IPv4 supported

## Credits

Thanks go out to <http://ipverse.net> who provide address block lists aggregated by country.

## JSON file layout

```json
{
    "_comment": "This is how a json file could look with exclusions/inclusions",
    "manual_entries": {
        "bad": [
                "192.168.200.0/22",
                "192.168.0.0/22",
                "10.0.1.0/22",
                "192.168.5.0/22"
        ]

    },
    "automatic_entries": {
        "example": {
            "foo": [
                "192.168.252.0/22",
                "192.168.10.0/22"
            ],
            "bar": [
                "10.0.0.0/22",
                "192.168.1.0/22"
            ]
        }
    }
}
```
