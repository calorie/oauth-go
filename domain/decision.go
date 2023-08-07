package domain

type DecisionRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Decision struct {
	Location string
}
