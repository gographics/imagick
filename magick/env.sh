export CGO_CFLAGS="-I`pkg-config --cflags MagickWand`"
export CGO_LDFLAGS="-I`pkg-config --libs MagickWand`"