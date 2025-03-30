# Let's Go Tutorial Book Repository

Welcome to the **Let's Go Tutorial Book Repository**! 🎉

This repository contains source code, examples, and exercises related to Alex Edwards' _Let's Go_ tutorial book, which focuses on building web applications using the Go programming language. For more information about the book, visit the official website: [https://lets-go.alexedwards.net/](https://lets-go.alexedwards.net/).

## About the Book

_Let's Go_ provides an in-depth guide for developers interested in learning and mastering Go for web development. It covers topics like:

- HTTP handlers and middleware
- Working with databases
- Structuring and securing web applications
- Writing tests and ensuring application reliability

## Getting Started

1. Run code examples using the Go compiler:

   ```bash
   go run ./cmd/web
   ```

2. You may need to generate the `/tls` certificate for localhost developing:

   ```bash
   cd tls
   go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
   ```

For detailed instructions and code explanations, refer to the book.

Happy coding! 🚀
