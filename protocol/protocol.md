# Protocol Documentation

## Table of Contents

- [sasl.proto](#sasl-proto)
    - [ChallengeResponse](#saslproto-ChallengeResponse)
    - [ClientInitiation](#saslproto-ClientInitiation)
    - [HandshakeAbortion](#saslproto-HandshakeAbortion)
    - [Message](#saslproto-Message)
    - [ServerDone](#saslproto-ServerDone)
    - [ServerMechanismAdvertisement](#saslproto-ServerMechanismAdvertisement)
  
    - [MessageType](#saslproto-MessageType)
    - [ServerDoneResult](#saslproto-ServerDoneResult)
  



<a name="sasl-proto"></a>

## sasl.proto
SASL protocol messages.

This file contains the SASL protocol definitions according to section 4 of
RFC-4422.


<a name="saslproto-ChallengeResponse"></a>

### ChallengeResponse
ChallengeResponse is the message used during the challenge-response
phase. It either contains a challenge sent by the server or a response
by the client (see section 3.4).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [bytes](#bytes) |  | Payload in bytes |






<a name="saslproto-ClientInitiation"></a>

### ClientInitiation
ClientInitiation initiates the SASL challenge-response exchange
between client and server by indicating the by the client selected
SASL mechanism and an optional initial response (see section 3.3)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mechanism | [string](#string) |  | Selected SASL mechanism |
| initial_reponse_is_nil | [bool](#bool) |  | Used to distinguish between nil and empty initial responses as described in section 4.3.b |
| initial_response | [bytes](#bytes) |  | Optional initial response sent by the client (depending on the mechanism) |






<a name="saslproto-HandshakeAbortion"></a>

### HandshakeAbortion
HandshakeAbortion is used to abort an ongoing handshake (authentication exchange)
as described in section 3.5


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  | Optional message which may be used to provide a abortion cause/ error code for the client |






<a name="saslproto-Message"></a>

### Message
based on section 3 in RFC4422


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_type | [MessageType](#saslproto-MessageType) |  | Type of the message indicating which of the other fields is set |
| server_mechanism_advertisement | [ServerMechanismAdvertisement](#saslproto-ServerMechanismAdvertisement) |  | Payload for message type MessageTypeServerMechanismAdvertisement |
| client_initiation | [ClientInitiation](#saslproto-ClientInitiation) |  | Payload for message type MessageTypeClientInitiation |
| challenge_response | [ChallengeResponse](#saslproto-ChallengeResponse) |  | Payload for message type MessageTypeChallengeResponse |
| handshake_abortion | [HandshakeAbortion](#saslproto-HandshakeAbortion) |  | Payload for message type MessageTypeHandshakeAbortion |
| server_done | [ServerDone](#saslproto-ServerDone) |  | Payload for message type MessageTypeServerDone |






<a name="saslproto-ServerDone"></a>

### ServerDone
ServerDone is used for describing the authentication outcome
and sent by the server after the client initiated a mechanism
or the last challenge of a mechanism has been answered by the client
(see section 3.6)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [ServerDoneResult](#saslproto-ServerDoneResult) |  | Result of the authentication mechanism which is either success or reject |
| message | [string](#string) |  | Optional message which may be used to provide additional information for the client |






<a name="saslproto-ServerMechanismAdvertisement"></a>

### ServerMechanismAdvertisement
ServerMechanismAdvertisement is the first message of the handshake
indicating which SASL mechanisms the server supports (see section 3.2)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mechanisms | [string](#string) | repeated | List of supported SASL mechanisms |





 <!-- end messages -->


<a name="saslproto-MessageType"></a>

### MessageType
MessageType is used for specifying the type of
message transmitted.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MessageTypeUnknown | 0 | Used for error cases or as fallback |
| MessageTypeServerMechanismAdvertisement | 1 | Used for ServerMechanismAdvertisement messages |
| MessageTypeClientInitiation | 2 | Used for ClientInitiation messages |
| MessageTypeChallengeResponse | 3 | Used for ChallengeResponse messages |
| MessageTypeHandshakeAbortion | 4 | Used for HandshakeAbortion messages |
| MessageTypeServerDone | 5 | Used for ServerDone messages |



<a name="saslproto-ServerDoneResult"></a>

### ServerDoneResult


| Name | Number | Description |
| ---- | ------ | ----------- |
| ResultUnknown | 0 | Used for error cases or as fallback |
| ResultSuccess | 1 | Indicating a successful authentication |
| ResultReject | 2 | Indicating that the authentication was not successful |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

