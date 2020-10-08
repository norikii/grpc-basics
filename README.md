# USER-AUTH-gRPC
## 1. gRPC 
 - uses HTTP/2 as its transfer protocol - utilizes multiplexing for 
sending multiple requests and responses in parallel ove a single TCP 
connection. 
 - allows server push where with one single response we can send 
 multiple responses
 - single TCP connection carries multiple bidirectional streams where
 each stream has a unique ID and carries multiple bidirectional 
 messages - each message (request/response) is broken down into multiple
 binary frames => smallest unit that carries different type of data: 
 headers, settings priority, data, etc. 
 ###1.1. Types:
 a) unary - one request and response like normal http api
 
 b) client streaming - client sends a stream a multiple messages
 and expects a server to send just one response
 
 c) server streaming - client send just one request but server 
 responses with a stream of multiple messages
 
 d) bidirectional streaming - client and server keep sending and
 receiving multiple messages in parallel with arbitrary order - with
 non blocking

 