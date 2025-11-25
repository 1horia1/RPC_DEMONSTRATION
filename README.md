# RPC Calculator in Go

## Description

This project demonstrates a simple **RPC server and client** in Go.  
The server exposes a `Calculator` service that can add two integers. The client connects to the server and calls the `Add` method to get the sum.

---

## Project Structure

- **Server** (`server.go`)  
  - Defines the `Calculator` type and the RPC method `Add`.  
  - Registers the service with RPC.  
  - Listens on TCP port 1234 and handles client connections using `rpc.ServeConn`.

- **Client** (`client.go`)  
  - Connects to the RPC server using `rpc.Dial`.  
  - Prepares the arguments for the `Add` method and a `reply` variable for the result.  
  - Calls the `Calculator.Add` RPC method and prints the result.

---

## How it works

1. The server creates an instance of the `Calculator` type.  
2. The `Add` method is defined with a pointer receiver (`*Calculator`) and receives:
   - `args [2]int` – the two integers sent by the client.go  
   - `reply *int` – pointer where the server writes the result for the client  
3. The server listens on port 1234 for TCP connections.  
4. The client connects to the server and calls the `Calculator.Add` method.  
5. RPC handles the serialization of arguments and the result between the client and the server.

---

## Usage Example

### Start the server:

```bash
go run server.go
```

### Output server:

```bash
Server RPC ascultă pe portul 1234...
```

### Start client:

```bash
go run client.go
```

### Output:

```bash
8
```
