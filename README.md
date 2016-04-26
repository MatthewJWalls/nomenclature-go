
Nomenclature-go
===============

Nomenclature is a web service written in Go which generates human readable names which
are guaranteed unique. Useful for naming servers, databases, etc.

This is a work in progress and is still missing key features for usability.

Usage:

    nomenclature -f1 prefixes.txt -f2 postfixes.txt

After which it will serve names on port 80
    
Todo:
* ~~name generation~~
* definable input lists
* persistence
* config management integration
