# 🚀 jwtx 🔐

<p align="center">
  <img src="./demo.png" alt="jwtx demo"/>
</p>

<p align="center">
  <em>JWT decoder/encoder for your terminal! 🎯</em>
</p>

<p align="center">
  <a href="#installation"><strong>Install</strong></a> •
  <a href="#features"><strong>Features</strong></a> •
  <a href="#usage"><strong>Usage</strong></a> •
  <a href="#keyboard-shortcuts"><strong>Shortcuts</strong></a>
</p>

<br/>

[![GitHub stars](https://img.shields.io/github/stars/gurleensethi/jwtx.svg?style=social&label=Star)](https://github.com/gurleensethi/jwtx)
[![License](https://img.shields.io/github/license/gurleensethi/jwtx.svg)](https://github.com/gurleensethi/jwtx/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gurleensethi/jwtx)](https://goreportcard.com/report/github.com/gurleensethi/jwtx)
[![Downloads](https://img.shields.io/github/downloads/gurleensethi/jwtx/total.svg)](https://github.com/gurleensethi/jwtx/releases)
[![Contributors](https://img.shields.io/github/contributors/gurleensethi/jwtx.svg)](https://github.com/gurleensethi/jwtx/graphs/contributors)

## ✨ Features

- 🛠️ **Real-time JWT Decoding** - Instantly decode JWT tokens as you type
- 🎨 **Beautiful TUI Interface** - Clean, intuitive terminal interface
- ⌨️ **Keyboard Navigation** - Full keyboard control with shortcuts
- 🔍 **Header & Payload Inspection** - View both header and payload separately
- 🛡️ **Signature Verification** - Validate JWT signatures with your secret
- 📱 **Responsive Layout** - Automatically adapts to your terminal size
- 🌟 **Visual Feedback** - Clear status indicators for validation results

## 🚀 Installation

### Option 1: Using Go

```bash
go install github.com/gurleensethi/jwtx@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/gurleensethi/jwtx.git
cd jwtx
go build -o jwtx .
./jwtx
```

## 🎯 Usage

Simply run the application and start exploring JWT tokens:

```bash
jwtx
```

Then paste your JWT token in the **JSON WEB TOKEN** field and your secret in the **SECRET** field. The decoded header and payload will appear instantly!

## ⌨️ Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Ctrl + T` | Focus on JWT Token field |
| `Ctrl + S` | Focus on Secret field |
| `Ctrl + H` | Focus on Decoded Header |
| `Ctrl + P` | Focus on Decoded Payload |
| `Ctrl + C` | Quit application |
| `Ctrl + Q` | Alternative quit |

## 📈 Stats

![GitHub Repo stars](https://img.shields.io/github/stars/gurleensethi/jwtx?style=social)
![GitHub forks](https://img.shields.io/github/forks/gurleensethi/jwtx?style=social)
![GitHub contributors](https://img.shields.io/github/contributors/gurleensethi/jwtx)
![GitHub issues](https://img.shields.io/github/issues/gurleensethi/jwtx)
![GitHub pull requests](https://img.shields.io/github/issues-pr/gurleensethi/jwtx)

## 📜 License

[MIT ©](https://github.com/gurleensethi/jwtx/blob/master/LICENSE)

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/gurleensethi">Gurleen Sethi</a>
</p>