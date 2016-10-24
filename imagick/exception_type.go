// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ExceptionType int

const (
	EXCEPTION_UNDEFINED          ExceptionType = C.UndefinedException
	EXCEPTION_WARNING            ExceptionType = C.WarningException
	WARNING_RESOURCE_LIMIT       ExceptionType = C.ResourceLimitWarning
	WARNING_TYPE                 ExceptionType = C.TypeWarning
	WARNING_OPTION               ExceptionType = C.OptionWarning
	WARNING_DELEGATE             ExceptionType = C.DelegateWarning
	WARNING_MISSING_DELEGATE     ExceptionType = C.MissingDelegateWarning
	WARNING_CORRUPT_IMAGE        ExceptionType = C.CorruptImageWarning
	WARNING_FILE_OPEN            ExceptionType = C.FileOpenWarning
	WARNING_BLOB                 ExceptionType = C.BlobWarning
	WARNING_STREAM               ExceptionType = C.StreamWarning
	WARNING_CACHE                ExceptionType = C.CacheWarning
	WARNING_CODER                ExceptionType = C.CoderWarning
	WARNING_FILTER               ExceptionType = C.FilterWarning
	WARNING_MODULE               ExceptionType = C.ModuleWarning
	WARNING_DRAW                 ExceptionType = C.DrawWarning
	WARNING_IMAGE                ExceptionType = C.ImageWarning
	WARNING_WAND                 ExceptionType = C.WandWarning
	WARNING_RANDOM               ExceptionType = C.RandomWarning
	WARNING_XSERVER              ExceptionType = C.XServerWarning
	WARNING_MONITOR              ExceptionType = C.MonitorWarning
	WARNING_REGISTRY             ExceptionType = C.RegistryWarning
	WARNING_CONFIGURE            ExceptionType = C.ConfigureWarning
	WARNING_POLICY               ExceptionType = C.PolicyWarning
	EXCEPTION_ERROR              ExceptionType = C.ErrorException
	ERROR_RESOURCE_LIMIT         ExceptionType = C.ResourceLimitError
	ERROR_TYPE                   ExceptionType = C.TypeError
	ERROR_OPTION                 ExceptionType = C.OptionError
	ERROR_DELEGATE               ExceptionType = C.DelegateError
	ERROR_MISSING_DELEGATE       ExceptionType = C.MissingDelegateError
	ERROR_CORRUPT_IMAGE          ExceptionType = C.CorruptImageError
	ERROR_FILE_OPEN              ExceptionType = C.FileOpenError
	ERROR_BLOB                   ExceptionType = C.BlobError
	ERROR_STREAM                 ExceptionType = C.StreamError
	ERROR_CACHE                  ExceptionType = C.CacheError
	ERROR_CODER                  ExceptionType = C.CoderError
	ERROR_FILTER                 ExceptionType = C.FilterError
	ERROR_MODULE                 ExceptionType = C.ModuleError
	ERROR_DRAW                   ExceptionType = C.DrawError
	ERROR_IMAGE                  ExceptionType = C.ImageError
	ERROR_WAND                   ExceptionType = C.WandError
	ERROR_RANDOM                 ExceptionType = C.RandomError
	ERROR_XSERVER                ExceptionType = C.XServerError
	ERROR_MONITOR                ExceptionType = C.MonitorError
	ERROR_REGISTRY               ExceptionType = C.RegistryError
	ERROR_CONFIGURE              ExceptionType = C.ConfigureError
	ERROR_POLICY                 ExceptionType = C.PolicyError
	EXCEPTION_FATAL_ERROR        ExceptionType = C.FatalErrorException
	FATAL_ERROR_RESOURCE_LIMIT   ExceptionType = C.ResourceLimitFatalError
	FATAL_ERROR_TYPE             ExceptionType = C.TypeFatalError
	FATAL_ERROR_OPTION           ExceptionType = C.OptionFatalError
	FATAL_ERROR_DELEGATE         ExceptionType = C.DelegateFatalError
	FATAL_ERROR_MISSING_DELEGATE ExceptionType = C.MissingDelegateFatalError
	FATAL_ERROR_CORRUPT_IMAGE    ExceptionType = C.CorruptImageFatalError
	FATAL_ERROR_FILE_OPEN        ExceptionType = C.FileOpenFatalError
	FATAL_ERROR_BLOB             ExceptionType = C.BlobFatalError
	FATAL_ERROR_STREAM           ExceptionType = C.StreamFatalError
	FATAL_ERROR_CACHE            ExceptionType = C.CacheFatalError
	FATAL_ERROR_CODER            ExceptionType = C.CoderFatalError
	FATAL_ERROR_FILTER           ExceptionType = C.FilterFatalError
	FATAL_ERROR_MODULE           ExceptionType = C.ModuleFatalError
	FATAL_ERROR_DRAW             ExceptionType = C.DrawFatalError
	FATAL_ERROR_IMAGE            ExceptionType = C.ImageFatalError
	FATAL_ERROR_WAND             ExceptionType = C.WandFatalError
	FATAL_ERROR_RANDOM           ExceptionType = C.RandomFatalError
	FATAL_ERROR_XSERVER          ExceptionType = C.XServerFatalError
	FATAL_ERROR_MONITOR          ExceptionType = C.MonitorFatalError
	FATAL_ERROR_REGISTRY         ExceptionType = C.RegistryFatalError
	FATAL_ERROR_CONFIGURE        ExceptionType = C.ConfigureFatalError
	FATAL_ERROR_POLICY           ExceptionType = C.PolicyFatalError
)

var exceptionTypeStrings = map[ExceptionType]string{
	EXCEPTION_UNDEFINED:          "EXCEPTION_UNDEFINED",
	EXCEPTION_WARNING:            "EXCEPTION_WARNING",
	WARNING_TYPE:                 "WARNING_TYPE",
	WARNING_OPTION:               "WARNING_OPTION",
	WARNING_DELEGATE:             "WARNING_DELEGATE",
	WARNING_MISSING_DELEGATE:     "WARNING_MISSING_DELEGATE",
	WARNING_CORRUPT_IMAGE:        "WARNING_CORRUPT_IMAGE",
	WARNING_FILE_OPEN:            "WARNING_FILE_OPEN",
	WARNING_BLOB:                 "WARNING_BLOB",
	WARNING_STREAM:               "WARNING_STREAM",
	WARNING_CACHE:                "WARNING_CACHE",
	WARNING_CODER:                "WARNING_CODER",
	WARNING_FILTER:               "WARNING_FILTER",
	WARNING_MODULE:               "WARNING_MODULE",
	WARNING_DRAW:                 "WARNING_DRAW",
	WARNING_IMAGE:                "WARNING_IMAGE",
	WARNING_WAND:                 "WARNING_WAND",
	WARNING_RANDOM:               "WARNING_RANDOM",
	WARNING_XSERVER:              "WARNING_XSERVER",
	WARNING_MONITOR:              "WARNING_MONITOR",
	WARNING_REGISTRY:             "WARNING_REGISTRY",
	WARNING_CONFIGURE:            "WARNING_CONFIGURE",
	WARNING_POLICY:               "WARNING_POLICY",
	EXCEPTION_ERROR:              "EXCEPTION_ERROR",
	ERROR_TYPE:                   "ERROR_TYPE",
	ERROR_OPTION:                 "ERROR_OPTION",
	ERROR_DELEGATE:               "ERROR_DELEGATE",
	ERROR_MISSING_DELEGATE:       "ERROR_MISSING_DELEGATE",
	ERROR_CORRUPT_IMAGE:          "ERROR_CORRUPT_IMAGE",
	ERROR_FILE_OPEN:              "ERROR_FILE_OPEN",
	ERROR_BLOB:                   "ERROR_BLOB",
	ERROR_STREAM:                 "ERROR_STREAM",
	ERROR_CACHE:                  "ERROR_CACHE",
	ERROR_CODER:                  "ERROR_CODER",
	ERROR_FILTER:                 "ERROR_FILTER",
	ERROR_MODULE:                 "ERROR_MODULE",
	ERROR_DRAW:                   "ERROR_DRAW",
	ERROR_IMAGE:                  "ERROR_IMAGE",
	ERROR_WAND:                   "ERROR_WAND",
	ERROR_RANDOM:                 "ERROR_RANDOM",
	ERROR_XSERVER:                "ERROR_XSERVER",
	ERROR_MONITOR:                "ERROR_MONITOR",
	ERROR_REGISTRY:               "ERROR_REGISTRY",
	ERROR_CONFIGURE:              "ERROR_CONFIGURE",
	ERROR_POLICY:                 "ERROR_POLICY",
	EXCEPTION_FATAL_ERROR:        "EXCEPTION_FATAL_ERROR",
	FATAL_ERROR_TYPE:             "FATAL_ERROR_TYPE",
	FATAL_ERROR_OPTION:           "FATAL_ERROR_OPTION",
	FATAL_ERROR_DELEGATE:         "FATAL_ERROR_DELEGATE",
	FATAL_ERROR_MISSING_DELEGATE: "FATAL_ERROR_MISSING_DELEGATE",
	FATAL_ERROR_CORRUPT_IMAGE:    "FATAL_ERROR_CORRUPT_IMAGE",
	FATAL_ERROR_FILE_OPEN:        "FATAL_ERROR_FILE_OPEN",
	FATAL_ERROR_BLOB:             "FATAL_ERROR_BLOB",
	FATAL_ERROR_STREAM:           "FATAL_ERROR_STREAM",
	FATAL_ERROR_CACHE:            "FATAL_ERROR_CACHE",
	FATAL_ERROR_CODER:            "FATAL_ERROR_CODER",
	FATAL_ERROR_FILTER:           "FATAL_ERROR_FILTER",
	FATAL_ERROR_MODULE:           "FATAL_ERROR_MODULE",
	FATAL_ERROR_DRAW:             "FATAL_ERROR_DRAW",
	FATAL_ERROR_IMAGE:            "FATAL_ERROR_IMAGE",
	FATAL_ERROR_WAND:             "FATAL_ERROR_WAND",
	FATAL_ERROR_RANDOM:           "FATAL_ERROR_RANDOM",
	FATAL_ERROR_XSERVER:          "FATAL_ERROR_XSERVER",
	FATAL_ERROR_MONITOR:          "FATAL_ERROR_MONITOR",
	FATAL_ERROR_REGISTRY:         "FATAL_ERROR_REGISTRY",
	FATAL_ERROR_CONFIGURE:        "FATAL_ERROR_CONFIGURE",
	FATAL_ERROR_POLICY:           "FATAL_ERROR_POLICY",
	//WARNING_RESOURCE_LIMIT:       "WARNING_RESOURCE_LIMIT",
	//ERROR_RESOURCE_LIMIT:         "ERROR_RESOURCE_LIMIT",
	//FATAL_ERROR_RESOURCE_LIMIT:   "FATAL_ERROR_RESOURCE_LIMIT",
}

func (et *ExceptionType) String() string {
	if v, ok := exceptionTypeStrings[ExceptionType(*et)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownError[%d]", *et)
}
