# Use the official Grafana base image
FROM grafana/grafana:latest

# Copy provisioning files
COPY ./grafana-configs/dashboards/main.yml /etc/grafana/provisioning/dashboards
COPY ./grafana-configs/datasources/prometheus_datasource.yaml /etc/grafana/provisioning/datasources/prometheus_datasource.yaml
COPY ./grafana/dashboards/dashboard-perbankan-1711812579563.json /etc/grafana/main-dashboards/dashboard-perbankan-1711812579563.json

ENV GF_SECURITY_ADMIN_PASSWORD=admin
ENV GF_USERS_ALLOW_SIGN_UP=false

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run Grafana
CMD ["grafana-server", "-config=/etc/grafana/grafana.ini", "-homepath=/usr/share/grafana", "cfg:default.paths.logs=/var/log/grafana", "cfg:default.paths.data=/var/lib/grafana", "cfg:default.paths.plugins=/var/lib/grafana/plugins"]