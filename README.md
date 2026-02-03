# jwtx

A terminal-based TUI application for decoding and validating JSON Web Tokens (JWTs).

## Features

- Decode JWT tokens to view headers and payload
- Validate JWT signatures using secrets or public keys
- Real-time validation feedback
- Interactive terminal interface with keyboard shortcuts

## Installation

```bash
go install github.com/gurleensethi/jwtx@latest
```

## Usage

Run the application:

```bash
jwtx
```

### Keyboard Shortcuts

- `Ctrl+1`: Focus on JWT token input field
- `Ctrl+2`: Focus on secret input field
- `Ctrl+C` or `Ctrl+Q`: Quit the application

### Interface

The TUI displays four panels:
- **JSON WEB TOKEN**: Input field for the JWT token
- **SECRET**: Input field for the signing secret/public key
- **DECODED HEADER**: Shows the decoded JWT header
- **DECODED PAYLOAD**: Shows the decoded JWT claims

Status indicators show validation results:
- Green: Valid JWT or signature verified
- Red: Invalid token or signature verification failed

## Development

Clone the repository:

```bash
git clone https://github.com/gurleensethi/jwtx.git
cd jwtx
go run .
```

## License

MIT