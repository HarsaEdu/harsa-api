# Use the official Golang image as the base image
FROM golang:1.20-bookworm

# Set the working directory inside the container
WORKDIR /app

RUN apt update

# Install wkhtmltox depedencies
RUN apt install fontconfig libjpeg62-turbo libx11-6 libxcb1 libxext6 libxrender1 xfonts-75dpi xfonts-base -y
RUN wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2_amd64.deb
RUN dpkg -i libssl1.1_1.1.1f-1ubuntu2_amd64.deb

# Install wkhtmltox
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