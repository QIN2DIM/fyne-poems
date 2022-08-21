package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

// Font Override fonts so that Chinese can be rendered normally
// 可忽略 `未解析的引用` 警告，不影响编译
func (*MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceYuWeiYingBiChangGuiTi1Ttf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceYuWeiYingBiChangGuiTi1Ttf
}

func (*MyTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 233, G: 233, B: 227, A: 255}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 10, G: 10, B: 8, A: 255}
	case theme.ColorNameHover:
		return color.NRGBA{R: 222, G: 219, B: 214, A: 255}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 199, G: 200, B: 196, A: 255}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}
