package docx

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

func decodeRaster(b []byte, ext string) (image.Image, error) {
	r := bytes.NewReader(b)
	switch strings.ToLower(ext) {
	case "png":
		return png.Decode(r)
	case "jpg", "jpeg":
		return jpeg.Decode(r)
	case "gif":
		if img, err := gif.Decode(bytes.NewReader(b)); err == nil {
			return img, nil
		}
		ga, err := gif.DecodeAll(bytes.NewReader(b))
		if err != nil || len(ga.Image) == 0 {
			return nil, err
		}
		return ga.Image[0], nil
	case "webp":
		return webp.Decode(r)
	case "bmp":
		return bmp.Decode(r)
	case "tif", "tiff":
		return tiff.Decode(r)
	default:
		return png.Decode(r)
	}
}

func resizeImage(img image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

func normalizeToPNG(data []byte, ext string, maxWidthPx int) (out []byte, w, h int, err error) {
	img, err := decodeRaster(data, ext)
	if err != nil {
		return nil, 0, 0, err
	}
	w = img.Bounds().Dx()
	h = img.Bounds().Dy()
	if w <= 0 || h <= 0 {
		return nil, 0, 0, err
	}
	if maxWidthPx > 0 && w > maxWidthPx {
		scale := float64(maxWidthPx) / float64(w)
		w = maxWidthPx
		h = int(float64(h) * scale)
		img = resizeImage(img, w, h)
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, 0, 0, err
	}
	return buf.Bytes(), w, h, nil
}
