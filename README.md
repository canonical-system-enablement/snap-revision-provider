Snap Revision Provider
=====================

This utility queries the snap store API in search of a snap revision in a 
particular channel. It returns the results for each architecture.

Usage
====

```
Usage of ./snap-revision-provider:
  -channel string
        Channel from which the snap version is queried (default "candidate")
  -snap string
        Snap to query (default "bluez")
```

Output
=====

Similar to rmadison:

```
% ./snap-revision-provider -channel beta -snap bluez
 bluez | beta | 10 | i386
 bluez | beta | 12 | armhf
 bluez | beta | 15 | amd64
 bluez | beta | 13 | arm64

```
