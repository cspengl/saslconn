package saslconn

import (
	"testing"

	"github.com/emersion/go-sasl"
)

func Test_validateRole(t *testing.T) {
	clientMechanism := &Mechanism{
		Client: sasl.NewAnonymousClient(""),
	}
	serverMechanism := &Mechanism{
		Server: sasl.NewAnonymousServer(func(trace string) error { return nil }),
	}
	if err := clientMechanism.validateRole(roleClient); err != nil {
		t.Error(err)
	}
	if err := clientMechanism.validateRole(roleServer); err == nil {
		t.Fatal("expected error")
	}
	if err := serverMechanism.validateRole(roleServer); err != nil {
		t.Error(err)
	}
	if err := serverMechanism.validateRole(roleClient); err == nil {
		t.Fatal("expected error")
	}
}

func Test_validateName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"PLAIN", false},
		{"A1234", false},
		{"A_123-", false},
		{"plain", true},
		{"THISISAVERYLONGMECHANISMNAME", true},
		{"", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := (&Mechanism{Name: test.name}).validateName()
			if err != nil && !test.wantErr {
				t.Error(err)
			} else if err == nil && test.wantErr {
				t.Error("did not get expected error")
			}
		})
	}
}
