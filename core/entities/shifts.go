package entities

import "flick_tickets/core/domain"

type ShiftAddReq struct {
	ShiftName  string  `form:"shift_name"`
	Salary     float64 `form:"salary"`
	ShiftStart int     `form:"shift_start"`
	ShiftEnd   int     `form:"shift_end"`
}
type ShiftAddResp struct {
	Result Result `json:"result"`
}
type ShiftGetAllResp struct {
	Result Result          `json:"result"`
	Shifts []*domain.Shift `json:"shifts"`
}
