#!/bin/bash

. /opt/prometheus/prometheus.env

start() {
  exec "$EXPORTER_EXE" $EXPORTER_OPTS 
}

stop() {
  echo "TODO - add node_exporter cleanup tasks here"
}

case $1 in
  start|stop) "$1" ;;
esac
