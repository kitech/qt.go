//  header block begin
// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QTextStreamManipulator struct {
	*qtrt.CObject
}

func (this *QTextStreamManipulator) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qtextstream.h:217
// index:0
// inline
// void exec(class QTextStream &)
func (this *QTextStreamManipulator) Exec(s unsafe.Pointer) {
	// 0: (, s QTextStream &), (s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QTextStreamManipulator4execER11QTextStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s)
	gopp.ErrPrint(err, rv)
}

//  body block end