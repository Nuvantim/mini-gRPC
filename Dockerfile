FROM alpine:latest

# Create required directories
RUN mkdir /source /app

# Set working directory for the project
WORKDIR /source

# Copy project files into the container
COPY . .

# Prepare .env file for running in /app
RUN mv .env.prod .env && \
    mv .env /app

# Move the built binary to the final directory
RUN mv ./bin/api /app/bin

# Set working directory to /app and run the application
WORKDIR /app
RUN chmod +x /app/bin/api
CMD ["./bin/api"]
