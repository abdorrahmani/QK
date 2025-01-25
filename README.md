# BB84 Quantum Key Distribution Protocol in Go - Dockerized

![Go Version](https://img.shields.io/badge/Go-1.23-blue?style=flat-square&logo=go)
![Docker](https://img.shields.io/badge/Docker-Supported-blue?style=flat-square&logo=docker)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://shields.io/)

This project implements the **BB84 Quantum Key Distribution Protocol** in Go. It is fully containerized using Docker, allowing easy setup and execution in any environment.

---

## 🚀 Features

- **BB84 Quantum Key Distribution Protocol**: Simulates the BB84 protocol for secure key exchange between two parties.
- **Dockerized Application**: Run the application with Docker and Docker Compose without worrying about setup or dependencies.
- **Secure Communication**: Simulates quantum key distribution for encryption and decryption of messages.

---

## 📘 What is BB84?

The BB84 protocol, introduced in 1984 by Charles Bennett and Gilles Brassard, is a quantum key distribution (QKD) protocol. It enables two parties to generate a shared encryption key securely, even in the presence of an eavesdropper.

Key concepts:
1. **Qubits and Bases**: Alice encodes bits using two types of bases (Z basis and X basis). Bob measures them using randomly chosen bases.
2. **Transmission**: Qubits are transmitted over a quantum channel.
3. **Key Sifting**: Alice and Bob compare their bases over a public channel. Bits measured with matching bases are kept for the shared key.

**Why BB84 is secure**: Any eavesdropping (e.g., by Eve) on the quantum channel introduces detectable anomalies in the transmission, enabling Alice and Bob to discard affected qubits.

---

## 🛠️ Prerequisites

- **Docker**: Install Docker and Docker Compose on your machine. Follow the installation guide from [Docker's official site](https://www.docker.com/get-started).
- **Git**: Clone the repository to your local machine.

---

## 🧑‍💻 Installation and Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/abdorrahmani/QK.git
    cd QK
    ```
2. Build and run the Docker container:
   ```bash
   docker-compose up --build
   ```

3. Run tests to ensure everything is working:
   ```bash
   docker exec -it qk go test ./... -v
   ```
---

## 📖 How It Works

1. **Protocol Initialization**: The program initializes the BB84 protocol with a specified number of bits (e.g., 128).
2. **Quantum Transmission**:
   - Alice prepares random bits and bases.
   - Qubits are transmitted over the quantum channel.
   - Bob measures the qubits using random bases.
3. **Key Sifting**: Alice and Bob compare bases to generate a shared key.
4. **Secure Communication**:
   - Alice encrypts a message using the shared key.
   - Bob decrypts the message using the same key.

---

## 📚 Simulation Details

The program simulates the entire BB84 protocol without physical quantum systems by:
- Generating random bits and bases for Alice and Bob.
- Simulating the quantum channel to transmit qubits.
- Producing a shared key based on matching bases.
- Using the shared key to encrypt and decrypt messages.

---

## 🧪 Testing

Run the included tests to verify the protocol and secure communication implementation:
```bash
go test ./... -v
```

The tests include:
- Validation of random bit and basis generation.
- End-to-end BB84 protocol simulation.
- Encryption and decryption functionality.

---

## 🤝 Contributions

Contributions are welcome! Feel free to submit issues or pull requests to improve the implementation.

---

### ⚡ References
- [BB84 Protocol on Wikipedia](https://en.wikipedia.org/wiki/BB84)
- [Quantum Key Distribution](https://www.ibm.com/quantum-computing/learn/what-is-quantum-key-distribution/)

