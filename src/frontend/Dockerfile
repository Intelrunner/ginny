# Use an official Python runtime as a parent image
FROM python:3.10-slim

# Set environment variables
ENV PYTHONDONTWRITEBYTECODE=1  
# Prevents Python from writing .pyc files
ENV PYTHONUNBUFFERED=1       
# Ensures stdout/stderr logs are printed immediately

# Set the working directory inside the container
WORKDIR /app

# Copy requirements.txt to the container
COPY requirements.txt /app/

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code into the container
COPY . /app/

# Expose the port your Flask app runs on
EXPOSE 8080

# Command to run the Flask app
CMD ["python", "main.py"]
