//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlabel.h
// #include <qlabel.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qlabel.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QMovie * movie() const

/*
Returns a pointer to the label's movie, or nullptr if no movie has been set.

See also setMovie().
*/
func (this *QLabel) Movie() *qtgui.QMovie /*777 QMovie **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel5movieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQMovieFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovie(QMovie *)

/*
Sets the label contents to movie. Any previous content is cleared. The label does NOT take ownership of the movie.

The buddy shortcut, if any, is disabled.

See also movie() and setBuddy().
*/
func (this *QLabel) SetMovie(movie qtgui.QMovie_ITF /*777 QMovie **/) {
	var convArg0 unsafe.Pointer
	if movie != nil && movie.QMovie_PTR() != nil {
		convArg0 = movie.QMovie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel8setMovieEP6QMovie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11264() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
