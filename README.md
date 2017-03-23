Snap Revision Provider
=====================

This utility queries the snap store API in search of a snap revision in a\
beta, edge, stable and candidate channels. It returns the results for all
architectures.

Usage
====

```
Usage of ./snap-revision-provider:
  -snap string
        Snap to query (default "bluez")
```

Output
=====

Similar to rmadison:

```
% ./snap-revision-provider -snap pulseaudio
 pulseaudio | stable    | 8.0-3 | 12 | armhf
 pulseaudio | stable    | 8.0-3 | 10 | i386
 pulseaudio | stable    | 8.0-3 | 9  | amd64
 pulseaudio | stable    | 8.0-3 | 11 | arm64
 pulseaudio | candidate | 8.0-3 | 9  | amd64
 pulseaudio | candidate | 8.0-3 | 10 | i386
 pulseaudio | candidate | 8.0-3 | 12 | armhf
 pulseaudio | candidate | 8.0-3 | 11 | arm64
 pulseaudio | beta      | 8.0-4 | 28 | amd64
 pulseaudio | beta      | 8.0-4 | 31 | i386
 pulseaudio | beta      | 8.0-4 | 30 | armhf
 pulseaudio | beta      | 8.0-4 | 29 | arm64
 pulseaudio | edge      | 8.0-4 | 39 | armhf
 pulseaudio | edge      | 8.0-4 | 37 | i386
 pulseaudio | edge      | 8.0-4 | 36 | amd64
 pulseaudio | edge      | 8.0-4 | 38 | arm64
```
