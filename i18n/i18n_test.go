package i18n_test

import (
	"annotation/i18n"
	"testing"
)

var basePath = "/home/eder.costa/go/src/github.com/edermanoel94"


func TestStart(t *testing.T) {

	t.Run("should start app with lang", func(t *testing.T) {
		err := i18n.Start(basePath, i18n.EnUS)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should not start app if use a lang which not supported yet", func(t *testing.T) {
		err := i18n.Start(basePath, i18n.EnUS)
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestMessage(t *testing.T) {

	t.Run("should apply a message `resource_not_found` for language pt-BR", func(t *testing.T) {
		err := i18n.Start(basePath, i18n.PtBR)

		if err != nil {
			t.Fatal(err)
		}

		messageNotFound := i18n.Message(i18n.ResourceNotFound)

		if messageNotFound != "recurso não encontrado." {
			t.Fatal("mensagem é diferente do esperado.")
		}
	})

	t.Run("should not apply a message for given key without exists in language pt-BR", func(t *testing.T) {

		err := i18n.Start(basePath, i18n.PtBR)

		if err != nil {
			t.Fatal(err)
		}

		message := i18n.Message("alguma_key")

		if message != "" {
			t.Fatal("mensagem é diferente do esperado.")
		}
	})

}

func TestError(t *testing.T) {

	t.Run("should apply a message `resource_not_found` for language pt-BR", func(t *testing.T) {
		err := i18n.Start(basePath, i18n.PtBR)

		if err != nil {
			t.Fatal(err)
		}

		messageNotFound := i18n.Error(i18n.ResourceNotFound)

		if messageNotFound != nil {
			t.Fatal(err)
		}

	})

	t.Run("should not apply a message for given key without exists in language pt-BR", func(t *testing.T) {

		err := i18n.Start(basePath, i18n.PtBR)

		if err != nil {
			t.Fatal(err)
		}

		message := i18n.Message("alguma_key")

		if message != "" {
			t.Fatal("mensagem é diferente do esperado.")
		}
	})
}