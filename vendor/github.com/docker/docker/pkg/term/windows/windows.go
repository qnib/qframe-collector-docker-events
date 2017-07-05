// These files implement ANSI-aware input and output streams for use by the Docker Windows client.
// When asked for the set of standard streams (e.g., stdin, stdout, stderr), the code will create
// and return pseudo-streams that convert ANSI sequences to / from Windows Console API calls.

<<<<<<< HEAD
package windows
=======
package windowsconsole
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

import (
	"io/ioutil"
	"os"
	"sync"

	ansiterm "github.com/Azure/go-ansiterm"
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger
var initOnce sync.Once

func initLogger() {
	initOnce.Do(func() {
		logFile := ioutil.Discard

		if isDebugEnv := os.Getenv(ansiterm.LogEnv); isDebugEnv == "1" {
			logFile, _ = os.Create("ansiReaderWriter.log")
		}

		logger = &logrus.Logger{
			Out:       logFile,
			Formatter: new(logrus.TextFormatter),
			Level:     logrus.DebugLevel,
		}
	})
}
