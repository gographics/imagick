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
	UndefinedException:        "UndefinedException",
	WarningException:          "WarningExceptionOrResourceLimitWarning",
	TypeWarning:               "TypeWarning",
	OptionWarning:             "OptionWarning",
	DelegateWarning:           "DelegateWarning",
	MissingDelegateWarning:    "MissingDelegateWarning",
	CorruptImageWarning:       "CorruptImageWarning",
	FileOpenWarning:           "FileOpenWarning",
	BlobWarning:               "BlobWarning",
	StreamWarning:             "StreamWarning",
	CacheWarning:              "CacheWarning",
	CoderWarning:              "CoderWarning",
	FilterWarning:             "FilterWarning",
	ModuleWarning:             "ModuleWarning",
	DrawWarning:               "DrawWarning",
	ImageWarning:              "ImageWarning",
	WandWarning:               "WandWarning",
	RandomWarning:             "RandomWarning",
	XServerWarning:            "XServerWarning",
	MonitorWarning:            "MonitorWarning",
	RegistryWarning:           "RegistryWarning",
	ConfigureWarning:          "ConfigureWarning",
	PolicyWarning:             "PolicyWarning",
	ErrorException:            "ErrorExceptionOrResourceLimitError",
	TypeError:                 "TypeError",
	OptionError:               "OptionError",
	DelegateError:             "DelegateError",
	MissingDelegateError:      "MissingDelegateError",
	CorruptImageError:         "CorruptImageError",
	FileOpenError:             "FileOpenError",
	BlobError:                 "BlobError",
	StreamError:               "StreamError",
	CacheError:                "CacheError",
	CoderError:                "CoderError",
	FilterError:               "FilterError",
	ModuleError:               "ModuleError",
	DrawError:                 "DrawError",
	ImageError:                "ImageError",
	WandError:                 "WandError",
	RandomError:               "RandomError",
	XServerError:              "XServerError",
	MonitorError:              "MonitorError",
	RegistryError:             "RegistryError",
	ConfigureError:            "ConfigureError",
	PolicyError:               "PolicyError",
	FatalErrorException:       "FatalErrorExceptionOrResourceLimitFatalError",
	TypeFatalError:            "TypeFatalError",
	OptionFatalError:          "OptionFatalError",
	DelegateFatalError:        "DelegateFatalError",
	MissingDelegateFatalError: "MissingDelegateFatalError",
	CorruptImageFatalError:    "CorruptImageFatalError",
	FileOpenFatalError:        "FileOpenFatalError",
	BlobFatalError:            "BlobFatalError",
	StreamFatalError:          "StreamFatalError",
	CacheFatalError:           "CacheFatalError",
	CoderFatalError:           "CoderFatalError",
	FilterFatalError:          "FilterFatalError",
	ModuleFatalError:          "ModuleFatalError",
	DrawFatalError:            "DrawFatalError",
	ImageFatalError:           "ImageFatalError",
	WandFatalError:            "WandFatalError",
	RandomFatalError:          "RandomFatalError",
	XServerFatalError:         "XServerFatalError",
	MonitorFatalError:         "MonitorFatalError",
	RegistryFatalError:        "RegistryFatalError",
	ConfigureFatalError:       "ConfigureFatalError",
	PolicyFatalError:          "PolicyFatalError",
}

func (et *ExceptionType) String() string {
	if v, ok := exceptionTypeStrings[ExceptionType(*et)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownError[%d]", *et)
}
