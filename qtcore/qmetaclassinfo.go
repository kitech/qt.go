package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 16
type QMetaClassInfo struct {
	*qtrt.CObject
}
type QMetaClassInfo_ITF interface {
	QMetaClassInfo_PTR() *QMetaClassInfo
}

func (ptr *QMetaClassInfo) QMetaClassInfo_PTR() *QMetaClassInfo { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QMetaClassInfoFromptr(cthis Voidptr) *QMetaClassInfo {
	return &QMetaClassInfo{&qtrt.CObject{cthis}}
}
func (*QMetaClassInfo) Fromptr(cthis Voidptr) *QMetaClassInfo {
	return QMetaClassInfoFromptr(cthis)
}

func DeleteQMetaClassInfo(this *QMetaClassInfo) {
	rv, err := qtrt.Qtcc1(0, "_ZN14QMetaClassInfoD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10037() {
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
