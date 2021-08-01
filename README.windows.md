# Install Instruction For Windows

## step 1: install msys2

Msys2 is required to build ImageMagick dll. You can download from [Msys2](https://www.msys2.org/).

Follow the install guide described in the page.

## step 2: install build tools and requirements.

Open msys2 shell.

Run commands below.

If you want to manually confirm installation, please remove the `--noconfirm` option.

1. Refresh package database

    ```shell
    pacman -Syuu
    ```

2. install gcc

    ```shell
    pacman -S --noconfirm mingw-w64-x86_64-gcc
    ```

3. install pkg-config

    ```shell
    pacman -S --noconfirm mingw-w64-x86_64-pkg-config
    ```

4. install zlib

    ```shell
    pacman -S --noconfirm mingw-w64-x86_64-zlib
    ```

5. install imagemagick

    ```shell
    pacman -S --noconfirm mingw-w64-x86_64-imagemagick
    ```

> we install the latest ImageMagick version here, which is **7.x** . This is very important.

## step 3: set system environment variables.

There are many ways to set env variables. Through Windows GUI or by command. In this instruction,
we use `PowerShell`. If you are not sure about the right command in your shell environment, please
use Windows GUI.

1. enable msys2 shell to inherit environment variables to enable calling `go build` command in msys2 shell from windows. set `MSYS2_PATH_TYPE` to `inherit`

    ```powershell
    $env:MSYS2_PATH_TYPE="inherit"
    ```

2. add `msys2_shell.cmd` to `Path` to enable invoking msys2 shell from anywhere. File `msys2_shell.cmd` is in the root of msys2 install directory, just add it to your system `Path` variable.


    ```powershell
    $env:Path += ";<your_msys2_root>"
    ```

3. set decoders variable. This is very important for building. The decoders need to be included at build time for your code to recognize different image formats.

    ```powershell
    $env:MAGICK_CODER_MODULE_PATH="<your_msys2_root>/mingw64/lib/ImageMagick-7.x.x/modules-Q16HDRI/coders"
    ```

    > If you fail to set this variable, no error would be reported at build time, but this might cause `no decode delegate for this image format` error at run time.

## step 4: start building.

1. invoke the mingw64 shell in msys2 in your `GOPATH` or any directory with `go.mod`.

    ```shell
    msys2_shell.cmd -mingw64 -here
    ```

**Commands below should be finished in the invoked mingw64 shell.**

2. check for `ImageMagick` version. Please note the version here. Go binding version differs from different
`ImageMagick` version. if any error occurs, please try to reinstall `ImageMagic`.

    ```shell
    convert --version
    ```

3. check for `pkg-config` recognizing the `ImageMagick` header file. if any error occurs, please try to reinstall
`ImageMagic` and `pkg-config` in step 2.

    ```
    pkg-config --cflags --libs MagickWand
    ```

4. build go binding.

    Please choose the correct binding version for your installed `ImageMagick` version.


    ```shell
    master (tag v2.x.x): 6.9.1-7 <= ImageMagick <= 6.9.9-35
    im-7   (tag v3.x.x): 7.x     <= ImageMagick <= 7.x
    legacy (tag v1.x.x): 6.7.x   <= ImageMagick <= 6.8.9-10
    ```

    ```shell
    gopkg.in/gographics/imagick.v2/imagick
    gopkg.in/gographics/imagick.v3/imagick
    gopkg.in/gographics/imagick.v1/imagick
    ```

    in this instruction, we should install `v3`, as it aligned with ImageMagick7.

    ```shell
    go build gopkg.in/gographics/imagick.v3/imagick
    ```

    If any reported header file error like

    ```shell
    fatal error: wand/MagickWand.h: No such file or directory
    ```

    Please check that you are using the correct binding version to match the installed `ImageMagick` version.

5. install the package.

    ```shell
    go get gopkg.in/gographics/imagick.v3/imagick
    ```

## Note

1. every time you build or run your code, you should do it in the invoked `mingw64` shell.

2. If you meet some errors like [this](https://github.com/golang/go/issues/37553)

   ```shell
   undefined references to __errno in errno_itself
   ```
   
   It may be caused by a badly cached cgo file. If you build the wrong version for your `ImageMagick`, you might meet this error. Please clean the cache:
   
   ```shell
   go clean -cache
   ```
   
   then choose the correct version to rebuild. option `-a` here is to force rebuild the package to prevent using the cached cgo file.
   
   ```shell
   go build -a gopkg.in/gographics/imagick.<right-version>/imagick
   ```
   
   then install again
   
   ```shell
   go get gopkg.in/gographics/imagick.<right-version>/imagick
   ```
