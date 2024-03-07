// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package inputdate

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/iamajoe/templui/utils"
	"time"
)

type styleguideItem struct {
	title       string
	description string
	usage       string
	opts        []Opt
}

var styleguideItems = []styleguideItem{
	{
		title: "Default",
		usage: `
@inputdate.New(
    inputdate.WithID("zing"),
    inputdate.WithClasses("foo"),
    inputdate.WithAttributes(map[string]any{"data-zed": "zung"}),
    // inputdate.WithDisabled(),
    inputdate.WithRequired(),
    inputdate.WithName("start-date"),
    inputdate.WithValue(time.Now()),
    inputdate.WithPlaceholder("Set your start date"),
    inputdate.WithMin(time.Now()),
    inputdate.WithMax(time.Now()),
    OptFn(func(c *inputdate.InputDate) {
        c.ClassNames = append(c.ClassNames, "bar")
    }),
)
    `,
		opts: []Opt{
			WithID("zing"),
			WithClasses("foo"),
			WithAttributes(map[string]any{"data-zed": "zung"}),
			// WithDisabled(),
			WithRequired(),
			WithName("start-date"),
			WithValue(time.Now()),
			WithPlaceholder("Set your start date"),
			WithMin(time.Now()),
			WithMax(time.Now()),
		},
	},
}

func Styleguide() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, component := range styleguideItems {
			templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Err = New(component.opts...).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = utils.Item(component.title, component.description, component.usage).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}