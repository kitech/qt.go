package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qprogressdialog.h
// dst-file: /src/widgets/qprogressdialog.go
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
  // proto:  void QProgressDialog::setLabel(QLabel * label);
extern void C_ZN15QProgressDialog8setLabelEP6QLabel(void* qthis, void* arg0); // 4
  // proto:  int QProgressDialog::minimumDuration();
extern void C_ZNK15QProgressDialog15minimumDurationEv(void* qthis); // 4
  // proto:  QString QProgressDialog::labelText();
extern void C_ZNK15QProgressDialog9labelTextEv(void* qthis); // 4
  // proto:  int QProgressDialog::minimum();
extern void C_ZNK15QProgressDialog7minimumEv(void* qthis); // 4
  // proto:  void QProgressDialog::cancel();
extern void C_ZN15QProgressDialog6cancelEv(void* qthis); // 4
  // proto:  void QProgressDialog::open(QObject * receiver, const char * member);
extern void C_ZN15QProgressDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  void QProgressDialog::setLabelText(const QString & text);
extern void C_ZN15QProgressDialog12setLabelTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QProgressDialog::autoReset();
extern void C_ZNK15QProgressDialog9autoResetEv(void* qthis); // 4
  // proto:  bool QProgressDialog::wasCanceled();
extern void C_ZNK15QProgressDialog11wasCanceledEv(void* qthis); // 4
  // proto:  void QProgressDialog::setCancelButtonText(const QString & text);
extern void C_ZN15QProgressDialog19setCancelButtonTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProgressDialog::setMinimumDuration(int ms);
extern void C_ZN15QProgressDialog18setMinimumDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setMaximum(int maximum);
extern void C_ZN15QProgressDialog10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setBar(QProgressBar * bar);
extern void C_ZN15QProgressDialog6setBarEP12QProgressBar(void* qthis, void* arg0); // 4
  // proto:  void QProgressDialog::setValue(int progress);
extern void C_ZN15QProgressDialog8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setCancelButton(QPushButton * button);
extern void C_ZN15QProgressDialog15setCancelButtonEP11QPushButton(void* qthis, void* arg0); // 4
  // proto:  bool QProgressDialog::autoClose();
extern void C_ZNK15QProgressDialog9autoCloseEv(void* qthis); // 4
  // proto:  const QMetaObject * QProgressDialog::metaObject();
extern void C_ZNK15QProgressDialog10metaObjectEv(void* qthis); // 4
  // proto:  void QProgressDialog::setRange(int minimum, int maximum);
extern void C_ZN15QProgressDialog8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QProgressDialog::reset();
extern void C_ZN15QProgressDialog5resetEv(void* qthis); // 4
  // proto:  QSize QProgressDialog::sizeHint();
extern void C_ZNK15QProgressDialog8sizeHintEv(void* qthis); // 4
  // proto:  void QProgressDialog::setAutoReset(bool reset);
extern void C_ZN15QProgressDialog12setAutoResetEb(void* qthis, bool arg0); // 4
  // proto:  int QProgressDialog::maximum();
extern void C_ZNK15QProgressDialog7maximumEv(void* qthis); // 4
  // proto:  int QProgressDialog::value();
extern void C_ZNK15QProgressDialog5valueEv(void* qthis); // 4
  // proto:  void QProgressDialog::setAutoClose(bool close);
extern void C_ZN15QProgressDialog12setAutoCloseEb(void* qthis, bool arg0); // 4
  // proto:  void QProgressDialog::~QProgressDialog();
extern void C_ZN15QProgressDialogD2Ev(void* qthis); // 4
  // proto:  void QProgressDialog::setMinimum(int minimum);
extern void C_ZN15QProgressDialog10setMinimumEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QProgressDialog)=1
type QProgressDialog struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
//  _canceled QProgressDialog_canceled_signal;
}

// setLabel(class QLabel *)
func (this *QProgressDialog) setLabel(args ...interface{}) () {
  // setLabel(class QLabel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLabel{}) // "QLabel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setLabelEP6QLabel
    // invoke: void setLabel(class QLabel *)
    var arg0 = args[0].(QLabel).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog8setLabelEP6QLabel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabel", args)
  }

}

// minimumDuration()
func (this *QProgressDialog) minimumDuration(args ...interface{}) () {
  // minimumDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog15minimumDurationEv
    // invoke: int minimumDuration()
    C.C_ZNK15QProgressDialog15minimumDurationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimumDuration", args)
  }

}

// labelText()
func (this *QProgressDialog) labelText(args ...interface{}) () {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9labelTextEv
    // invoke: QString labelText()
    C.C_ZNK15QProgressDialog9labelTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "labelText", args)
  }

}

// minimum()
func (this *QProgressDialog) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7minimumEv
    // invoke: int minimum()
    C.C_ZNK15QProgressDialog7minimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimum", args)
  }

}

// cancel()
func (this *QProgressDialog) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6cancelEv
    // invoke: void cancel()
    C.C_ZN15QProgressDialog6cancelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "cancel", args)
  }

}

// open(class QObject *, const char *)
func (this *QProgressDialog) open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN15QProgressDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressDialog", "open", args)
  }

}

// setLabelText(const class QString &)
func (this *QProgressDialog) setLabelText(args ...interface{}) () {
  // setLabelText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setLabelTextERK7QString
    // invoke: void setLabelText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setLabelTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabelText", args)
  }

}

// autoReset()
func (this *QProgressDialog) autoReset(args ...interface{}) () {
  // autoReset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoResetEv
    // invoke: bool autoReset()
    C.C_ZNK15QProgressDialog9autoResetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoReset", args)
  }

}

// wasCanceled()
func (this *QProgressDialog) wasCanceled(args ...interface{}) () {
  // wasCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog11wasCanceledEv
    // invoke: bool wasCanceled()
    C.C_ZNK15QProgressDialog11wasCanceledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "wasCanceled", args)
  }

}

// setCancelButtonText(const class QString &)
func (this *QProgressDialog) setCancelButtonText(args ...interface{}) () {
  // setCancelButtonText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog19setCancelButtonTextERK7QString
    // invoke: void setCancelButtonText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog19setCancelButtonTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButtonText", args)
  }

}

// setMinimumDuration(int)
func (this *QProgressDialog) setMinimumDuration(args ...interface{}) () {
  // setMinimumDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog18setMinimumDurationEi
    // invoke: void setMinimumDuration(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog18setMinimumDurationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimumDuration", args)
  }

}

// setMaximum(int)
func (this *QProgressDialog) setMaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMaximumEi
    // invoke: void setMaximum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog10setMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMaximum", args)
  }

}

// setBar(class QProgressBar *)
func (this *QProgressDialog) setBar(args ...interface{}) () {
  // setBar(class QProgressBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProgressBar{}) // "QProgressBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6setBarEP12QProgressBar
    // invoke: void setBar(class QProgressBar *)
    var arg0 = args[0].(QProgressBar).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog6setBarEP12QProgressBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setBar", args)
  }

}

// setValue(int)
func (this *QProgressDialog) setValue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setValueEi
    // invoke: void setValue(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog8setValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setValue", args)
  }

}

// setCancelButton(class QPushButton *)
func (this *QProgressDialog) setCancelButton(args ...interface{}) () {
  // setCancelButton(class QPushButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPushButton{}) // "QPushButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog15setCancelButtonEP11QPushButton
    // invoke: void setCancelButton(class QPushButton *)
    var arg0 = args[0].(QPushButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog15setCancelButtonEP11QPushButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButton", args)
  }

}

// autoClose()
func (this *QProgressDialog) autoClose(args ...interface{}) () {
  // autoClose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoCloseEv
    // invoke: bool autoClose()
    C.C_ZNK15QProgressDialog9autoCloseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoClose", args)
  }

}

// metaObject()
func (this *QProgressDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QProgressDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "metaObject", args)
  }

}

// setRange(int, int)
func (this *QProgressDialog) setRange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setRangeEii
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QProgressDialog8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setRange", args)
  }

}

// reset()
func (this *QProgressDialog) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog5resetEv
    // invoke: void reset()
    C.C_ZN15QProgressDialog5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "reset", args)
  }

}

// sizeHint()
func (this *QProgressDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK15QProgressDialog8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "sizeHint", args)
  }

}

// setAutoReset(_Bool)
func (this *QProgressDialog) setAutoReset(args ...interface{}) () {
  // setAutoReset(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoResetEb
    // invoke: void setAutoReset(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setAutoResetEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoReset", args)
  }

}

// maximum()
func (this *QProgressDialog) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7maximumEv
    // invoke: int maximum()
    C.C_ZNK15QProgressDialog7maximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "maximum", args)
  }

}

// value()
func (this *QProgressDialog) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog5valueEv
    // invoke: int value()
    C.C_ZNK15QProgressDialog5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "value", args)
  }

}

// setAutoClose(_Bool)
func (this *QProgressDialog) setAutoClose(args ...interface{}) () {
  // setAutoClose(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoCloseEb
    // invoke: void setAutoClose(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setAutoCloseEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoClose", args)
  }

}

// ~QProgressDialog()
func (this *QProgressDialog) FreeQProgressDialog(args ...interface{}) () {
  // ~QProgressDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialogD0Ev
    // invoke: void ~QProgressDialog()
    C.C_ZN15QProgressDialogD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "~QProgressDialog", args)
  }

}

// setMinimum(int)
func (this *QProgressDialog) setMinimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMinimumEi
    // invoke: void setMinimum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog10setMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimum", args)
  }

}

// <= body block end

