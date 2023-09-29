# Use the official Nginx base image
FROM nginx

# Copy the custom Nginx configuration file to the container
COPY default.conf /etc/nginx/conf.d/default.conf

# Expose port 80 for HTTP
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]