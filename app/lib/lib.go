package lib

type Library struct {
	Env     Env
	Logger  Logger
	Handler RequestHandler
}

func NewLibrary() Library {
	env := NewEnv()
	return Library{
		Env:     env,
		Logger:  GetLogger(),
		Handler: NewRequestHandler(),
	}
}
