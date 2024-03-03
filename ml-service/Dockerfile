# Use the official Python image as a base
FROM python:3.9-slim-buster

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files into the container
COPY . .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt