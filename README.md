# gotracer-prometheus

## Synopsis

Basically a cron job runs gotracer-prometheus every 2 minutes.  The output is placed in /opt/prometheus/textfile_collector (TODO change /opt to /var for output data).

Any Prometheus installation accessing the node exporter will have data like example at /opt/prometheus//textfile_collector/epsolar.prom

## Install

Assumes Debian/Ubuntu linus installation... Copy files from install folder into the root file system of the target installation.

1) Requires binary for node_exporter to be put in /opt/prometheus/node_exporter directory.
2) Build src/gotracer-prometheus.go and copy binary to /usr/local/bin of target installation.
3) Need to run systemd commands to register node_exporter as service

TODO - add details on downloading node_exporter and building gotracer-prometheus

