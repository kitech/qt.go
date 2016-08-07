package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QCommonStyle::polish(QWidget * widget);
extern void C_ZN12QCommonStyle6polishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::polish(QPalette & );
extern void C_ZN12QCommonStyle6polishER8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::polish(QApplication * app);
extern void C_ZN12QCommonStyle6polishEP12QApplication(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::QCommonStyle();
extern void* C_ZN12QCommonStyleC2Ev(); // 3
  // proto:  void QCommonStyle::~QCommonStyle();
extern void C_ZN12QCommonStyleD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QCommonStyle::metaObject();
extern void C_ZNK12QCommonStyle10metaObjectEv(void* qthis); // 4
  // proto:  void QCommonStyle::unpolish(QWidget * widget);
extern void C_ZN12QCommonStyle8unpolishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QCommonStyle::unpolish(QApplication * application);
extern void C_ZN12QCommonStyle8unpolishEP12QApplication(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QCommonStyle)=1
type QCommonStyle struct {
  /*qbase*/ QStyle;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// polish(class QWidget *)
func (this *QCommonStyle) Polish(args ...interface{}) () {
  // polish(class QWidget *)
  // polish(class QPalette &)
  // polish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QPalette{}) // "QPalette &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QApplication{}) // "QApplication *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QCommonStyle6polishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QCommonStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(*qtgui.QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QCommonStyle6polishER8QPalette(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN12QCommonStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QCommonStyle6polishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommonStyle", "polish", args)
  }

  return
}

// QCommonStyle()
func NewQCommonStyle(args ...interface{}) *QCommonStyle {
  // QCommonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyleC1Ev
    // invoke: void QCommonStyle()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QCommonStyleC2Ev()
    return &QCommonStyle{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCommonStyle", "QCommonStyle", args)
  }

  return nil // QCommonStyle{Qclsinst:qthis}
}

// ~QCommonStyle()
func (this *QCommonStyle) Freeqcommonstyle(args ...interface{}) () {
  // ~QCommonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyleD0Ev
    // invoke: void ~QCommonStyle()
    C.C_ZN12QCommonStyleD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommonStyle", "~QCommonStyle", args)
  }

  return
}

// metaObject()
func (this *QCommonStyle) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QCommonStyle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QCommonStyle10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommonStyle", "metaObject", args)
  }

  return
}

// unpolish(class QWidget *)
func (this *QCommonStyle) Unpolish(args ...interface{}) () {
  // unpolish(class QWidget *)
  // unpolish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QApplication{}) // "QApplication *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QCommonStyle8unpolishEP7QWidget
    // invoke: void unpolish(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QCommonStyle8unpolishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QCommonStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QCommonStyle8unpolishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommonStyle", "unpolish", args)
  }

  return
}

// <= body block end

