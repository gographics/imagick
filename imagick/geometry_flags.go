package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

const (
	NOVALUE  = C.NoValue  // = 0x0000
	XVALUE   = C.XValue   ////= 0x0001
	XIVALUE  = C.XiValue  //= 0x0001
	YVALUE   = C.YValue   //= 0x0002
	PSIVALUE = C.PsiValue //= 0x0002

	WIDTHVALUE = C.WidthValue //= 0x0004
	RHOVALUE   = C.RhoValue   //= 0x0004

	HEIGHTVALUE = C.HeightValue //= 0x0008
	SIGMAVALUE  = C.SigmaValue  //= 0x0008
	CHIVALUE    = C.ChiValue    //= 0x0010
	XINEGATIVE  = C.XiNegative  //= 0x0020

	XNEGATIVE               = C.XNegative               //= 0x0020
	PSINEGATIVE             = C.PsiNegative             //= 0x0040
	YNEGATIVE               = C.YNegative               //= 0x0040
	CHINEGATIVE             = C.ChiNegative             //= 0x0080
	PERCENTVALUE            = C.PercentValue            //= 0x1000   /* '%'  percentage of something */
	ASPECTVALUE             = C.AspectValue             //= 0x2000    /* '!'  resize no-aspect - special use flag */
	NORMALIZEVALUE          = C.NormalizeValue          //= 0x2000 /* '!'  ScaleKernelValue() in morphology.c */
	LESSVALUE               = C.LessValue               //= 0x4000      /* '<'  resize smaller - special use flag */
	GREATERVALUE            = C.GreaterValue            //= 0x8000   /* '>'  resize larger - spacial use flag */
	MINIMUMVALUE            = C.MinimumValue            //= 0x10000  /* '^'  special handling needed */
	CORRELATENORMALIZEVALUE = C.CorrelateNormalizeValue //= 0x10000 /* '^' see ScaleKernelValue() */
	AREAVALUE               = C.AreaValue               //= 0x20000     /* '@'  resize to area - special use flag */
	DECIMALVALUE            = C.DecimalValue            //= 0x40000  /* '.'  floating point numbers found */
	SEPARATORVALUE          = C.SeparatorValue          //= 0x80000  /* 'x'  separator found  */

	ALLVALUES = C.AllValues //= 0x7fffffff
)
