package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QString QApplication::styleSheet();
extern void _ZNK12QApplication10styleSheetEv(void* qthis);
  // proto: static QPalette QApplication::palette(const char * className);
extern void _ZN12QApplication7paletteEPKc(char* arg0);
  // proto: static QWidget * QApplication::activeWindow();
extern void _ZN12QApplication12activeWindowEv();
  // proto: static void QApplication::setKeyboardInputInterval(int );
extern void _ZN12QApplication24setKeyboardInputIntervalEi(int arg0);
  // proto: static QWidget * QApplication::focusWidget();
extern void _ZN12QApplication11focusWidgetEv();
  // proto: static QFontMetrics QApplication::fontMetrics();
extern void _ZN12QApplication11fontMetricsEv();
  // proto: static QFont QApplication::font(const char * className);
extern void _ZN12QApplication4fontEPKc(char* arg0);
  // proto: static QStyle * QApplication::style();
extern void _ZN12QApplication5styleEv();
  // proto: static QWidget * QApplication::widgetAt(const QPoint & p);
extern void _ZN12QApplication8widgetAtERK6QPoint(void* arg0);
  // proto: static void QApplication::setActiveWindow(QWidget * act);
extern void _ZN12QApplication15setActiveWindowEP7QWidget(void* arg0);
  // proto: static QFont QApplication::font();
extern void _ZN12QApplication4fontEv();
  // proto: static void QApplication::setWheelScrollLines(int );
extern void _ZN12QApplication19setWheelScrollLinesEi(int arg0);
  // proto:  void QApplication::setStyleSheet(const QString & sheet);
extern void _ZN12QApplication13setStyleSheetERK7QString(void* qthis, void* arg0);
  // proto:  void QApplication::setAutoSipEnabled(const bool enabled);
extern void _ZN12QApplication17setAutoSipEnabledEb(void* qthis, const bool arg0);
  // proto:  const QMetaObject * QApplication::metaObject();
extern void _ZNK12QApplication10metaObjectEv(void* qthis);
  // proto: static int QApplication::keyboardInputInterval();
extern void _ZN12QApplication21keyboardInputIntervalEv();
  // proto: static int QApplication::cursorFlashTime();
extern void _ZN12QApplication15cursorFlashTimeEv();
  // proto: static int QApplication::startDragDistance();
extern void _ZN12QApplication17startDragDistanceEv();
  // proto: static QDesktopWidget * QApplication::desktop();
extern void _ZN12QApplication7desktopEv();
  // proto: static void QApplication::setStartDragDistance(int l);
extern void _ZN12QApplication20setStartDragDistanceEi(int arg0);
  // proto: static QFont QApplication::font(const QWidget * );
extern void _ZN12QApplication4fontEPK7QWidget(void* arg0);
  // proto: static int QApplication::colorSpec();
extern void _ZN12QApplication9colorSpecEv();
  // proto: static void QApplication::setFont(const QFont & , const char * className);
extern void _ZN12QApplication7setFontERK5QFontPKc(void* arg0, char* arg1);
  // proto: static void QApplication::closeAllWindows();
extern void _ZN12QApplication15closeAllWindowsEv();
  // proto: static void QApplication::setCursorFlashTime(int );
extern void _ZN12QApplication18setCursorFlashTimeEi(int arg0);
  // proto: static QWidget * QApplication::widgetAt(int x, int y);
extern void demth_ZN12QApplication8widgetAtEii(int arg0, int arg1);
  // proto: static void QApplication::alert(QWidget * widget, int duration);
extern void _ZN12QApplication5alertEP7QWidgeti(void* arg0, int arg1);
  // proto: static QPalette QApplication::palette(const QWidget * );
extern void _ZN12QApplication7paletteEPK7QWidget(void* arg0);
  // proto: static int QApplication::wheelScrollLines();
extern void _ZN12QApplication16wheelScrollLinesEv();
  // proto:  void QApplication::QApplication(const QApplication & );
extern void* dector_ZN12QApplicationC1ERKS_(void* arg0);
extern void _ZN12QApplicationC1ERKS_(void* qthis, void* arg0);
  // proto: static void QApplication::aboutQt();
extern void _ZN12QApplication7aboutQtEv();
  // proto: static QWidget * QApplication::activeModalWidget();
extern void _ZN12QApplication17activeModalWidgetEv();
  // proto: static QWidget * QApplication::activePopupWidget();
extern void _ZN12QApplication17activePopupWidgetEv();
  // proto:  void QApplication::QApplication(int & argc, char ** argv, int );
extern void* dector_ZN12QApplicationC1ERiPPci(int* arg0, char* arg1, int arg2);
extern void _ZN12QApplicationC1ERiPPci(void* qthis, int* arg0, char* arg1, int arg2);
  // proto: static void QApplication::setStartDragTime(int ms);
extern void _ZN12QApplication16setStartDragTimeEi(int arg0);
  // proto: static QWidget * QApplication::topLevelAt(int x, int y);
extern void demth_ZN12QApplication10topLevelAtEii(int arg0, int arg1);
  // proto: static void QApplication::setStyle(QStyle * );
extern void _ZN12QApplication8setStyleEP6QStyle(void* arg0);
  // proto:  void QApplication::~QApplication();
extern void _ZN12QApplicationD0Ev(void* qthis);
  // proto: static void QApplication::setDoubleClickInterval(int );
extern void _ZN12QApplication22setDoubleClickIntervalEi(int arg0);
  // proto: static void QApplication::setGlobalStrut(const QSize & );
extern void _ZN12QApplication14setGlobalStrutERK5QSize(void* arg0);
  // proto: static void QApplication::setColorSpec(int );
extern void _ZN12QApplication12setColorSpecEi(int arg0);
  // proto: static QWidgetList QApplication::allWidgets();
extern void _ZN12QApplication10allWidgetsEv();
  // proto: static QSize QApplication::globalStrut();
extern void _ZN12QApplication11globalStrutEv();
  // proto: static void QApplication::setPalette(const QPalette & , const char * className);
extern void _ZN12QApplication10setPaletteERK8QPalettePKc(void* arg0, char* arg1);
  // proto: static QStyle * QApplication::setStyle(const QString & );
extern void _ZN12QApplication8setStyleERK7QString(void* arg0);
  // proto: static QWidgetList QApplication::topLevelWidgets();
extern void _ZN12QApplication15topLevelWidgetsEv();
  // proto: static int QApplication::exec();
extern void _ZN12QApplication4execEv();
  // proto: static void QApplication::setWindowIcon(const QIcon & icon);
extern void _ZN12QApplication13setWindowIconERK5QIcon(void* arg0);
  // proto: static void QApplication::beep();
extern void _ZN12QApplication4beepEv();
  // proto:  bool QApplication::notify(QObject * , QEvent * );
extern void _ZN12QApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto:  bool QApplication::autoSipEnabled();
extern void _ZNK12QApplication14autoSipEnabledEv(void* qthis);
  // proto: static QWidget * QApplication::topLevelAt(const QPoint & p);
extern void _ZN12QApplication10topLevelAtERK6QPoint(void* arg0);
  // proto: static int QApplication::startDragTime();
extern void _ZN12QApplication13startDragTimeEv();
  // proto: static int QApplication::doubleClickInterval();
extern void _ZN12QApplication19doubleClickIntervalEv();
  // proto: static QIcon QApplication::windowIcon();
extern void _ZN12QApplication10windowIconEv();
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
  qclsinst uint64 /* *mut c_void*/;
//  _focusChanged QApplication_focusChanged_signal;
}

  // proto:  QString QApplication::styleSheet();
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
  default:
    qtrt.ErrorResolve("QApplication", "styleSheet", args)
  }

}

  // proto: static QPalette QApplication::palette(const char * className);
func (this *QApplication) palette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "palette", args)
  }

}

  // proto: static QWidget * QApplication::activeWindow();
func (this *QApplication) activeWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "activeWindow", args)
  }

}

  // proto: static void QApplication::setKeyboardInputInterval(int );
func (this *QApplication) setKeyboardInputInterval_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setKeyboardInputInterval", args)
  }

}

  // proto: static QWidget * QApplication::focusWidget();
func (this *QApplication) focusWidget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "focusWidget", args)
  }

}

  // proto: static QFontMetrics QApplication::fontMetrics();
func (this *QApplication) fontMetrics_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "fontMetrics", args)
  }

}

  // proto: static QFont QApplication::font(const char * className);
func (this *QApplication) font_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "font", args)
  }

}

  // proto: static QStyle * QApplication::style();
func (this *QApplication) style_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "style", args)
  }

}

  // proto: static QWidget * QApplication::widgetAt(const QPoint & p);
func (this *QApplication) widgetAt_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "widgetAt", args)
  }

}

  // proto: static void QApplication::setActiveWindow(QWidget * act);
func (this *QApplication) setActiveWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setActiveWindow", args)
  }

}

  // proto: static void QApplication::setWheelScrollLines(int );
func (this *QApplication) setWheelScrollLines_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setWheelScrollLines", args)
  }

}

  // proto:  void QApplication::setStyleSheet(const QString & sheet);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QApplication", "setStyleSheet", args)
  }

}

  // proto:  void QApplication::setAutoSipEnabled(const bool enabled);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QApplication", "setAutoSipEnabled", args)
  }

}

  // proto:  const QMetaObject * QApplication::metaObject();
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
  default:
    qtrt.ErrorResolve("QApplication", "metaObject", args)
  }

}

  // proto: static int QApplication::keyboardInputInterval();
func (this *QApplication) keyboardInputInterval_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "keyboardInputInterval", args)
  }

}

  // proto: static int QApplication::cursorFlashTime();
func (this *QApplication) cursorFlashTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "cursorFlashTime", args)
  }

}

  // proto: static int QApplication::startDragDistance();
func (this *QApplication) startDragDistance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "startDragDistance", args)
  }

}

  // proto: static QDesktopWidget * QApplication::desktop();
func (this *QApplication) desktop_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "desktop", args)
  }

}

  // proto: static void QApplication::setStartDragDistance(int l);
func (this *QApplication) setStartDragDistance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragDistance", args)
  }

}

  // proto: static int QApplication::colorSpec();
func (this *QApplication) colorSpec_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "colorSpec", args)
  }

}

  // proto: static void QApplication::setFont(const QFont & , const char * className);
func (this *QApplication) setFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setFont", args)
  }

}

  // proto: static void QApplication::closeAllWindows();
func (this *QApplication) closeAllWindows_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "closeAllWindows", args)
  }

}

  // proto: static void QApplication::setCursorFlashTime(int );
func (this *QApplication) setCursorFlashTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setCursorFlashTime", args)
  }

}

  // proto: static void QApplication::alert(QWidget * widget, int duration);
func (this *QApplication) alert_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "alert", args)
  }

}

  // proto: static int QApplication::wheelScrollLines();
func (this *QApplication) wheelScrollLines_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "wheelScrollLines", args)
  }

}

  // proto:  void QApplication::QApplication(const QApplication & );
func NewQApplication(args ...interface{}) QApplication {
  return QApplication{}
}

  // proto: static void QApplication::aboutQt();
func (this *QApplication) aboutQt_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "aboutQt", args)
  }

}

  // proto: static QWidget * QApplication::activeModalWidget();
func (this *QApplication) activeModalWidget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "activeModalWidget", args)
  }

}

  // proto: static QWidget * QApplication::activePopupWidget();
func (this *QApplication) activePopupWidget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "activePopupWidget", args)
  }

}

  // proto: static void QApplication::setStartDragTime(int ms);
func (this *QApplication) setStartDragTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragTime", args)
  }

}

  // proto: static QWidget * QApplication::topLevelAt(int x, int y);
func (this *QApplication) topLevelAt_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "topLevelAt", args)
  }

}

  // proto: static void QApplication::setStyle(QStyle * );
func (this *QApplication) setStyle_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setStyle", args)
  }

}

  // proto:  void QApplication::~QApplication();
func (this *QApplication) FreeQApplication(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "~QApplication", args)
  }

}

  // proto: static void QApplication::setDoubleClickInterval(int );
func (this *QApplication) setDoubleClickInterval_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setDoubleClickInterval", args)
  }

}

  // proto: static void QApplication::setGlobalStrut(const QSize & );
func (this *QApplication) setGlobalStrut_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setGlobalStrut", args)
  }

}

  // proto: static void QApplication::setColorSpec(int );
func (this *QApplication) setColorSpec_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setColorSpec", args)
  }

}

  // proto: static QWidgetList QApplication::allWidgets();
func (this *QApplication) allWidgets_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "allWidgets", args)
  }

}

  // proto: static QSize QApplication::globalStrut();
func (this *QApplication) globalStrut_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "globalStrut", args)
  }

}

  // proto: static void QApplication::setPalette(const QPalette & , const char * className);
func (this *QApplication) setPalette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setPalette", args)
  }

}

  // proto: static QWidgetList QApplication::topLevelWidgets();
func (this *QApplication) topLevelWidgets_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "topLevelWidgets", args)
  }

}

  // proto: static int QApplication::exec();
func (this *QApplication) exec_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "exec", args)
  }

}

  // proto: static void QApplication::setWindowIcon(const QIcon & icon);
func (this *QApplication) setWindowIcon_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "setWindowIcon", args)
  }

}

  // proto: static void QApplication::beep();
func (this *QApplication) beep_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "beep", args)
  }

}

  // proto:  bool QApplication::notify(QObject * , QEvent * );
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QApplication", "notify", args)
  }

}

  // proto:  bool QApplication::autoSipEnabled();
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
  default:
    qtrt.ErrorResolve("QApplication", "autoSipEnabled", args)
  }

}

  // proto: static int QApplication::startDragTime();
func (this *QApplication) startDragTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "startDragTime", args)
  }

}

  // proto: static int QApplication::doubleClickInterval();
func (this *QApplication) doubleClickInterval_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "doubleClickInterval", args)
  }

}

  // proto: static QIcon QApplication::windowIcon();
func (this *QApplication) windowIcon_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QApplication", "windowIcon", args)
  }

}

// <= body block end

