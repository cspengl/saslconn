# Protocol and Handshake

This document proposes a protocol offering SASL services as described in [RFC-4422](https://www.rfc-editor.org/rfc/rfc4422).

It provides:

- A simple handshake for negotiating and executing SASL mechanisms
- As set of [messages](./protocol.md) for implementing the handshake


## Handshake

```txt
                              ,------.                      ,------.                    
                              |Client|                      |Server|                    
                              `--+---'                      `--+---'                    
                                 | ServerMechanismAdvertisement|                        
                                 | <- - - - - - - - - - - - - -                         
                                 |                             |                        
                                 |  ClientMechanismInitiation  |                        
                                 |  - - - - - - - - - - - - - ->                        
                                 |                             |                        
                                 |                             |                        
          _____________________________________________________________________________ 
          ! ALT  /               |                             |                       !
          !_____/                |                             |                       !
          !                      |                             |                       !
          !         _________________________________________________________          !
          !         ! LOOP  /    |                             |             !         !
          !         !______/     |                             |             !         !
          !         !            |      ChallengeResponse      |             !         !
          !         !            | <- - - - - - - - - - - - - -              !         !
          !         !            |                             |             !         !
          !         !            |      ChallengeResponse      |             !         !
          !         !            |  - - - - - - - - - - - - - ->             !         !
          !         !~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~!         !
          !~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~!
                                 |                             |                        
                                 |          ServerDone         |                        
                                 | <- - - - - - - - - - - - - -                         
                              ,--+---.                      ,--+---.                    
                              |Client|                      |Server|                    
                              `------'                      `------'                    

```

### Mechanism Negotiation

For negotiating a SASL mechanism the server initiates the handshake by advertising supported mechanisms by sending
a list of mechanism names (see [ServerMechanismAdvertisement](./protocol.md#servermechanismadvertisement)). The order of the list MAY represent
a priority for selecting a mechanism.

The client MUST select one of the provided mechanisms. Otherwise the server will abort the handshake. The selection itself may be implementation specific. 

### Request Authentication Exchange

After selecting one of the provided mechanisms the client initiates the authentication by sending a [Client Initiation](./protocol.md#clientinitiation) message. This message MUST contain the selected mechanism name.

If the mechanism uses a 'initial_response' it MUST be provided. If there is no 'initial response' the field 'initial_response_is_nil' MUST be `false`. This field has the purpose to distinguish between 'nil' and empty initial responses (see [RFC-4422, section 4](https://www.rfc-editor.org/rfc/rfc4422#section-4)).

### Challenges and Responses

If the selected mechanism needs further challenges the server sends them using [ChallengeResponse](./protocol.md#challengeresponse) messages which need to be answered by the client using the same message type.

### Authentication Outcomes

As soon as the mechansim is done the server sends a [ServerDone](./protocol.md#serverdone) message indicating the result of the authentication mechanism. It contains a result field which MUST be set to either `ResultSuccess` or `ResultReject`.

The [ServerDone](./protocol.md#serverdone) message futher MAY contain a message containing a rejection cause or success message.

### Aborting Authentication Exchanges

In case of any error the client and the server can abort the handshake by sending a [HandshakeAbortion](./protocol.md#handshakeabortion) message. This message SHOULD contain a reason for the abortion.

The handshake can get aborted in following scenarios:

- The client does not support any advertised protocol

- The client fails to initiate a selected mechanism

- The server fails to process the initial or any other response

- The client initiates a not supported mechanism (handshake aborted by server)

- The client fails to process a challenge by the server

## Message Transport

The protocol assumes that there is a underying streaming connection for exchanging messages. Since the protobuf format itself does not provide a mechanism for delimiting messages each message MUST be prefixed with it's length. The length is represented by a unsigned 64-bit big-endian integer in the first eight bytes of the message.

```txt
 0                            63
+-------------------------------+
|        Message Length         |
+-------------------------------+
|           Message
+-------------- ...
```

## Limitations

 - The protocol does not support multiple authentications (see [RFC-4422, section 3.8](https://www.rfc-editor.org/rfc/rfc4422#section-3.8))