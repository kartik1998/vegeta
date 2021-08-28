# Vegeta ![go reference](https://img.shields.io/badge/go-reference-blue) ![](https://img.shields.io/badge/process-memory%20tracking-green)

Vegeta is a system resource usage tracking tool built to regularly take snapshots of the memory and CPU load of one or more running processes, so as to dynamically build up a profile of their usage of system resources.

<p align="center"> <img src = "./assets/vlogo.jpeg" height=600 /> </p>

# Introduction
Vegeta supports multiple modes. You can set the "delay" after which the snapshot of the processes is taken by using the "-d" or "--delay" flag. You can also use the "-m" flag and log only the top m processes with maximum memory usage. *Support for process id tracing and graph creation to be added later*. In either case, the monitoring of the system resource usage is based on repeated calls to the system command [ps](https://en.wikipedia.org/wiki/Ps_(Unix))

# Prerequisites
Given that vegeta is basically wrapper and parser of "ps" output, naturally the second-most important precondition for vegeta to work on your system is to have the "ps" command available. This is almost certainly true for all POSIX or mostly POSIX-compliant systems, including various flavors of UNIX, Linux's, Apple Mac OS X's etc.

Apart from requiring "ps" or something similar on your system, vegeta has no other dependencies whatsoever: it is a single self-contained pure go pkg and uses nothing but the standard Python libraries.
