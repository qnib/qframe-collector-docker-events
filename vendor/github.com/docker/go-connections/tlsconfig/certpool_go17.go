// +build go1.7

package tlsconfig

import (
	"crypto/x509"
	"runtime"
<<<<<<< HEAD

	"github.com/Sirupsen/logrus"
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

// SystemCertPool returns a copy of the system cert pool,
// returns an error if failed to load or empty pool on windows.
func SystemCertPool() (*x509.CertPool, error) {
	certpool, err := x509.SystemCertPool()
	if err != nil && runtime.GOOS == "windows" {
<<<<<<< HEAD
		logrus.Infof("Unable to use system certificate pool: %v", err)
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return x509.NewCertPool(), nil
	}
	return certpool, err
}
