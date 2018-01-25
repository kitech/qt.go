package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QAccessibleEditableTextInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleEditableTextInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleEditableTextInterface) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleEditableTextInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleEditableTextInterface {
	return &QAccessibleEditableTextInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleEditableTextInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleEditableTextInterface {
	return NewQAccessibleEditableTextInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:556
// index:0
// Public virtual
// void ~QAccessibleEditableTextInterface()
func DeleteQAccessibleEditableTextInterface(*QAccessibleEditableTextInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:558
// index:0
// Public pure virtual
// void deleteText(int, int)
func (this *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10deleteTextEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:559
// index:0
// Public pure virtual
// void insertText(int, const class QString &)
func (this *QAccessibleEditableTextInterface) InsertText(offset int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), offset, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:560
// index:0
// Public pure virtual
// void replaceText(int, int, const class QString &)
func (this *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text *qtcore.QString) {
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset, convArg2)
	gopp.ErrPrint(err, rv)
}

//  body block end
