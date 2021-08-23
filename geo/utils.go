package geo

import "time"

func SetHTTPClientTimeout(clientTimeout time.Duration) {
	agentTimeout = clientTimeout
}