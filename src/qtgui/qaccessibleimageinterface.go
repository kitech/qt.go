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
func NewQAccessibleImageInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleImageInterface {
	return &QAccessibleImageInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:655
// index:0
// Public virtual
// void ~QAccessibleImageInterface()
func DeleteQAccessibleImageInterface(*QAccessibleImageInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleImageInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:657
// index:0
// Public pure virtual
// QString imageDescription()
func (this *QAccessibleImageInterface) ImageDescription() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface16imageDescriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:658
// index:0
// Public pure virtual
// QSize imageSize()
func (this *QAccessibleImageInterface) ImageSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface9imageSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:659
// index:0
// Public pure virtual
// QPoint imagePosition()
func (this *QAccessibleImageInterface) ImagePosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface13imagePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
