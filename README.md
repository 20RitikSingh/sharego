# ShareGo

A secure and efficient file sharing application that allows you to share files across devices on your local network using QR codes.

## 🚀 Features

- **Secure File Transfer**: Uses HTTPS with TLS encryption for secure file transfers
- **QR Code Integration**: Automatically generates QR codes for easy access from mobile devices
- **Multiple File Support**: Upload and share multiple files simultaneously
- **ZIP Download**: Download multiple files as a ZIP archive
- **Web Interface**: Clean and simple web interface for file management
- **Local Network**: Works within your local network for fast transfers

## 📋 Prerequisites

- Go 1.23.2 or higher
- TLS certificates (for HTTPS)
- Make (for using Makefile commands)

## 🛠️ Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/20ritiksingh/sharego.git
   cd sharego
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate TLS certificates:
   ```bash
   make tls
   ```

## 🚦 Usage

1. Build and run the server:
   ```bash
   make run
   ```

   Alternatively, you can build and run separately:
   ```bash
   make build
   ./bin/app
   ```

2. The application will display QR codes for each network interface along with their URLs
3. Scan the QR code with your mobile device or visit the displayed URL in your browser
4. Upload files through the web interface
5. Share the generated links or QR codes with others on your network

## 🔧 Configuration

The server runs on port 8088 by default. You can modify this in the `server.go` file if needed.

## 🛠️ Development

### Building

- Build for your current platform:
  ```bash
  make build
  ```

- Build for all supported platforms:
  ```bash
  make build-for-all
  ```

### Testing and Quality Assurance

- Run all tests:
  ```bash
  make test
  ```

- Run code vetting:
  ```bash
  make vet
  ```

- Run linter:
  ```bash
  make lint
  ```

- Run all checks (vet, lint, test, build):
  ```bash
  make all
  ```

### Cleanup

- Clean build artifacts and generated files:
  ```bash
  make clean
  ```

## 📁 Project Structure

```
sharego/
├── bin/                # Binary files
├── certs/             # SSL certificates
├── dist/              # Distribution files
├── handlers/          # HTTP request handlers
├── server.go          # Main server file
├── networkAddr.go     # Network interface handling
├── go.mod            # Go module file
└── go.sum            # Go module checksum
```

## 🔒 Security

- The application uses TLS encryption for all communications
- Files are transferred securely over HTTPS
- Runs only on your local network for added security

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Before submitting a PR, ensure all tests pass:
```bash
make all
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## ✨ Acknowledgments

- [go-qrcode](https://github.com/skip2/go-qrcode) for QR code generation

## 👥 Authors

- **Ritik Singh** - *Initial work* - [20ritiksingh](https://github.com/20ritiksingh)

## 🆘 Support

For support, please open an issue in the GitHub repository. 