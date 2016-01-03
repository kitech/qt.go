package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qgroupbox.h
// dst-file: /src/widgets/qgroupbox.go
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
  // proto:  bool QGroupBox::isCheckable();
extern void _ZNK9QGroupBox11isCheckableEv(void* qthis);
  // proto:  void QGroupBox::setCheckable(bool checkable);
extern void _ZN9QGroupBox12setCheckableEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QGroupBox::metaObject();
extern void _ZNK9QGroupBox10metaObjectEv(void* qthis);
  // proto:  bool QGroupBox::isFlat();
extern void _ZNK9QGroupBox6isFlatEv(void* qthis);
  // proto:  QSize QGroupBox::minimumSizeHint();
extern void _ZNK9QGroupBox15minimumSizeHintEv(void* qthis);
  // proto:  void QGroupBox::setFlat(bool flat);
extern void _ZN9QGroupBox7setFlatEb(void* qthis, bool arg0);
  // proto:  void QGroupBox::~QGroupBox();
extern void _ZN9QGroupBoxD0Ev(void* qthis);
  // proto:  void QGroupBox::QGroupBox(QWidget * parent);
extern void* dector_ZN9QGroupBoxC1EP7QWidget(void* arg0);
extern void _ZN9QGroupBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QGroupBox::isChecked();
extern void _ZNK9QGroupBox9isCheckedEv(void* qthis);
  // proto:  void QGroupBox::setChecked(bool checked);
extern void _ZN9QGroupBox10setCheckedEb(void* qthis, bool arg0);
  // proto:  void QGroupBox::QGroupBox(const QGroupBox & );
extern void* dector_ZN9QGroupBoxC1ERKS_(void* arg0);
extern void _ZN9QGroupBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QGroupBox::title();
extern void _ZNK9QGroupBox5titleEv(void* qthis);
  // proto:  void QGroupBox::setAlignment(int alignment);
extern void _ZN9QGroupBox12setAlignmentEi(void* qthis, int32_t arg0);
  // proto:  void QGroupBox::setTitle(const QString & title);
extern void _ZN9QGroupBox8setTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QGroupBox::QGroupBox(const QString & title, QWidget * parent);
extern void* dector_ZN9QGroupBoxC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN9QGroupBoxC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QGroupBox)=1
type QGroupBox struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _clicked QGroupBox_clicked_signal;
//  _toggled QGroupBox_toggled_signal;
}

  // proto:  bool QGroupBox::isCheckable();
func (this *QGroupBox) isCheckable(args ...interface{}) () {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox11isCheckableEv
    // invoke: bool isCheckable()
    C._ZNK9QGroupBox11isCheckableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isCheckable", args)
  }

}

  // proto:  void QGroupBox::setCheckable(bool checkable);
func (this *QGroupBox) setCheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox12setCheckableEb
    // invoke: void setCheckable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QGroupBox12setCheckableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setCheckable", args)
  }

}

  // proto:  const QMetaObject * QGroupBox::metaObject();
func (this *QGroupBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QGroupBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "metaObject", args)
  }

}

  // proto:  bool QGroupBox::isFlat();
func (this *QGroupBox) isFlat(args ...interface{}) () {
  // isFlat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox6isFlatEv
    // invoke: bool isFlat()
    C._ZNK9QGroupBox6isFlatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isFlat", args)
  }

}

  // proto:  QSize QGroupBox::minimumSizeHint();
func (this *QGroupBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK9QGroupBox15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "minimumSizeHint", args)
  }

}

  // proto:  void QGroupBox::setFlat(bool flat);
func (this *QGroupBox) setFlat(args ...interface{}) () {
  // setFlat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox7setFlatEb
    // invoke: void setFlat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QGroupBox7setFlatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setFlat", args)
  }

}

  // proto:  void QGroupBox::~QGroupBox();
func (this *QGroupBox) FreeQGroupBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGroupBox", "~QGroupBox", args)
  }

}

  // proto:  void QGroupBox::QGroupBox(QWidget * parent);
func NewQGroupBox(args ...interface{}) QGroupBox {
  return QGroupBox{}
}

  // proto:  bool QGroupBox::isChecked();
func (this *QGroupBox) isChecked(args ...interface{}) () {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox9isCheckedEv
    // invoke: bool isChecked()
    C._ZNK9QGroupBox9isCheckedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isChecked", args)
  }

}

  // proto:  void QGroupBox::setChecked(bool checked);
func (this *QGroupBox) setChecked(args ...interface{}) () {
  // setChecked(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox10setCheckedEb
    // invoke: void setChecked(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QGroupBox10setCheckedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setChecked", args)
  }

}

  // proto:  QString QGroupBox::title();
func (this *QGroupBox) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox5titleEv
    // invoke: QString title()
    C._ZNK9QGroupBox5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "title", args)
  }

}

  // proto:  void QGroupBox::setAlignment(int alignment);
func (this *QGroupBox) setAlignment(args ...interface{}) () {
  // setAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox12setAlignmentEi
    // invoke: void setAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QGroupBox12setAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setAlignment", args)
  }

}

  // proto:  void QGroupBox::setTitle(const QString & title);
func (this *QGroupBox) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QGroupBox8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setTitle", args)
  }

}

// <= body block end

