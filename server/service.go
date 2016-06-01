package scepserver

import (
	"github.com/micromdm/scep/scep"
	"golang.org/x/net/context"
)

// Service is the interface for all supported SCEP server operations.
type Service interface {
	// GetCACaps returns a list of options
	// which are supported by the server.
	GetCACaps(ctx context.Context) ([]byte, error)

	// GetCACert returns CA certificate or
	// a CA certificate chain with intermediates
	// in a PKCS#7 Degenerate Certificates format
	GetCACert(ctx context.Context) ([]byte, error)

	// PKIOperation handles incoming SCEP messages such as PKCSReq and
	// sends back a CertRep PKIMessag.
	PKIOperation(ctx context.Context, msg []byte) (*scep.PKIMessage, error)

	// GetNextCACert returns a replacement certificate or certificate chain
	// when the old one expires. The response format is a PKCS#7 Degenerate
	// Certificates type.
	GetNextCACert(ctx context.Context) ([]byte, error)
}

type service struct {
}

func (svc service) GetCACaps(ctx context.Context) ([]byte, error) {
	defaultCaps := []byte(`POSTPKIOperation`)
	return defaultCaps, nil
}

func (svc service) GetCACert(ctx context.Context) ([]byte, error) {
	panic("not implemented")
}

func (svc service) PKIOperation(ctx context.Context, msg []byte) (*scep.PKIMessage, error) {
	panic("not implemented")
}

func (svc service) GetNextCACert(ctx context.Context) ([]byte, error) {
	panic("not implemented")
}

// NewService creates a new scep service
func NewService() Service {
	return &service{}
}
