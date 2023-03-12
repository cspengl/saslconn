# saslconn

A connection based SASL protocol definition + implementation according to [RFC-4422](https://www.rfc-editor.org/rfc/rfc4422)

--- 

[![Go Reference](https://pkg.go.dev/badge/github.com/cspengl/saslconn.svg)](https://pkg.go.dev/github.com/cspengl/saslconn)

This repository serves two purposes:

 - Proposing a [SASL connection protocol](./protocol/) described via protobuf files
 - Containing a implementation of the proposed protocol in Golang based on the [go-sasl](github.com/emersion/go-sasl) library

Since the protocol is described through protobuf messages it should be relatively easy to implement the protocol in other programming languages.

## Usage

The implementation of the protocol is inspired by golang's TLS implementation. Hence establishing connections hopefully
feels very familiar.

> **Note**: The following example transfers credentials in plain-text via the network. To ensure confidentiality and integrity you should consider combining the SASL connection with other protective mechanisms like TLS (see next section).

**Client**
```golang
import (
    "github.com/cspengl/saslconn"
    "github.com/emersion/go-sasl"
)

func main() {
    config := &saslconn.Config{
        Mechanisms: []*saslconn.Mechanism{
            {
                Name: sasl.Plain,
                Client: sasl.NewPlainClient("userident", "username", "secret"),
            },
        },
    }

    conn, err := saslconn.Dial("tcp", ":9000", config)
    defer func() {
        if err := conn.Close(); err != nil {
            panic(err)
        }
    }()
    if err != nil {
        panic(err)
    }
    if _, err := conn.Write([]byte("Hello World!")); err != nil {
        panic(err)
    }
}
```

**Server**
```golang
import (
    "fmt"

    "github.com/cspengl/saslconn"
    "github.com/emersion/go-sasl"
)

func main() {
    config := &saslconn.Config{
        Mechanisms: []*saslconn.Mechanism{
            {
                Name: sasl.Plain,
                Server: sasl.NewPlainServer(
					func(identity, username, password string) error {
                        // always successful
						return nil
					},
				),
            },
            
        },
    }

    listener, err := saslconn.Listen("tcp", ":9000", config)
    defer func() {
        if err := listener.Close(); err != nil {
            panic(err)
        }
    }()
    conn, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    helloWorldFromClient := make([]byte, 12)
    if _, err := conn.Read(helloWorldFromClient); err != nil {
        panic(err)
    }
    fmt.Println(string(helloWorldFromClient))
}
```

## Combination with TLS

Neither SASL mechanisms nor the proposed connection protocol itself provide adequate integrity and/or confidentialty protection of the authentication exchange it is highly recommended to combine it with other protective mechanisms like TLS.

**Example**
```golang
import (
    "fmt"
    "crypto/tls"

    "github.com/cspengl/saslconn"
    "github.com/emersion/go-sasl"
)

func main() {
    config := &saslconn.Config{
        Mechanisms: []*saslconn.Mechanism{
            {
                Name: sasl.Plain,
                Server: sasl.NewPlainServer(
					func(identity, username, password string) error {
                        // always successful
						return nil
					},
				),
            },
            
        },
    }

    cert, err := tls.LoadX509KeyPair("example-cert.pem", "example-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

    listener, err := saslconn.ListenTLS("tcp", ":9000", tlsConfig, config)
    defer func() {
        if err := listener.Close(); err != nil {
            panic(err)
        }
    }()
    //...
}
```

## License

[MIT](LICENSE)