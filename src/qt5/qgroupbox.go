package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGroupBox::~QGroupBox();
extern void C_ZN9QGroupBoxD2Ev(void* qthis); // 4
  // proto:  void QGroupBox::setTitle(const QString & title);
extern void C_ZN9QGroupBox8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QGroupBox::isCheckable();
extern void C_ZNK9QGroupBox11isCheckableEv(void* qthis); // 4
  // proto:  Qt::Alignment QGroupBox::alignment();
extern void C_ZNK9QGroupBox9alignmentEv(void* qthis); // 4
  // proto:  void QGroupBox::setFlat(bool flat);
extern void C_ZN9QGroupBox7setFlatEb(void* qthis, bool arg0); // 4
  // proto:  void QGroupBox::QGroupBox(QWidget * parent);
extern void C_ZN9QGroupBoxC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QGroupBox::QGroupBox(const QString & title, QWidget * parent);
extern void C_ZN9QGroupBoxC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto:  QString QGroupBox::title();
extern void C_ZNK9QGroupBox5titleEv(void* qthis); // 4
  // proto:  bool QGroupBox::isChecked();
extern void C_ZNK9QGroupBox9isCheckedEv(void* qthis); // 4
  // proto:  void QGroupBox::setAlignment(int alignment);
extern void C_ZN9QGroupBox12setAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGroupBox::setChecked(bool checked);
extern void C_ZN9QGroupBox10setCheckedEb(void* qthis, bool arg0); // 4
  // proto:  void QGroupBox::setCheckable(bool checkable);
extern void C_ZN9QGroupBox12setCheckableEb(void* qthis, bool arg0); // 4
  // proto:  bool QGroupBox::isFlat();
extern void C_ZNK9QGroupBox6isFlatEv(void* qthis); // 4
  // proto:  const QMetaObject * QGroupBox::metaObject();
extern void C_ZNK9QGroupBox10metaObjectEv(void* qthis); // 4
  // proto:  QSize QGroupBox::minimumSizeHint();
extern void C_ZNK9QGroupBox15minimumSizeHintEv(void* qthis); // 4
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

// ~QGroupBox()
func (this *QGroupBox) FreeQGroupBox(args ...interface{}) () {
  // ~QGroupBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBoxD0Ev
    // invoke: void ~QGroupBox()
    C.C_ZN9QGroupBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "~QGroupBox", args)
  }

}

// setTitle(const class QString &)
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
    C.C_ZN9QGroupBox8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setTitle", args)
  }

}

// isCheckable()
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
    C.C_ZNK9QGroupBox11isCheckableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isCheckable", args)
  }

}

// alignment()
func (this *QGroupBox) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK9QGroupBox9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "alignment", args)
  }

}

// setFlat(_Bool)
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
    C.C_ZN9QGroupBox7setFlatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setFlat", args)
  }

}

// QGroupBox(class QWidget *)
func NewQGroupBox(args ...interface{}) QGroupBox {
  // QGroupBox(class QWidget *)
  // QGroupBox(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBoxC1EP7QWidget
    // invoke: void QGroupBox(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QGroupBoxC2EP7QWidget(qthis, arg0)
  case 1:
    // invoke: _ZN9QGroupBoxC1ERK7QStringP7QWidget
    // invoke: void QGroupBox(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QGroupBoxC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGroupBox", "QGroupBox", args)
  }

  return QGroupBox{}
}

// title()
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
    C.C_ZNK9QGroupBox5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "title", args)
  }

}

// isChecked()
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
    C.C_ZNK9QGroupBox9isCheckedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isChecked", args)
  }

}

// setAlignment(int)
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
    C.C_ZN9QGroupBox12setAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setAlignment", args)
  }

}

// setChecked(_Bool)
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
    C.C_ZN9QGroupBox10setCheckedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setChecked", args)
  }

}

// setCheckable(_Bool)
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
    C.C_ZN9QGroupBox12setCheckableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setCheckable", args)
  }

}

// isFlat()
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
    C.C_ZNK9QGroupBox6isFlatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "isFlat", args)
  }

}

// metaObject()
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
    C.C_ZNK9QGroupBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "metaObject", args)
  }

}

// minimumSizeHint()
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
    C.C_ZNK9QGroupBox15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "minimumSizeHint", args)
  }

}

// <= body block end

