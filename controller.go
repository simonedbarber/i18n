package i18n

import (
	"github.com/simonedbarber/admin"
	"github.com/simonedbarber/qor/utils"
)

type i18nController struct {
	*I18n
}

func (controller *i18nController) Index(context *admin.Context) {
	context.Execute("index", controller.I18n)
}

func (controller *i18nController) Update(context *admin.Context) {
	form := context.Request.Form
	translation := Translation{Key: form.Get("Key"), Locale: form.Get("Locale"), Value: utils.HTMLSanitizer.Sanitize(form.Get("Value"))}

	if err := controller.I18n.SaveTranslation(&translation); err == nil {
		context.Execute("index", controller.I18n)
	} else {
		context.Writer.WriteHeader(422)
		context.Execute("index", controller.I18n)
	}
}
