package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ExceptionType int

const (
	UndefinedException        ExceptionType = 0
	WarningException                        = 300
	ResourceLimitWarning                    = 300
	TypeWarning                             = 305
	OptionWarning                           = 310
	DelegateWarning                         = 315
	MissingDelegateWarning                  = 320
	CorruptImageWarning                     = 325
	FileOpenWarning                         = 330
	BlobWarning                             = 335
	StreamWarning                           = 340
	CacheWarning                            = 345
	CoderWarning                            = 350
	FilterWarning                           = 352
	ModuleWarning                           = 355
	DrawWarning                             = 360
	ImageWarning                            = 365
	WandWarning                             = 370
	RandomWarning                           = 375
	XServerWarning                          = 380
	MonitorWarning                          = 385
	RegistryWarning                         = 390
	ConfigureWarning                        = 395
	PolicyWarning                           = 399
	ErrorException                          = 400
	ResourceLimitError                      = 400
	TypeError                               = 405
	OptionError                             = 410
	DelegateError                           = 415
	MissingDelegateError                    = 420
	CorruptImageError                       = 425
	FileOpenError                           = 430
	BlobError                               = 435
	StreamError                             = 440
	CacheError                              = 445
	CoderError                              = 450
	FilterError                             = 452
	ModuleError                             = 455
	DrawError                               = 460
	ImageError                              = 465
	WandError                               = 470
	RandomError                             = 475
	XServerError                            = 480
	MonitorError                            = 485
	RegistryError                           = 490
	ConfigureError                          = 495
	PolicyError                             = 499
	FatalErrorException                     = 700
	ResourceLimitFatalError                 = 700
	TypeFatalError                          = 705
	OptionFatalError                        = 710
	DelegateFatalError                      = 715
	MissingDelegateFatalError               = 720
	CorruptImageFatalError                  = 725
	FileOpenFatalError                      = 730
	BlobFatalError                          = 735
	StreamFatalError                        = 740
	CacheFatalError                         = 745
	CoderFatalError                         = 750
	FilterFatalError                        = 752
	ModuleFatalError                        = 755
	DrawFatalError                          = 760
	ImageFatalError                         = 765
	WandFatalError                          = 770
	RandomFatalError                        = 775
	XServerFatalError                       = 780
	MonitorFatalError                       = 785
	RegistryFatalError                      = 790
	ConfigureFatalError                     = 795
	PolicyFatalError                        = 799
)

var exceptionTypeStrings = map[ExceptionType]string{
	0:   "UndefinedException",
	300: "WarningException",
	//300: "ResourceLimitWarning",
	305: "TypeWarning",
	310: "OptionWarning",
	315: "DelegateWarning",
	320: "MissingDelegateWarning",
	325: "CorruptImageWarning",
	330: "FileOpenWarning",
	335: "BlobWarning",
	340: "StreamWarning",
	345: "CacheWarning",
	350: "CoderWarning",
	352: "FilterWarning",
	355: "ModuleWarning",
	360: "DrawWarning",
	365: "ImageWarning",
	370: "WandWarning",
	375: "RandomWarning",
	380: "XServerWarning",
	385: "MonitorWarning",
	390: "RegistryWarning",
	395: "ConfigureWarning",
	399: "PolicyWarning",
	400: "ErrorException",
	//400: "ResourceLimitError",
	405: "TypeError",
	410: "OptionError",
	415: "DelegateError",
	420: "MissingDelegateError",
	425: "CorruptImageError",
	430: "FileOpenError",
	435: "BlobError",
	440: "StreamError",
	445: "CacheError",
	450: "CoderError",
	452: "FilterError",
	455: "ModuleError",
	460: "DrawError",
	465: "ImageError",
	470: "WandError",
	475: "RandomError",
	480: "XServerError",
	485: "MonitorError",
	490: "RegistryError",
	495: "ConfigureError",
	499: "PolicyError",
	700: "FatalErrorException",
	//700: "ResourceLimitFatalError",
	705: "TypeFatalError",
	710: "OptionFatalError",
	715: "DelegateFatalError",
	720: "MissingDelegateFatalError",
	725: "CorruptImageFatalError",
	730: "FileOpenFatalError",
	735: "BlobFatalError",
	740: "StreamFatalError",
	745: "CacheFatalError",
	750: "CoderFatalError",
	752: "FilterFatalError",
	755: "ModuleFatalError",
	760: "DrawFatalError",
	765: "ImageFatalError",
	770: "WandFatalError",
	775: "RandomFatalError",
	780: "XServerFatalError",
	785: "MonitorFatalError",
	790: "RegistryFatalError",
	795: "ConfigureFatalError",
	799: "PolicyFatalError",
}

func (et *ExceptionType) String() string {
	if v, ok := exceptionTypeStrings[ExceptionType(*et)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownError[%d]", *et)
}
