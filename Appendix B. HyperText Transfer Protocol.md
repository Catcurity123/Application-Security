## 🌐 What is **HTTP**?

**HTTP (HyperText Transfer Protocol)** is the foundational protocol for communication on the web. It defines how **requests** and **responses** are made between **clients (browsers, apps, etc.)** and **servers**.

### Basic Structure of HTTP

An HTTP **request** typically consists of:
1. **Request Line** (includes the HTTP method, e.g., `GET`, `POST`, the path, and HTTP version).
2. **Headers** (provide additional information about the request, like `User-Agent`, `Accept`, etc.).
3. **Body** (optional data sent along with certain request types like `POST` or `PUT`).

An HTTP **response** typically consists of:
1. **Status Line** (includes HTTP version, a status code like `200 OK`, and a reason phrase).
2. **Headers** (similar to request headers, like `Content-Type`, `Content-Length`, etc.).
3. **Body** (the actual content of the response — could be HTML, JSON, images, etc.).

### Example of an HTTP Request and Response

`Request:`
```
GET /home HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0
Accept: text/html
```

`Response:`
```
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 13

Hello, world!
```