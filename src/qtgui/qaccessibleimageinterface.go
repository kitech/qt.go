//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QAccessibleImageInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleImageInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qaccessible.h:655
// index:0
// virtual
// void ~QAccessibleImageInterface()
func DeleteQAccessibleImageInterface(*QAccessibleImageInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleImageInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:657
// index:0
// pure virtual
// QString imageDescription()
func (this *QAccessibleImageInterface) ImageDescription() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface16imageDescriptionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:658
// index:0
// pure virtual
// QSize imageSize()
func (this *QAccessibleImageInterface) ImageSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface9imageSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:659
// index:0
// pure virtual
// QPoint imagePosition()
func (this *QAccessibleImageInterface) ImagePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface13imagePositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
