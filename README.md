# Go Imagick

[![GoDoc](https://godoc.org/gopkg.in/gographics/imagick.v2/imagick?status.svg)](https://gopkg.in/gographics/imagick.v2/imagick)
[![Build](https://github.com/gographics/imagick/workflows/Test/badge.svg)](https://github.com/gographics/imagick/actions?query=workflow%3A%22Test%22)

Go Imagick is a Go bind to ImageMagick's MagickWand C API.

We support two compatibility branches:

```
im-7   (tag v3.x.x): 7.x     <= ImageMagick <= 7.x
master (tag v2.x.x): 6.9.1-7 <= ImageMagick <= 6.9.9-35
legacy (tag v1.x.x): 6.7.x   <= ImageMagick <= 6.8.9-10
```

They map, respectively, through gopkg.in:

```
gopkg.in/gographics/imagick.v3/imagick
gopkg.in/gographics/imagick.v2/imagick
gopkg.in/gographics/imagick.v1/imagick
```

# Install

## Docker

See [examples/docker](examples/docker)

## Mac OS X

### MacPorts

```
brew install imagemagick
```

## Ubuntu / Debian

```
sudo apt-get install libmagickwand-dev
```

## Windows

*Thanks @vprus*

+ Install [msys2-x86_64](http://www.msys2.org/)
+ In msys shell, do:
```
pacman -Syuu
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-pkg-config
pacman -S mingw-w64-x86_64-zlib
pacman -S mingw-w64-x86_64-imagemagick
```

+ Switch to cmd.exe shell, and do:
```
set PATH=<msys64>\mingw64\bin;%PATH%
set PKG_CONFIG_PATH=<msys64>\mingw64\lib\pkgconfig
set MAGICK_CODER_MODULE_PATH=<msys64>\mingw64\lib\ImageMagick-7.0.6\modules-Q16HDRI\coders
go build gopkg.in/gographics/imagick.v3/imagick
```

The default installation path of `msys2` is `C:\msys64` and you must change
`<msys64>` to your installation path of `msys2`.

The `MAGICK_CODER_MODULE_PATH` environment variable tells ImageMagick where to
find the decoders. If you still get the `NoDecodeDelegateForThisImageFormat`
error, then make sure the version number and folders are correct.

## Common

Check if pkg-config is able to find the right ImageMagick include and libs:

```
pkg-config --cflags --libs MagickWand
```

Then go get it:

```
go get gopkg.in/gographics/imagick.v2/imagick
```

Per the security update https://groups.google.com/forum/#!topic/golang-announce/X7N1mvntnoU
you may need whitelist the -Xpreprocessor flag in your environment.

```
export CGO_CFLAGS_ALLOW='-Xpreprocessor'
```

### Build tags

If you want to specify CGO_CFLAGS/CGO_LDFLAGS manually at build time, such as for building statically or without pkg-config, you can use the "no_pkgconfig" build tag:

```
go build -tags no_pkgconfig gopkg.in/gographics/imagick.v2/imagick
```

# Examples

The examples folder is full with usage examples ported from C ones found in here: http://www.imagemagick.org/MagickWand/

# Quick and partial example

```go
package main

import "gopkg.in/gographics/imagick.v2/imagick"

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()

    ...
}
```
## `Initialize()` and `Terminate`

As per the ImageMagick C API, `Initialize()` should be called only once to set up the resources for using ImageMagick. This is typically done in your `main()` or `init()` for the entire application or library. Applications can defer a call to `Terminate()` to tear down the ImageMagick resources.

It is an error to `Initialize` and `Terminate` multiple times in specific functions and leads to common problems such as crashes or missing delegates. Do not use `Terminate` anywhere other than the absolute end of your need for ImageMagick within the program.

## Managing memory

Since this is a CGO binding, and the Go GC does not manage memory allocated by the C API, it is then necessary to use the Terminate() and Destroy() methods.

Types which are created via `New*` constructors (MagickWand, DrawingWand, PixelIterator, PixelWand,...), or created and returned by methods like `Clone()` are managed by Go GC through the use of finalizers. They are not guaranteed to be cleaned up immediately after the object is not longer in use, but rather at some point later when GC actually executes its finalizers.

If you use struct literals, you should free resources manually:

```go
package main

import "github.com/gographics/imagick/imagick"

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.MagickWand{...}
    defer mw.Destroy()

    ...
}
```

Both methods are compatible if constructor methods used:

```go
package main

import "github.com/gographics/imagick/imagick"

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()
    defer mw.Destroy()

    ...
}
```

But you should NOT mix two ways of object creation:
```go
package main

import "github.com/gographics/imagick/imagick"

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw1 := imagick.MagickWand{...}
    defer mw1.Destroy()

    mw2 := imagick.NewMagickWand()

    ...
}
```

Calling `Destroy()` on types that are either created via `New*` or returned from other functions calls forces the cleanup of the item immediately as opposed to later after garbage collection triggers the finalizer for the object. It would be good practice to explicitely call `Destroy()` to ensure C memory is freed sooner rather than later, depending on how often the GC is triggered.

# License

Copyright (c) 2013-2016, The GoGraphics Team
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
 * Neither the name of the organization nor the
   names of its contributors may be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL HERBERT G. FISCHER BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
