package imagick

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
	UndefinedException        ExceptionType = C.UndefinedException
	WarningException          ExceptionType = C.WarningException
	ResourceLimitWarning      ExceptionType = C.ResourceLimitWarning
	TypeWarning               ExceptionType = C.TypeWarning
	OptionWarning             ExceptionType = C.OptionWarning
	DelegateWarning           ExceptionType = C.DelegateWarning
	MissingDelegateWarning    ExceptionType = C.MissingDelegateWarning
	CorruptImageWarning       ExceptionType = C.CorruptImageWarning
	FileOpenWarning           ExceptionType = C.FileOpenWarning
	BlobWarning               ExceptionType = C.BlobWarning
	StreamWarning             ExceptionType = C.StreamWarning
	CacheWarning              ExceptionType = C.CacheWarning
	CoderWarning              ExceptionType = C.CoderWarning
	FilterWarning             ExceptionType = C.FilterWarning
	ModuleWarning             ExceptionType = C.ModuleWarning
	DrawWarning               ExceptionType = C.DrawWarning
	ImageWarning              ExceptionType = C.ImageWarning
	WandWarning               ExceptionType = C.WandWarning
	RandomWarning             ExceptionType = C.RandomWarning
	XServerWarning            ExceptionType = C.XServerWarning
	MonitorWarning            ExceptionType = C.MonitorWarning
	RegistryWarning           ExceptionType = C.RegistryWarning
	ConfigureWarning          ExceptionType = C.ConfigureWarning
	PolicyWarning             ExceptionType = C.PolicyWarning
	ErrorException            ExceptionType = C.ErrorException
	ResourceLimitError        ExceptionType = C.ResourceLimitError
	TypeError                 ExceptionType = C.TypeError
	OptionError               ExceptionType = C.OptionError
	DelegateError             ExceptionType = C.DelegateError
	MissingDelegateError      ExceptionType = C.MissingDelegateError
	CorruptImageError         ExceptionType = C.CorruptImageError
	FileOpenError             ExceptionType = C.FileOpenError
	BlobError                 ExceptionType = C.BlobError
	StreamError               ExceptionType = C.StreamError
	CacheError                ExceptionType = C.CacheError
	CoderError                ExceptionType = C.CoderError
	FilterError               ExceptionType = C.FilterError
	ModuleError               ExceptionType = C.ModuleError
	DrawError                 ExceptionType = C.DrawError
	ImageError                ExceptionType = C.ImageError
	WandError                 ExceptionType = C.WandError
	RandomError               ExceptionType = C.RandomError
	XServerError              ExceptionType = C.XServerError
	MonitorError              ExceptionType = C.MonitorError
	RegistryError             ExceptionType = C.RegistryError
	ConfigureError            ExceptionType = C.ConfigureError
	PolicyError               ExceptionType = C.PolicyError
	FatalErrorException       ExceptionType = C.FatalErrorException
	ResourceLimitFatalError   ExceptionType = C.ResourceLimitFatalError
	TypeFatalError            ExceptionType = C.TypeFatalError
	OptionFatalError          ExceptionType = C.OptionFatalError
	DelegateFatalError        ExceptionType = C.DelegateFatalError
	MissingDelegateFatalError ExceptionType = C.MissingDelegateFatalError
	CorruptImageFatalError    ExceptionType = C.CorruptImageFatalError
	FileOpenFatalError        ExceptionType = C.FileOpenFatalError
	BlobFatalError            ExceptionType = C.BlobFatalError
	StreamFatalError          ExceptionType = C.StreamFatalError
	CacheFatalError           ExceptionType = C.CacheFatalError
	CoderFatalError           ExceptionType = C.CoderFatalError
	FilterFatalError          ExceptionType = C.FilterFatalError
	ModuleFatalError          ExceptionType = C.ModuleFatalError
	DrawFatalError            ExceptionType = C.DrawFatalError
	ImageFatalError           ExceptionType = C.ImageFatalError
	WandFatalError            ExceptionType = C.WandFatalError
	RandomFatalError          ExceptionType = C.RandomFatalError
	XServerFatalError         ExceptionType = C.XServerFatalError
	MonitorFatalError         ExceptionType = C.MonitorFatalError
	RegistryFatalError        ExceptionType = C.RegistryFatalError
	ConfigureFatalError       ExceptionType = C.ConfigureFatalError
	PolicyFatalError          ExceptionType = C.PolicyFatalError
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
