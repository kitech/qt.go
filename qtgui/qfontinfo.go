package qtgui

// /usr/include/qt/QtGui/qfontinfo.h
// #include <qfontinfo.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

/*
 */
// size 8
type QFontInfo struct {
	*qtrt.CObject
}
type QFontInfo_ITF interface {
	QFontInfo_PTR() *QFontInfo
}

func (ptr *QFontInfo) QFontInfo_PTR() *QFontInfo { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QFontInfoFromptr(cthis Voidptr) *QFontInfo {
	return &QFontInfo{&qtrt.CObject{cthis}}
}
func (*QFontInfo) Fromptr(cthis Voidptr) *QFontInfo {
	return QFontInfoFromptr(cthis)
}

func DeleteQFontInfo(this *QFontInfo) {
	rv, err := qtrt.Qtcc3(3740458797, "_ZN9QFontInfoD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10167() {
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
