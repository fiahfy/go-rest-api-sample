package handler

type AppHandler interface {
	HelloHandler
	TodoHandler
}

type appHandler struct {
	HelloHandler
	TodoHandler
}

func NewAppHandler(helloHandler HelloHandler, todoHandler TodoHandler) AppHandler {
	return &appHandler{
		helloHandler,
		todoHandler,
	}
}
