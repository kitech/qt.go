//  header block begin
// /usr/include/qt/QtWidgets/qscrollerproperties.h
// #include <qscrollerproperties.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QScrollerProperties struct {
	*qtrt.CObject
}

func (this *QScrollerProperties) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:60
// index:0
// void QScrollerProperties()
func NewQScrollerProperties() *QScrollerProperties {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerPropertiesC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollerPropertiesFromPointer(cthis)
	return gothis
}
func NewQScrollerPropertiesFromPointer(cthis unsafe.Pointer) *QScrollerProperties {
	return &QScrollerProperties{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:63
// index:0
// virtual
// void ~QScrollerProperties()
func DeleteQScrollerProperties(*QScrollerProperties) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerPropertiesD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:68
// index:0
// static
// void setDefaultScrollerProperties(const class QScrollerProperties &)
func (this *QScrollerProperties) SetDefaultScrollerProperties(sp unsafe.Pointer) {
	// 0: (sp const QScrollerProperties &), (sp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScrollerProperties_SetDefaultScrollerProperties(sp unsafe.Pointer) {
	// 0: (sp const QScrollerProperties &), (sp)
	var nilthis *QScrollerProperties
	nilthis.SetDefaultScrollerProperties(sp)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:69
// index:0
// static
// void unsetDefaultScrollerProperties()
func (this *QScrollerProperties) UnsetDefaultScrollerProperties() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScrollerProperties_UnsetDefaultScrollerProperties() {
	// 0: (), ()
	var nilthis *QScrollerProperties
	nilthis.UnsetDefaultScrollerProperties()
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:117
// index:0
// QVariant scrollMetric(enum QScrollerProperties::ScrollMetric)
func (this *QScrollerProperties) ScrollMetric(metric int) {
	// 0: (, metric QScrollerProperties::ScrollMetric), (&metric)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollerProperties12scrollMetricENS_12ScrollMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:118
// index:0
// void setScrollMetric(enum QScrollerProperties::ScrollMetric, const class QVariant &)
func (this *QScrollerProperties) SetScrollMetric(metric int, value unsafe.Pointer) {
	// 0: (, metric QScrollerProperties::ScrollMetric, value const QVariant &), (&metric, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties15setScrollMetricENS_12ScrollMetricERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric, value)
	gopp.ErrPrint(err, rv)
}

//  body block end
