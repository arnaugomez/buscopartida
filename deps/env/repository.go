package env

type envRepository struct {
	env Env
}

func (r envRepository) GetEnv() Env {
	return r.env
}