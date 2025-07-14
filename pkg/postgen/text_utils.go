package postgen

import (
	"errors"
	"image"
	"image/color"

	"github.com/haashemi/writer"
)

var ErrInvalidFormat = errors.New("invalid format")

func GenerateTextWriter(textContent string, fontFamilyPath string, fontColor string, fontSize int32) (*writer.Writer, *image.Uniform, error) {
	face := writer.NewFaceFromFile(fontFamilyPath)
	font := writer.NewFont(face, fontSize)
	w, err := writer.NewWriter(font, textContent, writer.DefaultOptions)
	if err != nil {
		return nil, nil, err
	}
	fontColorRGBA, err := ParseHexColor(fontColor)
	if err != nil {
		return nil, nil, err
	}
	fontColorText := image.NewUniform(fontColorRGBA)
	return w, fontColorText, nil
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	if len(s) == 0 || s[0] != '#' {
		return c, ErrInvalidFormat
	}
	c.A = 0xff
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = ErrInvalidFormat
		return 0
	}
	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = ErrInvalidFormat
	}
	return
}
