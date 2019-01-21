package handler

type AppHandler interface {
	IndexHandler
	TodoHandler
}

type appHandler struct {
	IndexHandler
	TodoHandler
}

func NewAppHandler(indexHandler IndexHandler, todoHandler TodoHandler) AppHandler {
	return &appHandler{
		indexHandler,
		todoHandler,
	}
}
