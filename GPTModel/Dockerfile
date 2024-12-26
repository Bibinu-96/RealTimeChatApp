# Use Python Alpine base image
FROM python:3.11-alpine

# Set environment variables to improve performance
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

# Set the working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache --virtual .build-deps gcc musl-dev libffi-dev \
    && apk add --no-cache libffi libstdc++ \
    && pip install --no-cache-dir --upgrade pip

# Copy requirements.txt first to install dependencies
COPY requirements.txt /app/

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Remove build dependencies to reduce image size
RUN apk del .build-deps

# Copy application code after installing dependencies
COPY . /app/

# Expose Flask server port
EXPOSE 8000

# Run the application
CMD ["gunicorn", "-w", "1", "-b", "0.0.0.0:8000", "--timeout", "120", "app:app"]
