######## Start a new stage from scratch #######
FROM alpine:latest

CMD ["pwd"]

# Copy the Pre-built binary file from the previous stage
COPY ./bin/article-service .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./article-service"]