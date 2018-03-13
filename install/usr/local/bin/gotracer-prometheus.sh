#!/bin/bash
OUTPUT_FILE=/opt/prometheus/textfile_collector/epsolar.prom
: > "$OUTPUT_FILE"
/usr/local/bin/gotracer-prometheus > "$OUTPUT_FILE".tmp
mv "$OUTPUT_FILE".tmp "$OUTPUT_FILE"

