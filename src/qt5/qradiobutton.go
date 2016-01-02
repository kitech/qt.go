package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qradiobutton.h
// dst-file: /src/widgets/qradiobutton.go
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QRadioButton::metaObject();
extern void _ZNK12QRadioButton10metaObjectEv(void* qthis);
  // proto:  void QRadioButton::QRadioButton(QWidget * parent);
extern void* dector_ZN12QRadioButtonC1EP7QWidget(void* arg0);
extern void _ZN12QRadioButtonC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QRadioButton::sizeHint();
extern void _ZNK12QRadioButton8sizeHintEv(void* qthis);
  // proto:  QSize QRadioButton::minimumSizeHint();
extern void _ZNK12QRadioButton15minimumSizeHintEv(void* qthis);
  // proto:  void QRadioButton::~QRadioButton();
extern void _ZN12QRadioButtonD0Ev(void* qthis);
  // proto:  void QRadioButton::QRadioButton(const QRadioButton & );
extern void* dector_ZN12QRadioButtonC1ERKS_(void* arg0);
extern void _ZN12QRadioButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QRadioButton::QRadioButton(const QString & text, QWidget * parent);
extern void* dector_ZN12QRadioButtonC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN12QRadioButtonC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QRadioButton)=1
type QRadioButton struct {
  /*qbase*/ QAbstractButton;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  const QMetaObject * QRadioButton::metaObject();
func (this *QRadioButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton10metaObjectEv
  default:
    qtrt.ErrorResolve("QRadioButton", "metaObject", args)
  }

}

  // proto:  void QRadioButton::QRadioButton(QWidget * parent);
func NewQRadioButton(args ...interface{}) QRadioButton {
  return QRadioButton{}
}

  // proto:  QSize QRadioButton::sizeHint();
func (this *QRadioButton) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton8sizeHintEv
  default:
    qtrt.ErrorResolve("QRadioButton", "sizeHint", args)
  }

}

  // proto:  QSize QRadioButton::minimumSizeHint();
func (this *QRadioButton) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QRadioButton", "minimumSizeHint", args)
  }

}

  // proto:  void QRadioButton::~QRadioButton();
func (this *QRadioButton) FreeQRadioButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRadioButton", "~QRadioButton", args)
  }

}

// <= body block end

