package qtgui

// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
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
// size 8
type QFontMetrics struct {
	*qtrt.CObject
}
type QFontMetrics_ITF interface {
	QFontMetrics_PTR() *QFontMetrics
}

func (ptr *QFontMetrics) QFontMetrics_PTR() *QFontMetrics { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QFontMetricsFromptr(cthis Voidptr) *QFontMetrics {
	return &QFontMetrics{&qtrt.CObject{cthis}}
}
func (*QFontMetrics) Fromptr(cthis Voidptr) *QFontMetrics {
	return QFontMetricsFromptr(cthis)
}

func DeleteQFontMetrics(this *QFontMetrics) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QFontMetricsD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10073() {
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
