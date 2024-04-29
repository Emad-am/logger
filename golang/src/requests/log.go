package requests

type StoreLogRequest struct {
	LogType    string      `json:"log_type" validate:"required"`
	Actortype  string      `json:"actor_type" validate:"required"` // admin or user or system
	Actor      interface{} `json:"actor" validate:"required"`  // user object from client db
	TargetType string      `json:"target_type" validate:"required"` // admin or user or system
	Target     string      `json:"target" validate:"required"` // user object from client db
	Message    string      `json:"message" validate:"required"`
	Payload    interface{} `json:"payload" validate:"required"`
}
