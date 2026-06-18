package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func Generate(alphabet string, size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}

func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}
	return id
}
