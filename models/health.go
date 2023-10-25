package models

import "time"

type HealthCheckResponse struct {
	Status string        `json:"status"`
	Uptime time.Duration `json:"uptime"`
}
