package saslconn

import (
	"regexp"

	"github.com/emersion/go-sasl"
	"github.com/pkg/errors"
)

type (
	//MechanismClient is redefined from the go-sasl library
	MechanismClient sasl.Client
	//MechanismServer is redefined from the go-sasl library
	MechanismServer sasl.Server
)

// Mechanism describes a SASL mechanism as describes in RFC4422
type Mechanism struct {
	// Name is the name of the mechanism which must comply the
	// naming scheme defined in RFC4422 section 3.1
	Name string
	// Client specifies the client implementation for this mechanism
	Client MechanismClient
	// Server specifies the server implementation for this mechanism
	Server MechanismServer
}

const (
	mechanismNameExpression = "^[A-Z0-9-_]{1,20}$"
)

var (
	errInvalidMechanism     = errors.New("invalid mechanism %s")
	errInvalidMechanismName = errors.New("invalid mechanism name %s")
	errIncompleteMechanism  = errors.New("no implementation given for role")
)

type mechanismRoleType string

const (
	roleClient mechanismRoleType = "client"
	roleServer mechanismRoleType = "server"
)

func (m *Mechanism) validateRole(role mechanismRoleType) (err error) {
	var ok bool
	switch role {
	case roleClient:
		ok = m.Client != nil
	case roleServer:
		ok = m.Server != nil
	}
	if !ok {
		err = errors.Wrap(errIncompleteMechanism, string(role))
	}
	return
}

func (m *Mechanism) validateName() (err error) {
	expr := regexp.MustCompile(mechanismNameExpression)
	if !expr.MatchString(m.Name) {
		err = errors.Errorf(errInvalidMechanismName.Error(), m.Name)
	}
	return
}

func (m *Mechanism) validate(role mechanismRoleType) error {
	err := m.validateName()
	if err != nil {
		return err
	}
	return m.validateRole(role)
}

// Config describes the the configuration used for establishing
// a SASL connection. It contains a list of supported mechanisms.
type Config struct {
	Mechanisms []*Mechanism
}

func (c *Config) validate(role mechanismRoleType) error {
	for _, mechanism := range c.Mechanisms {
		err := mechanism.validate(role)
		if err != nil {
			returnErr := errors.Errorf(errInvalidMechanism.Error(), mechanism.Name)
			return errors.Wrap(err, returnErr.Error())
		}
	}
	return nil
}

func (c *Config) mechanismList() (mechanisms []string) {
	for _, mechanism := range c.Mechanisms {
		mechanisms = append(mechanisms, mechanism.Name)
	}
	return
}

func (c *Config) mechanismIsSupported(mechanismName string) (supported bool) {
	mechanisms := c.mechanismList()
	for _, mechanism := range mechanisms {
		if mechanism == mechanismName {
			supported = true
			return
		}
	}
	return
}

func (c *Config) getMechanism(mechanismName string) *Mechanism {
	if !c.mechanismIsSupported(mechanismName) {
		return nil
	}
	for i, mechanism := range c.mechanismList() {
		if mechanism == mechanismName {
			return c.Mechanisms[i]
		}
	}
	return nil
}
