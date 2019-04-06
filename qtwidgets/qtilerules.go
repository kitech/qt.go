package qtwidgets

// /usr/include/qt/QtWidgets/qdrawutil.h
// #include <qdrawutil.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QTileRules struct {
	*qtrt.CObject
}
type QTileRules_ITF interface {
	QTileRules_PTR() *QTileRules
}

func (ptr *QTileRules) QTileRules_PTR() *QTileRules { return ptr }

func (this *QTileRules) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTileRules) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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

/*

 */
func (*QTileRules) NewForInherit(horizontalRule int, verticalRule int) *QTileRules {
	return NewQTileRules(horizontalRule, verticalRule)
}
func NewQTileRules(horizontalRule int, verticalRule int) *QTileRules {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTileRulesC2EN2Qt8TileRuleES1_", qtrt.FFI_TYPE_POINTER, horizontalRule, verticalRule)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTileRulesFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTileRules)
	return gothis
}

// /usr/include/qt/QtWidgets/qdrawutil.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QTileRules(Qt::TileRule)

/*

 */
func (*QTileRules) NewForInherit1(rule int) *QTileRules {
	return NewQTileRules1(rule)
}
func NewQTileRules1(rule int) *QTileRules {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTileRulesC2EN2Qt8TileRuleE", qtrt.FFI_TYPE_POINTER, rule)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTileRulesFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTileRules)
	return gothis
}

// /usr/include/qt/QtWidgets/qdrawutil.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QTileRules(Qt::TileRule)

/*

 */
func (*QTileRules) NewForInherit1p() *QTileRules {
	return NewQTileRules1p()
}
func NewQTileRules1p() *QTileRules {
	// arg: 0, Qt::TileRule=Elaborated, Qt::TileRule=Enum, , Invalid
	rule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTileRulesC2EN2Qt8TileRuleE", qtrt.FFI_TYPE_POINTER, rule)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTileRulesFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTileRules)
	return gothis
}

func DeleteQTileRules(this *QTileRules) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTileRulesD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11133() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
