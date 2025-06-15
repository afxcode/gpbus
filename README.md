# 🛠️ gpbus

![License: MIT License](https://img.shields.io/badge/LICENSE-MIT-blue)
![Go: 1.24](https://img.shields.io/badge/Go-1.24-blue)
![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-green)

**gpbus** is a lightweight Modbus TCP server that exposes Raspberry Pi (and compatible SBC like Orange Pi) **GPIOs as Modbus discrete inputs and coils**.  
Built in Go, it enables industrial tools and SCADA systems to interact with physical GPIO pins over a standard Ethernet network.

---

## ✨ Features

- 🔌 Maps GPIO pins to **Modbus registers**
- 📡 Supports **Modbus TCP** on configurable port (default: `1502`)
- ⚙️ Clean JSON-style configuration
- ⏱️ Fast and efficient GPIO polling every **_5ms_**
- 📦 Single binary deployment (`gpbus`)
- 🧩 Suitable for **Raspberry Pi**, partial support for **Orange Pi**
- 🛠 Written in [Go](https://golang.org/) with [periph.io](https://periph.io)
---

## ⚙️ Config file
Example config file.
```json
{
  "inputs": [
    {
      "name": "GPIO1",
      "register": 101
    },
    {
      "name": "GPIO2",
      "register": 102
    }
  ],
  "outputs": [
    {
      "name": "GPIO13",
      "register": 2013
    },
    {
      "name": "GPIO14",
      "register": 2014
    }
  ],
  "port": 1502
}
```

📌 _**Input pins will be readable via Modbus as discrete inputs, while output pins are mapped as coils.**_

---

## 🚀 Running

Start the server on default port `1502`:
```bash
./gpbus
``` 

Or Start with config file in different location
```bash
./gpbus -config=myproject/ioconfig.json
``` 

---
## 🙏 Credits
- GPIO handling powered by [periph.io]() — excellent hardware library for Go

- Modbus server functionality inspired by and based on [github.com/tbrandon/mbserver]() — a robust Modbus TCP/RTU server implementation in Go
