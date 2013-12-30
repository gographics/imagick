#include <wand/MagickWand.h>

#include "wand.h"

void getImageFloatRGB(MagickWand *mw, float *results, FloatFormat channels) {
	PixelIterator *pi;
	PixelWand **pw;

	int pos;
	unsigned long y, width, height;
	register long x;

	width = MagickGetImageWidth(mw);
	height = MagickGetImageHeight(mw);

	pi = NewPixelIterator(mw);

	pos = 0;

	for (y=0; y < height; y++) {
		pw = PixelGetNextIteratorRow(pi, &width);
		for (x=0; x < (long) width; x++) {

			if (channels & FLOAT_FORMAT_R) {
				results[pos++] = (float) PixelGetRed(pw[x]);
			}

			if (channels & FLOAT_FORMAT_G) {
				results[pos++] = (float) PixelGetGreen(pw[x]);
			}

			if (channels & FLOAT_FORMAT_B) {
				results[pos++] = (float) PixelGetBlue(pw[x]);
			}

			if (channels & FLOAT_FORMAT_A) {
				results[pos++] = (float) PixelGetAlpha(pw[x]);
			}
		}
	}

	DestroyPixelIterator(pi);
}
