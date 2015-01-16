# geomys
Dynamic Routing Protocol Suite built exclusively in GO. The intent is to
bring the established and mature routing protocols into line with current
development and automation practices.

Geomys operates on a few key principles:
- Everything configurable via REST. 
- No reload necessary on config change - neighbors survive unless you screw up
- Each daemon is stateless, and assumes no config until instructed otherwise on each startup

# Daemons
- geomysd: Watchdog service for starting daemons and crash handling. Provides
    REST API to other processes
- georipd: Routing Information Protocol Daemon
- geoospfd: Open Shortest Path First Daemon
- geoospf3d: Open Shortest Path First v3 Daemon
- geobgpd: Bridge Gateway Protocol Daemon

# Utilities
- geo-config: Utility for configuring Geomys.
- geo-state: A utility for showing the running state of the daemon (neighbor tables, etc)
- geo-debug: A utility for displaying debug information about running threads
