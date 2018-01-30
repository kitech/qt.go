package qtwidgets

// /usr/include/qt/QtWidgets/qdrawutil.h
// #include <qdrawutil.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QTileRules struct {
	*qtrt.CObject
}

func (this *QTileRules) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTileRules) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTileRulesFromPointer(cthis unsafe.Pointer) *QTileRules {
	return &QTileRules{&qtrt.CObject{cthis}}
}
func (*QTileRules) NewFromPointer(cthis unsafe.Pointer) *QTileRules {
	return NewQTileRulesFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTileRules(Qt::TileRule, Qt::TileRule)
func NewQTileRules(horizontalRule int, verticalRule int) *QTileRules {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTileRulesC2EN2Qt8TileRuleES1_", ffiqt.FFI_TYPE_POINTER, horizontalRule, verticalRule)
	gopp.ErrPrint(err, rv)
	gothis := NewQTileRulesFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdrawutil.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QTileRules(Qt::TileRule)
func NewQTileRules_1(rule int) *QTileRules {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTileRulesC2EN2Qt8TileRuleE", ffiqt.FFI_TYPE_POINTER, rule)
	gopp.ErrPrint(err, rv)
	gothis := NewQTileRulesFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

//  body block end
