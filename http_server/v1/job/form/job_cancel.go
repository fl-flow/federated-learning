package form


type JobCancelForm struct {
  JobID         uint                  `json:"job_id" binding:"required"`
}
