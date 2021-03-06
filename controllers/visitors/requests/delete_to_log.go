package requests

import "backend/business/visitors"

type DeleteToLogRequest struct {
	PatientsId  uint   `json:"patients_id"`
	SchedulesId uint   `json:"schedules_id"`
	Solusi      string `json:"solusi"`
	Message     string `json:"message"`
}

func (uplog *DeleteToLogRequest) ToDomainLog() visitors.Log {
	return visitors.Log{
		PatientsId:  uplog.PatientsId,
		SchedulesId: uplog.SchedulesId,
		Solusi:      uplog.Solusi,
		Message:     uplog.Message,
	}
}
