package lib

type Library struct {
	Env     Env
	Logger  Logger
	Handler RequestHandler
}

func NewLibrary() Library {
	return Library{
		Env:     NewEnv(),
		Logger:  GetLogger(),
		Handler: NewRequestHandler(),
	}
}
