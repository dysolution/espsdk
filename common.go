package espsdk

import (
	"time"

	"github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Log is an instance of logrus.(*Logger) that allows easy access to
// nicely-formatted log output that includes properties of completed HTTP
// requests such as the response time and status code.
var Log = logrus.New()

func init() {
	Log.Formatter = &prefixed.TextFormatter{TimestampFormat: time.RFC3339}
}

// A Serializable object can be serialized to a byte stream such as JSON.
type serializable interface {
	Marshal() ([]byte, error)
}
