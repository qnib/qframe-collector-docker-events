/*
Package logrus is a structured logger for Go, completely API compatible with the standard library logger.


The simplest way to use Logrus is simply the package-level exported logger:

  package main

  import (
<<<<<<< HEAD
    log "github.com/Sirupsen/logrus"
=======
    log "github.com/sirupsen/logrus"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
  )

  func main() {
    log.WithFields(log.Fields{
      "animal": "walrus",
      "number": 1,
      "size":   10,
    }).Info("A walrus appears")
  }

Output:
  time="2015-09-07T08:48:33Z" level=info msg="A walrus appears" animal=walrus number=1 size=10

<<<<<<< HEAD
For a full guide visit https://github.com/Sirupsen/logrus
=======
For a full guide visit https://github.com/sirupsen/logrus
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
*/
package logrus
