package qtcore

// /usr/include/qt/QtCore/qwaitcondition.h
// #include <qwaitcondition.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QWaitCondition struct {
	*qtrt.CObject
}
type QWaitCondition_ITF interface {
	QWaitCondition_PTR() *QWaitCondition
}

func (ptr *QWaitCondition) QWaitCondition_PTR() *QWaitCondition { return ptr }

func (this *QWaitCondition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWaitCondition) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWaitConditionFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return &QWaitCondition{&qtrt.CObject{cthis}}
}
func (*QWaitCondition) NewFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return NewQWaitConditionFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10611() {
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
}

//  keep block end
