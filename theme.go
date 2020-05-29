package main

import (
	"image/color"

	"fyne.io/fyne"
)

type AppTheme struct {
	defaultTheme fyne.Theme
}

func (t *AppTheme) BackgroundColor() color.Color {
	return t.defaultTheme.BackgroundColor()
}

func (t *AppTheme) ButtonColor() color.Color {
	return t.defaultTheme.ButtonColor()
}
func (t *AppTheme) DisabledButtonColor() color.Color {
	return t.defaultTheme.DisabledButtonColor()
}
func (t *AppTheme) HyperlinkColor() color.Color {
	return t.defaultTheme.HyperlinkColor()
}
func (t *AppTheme) TextColor() color.Color {
	return t.defaultTheme.TextColor()
}
func (t *AppTheme) DisabledTextColor() color.Color {
	return t.defaultTheme.DisabledTextColor()
}
func (t *AppTheme) IconColor() color.Color {
	return t.defaultTheme.IconColor()
}
func (t *AppTheme) DisabledIconColor() color.Color {
	return t.defaultTheme.DisabledIconColor()
}
func (t *AppTheme) PlaceHolderColor() color.Color {
	return t.defaultTheme.PlaceHolderColor()
}
func (t *AppTheme) PrimaryColor() color.Color {
	return t.defaultTheme.PrimaryColor()
}
func (t *AppTheme) HoverColor() color.Color {
	return t.defaultTheme.HoverColor()
}
func (t *AppTheme) FocusColor() color.Color {
	return t.defaultTheme.FocusColor()
}
func (t *AppTheme) ScrollBarColor() color.Color {
	return t.defaultTheme.ScrollBarColor()
}
func (t *AppTheme) ShadowColor() color.Color {
	return t.defaultTheme.ShadowColor()
}

func (t *AppTheme) TextSize() int {
	return 72 // t.defaultTheme.TextSize()
}
func (t *AppTheme) TextFont() fyne.Resource {
	return t.defaultTheme.TextFont()
}
func (t *AppTheme) TextBoldFont() fyne.Resource {
	return t.defaultTheme.TextBoldFont()
}
func (t *AppTheme) TextItalicFont() fyne.Resource {
	return t.defaultTheme.TextItalicFont()
}
func (t *AppTheme) TextBoldItalicFont() fyne.Resource {
	return t.defaultTheme.TextBoldItalicFont()
}
func (t *AppTheme) TextMonospaceFont() fyne.Resource {
	return t.defaultTheme.TextMonospaceFont()
}

func (t *AppTheme) Padding() int {
	return t.defaultTheme.Padding()
}
func (t *AppTheme) IconInlineSize() int {
	return t.defaultTheme.IconInlineSize()
}
func (t *AppTheme) ScrollBarSize() int {
	return t.defaultTheme.ScrollBarSize()
}
func (t *AppTheme) ScrollBarSmallSize() int {
	return t.defaultTheme.ScrollBarSmallSize()
}
