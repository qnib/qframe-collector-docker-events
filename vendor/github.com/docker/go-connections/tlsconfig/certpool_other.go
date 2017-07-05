// +build !go1.7

package tlsconfig

import (
	"crypto/x509"

<<<<<<< HEAD
	"github.com/Sirupsen/logrus"
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

// SystemCertPool returns an new empty cert pool,
// accessing system cert pool is supported in go 1.7
func SystemCertPool() (*x509.CertPool, error) {
<<<<<<< HEAD
	logrus.Warn("Unable to use system certificate pool: requires building with go 1.7 or later")
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return x509.NewCertPool(), nil
}
