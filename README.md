# Go Imagick

[![GoDoc](https://gopkg.in/gographics/imagick.v2/imagick?status.svg)](https://gopkg.in/gographics/imagick.v2/imagick)

Go Imagick is a Go bind to ImageMagick's MagickWand C API.

We support two compatibility branches:

```
master (tag v2.x.x): >= ImageMagick 6.8.9
legacy (tag v1.x.x): <= ImageMagick 6.8.8
```

The map, respectively, through gopkg.in:

```
gopkg.in/gographics/imagick.v2/imagick
gopkg.in/gographics/imagick.v1/imagick
```

# Install

## Mac OS X

### MacPorts

```
sudo port install ImageMagick
```

## Ubuntu / Debian

```
sudo apt-get install libmagickwand-dev
```

## Common

Check if pkg-config is able to find the right ImageMagick include and libs:

```
pkg-config --cflags --libs MagickWand
```

Then go get it:

```
go get gopkg.in/gographics/imagick.v2/imagick
```

### Build tags

If you want to specify CGO_CFLAGS/CGO_LDFLAGS manually at build time, such as for building statically or without pkg-config, you can use the "no_pkgconfig" build tag:

```
go build -tags no_pkgconfig gopkg.in/gographics/imagick.v2/imagick
```

# Examples

The examples folder is full with usage examples ported from C ones found in here: http://members.shaw.ca/el.supremo/MagickWand/

# Quick and partial example

Since this is a CGO binding, Go GC does not manage memory allocated by the C API then is necessary to use Terminate() and Destroy() methods.

```
package main

import "gopkg.in/gographics/imagick.v2/imagick"

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()
    defer mw.Destroy()

    ...
}
```

# License

Copyright (c) 2013-2014, The GoGraphics Team
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
