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

// /usr/include/qt/QtGui/qtextoption.h:85
// index:0
// void QTextOption()
func NewQTextOption() *QTextOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOptionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextOptionFromPointer(cthis)
	return gothis
}
func NewQTextOptionFromPointer(cthis unsafe.Pointer) *QTextOption {
	return &QTextOption{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextoption.h:86
// index:1
// void QTextOption(Qt::Alignment)
func NewQTextOption_1(alignment int) *QTextOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOptionC2E6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, cthis, &alignment)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextoption.h:87
// index:0
// void ~QTextOption()
func DeleteQTextOption(*QTextOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:92
// index:0
// inline
// void setAlignment(Qt::Alignment)
func (this *QTextOption) SetAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:93
// index:0
// inline
// Qt::Alignment alignment()
func (this *QTextOption) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:95
// index:0
// inline
// void setTextDirection(Qt::LayoutDirection)
func (this *QTextOption) SetTextDirection(aDirection int) {
	// 0: (, aDirection Qt::LayoutDirection), (&aDirection)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption16setTextDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &aDirection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:96
// index:0
// inline
// Qt::LayoutDirection textDirection()
func (this *QTextOption) TextDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption13textDirectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:105
// index:0
// inline
// void setWrapMode(enum QTextOption::WrapMode)
func (this *QTextOption) SetWrapMode(wrap int) {
	// 0: (, wrap QTextOption::WrapMode), (&wrap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption11setWrapModeENS_8WrapModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &wrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:106
// index:0
// inline
// QTextOption::WrapMode wrapMode()
func (this *QTextOption) WrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption8wrapModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:117
// index:0
// inline
// void setFlags(QTextOption::Flags)
func (this *QTextOption) SetFlags(flags int) {
	// 0: (, QFlags<QTextOption::Flag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption8setFlagsE6QFlagsINS_4FlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:118
// index:0
// inline
// QTextOption::Flags flags()
func (this *QTextOption) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption5flagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:121
// index:0
// inline
// void setTabStop(qreal)
func (this *QTextOption) SetTabStop(tabStop float64) {
	// 0: (, tabStop qreal), (&tabStop)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption10setTabStopEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tabStop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:122
// index:0
// inline
// qreal tabStop()
func (this *QTextOption) TabStop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption7tabStopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:125
// index:0
// inline
// void setTabStopDistance(qreal)
func (this *QTextOption) SetTabStopDistance(tabStopDistance float64) {
	// 0: (, tabStopDistance qreal), (&tabStopDistance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption18setTabStopDistanceEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tabStopDistance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:126
// index:0
// inline
// qreal tabStopDistance()
func (this *QTextOption) TabStopDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption15tabStopDistanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:129
// index:0
// QList<qreal> tabArray()
func (this *QTextOption) TabArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption8tabArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:132
// index:0
// QList<QTextOption::Tab> tabs()
func (this *QTextOption) Tabs() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption4tabsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:134
// index:0
// inline
// void setUseDesignMetrics(_Bool)
func (this *QTextOption) SetUseDesignMetrics(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextOption19setUseDesignMetricsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:135
// index:0
// inline
// bool useDesignMetrics()
func (this *QTextOption) UseDesignMetrics() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextOption16useDesignMetricsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
