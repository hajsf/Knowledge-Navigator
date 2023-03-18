package translation

import (
	"fmt"
	"wa/global"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func HelloPerson(lang string, name string) string {
	fmt.Println(lang)
	fmt.Println(name)

	// Create a new localizer.
	//var localizer = i18n.NewLocalizer(global.Bundle, lang)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	// bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("lang/active.ar.toml")
	localizer := i18n.NewLocalizer(bundle, lang)

	// Set title message.
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "HelloPerson", // set translation ID
			Other: "Hello *{{.Name}}* \n" +
				"Welcome at King Khalid University", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
		PluralCount: nil,
	})
}

var WhoIsThis = func(lang string, name string) string {
	// Create a new localizer.
	var localizer = i18n.NewLocalizer(global.Bundle, lang)

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "WhoIsThis",    // set translation ID
			Other: "Who are you?", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})

	//var content strings.Builder
	//	content.WriteString(fmt.Sprintf("مرحبا *%v* \n", name))
	//content.WriteString(helloPerson)
}
