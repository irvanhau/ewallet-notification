package services

import "ewallet-notification/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepo
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service healthy", nil
}
