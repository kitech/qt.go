//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qaccessible.h:556
// index:0
// virtual
// void ~QAccessibleEditableTextInterface()
func DeleteQAccessibleEditableTextInterface(*QAccessibleEditableTextInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:558
// index:0
// pure virtual
// void deleteText(int, int)
func (this *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	// 0: (, startOffset int, endOffset int), (&startOffset, &endOffset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10deleteTextEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:559
// index:0
// pure virtual
// void insertText(int, const class QString &)
func (this *QAccessibleEditableTextInterface) InsertText(offset int, text unsafe.Pointer) {
	// 0: (, offset int, text const QString &), (&offset, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &offset, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:560
// index:0
// pure virtual
// void replaceText(int, int, const class QString &)
func (this *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text unsafe.Pointer) {
	// 0: (, startOffset int, endOffset int, text const QString &), (&startOffset, &endOffset, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startOffset, &endOffset, text)
	gopp.ErrPrint(err, rv)
}

//  body block end
