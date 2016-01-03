package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qdialog.h
// dst-file: /src/widgets/qdialog.go
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
  // proto:  void QDialog::setExtension(QWidget * extension);
extern void _ZN7QDialog12setExtensionEP7QWidget(void* qthis, void* arg0);
  // proto:  int QDialog::result();
extern void _ZNK7QDialog6resultEv(void* qthis);
  // proto:  void QDialog::done(int );
extern void _ZN7QDialog4doneEi(void* qthis, int32_t arg0);
  // proto:  void QDialog::open();
extern void _ZN7QDialog4openEv(void* qthis);
  // proto:  void QDialog::~QDialog();
extern void _ZN7QDialogD0Ev(void* qthis);
  // proto:  void QDialog::setResult(int r);
extern void _ZN7QDialog9setResultEi(void* qthis, int32_t arg0);
  // proto:  void QDialog::setSizeGripEnabled(bool );
extern void _ZN7QDialog18setSizeGripEnabledEb(void* qthis, bool arg0);
  // proto:  void QDialog::showExtension(bool );
extern void _ZN7QDialog13showExtensionEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QDialog::metaObject();
extern void _ZNK7QDialog10metaObjectEv(void* qthis);
  // proto:  QSize QDialog::minimumSizeHint();
extern void _ZNK7QDialog15minimumSizeHintEv(void* qthis);
  // proto:  QSize QDialog::sizeHint();
extern void _ZNK7QDialog8sizeHintEv(void* qthis);
  // proto:  void QDialog::accept();
extern void _ZN7QDialog6acceptEv(void* qthis);
  // proto:  void QDialog::setVisible(bool visible);
extern void _ZN7QDialog10setVisibleEb(void* qthis, bool arg0);
  // proto:  QWidget * QDialog::extension();
extern void _ZNK7QDialog9extensionEv(void* qthis);
  // proto:  int QDialog::exec();
extern void _ZN7QDialog4execEv(void* qthis);
  // proto:  void QDialog::reject();
extern void _ZN7QDialog6rejectEv(void* qthis);
  // proto:  void QDialog::QDialog(const QDialog & );
extern void* dector_ZN7QDialogC1ERKS_(void* arg0);
extern void _ZN7QDialogC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QDialog::isSizeGripEnabled();
extern void _ZNK7QDialog17isSizeGripEnabledEv(void* qthis);
  // proto:  void QDialog::setModal(bool modal);
extern void _ZN7QDialog8setModalEb(void* qthis, bool arg0);
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

// class sizeof(QDialog)=1
type QDialog struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _finished QDialog_finished_signal;
//  _accepted QDialog_accepted_signal;
//  _rejected QDialog_rejected_signal;
}

  // proto:  void QDialog::setExtension(QWidget * extension);
func (this *QDialog) setExtension(args ...interface{}) () {
  // setExtension(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog12setExtensionEP7QWidget
    // invoke: void setExtension(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QDialog12setExtensionEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "setExtension", args)
  }

}

  // proto:  int QDialog::result();
func (this *QDialog) result(args ...interface{}) () {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog6resultEv
    // invoke: int result()
    C._ZNK7QDialog6resultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "result", args)
  }

}

  // proto:  void QDialog::done(int );
func (this *QDialog) done(args ...interface{}) () {
  // done(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4doneEi
    // invoke: void done(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog4doneEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "done", args)
  }

}

  // proto:  void QDialog::open();
func (this *QDialog) open(args ...interface{}) () {
  // open()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4openEv
    // invoke: void open()
    C._ZN7QDialog4openEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "open", args)
  }

}

  // proto:  void QDialog::~QDialog();
func (this *QDialog) FreeQDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDialog", "~QDialog", args)
  }

}

  // proto:  void QDialog::setResult(int r);
func (this *QDialog) setResult(args ...interface{}) () {
  // setResult(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog9setResultEi
    // invoke: void setResult(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog9setResultEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "setResult", args)
  }

}

  // proto:  void QDialog::setSizeGripEnabled(bool );
func (this *QDialog) setSizeGripEnabled(args ...interface{}) () {
  // setSizeGripEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog18setSizeGripEnabledEb
    // invoke: void setSizeGripEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog18setSizeGripEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "setSizeGripEnabled", args)
  }

}

  // proto:  void QDialog::showExtension(bool );
func (this *QDialog) showExtension(args ...interface{}) () {
  // showExtension(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog13showExtensionEb
    // invoke: void showExtension(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog13showExtensionEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "showExtension", args)
  }

}

  // proto:  const QMetaObject * QDialog::metaObject();
func (this *QDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK7QDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "metaObject", args)
  }

}

  // proto:  QSize QDialog::minimumSizeHint();
func (this *QDialog) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK7QDialog15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "minimumSizeHint", args)
  }

}

  // proto:  QSize QDialog::sizeHint();
func (this *QDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK7QDialog8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "sizeHint", args)
  }

}

  // proto:  void QDialog::accept();
func (this *QDialog) accept(args ...interface{}) () {
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog6acceptEv
    // invoke: void accept()
    C._ZN7QDialog6acceptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "accept", args)
  }

}

  // proto:  void QDialog::setVisible(bool visible);
func (this *QDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "setVisible", args)
  }

}

  // proto:  QWidget * QDialog::extension();
func (this *QDialog) extension(args ...interface{}) () {
  // extension()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog9extensionEv
    // invoke: QWidget * extension()
    C._ZNK7QDialog9extensionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "extension", args)
  }

}

  // proto:  int QDialog::exec();
func (this *QDialog) exec(args ...interface{}) () {
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4execEv
    // invoke: int exec()
    C._ZN7QDialog4execEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "exec", args)
  }

}

  // proto:  void QDialog::reject();
func (this *QDialog) reject(args ...interface{}) () {
  // reject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog6rejectEv
    // invoke: void reject()
    C._ZN7QDialog6rejectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "reject", args)
  }

}

  // proto:  void QDialog::QDialog(const QDialog & );
func NewQDialog(args ...interface{}) QDialog {
  return QDialog{}
}

  // proto:  bool QDialog::isSizeGripEnabled();
func (this *QDialog) isSizeGripEnabled(args ...interface{}) () {
  // isSizeGripEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog17isSizeGripEnabledEv
    // invoke: bool isSizeGripEnabled()
    C._ZNK7QDialog17isSizeGripEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDialog", "isSizeGripEnabled", args)
  }

}

  // proto:  void QDialog::setModal(bool modal);
func (this *QDialog) setModal(args ...interface{}) () {
  // setModal(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog8setModalEb
    // invoke: void setModal(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QDialog8setModalEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialog", "setModal", args)
  }

}

// <= body block end

