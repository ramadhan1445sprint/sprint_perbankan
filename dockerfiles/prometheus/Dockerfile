# Use the official Prometheus base image
FROM prom/prometheus:latest

# Copy custom configuration files
COPY ./prometheus_deploy.yml /etc/prometheus/prometheus.yml

# Expose port 9090 to the outside worldss
EXPOSE 9090

# Command to run Prometheus
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/prometheus", "--web.console.templates=/etc/prometheus/consoles", "--web.console.libraries=/etc/prometheus/console_libraries"]