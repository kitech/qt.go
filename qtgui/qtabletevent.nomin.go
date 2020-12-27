//  header block begin

// +build !minimal

package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qevent.h:306
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int x() const

/*
 */
func (this *QTabletEvent) X() int {
	rv, err := qtrt.Qtcc3(40556910, "_ZNK12QTabletEvent1xEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:307
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int y() const

/*
 */
func (this *QTabletEvent) Y() int {
	rv, err := qtrt.Qtcc3(61387609, "_ZNK12QTabletEvent1yEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:308
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int globalX() const

/*
 */
func (this *QTabletEvent) GlobalX() int {
	rv, err := qtrt.Qtcc3(4236764281, "_ZNK12QTabletEvent7globalXEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:309
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int globalY() const

/*
 */
func (this *QTabletEvent) GlobalY() int {
	rv, err := qtrt.Qtcc3(4249194062, "_ZNK12QTabletEvent7globalYEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

func DeleteQTabletEvent(this *QTabletEvent) {
	rv, err := qtrt.Qtcc3(1745538712, "_ZN12QTabletEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10078() {
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
}

//  keep block end
