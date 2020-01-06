# Dummy-chat

This project is the result of me learning Golang over few weeks. It was developed for educational purpose. I hope this
project will help me improve and tempt Golang beginner's to continue learning this amazing language!

## Getting Started

### Prerequisites

This project has be developed using:
- go version go1.13.5
- libprotoc 3.11.2 (protoc)

Golang can be installed from the official [go website](https://golang.org/dl/).

For libprotoc, the easiest way to install it is by retrieving the right version directly
from [protobuf's GitHub](https://github.com/protocolbuffers/protobuf/releases)

### Installing

It's quite simple to start the project. A basic CLI made with the standard `flag` package was made to specify common
options like the port the server should listen to.

```
$ dummy-chat-server -port 1234
```

For the client, it's very similar. You can specify the host and port you are attacking by using:
```
$ dummy-chat-client -host localhost -port 1234
```

If you don't specify any options, by default the server will start listening on port `50051` (gRPC default) and
the client will connect to `localhost:50051`

When you're connected; simply input on stdin to send messages to all clients connected to the same server. 

## Contributing

I'll be more than happy to have feedbacks on the way I designed this application. Things can always be done better and
I m eager to learn what could be improved on my code.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details