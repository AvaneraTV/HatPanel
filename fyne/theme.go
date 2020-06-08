package fyne

import (
	"image/color"

	"fyne.io/fyne"
)

// NewAppTheme is
func newAppTheme(fallbackTheme fyne.Theme) fyne.Theme {
	return &appTheme{
		defaultTheme: fallbackTheme,
	}
}

type appTheme struct {
	defaultTheme fyne.Theme

	textSize *int
	padding  *int
}

func (t *appTheme) BackgroundColor() color.Color {
	return t.defaultTheme.BackgroundColor()
}
func (t *appTheme) ButtonColor() color.Color {
	return t.defaultTheme.ButtonColor()
}
func (t *appTheme) DisabledButtonColor() color.Color {
	return t.defaultTheme.DisabledButtonColor()
}
func (t *appTheme) HyperlinkColor() color.Color {
	return t.defaultTheme.HyperlinkColor()
}
func (t *appTheme) TextColor() color.Color {
	return t.defaultTheme.TextColor()
}
func (t *appTheme) DisabledTextColor() color.Color {
	return t.defaultTheme.DisabledTextColor()
}
func (t *appTheme) IconColor() color.Color {
	return t.defaultTheme.IconColor()
}
func (t *appTheme) DisabledIconColor() color.Color {
	return t.defaultTheme.DisabledIconColor()
}
func (t *appTheme) PlaceHolderColor() color.Color {
	return t.defaultTheme.PlaceHolderColor()
}
func (t *appTheme) PrimaryColor() color.Color {
	return t.defaultTheme.PrimaryColor()
}
func (t *appTheme) HoverColor() color.Color {
	return t.defaultTheme.HoverColor()
}
func (t *appTheme) FocusColor() color.Color {
	return t.defaultTheme.FocusColor()
}
func (t *appTheme) ScrollBarColor() color.Color {
	return t.defaultTheme.ScrollBarColor()
}
func (t *appTheme) ShadowColor() color.Color {
	return t.defaultTheme.ShadowColor()
}
func (t *appTheme) TextSize() int {
	if t.textSize != nil {
		return *t.textSize
	}
	return 32
}
func (t *appTheme) TextFont() fyne.Resource {
	return t.defaultTheme.TextFont()
}
func (t *appTheme) TextBoldFont() fyne.Resource {
	return t.defaultTheme.TextBoldFont()
}
func (t *appTheme) TextItalicFont() fyne.Resource {
	return t.defaultTheme.TextItalicFont()
}
func (t *appTheme) TextBoldItalicFont() fyne.Resource {
	return t.defaultTheme.TextBoldItalicFont()
}
func (t *appTheme) TextMonospaceFont() fyne.Resource {
	return t.defaultTheme.TextMonospaceFont()
}
func (t *appTheme) Padding() int {
	if t.padding != nil {
		return *t.padding
	}
	return 32
}
func (t *appTheme) IconInlineSize() int {
	return t.defaultTheme.IconInlineSize()
}
func (t *appTheme) ScrollBarSize() int {
	return t.defaultTheme.ScrollBarSize()
}
func (t *appTheme) ScrollBarSmallSize() int {
	return t.defaultTheme.ScrollBarSmallSize()
}
