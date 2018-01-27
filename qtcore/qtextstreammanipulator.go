package qtcore

// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextStreamManipulator) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextStreamManipulatorFromPointer(cthis unsafe.Pointer) *QTextStreamManipulator {
	return &QTextStreamManipulator{&qtrt.CObject{cthis}}
}
func (*QTextStreamManipulator) NewFromPointer(cthis unsafe.Pointer) *QTextStreamManipulator {
	return NewQTextStreamManipulatorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextstream.h:217
// index:0
// Public inline
// void exec(QTextStream &)
func (this *QTextStreamManipulator) Exec(s *QTextStream) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QTextStreamManipulator4execER11QTextStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
