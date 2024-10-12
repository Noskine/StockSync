package internal

import "os"

type PreInitialization struct {
	Port string
	Host string
}

var Con = new(PreInitialization)

func Conf() *PreInitialization {
	return Con
}

func init() {
	Con.Port = os.Getenv("PORT")
	Con.Host = os.Getenv("HOST")
}
