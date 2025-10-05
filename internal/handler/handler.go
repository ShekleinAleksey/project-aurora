package handler

type Deps struct {
	ExpenseService ExpenseService
}

type Handler struct {
	ExpenseHandler *ExpenseHandler
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		ExpenseHandler: NewExpenseHandler(deps.ExpenseService),
	}
}
