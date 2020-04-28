package main

import (
	"os"

	"github.com/ahmdrz/goinsta"
)

func main() {
	var (
		insta *goinsta.Instagram
		err   error
	)

	insta = goinsta.New(os.Getenv("IG_USERNAME"), os.Getenv("IG_PASSWORD"))

	if err = insta.Login(); err != nil {
		// Panic es forma sencilla de caputar errores desconocidos
		panic(err)
	}

	if err != nil {
		panic(err)
	}

}
