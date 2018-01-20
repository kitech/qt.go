//  header block begin
// /usr/include/qt/QtGui/qtextoption.h
// #include <qtextoption.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QTextOption struct {
	*qtrt.CObject
}

func (this *QTextOption) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextOptionFromPointer(cthis unsafe.Pointer) *QTextOption {
	return &QTextOption{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextoption.h:85
// index:0
// Public
// void QTextOption()
func NewQTextOption() *QTextOption {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOptionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextoption.h:87
// index:0
// Public
// void ~QTextOption()
func DeleteQTextOption(*QTextOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:93
// index:0
// Public inline
// Qt::Alignment alignment()
func (this *QTextOption) Alignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:95
// index:0
// Public inline
// void setTextDirection(Qt::LayoutDirection)
func (this *QTextOption) SetTextDirection(aDirection int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption16setTextDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &aDirection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:96
// index:0
// Public inline
// Qt::LayoutDirection textDirection()
func (this *QTextOption) TextDirection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption13textDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:105
// index:0
// Public inline
// void setWrapMode(enum QTextOption::WrapMode)
func (this *QTextOption) SetWrapMode(wrap int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption11setWrapModeENS_8WrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &wrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:106
// index:0
// Public inline
// QTextOption::WrapMode wrapMode()
func (this *QTextOption) WrapMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption8wrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:118
// index:0
// Public inline
// QTextOption::Flags flags()
func (this *QTextOption) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:121
// index:0
// Public inline
// void setTabStop(qreal)
func (this *QTextOption) SetTabStop(tabStop float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption10setTabStopEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tabStop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:122
// index:0
// Public inline
// qreal tabStop()
func (this *QTextOption) TabStop() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption7tabStopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:125
// index:0
// Public inline
// void setTabStopDistance(qreal)
func (this *QTextOption) SetTabStopDistance(tabStopDistance float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption18setTabStopDistanceEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tabStopDistance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:126
// index:0
// Public inline
// qreal tabStopDistance()
func (this *QTextOption) TabStopDistance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption15tabStopDistanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:129
// index:0
// Public
// QList<qreal> tabArray()
func (this *QTextOption) TabArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption8tabArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:132
// index:0
// Public
// QList<QTextOption::Tab> tabs()
func (this *QTextOption) Tabs() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption4tabsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextoption.h:134
// index:0
// Public inline
// void setUseDesignMetrics(_Bool)
func (this *QTextOption) SetUseDesignMetrics(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption19setUseDesignMetricsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:135
// index:0
// Public inline
// bool useDesignMetrics()
func (this *QTextOption) UseDesignMetrics() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption16useDesignMetricsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
