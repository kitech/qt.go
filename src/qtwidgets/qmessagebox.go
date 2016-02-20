package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qmessagebox.h
// dst-file: /src/widgets/qmessagebox.go
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
  // proto:  void QMessageBox::setCheckBox(QCheckBox * cb);
extern void C_ZN11QMessageBox11setCheckBoxEP9QCheckBox(void* qthis, void* arg0); // 4
  // proto:  Qt::TextInteractionFlags QMessageBox::textInteractionFlags();
extern void C_ZNK11QMessageBox20textInteractionFlagsEv(void* qthis); // 4
  // proto:  void QMessageBox::setDetailedText(const QString & text);
extern void C_ZN11QMessageBox15setDetailedTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QMessageBox::text();
extern void* C_ZNK11QMessageBox4textEv(void* qthis); // 4
  // proto:  void QMessageBox::setButtonText(int button, const QString & text);
extern void C_ZN11QMessageBox13setButtonTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QPixmap QMessageBox::iconPixmap();
extern void* C_ZNK11QMessageBox10iconPixmapEv(void* qthis); // 4
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern int32_t C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern int32_t C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto:  QString QMessageBox::detailedText();
extern void* C_ZNK11QMessageBox12detailedTextEv(void* qthis); // 4
  // proto:  QMessageBox::StandardButton QMessageBox::standardButton(QAbstractButton * button);
extern void C_ZNK11QMessageBox14standardButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::removeButton(QAbstractButton * button);
extern void C_ZN11QMessageBox12removeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::open(QObject * receiver, const char * member);
extern void C_ZN11QMessageBox4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto: static void QMessageBox::aboutQt(QWidget * parent, const QString & title);
extern void C_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString(void* arg0, void* arg1); // 4
  // proto:  QString QMessageBox::informativeText();
extern void* C_ZNK11QMessageBox15informativeTextEv(void* qthis); // 4
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern int32_t C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern int32_t C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto:  QMessageBox::ButtonRole QMessageBox::buttonRole(QAbstractButton * button);
extern void C_ZNK11QMessageBox10buttonRoleEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QList<QAbstractButton *> QMessageBox::buttons();
extern void C_ZNK11QMessageBox7buttonsEv(void* qthis); // 4
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern int32_t C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern int32_t C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto:  QAbstractButton * QMessageBox::escapeButton();
extern void C_ZNK11QMessageBox12escapeButtonEv(void* qthis); // 4
  // proto:  void QMessageBox::setIconPixmap(const QPixmap & pixmap);
extern void C_ZN11QMessageBox13setIconPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::setWindowTitle(const QString & title);
extern void C_ZN11QMessageBox14setWindowTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAbstractButton * QMessageBox::clickedButton();
extern void C_ZNK11QMessageBox13clickedButtonEv(void* qthis); // 4
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern int32_t C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern int32_t C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto:  void QMessageBox::setInformativeText(const QString & text);
extern void C_ZN11QMessageBox18setInformativeTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QCheckBox * QMessageBox::checkBox();
extern void* C_ZNK11QMessageBox8checkBoxEv(void* qthis); // 4
  // proto:  void QMessageBox::setEscapeButton(QAbstractButton * button);
extern void C_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QMessageBox::Icon QMessageBox::icon();
extern void C_ZNK11QMessageBox4iconEv(void* qthis); // 4
  // proto: static void QMessageBox::about(QWidget * parent, const QString & title, const QString & text);
extern void C_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  const QMetaObject * QMessageBox::metaObject();
extern void C_ZNK11QMessageBox10metaObjectEv(void* qthis); // 4
  // proto:  void QMessageBox::~QMessageBox();
extern void C_ZN11QMessageBoxD2Ev(void* qthis); // 4
  // proto:  void QMessageBox::setText(const QString & text);
extern void C_ZN11QMessageBox7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QMessageBox::buttonText(int button);
extern void* C_ZNK11QMessageBox10buttonTextEi(void* qthis, int32_t arg0); // 4
  // proto:  StandardButtons QMessageBox::standardButtons();
extern void C_ZNK11QMessageBox15standardButtonsEv(void* qthis); // 4
  // proto:  QPushButton * QMessageBox::defaultButton();
extern void* C_ZNK11QMessageBox13defaultButtonEv(void* qthis); // 4
  // proto:  void QMessageBox::setDefaultButton(QPushButton * button);
extern void C_ZN11QMessageBox16setDefaultButtonEP11QPushButton(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::QMessageBox(QWidget * parent);
extern void* C_ZN11QMessageBoxC2EP7QWidget(void* arg0); // 3
  // proto:  Qt::TextFormat QMessageBox::textFormat();
extern void C_ZNK11QMessageBox10textFormatEv(void* qthis); // 4
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

// class sizeof(QMessageBox)=1
type QMessageBox struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _buttonClicked QMessageBox_buttonClicked_signal;
}

// setCheckBox(class QCheckBox *)
func (this *QMessageBox) Setcheckbox(args ...interface{}) () {
  // setCheckBox(class QCheckBox *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCheckBox{}) // "QCheckBox *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox11setCheckBoxEP9QCheckBox
    // invoke: void setCheckBox(class QCheckBox *)
    var arg0 = args[0].(*QCheckBox).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox11setCheckBoxEP9QCheckBox(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setCheckBox", args)
  }

  return
}

// textInteractionFlags()
func (this *QMessageBox) Textinteractionflags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK11QMessageBox20textInteractionFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "textInteractionFlags", args)
  }

  return
}

// setDetailedText(const class QString &)
func (this *QMessageBox) Setdetailedtext(args ...interface{}) () {
  // setDetailedText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox15setDetailedTextERK7QString
    // invoke: void setDetailedText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox15setDetailedTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setDetailedText", args)
  }

  return
}

// text()
func (this *QMessageBox) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK11QMessageBox4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "text", args)
  }

  return
}

// setButtonText(int, const class QString &)
func (this *QMessageBox) Setbuttontext(args ...interface{}) () {
  // setButtonText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setButtonTextEiRK7QString
    // invoke: void setButtonText(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMessageBox13setButtonTextEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "setButtonText", args)
  }

  return
}

// iconPixmap()
func (this *QMessageBox) Iconpixmap(args ...interface{}) (ret interface{}) {
  // iconPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10iconPixmapEv
    // invoke: QPixmap iconPixmap()
    var ret0 = C.C_ZNK11QMessageBox10iconPixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "iconPixmap", args)
  }

  return
}

// warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Warning_S(args ...interface{}) (ret interface{}) {
  // warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  // warning(class QWidget *, const class QString &, const class QString &, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][6] = qtrt.Int32Ty(false) // "int"
  vtys[0][7] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(qtrt.PrimConv(args[7], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii
    // invoke: int warning(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "warning", args)
  }

  return
}

// detailedText()
func (this *QMessageBox) Detailedtext(args ...interface{}) (ret interface{}) {
  // detailedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox12detailedTextEv
    // invoke: QString detailedText()
    var ret0 = C.C_ZNK11QMessageBox12detailedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "detailedText", args)
  }

  return
}

// standardButton(class QAbstractButton *)
func (this *QMessageBox) Standardbutton(args ...interface{}) () {
  // standardButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox14standardButtonEP15QAbstractButton
    // invoke: QMessageBox::StandardButton standardButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMessageBox14standardButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "standardButton", args)
  }

  return
}

// removeButton(class QAbstractButton *)
func (this *QMessageBox) Removebutton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox12removeButtonEP15QAbstractButton
    // invoke: void removeButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox12removeButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "removeButton", args)
  }

  return
}

// open(class QObject *, const char *)
func (this *QMessageBox) Open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN11QMessageBox4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "open", args)
  }

  return
}

// aboutQt(class QWidget *, const class QString &)
func (this *QMessageBox) Aboutqt_S(args ...interface{}) () {
  // aboutQt(class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7aboutQtEP7QWidgetRK7QString
    // invoke: void aboutQt(class QWidget *, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString(arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "aboutQt", args)
  }

  return
}

// informativeText()
func (this *QMessageBox) Informativetext(args ...interface{}) (ret interface{}) {
  // informativeText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox15informativeTextEv
    // invoke: QString informativeText()
    var ret0 = C.C_ZNK11QMessageBox15informativeTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "informativeText", args)
  }

  return
}

// question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Question_S(args ...interface{}) (ret interface{}) {
  // question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  // question(class QWidget *, const class QString &, const class QString &, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][6] = qtrt.Int32Ty(false) // "int"
  vtys[0][7] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(qtrt.PrimConv(args[7], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii
    // invoke: int question(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "question", args)
  }

  return
}

// buttonRole(class QAbstractButton *)
func (this *QMessageBox) Buttonrole(args ...interface{}) () {
  // buttonRole(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10buttonRoleEP15QAbstractButton
    // invoke: QMessageBox::ButtonRole buttonRole(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMessageBox10buttonRoleEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonRole", args)
  }

  return
}

// buttons()
func (this *QMessageBox) Buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox7buttonsEv
    // invoke: QList<QAbstractButton *> buttons()
    C.C_ZNK11QMessageBox7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "buttons", args)
  }

  return
}

// critical(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Critical_S(args ...interface{}) (ret interface{}) {
  // critical(class QWidget *, const class QString &, const class QString &, int, int, int)
  // critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][4] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[1][7] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii
    // invoke: int critical(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(qtrt.PrimConv(args[7], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "critical", args)
  }

  return
}

// escapeButton()
func (this *QMessageBox) Escapebutton(args ...interface{}) () {
  // escapeButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox12escapeButtonEv
    // invoke: QAbstractButton * escapeButton()
    C.C_ZNK11QMessageBox12escapeButtonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "escapeButton", args)
  }

  return
}

// setIconPixmap(const class QPixmap &)
func (this *QMessageBox) Seticonpixmap(args ...interface{}) () {
  // setIconPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setIconPixmapERK7QPixmap
    // invoke: void setIconPixmap(const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox13setIconPixmapERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setIconPixmap", args)
  }

  return
}

// setWindowTitle(const class QString &)
func (this *QMessageBox) Setwindowtitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox14setWindowTitleERK7QString
    // invoke: void setWindowTitle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox14setWindowTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setWindowTitle", args)
  }

  return
}

// clickedButton()
func (this *QMessageBox) Clickedbutton(args ...interface{}) () {
  // clickedButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox13clickedButtonEv
    // invoke: QAbstractButton * clickedButton()
    C.C_ZNK11QMessageBox13clickedButtonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "clickedButton", args)
  }

  return
}

// information(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Information_S(args ...interface{}) (ret interface{}) {
  // information(class QWidget *, const class QString &, const class QString &, int, int, int)
  // information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][4] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[1][7] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii
    // invoke: int information(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(qtrt.PrimConv(args[7], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "information", args)
  }

  return
}

// setInformativeText(const class QString &)
func (this *QMessageBox) Setinformativetext(args ...interface{}) () {
  // setInformativeText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox18setInformativeTextERK7QString
    // invoke: void setInformativeText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox18setInformativeTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setInformativeText", args)
  }

  return
}

// checkBox()
func (this *QMessageBox) Checkbox(args ...interface{}) (ret interface{}) {
  // checkBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox8checkBoxEv
    // invoke: QCheckBox * checkBox()
    var ret0 = C.C_ZNK11QMessageBox8checkBoxEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCheckBox{}) // "QCheckBox *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "checkBox", args)
  }

  return
}

// setEscapeButton(class QAbstractButton *)
func (this *QMessageBox) Setescapebutton(args ...interface{}) () {
  // setEscapeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox15setEscapeButtonEP15QAbstractButton
    // invoke: void setEscapeButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setEscapeButton", args)
  }

  return
}

// icon()
func (this *QMessageBox) Icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox4iconEv
    // invoke: QMessageBox::Icon icon()
    C.C_ZNK11QMessageBox4iconEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "icon", args)
  }

  return
}

// about(class QWidget *, const class QString &, const class QString &)
func (this *QMessageBox) About_S(args ...interface{}) () {
  // about(class QWidget *, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_
    // invoke: void about(class QWidget *, const class QString &, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QMessageBox", "about", args)
  }

  return
}

// metaObject()
func (this *QMessageBox) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QMessageBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "metaObject", args)
  }

  return
}

// ~QMessageBox()
func (this *QMessageBox) Freeqmessagebox(args ...interface{}) () {
  // ~QMessageBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBoxD0Ev
    // invoke: void ~QMessageBox()
    C.C_ZN11QMessageBoxD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "~QMessageBox", args)
  }

  return
}

// setText(const class QString &)
func (this *QMessageBox) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setText", args)
  }

  return
}

// buttonText(int)
func (this *QMessageBox) Buttontext(args ...interface{}) (ret interface{}) {
  // buttonText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10buttonTextEi
    // invoke: QString buttonText(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMessageBox10buttonTextEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonText", args)
  }

  return
}

// standardButtons()
func (this *QMessageBox) Standardbuttons(args ...interface{}) () {
  // standardButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox15standardButtonsEv
    // invoke: StandardButtons standardButtons()
    C.C_ZNK11QMessageBox15standardButtonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "standardButtons", args)
  }

  return
}

// defaultButton()
func (this *QMessageBox) Defaultbutton(args ...interface{}) (ret interface{}) {
  // defaultButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox13defaultButtonEv
    // invoke: QPushButton * defaultButton()
    var ret0 = C.C_ZNK11QMessageBox13defaultButtonEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPushButton{}) // "QPushButton *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageBox", "defaultButton", args)
  }

  return
}

// setDefaultButton(class QPushButton *)
func (this *QMessageBox) Setdefaultbutton(args ...interface{}) () {
  // setDefaultButton(class QPushButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPushButton{}) // "QPushButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox16setDefaultButtonEP11QPushButton
    // invoke: void setDefaultButton(class QPushButton *)
    var arg0 = args[0].(*QPushButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox16setDefaultButtonEP11QPushButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setDefaultButton", args)
  }

  return
}

// QMessageBox(class QWidget *)
func NewQMessageBox(args ...interface{}) *QMessageBox {
  // QMessageBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBoxC1EP7QWidget
    // invoke: void QMessageBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QMessageBoxC2EP7QWidget(arg0)
    return &QMessageBox{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMessageBox", "QMessageBox", args)
  }

  return nil // QMessageBox{Qclsinst:qthis}
}

// textFormat()
func (this *QMessageBox) Textformat(args ...interface{}) () {
  // textFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10textFormatEv
    // invoke: Qt::TextFormat textFormat()
    C.C_ZNK11QMessageBox10textFormatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "textFormat", args)
  }

  return
}

// <= body block end

