Template Color is a library for Go text/template that adds template functions to render text in with ansi terminal colors and style.

Usage:
```
template.New("name").Funcs(texttint.FuncMap())
```

Template examples:
```
{{ "hi" | fg.Red }}
{{ "hi" | fg.HiRed | bg.Red | bold | strike }}
{{ "hi" | bold | fg.Red }}
```

Template Functions:
* `fg.Black` - foreground black text
* `fg.Red` - foreground red text
* `fg.Green` - foreground green text
* `fg.Yellow` - foreground yellow text
* `fg.Blue` - foreground blue text
* `fg.Magenta` - foreground magenta text
* `fg.Cyan` - foreground cyan text
* `fg.White` - foreground white text
* `bg.Black` - black background
* `bg.Red` - red background
* `bg.Green` - green background
* `bg.Yellow` - yellow background
* `bg.Blue` - blue background
* `bg.Magenta` - magenta background
* `bg.Cyan` - cyan background
* `bg.White` - white background

* `fg.HiBlack` - foreground high intensity black text
* `fg.HiRed` - foreground high intensity red text
* `fg.HiGreen` - foreground high intensity green text
* `fg.HiYellow` - foreground high intensity yellow text
* `fg.HiBlue` - foreground high intensity blue text
* `fg.HiMagenta` - foreground high intensity magenta text
* `fg.HiCyan` - foreground high intensity cyan text
* `fg.HiWhite` - foreground high intensity white text
* `bg.HiBlack` - high intensity black background
* `bg.HiRed` - high intensity red background
* `bg.HiGreen` - high intensity green background
* `bg.HiYellow` - high intensity yellow background
* `bg.HiBlue` - high intensity blue background
* `bg.HiMagenta` - high intensity magenta background
* `bg.HiCyan` - high intensity cyan background
* `bg.HiWhite` - high intensity white background

* `bold` - increased intensity for text
* `dim` - decreased intensity for text
* `italics` - italicize text
* `underline` - draw line under text
* `slowBlink` - blink text slowly (ignored in many terminals)
* `rapidBlink` - blink text rapidly (ignored in many terminals)
* `invert` - swap foreground and background colors
* `hide` - mask text
* `strike` - text crossed out
