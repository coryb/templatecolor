package templatecolor

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const esc = '\x1b'

func TestTemplate(t *testing.T) {
	tt := template.New("template").Funcs(FuncMap())

	buf := bytes.Buffer{}
	tt, err := tt.Parse(`{{ "hi" | fg.Red }}`)
	require.NoError(t, err)
	err = tt.Execute(&buf, nil)
	require.NoError(t, err)
	// 31 fg.red
	require.Equal(t, []byte{esc, '[', '3', '1', 'm', 'h', 'i', esc, '[', '0', 'm'}, buf.Bytes())
	buf.Reset()

	tt, err = tt.Parse(`{{ "hi" | fg.HiRed | bg.Red | bold | strike }}`)
	require.NoError(t, err)
	err = tt.Execute(&buf, nil)
	require.NoError(t, err)
	// 91 fg.hiRed, 41 bg.red, 9 strike, 1 bold
	require.Equal(t, []byte{esc, '[', '9', '1', ';', '4', '1', ';', '9', ';', '1', 'm', 'h', 'i', esc, '[', '0', 'm'}, buf.Bytes())
	buf.Reset()

	tt, err = tt.Parse(`{{ "hi" | bold | fg.Red }}`)
	require.NoError(t, err)
	err = tt.Execute(&buf, nil)
	require.NoError(t, err)
	// 31 fg.red, 1 bold
	require.Equal(t, []byte{esc, '[', '3', '1', ';', '1', 'm', 'h', 'i', esc, '[', '0', 'm'}, buf.Bytes())
	buf.Reset()
}
