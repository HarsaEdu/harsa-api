# Use the official Golang image as the base image
FROM golang:1.20-bookworm

# Set the working directory inside the container
WORKDIR /app

# Install wkhtmltox binary
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6.1-3/wkhtmltox_0.12.6.1-3.bookworm_amd64.deb
RUN apt install ./wkhtmltox_0.12.6.1-3.bookworm_amd64.deb
ENV WKHTMLTOPDF_PATH /usr/bin/wkhtmltopdf

# Copy the necessary files into the container
COPY . .

# Build the Go application
RUN make build CGO_ENABLED=1 LDFLAGS="-lwkhtmltox"

# Expose the port that your API will run on
EXPOSE 8080

# Run the application
CMD ["./harsa-api.exe"]