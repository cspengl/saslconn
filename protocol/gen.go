package saslproto

//go:generate protoc --go_out=../internal/saslproto  --doc_out=. --doc_opt=util/md.tmpl,protocol.md sasl.proto
