# Use the official Elasticsearch image as the base image
FROM docker.elastic.co/elasticsearch/elasticsearch:8.10.4

# Install the ICU Analysis Plugin
RUN bin/elasticsearch-plugin install analysis-icu

COPY elasticsearch.yml /usr/share/elasticsearch/config/elasticsearch.yml
