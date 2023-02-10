package builder

import "strings"

type CoreBuilder struct {
	*strings.Builder
}

func (c *CoreBuilder) WriteString(s string) {
	if c.Builder == nil {
		c.Builder = &strings.Builder{}
	}
	c.Builder.WriteString(s)
}

func (c *CoreBuilder) WriteByte(v byte) {
	if c.Builder == nil {
		c.Builder = &strings.Builder{}
	}
	c.Builder.WriteByte(v)
}

func (c *CoreBuilder) ReadString() string {
	if c.Builder != nil {
		return c.Builder.String()
	}
	return ""
}

func (c *CoreBuilder) Comma() {
	c.WriteString(",")
}

func (c *CoreBuilder) CommaSpace() {
	c.Comma()
	c.WriteString(" ")
}


