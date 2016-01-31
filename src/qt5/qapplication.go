package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qapplication.h
// dst-file: /src/widgets/qapplication.go
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
  // proto: static QWidget * QApplication::activePopupWidget();
extern void C_ZN12QApplication17activePopupWidgetEv(); // 4
  // proto: static void QApplication::closeAllWindows();
extern void C_ZN12QApplication15closeAllWindowsEv(); // 4
  // proto: static int QApplication::startDragDistance();
extern void C_ZN12QApplication17startDragDistanceEv(); // 4
  // proto: static int QApplication::wheelScrollLines();
extern void C_ZN12QApplication16wheelScrollLinesEv(); // 4
  // proto: static void QApplication::setColorSpec(int );
extern void C_ZN12QApplication12setColorSpecEi(int32_t arg0); // 4
  // proto:  bool QApplication::autoSipEnabled();
extern void C_ZNK12QApplication14autoSipEnabledEv(void* qthis); // 4
  // proto: static QStyle * QApplication::setStyle(const QString & );
extern void C_ZN12QApplication8setStyleERK7QString(void* arg0); // 4
  // proto: static void QApplication::setStyle(QStyle * );
extern void C_ZN12QApplication8setStyleEP6QStyle(void* arg0); // 4
  // proto:  bool QApplication::notify(QObject * , QEvent * );
extern void C_ZN12QApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QWidget * QApplication::focusWidget();
extern void C_ZN12QApplication11focusWidgetEv(); // 4
  // proto: static QWidgetList QApplication::allWidgets();
extern void C_ZN12QApplication10allWidgetsEv(); // 4
  // proto: static QWidget * QApplication::widgetAt(int x, int y);
extern void C_ZN12QApplication8widgetAtEii(int32_t arg0, int32_t arg1); // 2
  // proto: static QWidget * QApplication::widgetAt(const QPoint & p);
extern void C_ZN12QApplication8widgetAtERK6QPoint(void* arg0); // 4
  // proto: static void QApplication::aboutQt();
extern void C_ZN12QApplication7aboutQtEv(); // 4
  // proto: static void QApplication::setPalette(const QPalette & , const char * className);
extern void C_ZN12QApplication10setPaletteERK8QPalettePKc(void* arg0, unsigned char* arg1); // 4
  // proto: static QPalette QApplication::palette(const char * className);
extern void C_ZN12QApplication7paletteEPKc(unsigned char* arg0); // 4
  // proto: static QPalette QApplication::palette(const QWidget * );
extern void C_ZN12QApplication7paletteEPK7QWidget(void* arg0); // 4
  // proto: static QFont QApplication::font();
extern void C_ZN12QApplication4fontEv(); // 4
  // proto: static QFont QApplication::font(const char * className);
extern void C_ZN12QApplication4fontEPKc(unsigned char* arg0); // 4
  // proto: static QFont QApplication::font(const QWidget * );
extern void C_ZN12QApplication4fontEPK7QWidget(void* arg0); // 4
  // proto: static void QApplication::setDoubleClickInterval(int );
extern void C_ZN12QApplication22setDoubleClickIntervalEi(int32_t arg0); // 4
  // proto: static QWidgetList QApplication::topLevelWidgets();
extern void C_ZN12QApplication15topLevelWidgetsEv(); // 4
  // proto: static QDesktopWidget * QApplication::desktop();
extern void C_ZN12QApplication7desktopEv(); // 4
  // proto: static void QApplication::beep();
extern void C_ZN12QApplication4beepEv(); // 4
  // proto:  const QMetaObject * QApplication::metaObject();
extern void C_ZNK12QApplication10metaObjectEv(void* qthis); // 4
  // proto: static void QApplication::setCursorFlashTime(int );
extern void C_ZN12QApplication18setCursorFlashTimeEi(int32_t arg0); // 4
  // proto: static QWidget * QApplication::activeWindow();
extern void C_ZN12QApplication12activeWindowEv(); // 4
  // proto: static QWidget * QApplication::topLevelAt(const QPoint & p);
extern void C_ZN12QApplication10topLevelAtERK6QPoint(void* arg0); // 4
  // proto: static QWidget * QApplication::topLevelAt(int x, int y);
extern void C_ZN12QApplication10topLevelAtEii(int32_t arg0, int32_t arg1); // 2
  // proto: static QFontMetrics QApplication::fontMetrics();
extern void C_ZN12QApplication11fontMetricsEv(); // 4
  // proto: static void QApplication::setKeyboardInputInterval(int );
extern void C_ZN12QApplication24setKeyboardInputIntervalEi(int32_t arg0); // 4
  // proto: static int QApplication::colorSpec();
extern void C_ZN12QApplication9colorSpecEv(); // 4
  // proto: static QStyle * QApplication::style();
extern void C_ZN12QApplication5styleEv(); // 4
  // proto: static void QApplication::alert(QWidget * widget, int duration);
extern void C_ZN12QApplication5alertEP7QWidgeti(void* arg0, int32_t arg1); // 4
  // proto: static QWidget * QApplication::activeModalWidget();
extern void C_ZN12QApplication17activeModalWidgetEv(); // 4
  // proto:  QString QApplication::styleSheet();
extern void C_ZNK12QApplication10styleSheetEv(void* qthis); // 4
  // proto: static QIcon QApplication::windowIcon();
extern void C_ZN12QApplication10windowIconEv(); // 4
  // proto: static int QApplication::cursorFlashTime();
extern void C_ZN12QApplication15cursorFlashTimeEv(); // 4
  // proto: static void QApplication::setGlobalStrut(const QSize & );
extern void C_ZN12QApplication14setGlobalStrutERK5QSize(void* arg0); // 4
  // proto: static void QApplication::setActiveWindow(QWidget * act);
extern void C_ZN12QApplication15setActiveWindowEP7QWidget(void* arg0); // 4
  // proto:  void QApplication::~QApplication();
extern void C_ZN12QApplicationD2Ev(void* qthis); // 4
  // proto: static QSize QApplication::globalStrut();
extern void C_ZN12QApplication11globalStrutEv(); // 4
  // proto: static void QApplication::setStartDragDistance(int l);
extern void C_ZN12QApplication20setStartDragDistanceEi(int32_t arg0); // 4
  // proto: static void QApplication::setWheelScrollLines(int );
extern void C_ZN12QApplication19setWheelScrollLinesEi(int32_t arg0); // 4
  // proto: static int QApplication::keyboardInputInterval();
extern void C_ZN12QApplication21keyboardInputIntervalEv(); // 4
  // proto: static void QApplication::setStartDragTime(int ms);
extern void C_ZN12QApplication16setStartDragTimeEi(int32_t arg0); // 4
  // proto:  void QApplication::QApplication(int & argc, char ** argv, int );
extern void C_ZN12QApplicationC2ERiPPci(void* qthis, int32_t* arg0, unsigned char* arg1, int32_t arg2); // 3
  // proto: static int QApplication::startDragTime();
extern void C_ZN12QApplication13startDragTimeEv(); // 4
  // proto: static void QApplication::setWindowIcon(const QIcon & icon);
extern void C_ZN12QApplication13setWindowIconERK5QIcon(void* arg0); // 4
  // proto:  void QApplication::setAutoSipEnabled(const bool enabled);
extern void C_ZN12QApplication17setAutoSipEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QApplication::setStyleSheet(const QString & sheet);
extern void C_ZN12QApplication13setStyleSheetERK7QString(void* qthis, void* arg0); // 4
  // proto: static int QApplication::exec();
extern void C_ZN12QApplication4execEv(); // 4
  // proto: static void QApplication::setFont(const QFont & , const char * className);
extern void C_ZN12QApplication7setFontERK5QFontPKc(void* arg0, unsigned char* arg1); // 4
  // proto: static int QApplication::doubleClickInterval();
extern void C_ZN12QApplication19doubleClickIntervalEv(); // 4
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

// class sizeof(QApplication)=1
type QApplication struct {
  /*qbase*/ QGuiApplication;
  qclsinst unsafe.Pointer /* *C.void */;
//  _focusChanged QApplication_focusChanged_signal;
}

// activePopupWidget()
func (this *QApplication) activePopupWidget_s(args ...interface{}) () {
  // activePopupWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication17activePopupWidgetEv
    // invoke: QWidget * activePopupWidget()
    C.C_ZN12QApplication17activePopupWidgetEv()
  default:
    qtrt.ErrorResolve("QApplication", "activePopupWidget", args)
  }

}

// closeAllWindows()
func (this *QApplication) closeAllWindows_s(args ...interface{}) () {
  // closeAllWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication15closeAllWindowsEv
    // invoke: void closeAllWindows()
    C.C_ZN12QApplication15closeAllWindowsEv()
  default:
    qtrt.ErrorResolve("QApplication", "closeAllWindows", args)
  }

}

// startDragDistance()
func (this *QApplication) startDragDistance_s(args ...interface{}) () {
  // startDragDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication17startDragDistanceEv
    // invoke: int startDragDistance()
    C.C_ZN12QApplication17startDragDistanceEv()
  default:
    qtrt.ErrorResolve("QApplication", "startDragDistance", args)
  }

}

// wheelScrollLines()
func (this *QApplication) wheelScrollLines_s(args ...interface{}) () {
  // wheelScrollLines()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication16wheelScrollLinesEv
    // invoke: int wheelScrollLines()
    C.C_ZN12QApplication16wheelScrollLinesEv()
  default:
    qtrt.ErrorResolve("QApplication", "wheelScrollLines", args)
  }

}

// setColorSpec(int)
func (this *QApplication) setColorSpec_s(args ...interface{}) () {
  // setColorSpec(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication12setColorSpecEi
    // invoke: void setColorSpec(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication12setColorSpecEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setColorSpec", args)
  }

}

// autoSipEnabled()
func (this *QApplication) autoSipEnabled(args ...interface{}) () {
  // autoSipEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QApplication14autoSipEnabledEv
    // invoke: bool autoSipEnabled()
    C.C_ZNK12QApplication14autoSipEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "autoSipEnabled", args)
  }

}

// setStyle(const class QString &)
func (this *QApplication) setStyle_s(args ...interface{}) () {
  // setStyle(const class QString &)
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication8setStyleERK7QString
    // invoke: QStyle * setStyle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication8setStyleERK7QString(arg0)
  case 1:
    // invoke: _ZN12QApplication8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication8setStyleEP6QStyle(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStyle", args)
  }

}

// notify(class QObject *, class QEvent *)
func (this *QApplication) notify(args ...interface{}) () {
  // notify(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication6notifyEP7QObjectP6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "notify", args)
  }

}

// focusWidget()
func (this *QApplication) focusWidget_s(args ...interface{}) () {
  // focusWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication11focusWidgetEv
    // invoke: QWidget * focusWidget()
    C.C_ZN12QApplication11focusWidgetEv()
  default:
    qtrt.ErrorResolve("QApplication", "focusWidget", args)
  }

}

// allWidgets()
func (this *QApplication) allWidgets_s(args ...interface{}) () {
  // allWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10allWidgetsEv
    // invoke: QWidgetList allWidgets()
    C.C_ZN12QApplication10allWidgetsEv()
  default:
    qtrt.ErrorResolve("QApplication", "allWidgets", args)
  }

}

// widgetAt(int, int)
func (this *QApplication) widgetAt_s(args ...interface{}) () {
  // widgetAt(int, int)
  // widgetAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication8widgetAtEii
    // invoke: QWidget * widgetAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication8widgetAtEii(arg0, arg1)
  case 1:
    // invoke: _ZN12QApplication8widgetAtERK6QPoint
    // invoke: QWidget * widgetAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication8widgetAtERK6QPoint(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "widgetAt", args)
  }

}

// aboutQt()
func (this *QApplication) aboutQt_s(args ...interface{}) () {
  // aboutQt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7aboutQtEv
    // invoke: void aboutQt()
    C.C_ZN12QApplication7aboutQtEv()
  default:
    qtrt.ErrorResolve("QApplication", "aboutQt", args)
  }

}

// setPalette(const class QPalette &, const char *)
func (this *QApplication) setPalette_s(args ...interface{}) () {
  // setPalette(const class QPalette &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10setPaletteERK8QPalettePKc
    // invoke: void setPalette(const class QPalette &, const char *)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication10setPaletteERK8QPalettePKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "setPalette", args)
  }

}

// palette(const char *)
func (this *QApplication) palette_s(args ...interface{}) () {
  // palette(const char *)
  // palette(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7paletteEPKc
    // invoke: QPalette palette(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication7paletteEPKc(arg0)
  case 1:
    // invoke: _ZN12QApplication7paletteEPK7QWidget
    // invoke: QPalette palette(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication7paletteEPK7QWidget(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "palette", args)
  }

}

// font()
func (this *QApplication) font_s(args ...interface{}) () {
  // font()
  // font(const char *)
  // font(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication4fontEv
    // invoke: QFont font()
    C.C_ZN12QApplication4fontEv()
  case 1:
    // invoke: _ZN12QApplication4fontEPKc
    // invoke: QFont font(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication4fontEPKc(arg0)
  case 2:
    // invoke: _ZN12QApplication4fontEPK7QWidget
    // invoke: QFont font(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication4fontEPK7QWidget(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "font", args)
  }

}

// setDoubleClickInterval(int)
func (this *QApplication) setDoubleClickInterval_s(args ...interface{}) () {
  // setDoubleClickInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication22setDoubleClickIntervalEi
    // invoke: void setDoubleClickInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication22setDoubleClickIntervalEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setDoubleClickInterval", args)
  }

}

// topLevelWidgets()
func (this *QApplication) topLevelWidgets_s(args ...interface{}) () {
  // topLevelWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication15topLevelWidgetsEv
    // invoke: QWidgetList topLevelWidgets()
    C.C_ZN12QApplication15topLevelWidgetsEv()
  default:
    qtrt.ErrorResolve("QApplication", "topLevelWidgets", args)
  }

}

// desktop()
func (this *QApplication) desktop_s(args ...interface{}) () {
  // desktop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7desktopEv
    // invoke: QDesktopWidget * desktop()
    C.C_ZN12QApplication7desktopEv()
  default:
    qtrt.ErrorResolve("QApplication", "desktop", args)
  }

}

// beep()
func (this *QApplication) beep_s(args ...interface{}) () {
  // beep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication4beepEv
    // invoke: void beep()
    C.C_ZN12QApplication4beepEv()
  default:
    qtrt.ErrorResolve("QApplication", "beep", args)
  }

}

// metaObject()
func (this *QApplication) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QApplication10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QApplication10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "metaObject", args)
  }

}

// setCursorFlashTime(int)
func (this *QApplication) setCursorFlashTime_s(args ...interface{}) () {
  // setCursorFlashTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication18setCursorFlashTimeEi
    // invoke: void setCursorFlashTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication18setCursorFlashTimeEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setCursorFlashTime", args)
  }

}

// activeWindow()
func (this *QApplication) activeWindow_s(args ...interface{}) () {
  // activeWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication12activeWindowEv
    // invoke: QWidget * activeWindow()
    C.C_ZN12QApplication12activeWindowEv()
  default:
    qtrt.ErrorResolve("QApplication", "activeWindow", args)
  }

}

// topLevelAt(const class QPoint &)
func (this *QApplication) topLevelAt_s(args ...interface{}) () {
  // topLevelAt(const class QPoint &)
  // topLevelAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10topLevelAtERK6QPoint
    // invoke: QWidget * topLevelAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication10topLevelAtERK6QPoint(arg0)
  case 1:
    // invoke: _ZN12QApplication10topLevelAtEii
    // invoke: QWidget * topLevelAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication10topLevelAtEii(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "topLevelAt", args)
  }

}

// fontMetrics()
func (this *QApplication) fontMetrics_s(args ...interface{}) () {
  // fontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication11fontMetricsEv
    // invoke: QFontMetrics fontMetrics()
    C.C_ZN12QApplication11fontMetricsEv()
  default:
    qtrt.ErrorResolve("QApplication", "fontMetrics", args)
  }

}

// setKeyboardInputInterval(int)
func (this *QApplication) setKeyboardInputInterval_s(args ...interface{}) () {
  // setKeyboardInputInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication24setKeyboardInputIntervalEi
    // invoke: void setKeyboardInputInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication24setKeyboardInputIntervalEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setKeyboardInputInterval", args)
  }

}

// colorSpec()
func (this *QApplication) colorSpec_s(args ...interface{}) () {
  // colorSpec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication9colorSpecEv
    // invoke: int colorSpec()
    C.C_ZN12QApplication9colorSpecEv()
  default:
    qtrt.ErrorResolve("QApplication", "colorSpec", args)
  }

}

// style()
func (this *QApplication) style_s(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication5styleEv
    // invoke: QStyle * style()
    C.C_ZN12QApplication5styleEv()
  default:
    qtrt.ErrorResolve("QApplication", "style", args)
  }

}

// alert(class QWidget *, int)
func (this *QApplication) alert_s(args ...interface{}) () {
  // alert(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication5alertEP7QWidgeti
    // invoke: void alert(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication5alertEP7QWidgeti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "alert", args)
  }

}

// activeModalWidget()
func (this *QApplication) activeModalWidget_s(args ...interface{}) () {
  // activeModalWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication17activeModalWidgetEv
    // invoke: QWidget * activeModalWidget()
    C.C_ZN12QApplication17activeModalWidgetEv()
  default:
    qtrt.ErrorResolve("QApplication", "activeModalWidget", args)
  }

}

// styleSheet()
func (this *QApplication) styleSheet(args ...interface{}) () {
  // styleSheet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QApplication10styleSheetEv
    // invoke: QString styleSheet()
    C.C_ZNK12QApplication10styleSheetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "styleSheet", args)
  }

}

// windowIcon()
func (this *QApplication) windowIcon_s(args ...interface{}) () {
  // windowIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10windowIconEv
    // invoke: QIcon windowIcon()
    C.C_ZN12QApplication10windowIconEv()
  default:
    qtrt.ErrorResolve("QApplication", "windowIcon", args)
  }

}

// cursorFlashTime()
func (this *QApplication) cursorFlashTime_s(args ...interface{}) () {
  // cursorFlashTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication15cursorFlashTimeEv
    // invoke: int cursorFlashTime()
    C.C_ZN12QApplication15cursorFlashTimeEv()
  default:
    qtrt.ErrorResolve("QApplication", "cursorFlashTime", args)
  }

}

// setGlobalStrut(const class QSize &)
func (this *QApplication) setGlobalStrut_s(args ...interface{}) () {
  // setGlobalStrut(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication14setGlobalStrutERK5QSize
    // invoke: void setGlobalStrut(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication14setGlobalStrutERK5QSize(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setGlobalStrut", args)
  }

}

// setActiveWindow(class QWidget *)
func (this *QApplication) setActiveWindow_s(args ...interface{}) () {
  // setActiveWindow(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication15setActiveWindowEP7QWidget
    // invoke: void setActiveWindow(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication15setActiveWindowEP7QWidget(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setActiveWindow", args)
  }

}

// ~QApplication()
func (this *QApplication) FreeQApplication(args ...interface{}) () {
  // ~QApplication()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplicationD0Ev
    // invoke: void ~QApplication()
    C.C_ZN12QApplicationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "~QApplication", args)
  }

}

// globalStrut()
func (this *QApplication) globalStrut_s(args ...interface{}) () {
  // globalStrut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication11globalStrutEv
    // invoke: QSize globalStrut()
    C.C_ZN12QApplication11globalStrutEv()
  default:
    qtrt.ErrorResolve("QApplication", "globalStrut", args)
  }

}

// setStartDragDistance(int)
func (this *QApplication) setStartDragDistance_s(args ...interface{}) () {
  // setStartDragDistance(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication20setStartDragDistanceEi
    // invoke: void setStartDragDistance(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication20setStartDragDistanceEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragDistance", args)
  }

}

// setWheelScrollLines(int)
func (this *QApplication) setWheelScrollLines_s(args ...interface{}) () {
  // setWheelScrollLines(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication19setWheelScrollLinesEi
    // invoke: void setWheelScrollLines(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication19setWheelScrollLinesEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setWheelScrollLines", args)
  }

}

// keyboardInputInterval()
func (this *QApplication) keyboardInputInterval_s(args ...interface{}) () {
  // keyboardInputInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication21keyboardInputIntervalEv
    // invoke: int keyboardInputInterval()
    C.C_ZN12QApplication21keyboardInputIntervalEv()
  default:
    qtrt.ErrorResolve("QApplication", "keyboardInputInterval", args)
  }

}

// setStartDragTime(int)
func (this *QApplication) setStartDragTime_s(args ...interface{}) () {
  // setStartDragTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication16setStartDragTimeEi
    // invoke: void setStartDragTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication16setStartDragTimeEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragTime", args)
  }

}

// QApplication(int &, char **, int)
func NewQApplication(args ...interface{}) QApplication {
  // QApplication(int &, char **, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int &"
  vtys[0][1] = qtrt.ByteTy(true) // "char **"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplicationC1ERiPPci
    // invoke: void QApplication(int &, char **, int)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QApplicationC2ERiPPci(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QApplication", "QApplication", args)
  }

  return QApplication{}
}

// startDragTime()
func (this *QApplication) startDragTime_s(args ...interface{}) () {
  // startDragTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication13startDragTimeEv
    // invoke: int startDragTime()
    C.C_ZN12QApplication13startDragTimeEv()
  default:
    qtrt.ErrorResolve("QApplication", "startDragTime", args)
  }

}

// setWindowIcon(const class QIcon &)
func (this *QApplication) setWindowIcon_s(args ...interface{}) () {
  // setWindowIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication13setWindowIconERK5QIcon
    // invoke: void setWindowIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication13setWindowIconERK5QIcon(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setWindowIcon", args)
  }

}

// setAutoSipEnabled(const _Bool)
func (this *QApplication) setAutoSipEnabled(args ...interface{}) () {
  // setAutoSipEnabled(const _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "const bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication17setAutoSipEnabledEb
    // invoke: void setAutoSipEnabled(const _Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication17setAutoSipEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setAutoSipEnabled", args)
  }

}

// setStyleSheet(const class QString &)
func (this *QApplication) setStyleSheet(args ...interface{}) () {
  // setStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication13setStyleSheetERK7QString
    // invoke: void setStyleSheet(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication13setStyleSheetERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStyleSheet", args)
  }

}

// exec()
func (this *QApplication) exec_s(args ...interface{}) () {
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication4execEv
    // invoke: int exec()
    C.C_ZN12QApplication4execEv()
  default:
    qtrt.ErrorResolve("QApplication", "exec", args)
  }

}

// setFont(const class QFont &, const char *)
func (this *QApplication) setFont_s(args ...interface{}) () {
  // setFont(const class QFont &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7setFontERK5QFontPKc
    // invoke: void setFont(const class QFont &, const char *)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication7setFontERK5QFontPKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "setFont", args)
  }

}

// doubleClickInterval()
func (this *QApplication) doubleClickInterval_s(args ...interface{}) () {
  // doubleClickInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication19doubleClickIntervalEv
    // invoke: int doubleClickInterval()
    C.C_ZN12QApplication19doubleClickIntervalEv()
  default:
    qtrt.ErrorResolve("QApplication", "doubleClickInterval", args)
  }

}

// <= body block end

