package magick

type FontMetrics struct {
	CharacterWidth           float64
	CharacterHeight          float64
	Ascender                 float64
	Descender                float64
	TextWidth                float64
	TextHeight               float64
	MaximumHorizontalAdvance float64
	BoundingBoxX1            float64
	BoundingBoxY1            float64
	BoundingBoxX2            float64
	BoundingBoxY2            float64
	OriginX                  float64
	OriginY                  float64
}

func NewFontMetricsFromArray(arr [13]float64) *FontMetrics {
	return &FontMetrics{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12]}
}
