package qtgui

// /usr/include/qt/QtGui/qguiapplication.h
// #include <qguiapplication.h>
// #include <QtGui>

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
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 16
type QGuiApplication struct {
	*qtcore.QCoreApplication
}
type QGuiApplication_ITF interface {
	qtcore.QCoreApplication_ITF
	QGuiApplication_PTR() *QGuiApplication
}

func (ptr *QGuiApplication) QGuiApplication_PTR() *QGuiApplication { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QGuiApplicationFromptr(cthis Voidptr) *QGuiApplication {
	bcthis0 := qtcore.QCoreApplicationFromptr(cthis)
	return &QGuiApplication{bcthis0}
}
func (*QGuiApplication) Fromptr(cthis Voidptr) *QGuiApplication {
	return QGuiApplicationFromptr(cthis)
}

func DeleteQGuiApplication(this *QGuiApplication) {
	rv, err := qtrt.Qtcc3(2140321792, "_ZN15QGuiApplicationD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10169() {
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
