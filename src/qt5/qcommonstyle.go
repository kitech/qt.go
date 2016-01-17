package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qcommonstyle.h
// dst-file: /src/widgets/qcommonstyle.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QCommonStyle::polish(QWidget * widget);
extern void _ZN12QCommonStyle6polishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::polish(QPalette & );
extern void _ZN12QCommonStyle6polishER8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::polish(QApplication * app);
extern void _ZN12QCommonStyle6polishEP12QApplication(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::QCommonStyle();
extern void _ZN12QCommonStyleC2Ev(void* qthis); // 3
  // proto:  void QCommonStyle::~QCommonStyle();
extern void _ZN12QCommonStyleD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QCommonStyle::metaObject();
extern void _ZNK12QCommonStyle10metaObjectEv(void* qthis); // 4
  // proto:  void QCommonStyle::unpolish(QWidget * widget);
extern void _ZN12QCommonStyle8unpolishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::unpolish(QApplication * application);
extern void _ZN12QCommonStyle8unpolishEP12QApplication(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QCommonStyle)=1
type QCommonStyle struct {
  /*qbase*/ QStyle;
  qclsinst unsafe.Pointer /* *C.void */;
}

// polish(class QWidget *)
func (this *QCommonStyle) polish(args ...interface{}) () {
  // polish(class QWidget *)
  // polish(class QPalette &)
  // polish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPalette{}) // "QPalette &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QCommonStyle6polishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QCommonStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QCommonStyle6polishER8QPalette(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN12QCommonStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QCommonStyle6polishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommonStyle", "polish", args)
  }

}

// QCommonStyle()
func NewQCommonStyle(args ...interface{}) QCommonStyle {
  // QCommonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyleC1Ev
    // invoke: void QCommonStyle()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QCommonStyleC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QCommonStyle", "QCommonStyle", args)
  }

  return QCommonStyle{}
}

// ~QCommonStyle()
func (this *QCommonStyle) FreeQCommonStyle(args ...interface{}) () {
  // ~QCommonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyleD0Ev
    // invoke: void ~QCommonStyle()
    C._ZN12QCommonStyleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommonStyle", "~QCommonStyle", args)
  }

}

// metaObject()
func (this *QCommonStyle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QCommonStyle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK12QCommonStyle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommonStyle", "metaObject", args)
  }

}

// unpolish(class QWidget *)
func (this *QCommonStyle) unpolish(args ...interface{}) () {
  // unpolish(class QWidget *)
  // unpolish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyle8unpolishEP7QWidget
    // invoke: void unpolish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QCommonStyle8unpolishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QCommonStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QCommonStyle8unpolishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommonStyle", "unpolish", args)
  }

}

// <= body block end

