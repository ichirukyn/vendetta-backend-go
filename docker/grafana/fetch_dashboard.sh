#!/bin/bash
mkdir -p /var/lib/grafana/dashboards

curl https://grafana.com/api/dashboards/10826/revisions/latest/download -o /var/lib/grafana/dashboards/go_metrics_dashboard.json
