package espsdk

import "github.com/Sirupsen/logrus"

// Log is a logger that allows easy access to nicely-formatted log output that
// includes properties of completed HTTP requests such as the response time
// and status code.
var Log = logrus.New()
