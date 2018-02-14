package qtcore

// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QTextStreamManipulator struct {
	*qtrt.CObject
}
type QTextStreamManipulator_ITF interface {
	QTextStreamManipulator_PTR() *QTextStreamManipulator
}

func (ptr *QTextStreamManipulator) QTextStreamManipulator_PTR() *QTextStreamManipulator { return ptr }

func (this *QTextStreamManipulator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextStreamManipulator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextStreamManipulatorFromPointer(cthis unsafe.Pointer) *QTextStreamManipulator {
	return &QTextStreamManipulator{&qtrt.CObject{cthis}}
}
func (*QTextStreamManipulator) NewFromPointer(cthis unsafe.Pointer) *QTextStreamManipulator {
	return NewQTextStreamManipulatorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextstream.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void exec(QTextStream &)
func (this *QTextStreamManipulator) Exec(s QTextStream_ITF) {
	var convArg0 = s.QTextStream_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QTextStreamManipulator4execER11QTextStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQTextStreamManipulator(this *QTextStreamManipulator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QTextStreamManipulatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
