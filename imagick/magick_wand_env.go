// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"
import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	initOnce      sync.Once
	terminateOnce *sync.Once

	// Indicates that terminate method can be called (there are no any ImageMagick objects)
	canTerminate = make(chan struct{}, 1)

	envSemaphore = make(chan struct{}, 1)

	// Ref counters
	magickWandCounter    int64
	drawingWandCounter   int64
	pixelIteratorCounter int64
	pixelWandCounter     int64
)

// Initializes the MagickWand environment
func Initialize() {
	envSemaphore <- struct{}{}
	defer func() {
		<-envSemaphore
	}()

	initOnce.Do(func() {
		C.MagickWandGenesis()
		terminateOnce = &sync.Once{}
		setCanTerminate()
	})
}

// Terminates the MagickWand environment
// wait until all imageMagick objects destroyed
func Terminate() {
	envSemaphore <- struct{}{}
	defer func() {
		<-envSemaphore
	}()

	if terminateOnce != nil {
		terminateOnce.Do(func() {
			runtime.GC()
			terminate()
		})
	}
}

// Guarantees that C.MagickWandTerminus run after all ImageMagick objects are destroyed
// or we got the panic from C.MagickWandTerminus
func terminate() {
	<-canTerminate
	C.MagickWandTerminus()
	initOnce = sync.Once{}
}

// Set status "terminate can be called"
func setCanTerminate() {
	if isImageMagickCleaned() {
		select {
		case canTerminate <- struct{}{}:
			// Now we can terminate
		default:
			// Nothing to do
		}
	}
}

// Set status "terminate can`t be called"
func unsetCanTerminate() {
	select {
	case <-canTerminate:
		// Now we can`t terminate
	default:
		// Nothing to do
	}
}

// Check are all IM objects are collected by GC
func isImageMagickCleaned() bool {
	if atomic.LoadInt64(&magickWandCounter) != 0 || atomic.LoadInt64(&drawingWandCounter) != 0 || atomic.LoadInt64(&pixelIteratorCounter) != 0 || atomic.LoadInt64(&pixelWandCounter) != 0 {
		return false
	}

	return true
}
