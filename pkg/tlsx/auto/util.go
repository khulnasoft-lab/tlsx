package auto

import (
	sliceutil "github.com/khulnasoft-lab/utils/slice"
	"github.com/khulnasoft-lab/tlsx/pkg/tlsx/openssl"
	"github.com/khulnasoft-lab/tlsx/pkg/tlsx/tls"
	"github.com/khulnasoft-lab/tlsx/pkg/tlsx/ztls"
)

var (
	allCiphersNames      []string
	supportedTlsVersions []string
)

func init() {
	allCiphersNames = append(tls.AllCiphersNames, ztls.AllCiphersNames...)
	allCiphersNames = append(allCiphersNames, openssl.AllCiphersNames...)
	supportedTlsVersions = append(tls.SupportedTlsVersions, ztls.SupportedTlsVersions...)
	supportedTlsVersions = append(supportedTlsVersions, openssl.SupportedTLSVersions...)
	allCiphersNames = sliceutil.Dedupe(allCiphersNames)
	supportedTlsVersions = sliceutil.Dedupe(supportedTlsVersions)
}
