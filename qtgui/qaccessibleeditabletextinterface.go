package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QAccessibleEditableTextInterface struct {
	*qtrt.CObject
}
type QAccessibleEditableTextInterface_ITF interface {
	QAccessibleEditableTextInterface_PTR() *QAccessibleEditableTextInterface
}

func (ptr *QAccessibleEditableTextInterface) QAccessibleEditableTextInterface_PTR() *QAccessibleEditableTextInterface {
	return ptr
}

func (this *QAccessibleEditableTextInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleEditableTextInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleEditableTextInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleEditableTextInterface {
	return &QAccessibleEditableTextInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleEditableTextInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleEditableTextInterface {
	return NewQAccessibleEditableTextInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:556
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleEditableTextInterface()
func DeleteQAccessibleEditableTextInterface(this *QAccessibleEditableTextInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:558
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void deleteText(int, int)
func (this *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10deleteTextEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:559
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void insertText(int, const QString &)
func (this *QAccessibleEditableTextInterface) InsertText(offset int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:560
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void replaceText(int, int, const QString &)
func (this *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text string) {
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset, convArg2)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
