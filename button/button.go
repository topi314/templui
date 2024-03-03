package button

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type Kind string

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

type (
	OptsFn func(*Button) error
	Button struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Disabled bool
		Kind     Kind

		Opts []OptsFn
	}
)

func WithID(id string) OptsFn {
	return func(element *Button) error {
		element.ID = id
		return nil
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *Button) error {
		element.ClassNames = append(element.ClassNames, classes...)
		return nil
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *Button) error {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}

		return nil
	}
}

func WithDisabled() OptsFn {
	return func(element *Button) error {
		element.Disabled = true
		return nil
	}
}

func WithKind(kind Kind) OptsFn {
	return func(element *Button) error {
		element.Kind = kind
		return nil
	}
}

func (c *Button) WithFunction(fn func(*Button) error) *Button {
	c.Opts = append(c.Opts, fn)
	return c
}

func (c *Button) WithID(id string) *Button {
	c.Opts = append(c.Opts, WithID(id))
	return c
}

func (c *Button) WithClasses(classes ...string) *Button {
	c.Opts = append(c.Opts, WithClasses(classes...))
	return c
}

func (c *Button) WithAttributes(attributes map[string]any) *Button {
	c.Opts = append(c.Opts, WithAttributes(attributes))
	return c
}

func (c *Button) WithDisabled() *Button {
	c.Opts = append(c.Opts, WithDisabled())
	return c
}

func (c *Button) WithKind(kind Kind) *Button {
	c.Opts = append(c.Opts, WithDisabled())
	return c
}

func (c *Button) Render(ctx context.Context, w io.Writer) error {
	for _, opt := range c.Opts {
		if err := opt(c); err != nil {
			// TODO: need to figure how to handle errors
			return templ.NopComponent.Render(ctx, w)
		}
	}

	return render(*c).Render(ctx, w)
}

func New(opts ...OptsFn) *Button {
	return &Button{
		Opts: opts,
	}
}
