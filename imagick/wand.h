#ifndef _IMAGICK_WAND_H_
#define _IMAGICK_WAND_H_

#include <wand/MagickWand.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
	FLOAT_FORMAT_R    = 1 << 0,
	FLOAT_FORMAT_G    = 1 << 1,
	FLOAT_FORMAT_B    = 1 << 2,
	FLOAT_FORMAT_A    = 1 << 3,
	FLOAT_FORMAT_RGB  = FLOAT_FORMAT_R | FLOAT_FORMAT_G | FLOAT_FORMAT_B,
	FLOAT_FORMAT_RGBA = FLOAT_FORMAT_RGB | FLOAT_FORMAT_A
} FloatFormat;

void getImageFloatRGB(MagickWand *mw, float *results, FloatFormat channels);

#ifdef __cplusplus
}
#endif
#endif