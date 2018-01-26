package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleImageInterface) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleImageInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleImageInterface {
	return &QAccessibleImageInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleImageInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleImageInterface {
	return NewQAccessibleImageInterfaceFromPointer(cthis)
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
func (this *QAccessibleImageInterface) ImageDescription() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface16imageDescriptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:658
// index:0
// Public pure virtual
// QSize imageSize()
func (this *QAccessibleImageInterface) ImageSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface9imageSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:659
// index:0
// Public pure virtual
// QPoint imagePosition()
func (this *QAccessibleImageInterface) ImagePosition() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface13imagePositionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
