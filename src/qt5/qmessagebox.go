package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK11QMessageBox4textEv(void* qthis); // 4
  // proto:  void QMessageBox::setButtonText(int button, const QString & text);
extern void C_ZN11QMessageBox13setButtonTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QPixmap QMessageBox::iconPixmap();
extern void C_ZNK11QMessageBox10iconPixmapEv(void* qthis); // 4
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto:  QString QMessageBox::detailedText();
extern void C_ZNK11QMessageBox12detailedTextEv(void* qthis); // 4
  // proto:  QMessageBox::StandardButton QMessageBox::standardButton(QAbstractButton * button);
extern void C_ZNK11QMessageBox14standardButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::removeButton(QAbstractButton * button);
extern void C_ZN11QMessageBox12removeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::open(QObject * receiver, const char * member);
extern void C_ZN11QMessageBox4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto: static void QMessageBox::aboutQt(QWidget * parent, const QString & title);
extern void C_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString(void* arg0, void* arg1); // 4
  // proto:  QString QMessageBox::informativeText();
extern void C_ZNK11QMessageBox15informativeTextEv(void* qthis); // 4
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto:  QMessageBox::ButtonRole QMessageBox::buttonRole(QAbstractButton * button);
extern void C_ZNK11QMessageBox10buttonRoleEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QList<QAbstractButton *> QMessageBox::buttons();
extern void C_ZNK11QMessageBox7buttonsEv(void* qthis); // 4
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto:  QAbstractButton * QMessageBox::escapeButton();
extern void C_ZNK11QMessageBox12escapeButtonEv(void* qthis); // 4
  // proto:  void QMessageBox::setIconPixmap(const QPixmap & pixmap);
extern void C_ZN11QMessageBox13setIconPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QMessageBox::setWindowTitle(const QString & title);
extern void C_ZN11QMessageBox14setWindowTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAbstractButton * QMessageBox::clickedButton();
extern void C_ZNK11QMessageBox13clickedButtonEv(void* qthis); // 4
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int32_t arg6, int32_t arg7); // 4
  // proto:  void QMessageBox::setInformativeText(const QString & text);
extern void C_ZN11QMessageBox18setInformativeTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QCheckBox * QMessageBox::checkBox();
extern void C_ZNK11QMessageBox8checkBoxEv(void* qthis); // 4
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
extern void C_ZNK11QMessageBox10buttonTextEi(void* qthis, int32_t arg0); // 4
  // proto:  StandardButtons QMessageBox::standardButtons();
extern void C_ZNK11QMessageBox15standardButtonsEv(void* qthis); // 4
  // proto:  QPushButton * QMessageBox::defaultButton();
extern void C_ZNK11QMessageBox13defaultButtonEv(void* qthis); // 4
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
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QMessageBox)=1
type QMessageBox struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
//  _buttonClicked QMessageBox_buttonClicked_signal;
}

// setCheckBox(class QCheckBox *)
func (this *QMessageBox) setCheckBox(args ...interface{}) () {
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
    var arg0 = args[0].(QCheckBox).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox11setCheckBoxEP9QCheckBox(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setCheckBox", args)
  }

}

// textInteractionFlags()
func (this *QMessageBox) textInteractionFlags(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox20textInteractionFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "textInteractionFlags", args)
  }

}

// setDetailedText(const class QString &)
func (this *QMessageBox) setDetailedText(args ...interface{}) () {
  // setDetailedText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox15setDetailedTextERK7QString
    // invoke: void setDetailedText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox15setDetailedTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setDetailedText", args)
  }

}

// text()
func (this *QMessageBox) text(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "text", args)
  }

}

// setButtonText(int, const class QString &)
func (this *QMessageBox) setButtonText(args ...interface{}) () {
  // setButtonText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setButtonTextEiRK7QString
    // invoke: void setButtonText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMessageBox13setButtonTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "setButtonText", args)
  }

}

// iconPixmap()
func (this *QMessageBox) iconPixmap(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox10iconPixmapEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "iconPixmap", args)
  }

}

// warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) warning_s(args ...interface{}) () {
  // warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  // warning(class QWidget *, const class QString &, const class QString &, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][6] = qtrt.Int32Ty(false) // "int"
  vtys[0][7] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var ret = C.C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii
    // invoke: int warning(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var ret = C.C_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "warning", args)
  }

}

// detailedText()
func (this *QMessageBox) detailedText(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox12detailedTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "detailedText", args)
  }

}

// standardButton(class QAbstractButton *)
func (this *QMessageBox) standardButton(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMessageBox14standardButtonEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "standardButton", args)
  }

}

// removeButton(class QAbstractButton *)
func (this *QMessageBox) removeButton(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox12removeButtonEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "removeButton", args)
  }

}

// open(class QObject *, const char *)
func (this *QMessageBox) open(args ...interface{}) () {
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
    // invoke: _ZN11QMessageBox4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN11QMessageBox4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "open", args)
  }

}

// aboutQt(class QWidget *, const class QString &)
func (this *QMessageBox) aboutQt_s(args ...interface{}) () {
  // aboutQt(class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7aboutQtEP7QWidgetRK7QString
    // invoke: void aboutQt(class QWidget *, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString(arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageBox", "aboutQt", args)
  }

}

// informativeText()
func (this *QMessageBox) informativeText(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox15informativeTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "informativeText", args)
  }

}

// question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) question_s(args ...interface{}) () {
  // question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  // question(class QWidget *, const class QString &, const class QString &, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][6] = qtrt.Int32Ty(false) // "int"
  vtys[0][7] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var ret = C.C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii
    // invoke: int question(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var ret = C.C_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "question", args)
  }

}

// buttonRole(class QAbstractButton *)
func (this *QMessageBox) buttonRole(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMessageBox10buttonRoleEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonRole", args)
  }

}

// buttons()
func (this *QMessageBox) buttons(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "buttons", args)
  }

}

// critical(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) critical_s(args ...interface{}) () {
  // critical(class QWidget *, const class QString &, const class QString &, int, int, int)
  // critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[1][7] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii
    // invoke: int critical(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var ret = C.C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var ret = C.C_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "critical", args)
  }

}

// escapeButton()
func (this *QMessageBox) escapeButton(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox12escapeButtonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "escapeButton", args)
  }

}

// setIconPixmap(const class QPixmap &)
func (this *QMessageBox) setIconPixmap(args ...interface{}) () {
  // setIconPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setIconPixmapERK7QPixmap
    // invoke: void setIconPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox13setIconPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setIconPixmap", args)
  }

}

// setWindowTitle(const class QString &)
func (this *QMessageBox) setWindowTitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox14setWindowTitleERK7QString
    // invoke: void setWindowTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox14setWindowTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setWindowTitle", args)
  }

}

// clickedButton()
func (this *QMessageBox) clickedButton(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox13clickedButtonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "clickedButton", args)
  }

}

// information(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) information_s(args ...interface{}) () {
  // information(class QWidget *, const class QString &, const class QString &, int, int, int)
  // information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[1][7] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii
    // invoke: int information(class QWidget *, const class QString &, const class QString &, int, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var ret = C.C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii(arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii
    // invoke: int information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var ret = C.C_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "information", args)
  }

}

// setInformativeText(const class QString &)
func (this *QMessageBox) setInformativeText(args ...interface{}) () {
  // setInformativeText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox18setInformativeTextERK7QString
    // invoke: void setInformativeText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox18setInformativeTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setInformativeText", args)
  }

}

// checkBox()
func (this *QMessageBox) checkBox(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox8checkBoxEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "checkBox", args)
  }

}

// setEscapeButton(class QAbstractButton *)
func (this *QMessageBox) setEscapeButton(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setEscapeButton", args)
  }

}

// icon()
func (this *QMessageBox) icon(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "icon", args)
  }

}

// about(class QWidget *, const class QString &, const class QString &)
func (this *QMessageBox) about_s(args ...interface{}) () {
  // about(class QWidget *, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_
    // invoke: void about(class QWidget *, const class QString &, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QMessageBox", "about", args)
  }

}

// metaObject()
func (this *QMessageBox) metaObject(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "metaObject", args)
  }

}

// ~QMessageBox()
func (this *QMessageBox) FreeQMessageBox(args ...interface{}) () {
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
    C.C_ZN11QMessageBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "~QMessageBox", args)
  }

}

// setText(const class QString &)
func (this *QMessageBox) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setText", args)
  }

}

// buttonText(int)
func (this *QMessageBox) buttonText(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QMessageBox10buttonTextEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonText", args)
  }

}

// standardButtons()
func (this *QMessageBox) standardButtons(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox15standardButtonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "standardButtons", args)
  }

}

// defaultButton()
func (this *QMessageBox) defaultButton(args ...interface{}) () {
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
    var ret = C.C_ZNK11QMessageBox13defaultButtonEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMessageBox", "defaultButton", args)
  }

}

// setDefaultButton(class QPushButton *)
func (this *QMessageBox) setDefaultButton(args ...interface{}) () {
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
    var arg0 = args[0].(QPushButton).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMessageBox16setDefaultButtonEP11QPushButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageBox", "setDefaultButton", args)
  }

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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QMessageBoxC2EP7QWidget(arg0)
    return &QMessageBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMessageBox", "QMessageBox", args)
  }

  return nil // QMessageBox{qclsinst:qthis}
}

// textFormat()
func (this *QMessageBox) textFormat(args ...interface{}) () {
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
    C.C_ZNK11QMessageBox10textFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageBox", "textFormat", args)
  }

}

// <= body block end

