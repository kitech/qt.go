package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qguiapplication.h
// dst-file: /src/gui/qguiapplication.go
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
  // proto:  void QGuiApplication::~QGuiApplication();
extern void _ZN15QGuiApplicationD0Ev(void* qthis);
  // proto: static void QGuiApplication::setFont(const QFont & );
extern void _ZN15QGuiApplication7setFontERK5QFont(void* arg0);
  // proto: static QString QGuiApplication::platformName();
extern void _ZN15QGuiApplication12platformNameEv();
  // proto: static QList<QScreen *> QGuiApplication::screens();
extern void _ZN15QGuiApplication7screensEv();
  // proto: static void QGuiApplication::setPalette(const QPalette & pal);
extern void _ZN15QGuiApplication10setPaletteERK8QPalette(void* arg0);
  // proto: static QInputMethod * QGuiApplication::inputMethod();
extern void _ZN15QGuiApplication11inputMethodEv();
  // proto:  void QGuiApplication::QGuiApplication(const QGuiApplication & );
extern void* dector_ZN15QGuiApplicationC1ERKS_(void* arg0);
extern void _ZN15QGuiApplicationC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QGuiApplication::isSavingSession();
extern void _ZNK15QGuiApplication15isSavingSessionEv(void* qthis);
  // proto: static QFont QGuiApplication::font();
extern void _ZN15QGuiApplication4fontEv();
  // proto:  bool QGuiApplication::isSessionRestored();
extern void _ZNK15QGuiApplication17isSessionRestoredEv(void* qthis);
  // proto:  QString QGuiApplication::sessionKey();
extern void _ZNK15QGuiApplication10sessionKeyEv(void* qthis);
  // proto: static bool QGuiApplication::desktopSettingsAware();
extern void _ZN15QGuiApplication20desktopSettingsAwareEv();
  // proto: static void QGuiApplication::sync();
extern void _ZN15QGuiApplication4syncEv();
  // proto: static void QGuiApplication::setQuitOnLastWindowClosed(bool quit);
extern void _ZN15QGuiApplication25setQuitOnLastWindowClosedEb(bool arg0);
  // proto: static QScreen * QGuiApplication::primaryScreen();
extern void _ZN15QGuiApplication13primaryScreenEv();
  // proto: static QCursor * QGuiApplication::overrideCursor();
extern void _ZN15QGuiApplication14overrideCursorEv();
  // proto: static QIcon QGuiApplication::windowIcon();
extern void _ZN15QGuiApplication10windowIconEv();
  // proto: static QStyleHints * QGuiApplication::styleHints();
extern void _ZN15QGuiApplication10styleHintsEv();
  // proto: static QClipboard * QGuiApplication::clipboard();
extern void _ZN15QGuiApplication9clipboardEv();
  // proto: static QPalette QGuiApplication::palette();
extern void _ZN15QGuiApplication7paletteEv();
  // proto:  bool QGuiApplication::notify(QObject * , QEvent * );
extern void _ZN15QGuiApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto: static QWindowList QGuiApplication::topLevelWindows();
extern void _ZN15QGuiApplication15topLevelWindowsEv();
  // proto: static bool QGuiApplication::isRightToLeft();
extern void demth_ZN15QGuiApplication13isRightToLeftEv();
  // proto: static void QGuiApplication::changeOverrideCursor(const QCursor & );
extern void _ZN15QGuiApplication20changeOverrideCursorERK7QCursor(void* arg0);
  // proto: static QWindowList QGuiApplication::allWindows();
extern void _ZN15QGuiApplication10allWindowsEv();
  // proto: static void QGuiApplication::setOverrideCursor(const QCursor & );
extern void _ZN15QGuiApplication17setOverrideCursorERK7QCursor(void* arg0);
  // proto: static void QGuiApplication::setWindowIcon(const QIcon & icon);
extern void _ZN15QGuiApplication13setWindowIconERK5QIcon(void* arg0);
  // proto:  QString QGuiApplication::sessionId();
extern void _ZNK15QGuiApplication9sessionIdEv(void* qthis);
  // proto: static void QGuiApplication::setApplicationDisplayName(const QString & name);
extern void _ZN15QGuiApplication25setApplicationDisplayNameERK7QString(void* arg0);
  // proto: static bool QGuiApplication::isLeftToRight();
extern void demth_ZN15QGuiApplication13isLeftToRightEv();
  // proto: static QWindow * QGuiApplication::topLevelAt(const QPoint & pos);
extern void _ZN15QGuiApplication10topLevelAtERK6QPoint(void* arg0);
  // proto:  void QGuiApplication::QGuiApplication(int & argc, char ** argv, int );
extern void* dector_ZN15QGuiApplicationC1ERiPPci(int* arg0, char* arg1, int arg2);
extern void _ZN15QGuiApplicationC1ERiPPci(void* qthis, int* arg0, char* arg1, int arg2);
  // proto: static void QGuiApplication::setDesktopSettingsAware(bool on);
extern void _ZN15QGuiApplication23setDesktopSettingsAwareEb(bool arg0);
  // proto: static QWindow * QGuiApplication::modalWindow();
extern void _ZN15QGuiApplication11modalWindowEv();
  // proto: static QString QGuiApplication::applicationDisplayName();
extern void _ZN15QGuiApplication22applicationDisplayNameEv();
  // proto: static int QGuiApplication::exec();
extern void _ZN15QGuiApplication4execEv();
  // proto: static bool QGuiApplication::quitOnLastWindowClosed();
extern void _ZN15QGuiApplication22quitOnLastWindowClosedEv();
  // proto: static void QGuiApplication::restoreOverrideCursor();
extern void _ZN15QGuiApplication21restoreOverrideCursorEv();
  // proto: static QPlatformNativeInterface * QGuiApplication::platformNativeInterface();
extern void _ZN15QGuiApplication23platformNativeInterfaceEv();
  // proto:  const QMetaObject * QGuiApplication::metaObject();
extern void _ZNK15QGuiApplication10metaObjectEv(void* qthis);
  // proto: static QObject * QGuiApplication::focusObject();
extern void _ZN15QGuiApplication11focusObjectEv();
  // proto: static QWindow * QGuiApplication::focusWindow();
extern void _ZN15QGuiApplication11focusWindowEv();
  // proto:  qreal QGuiApplication::devicePixelRatio();
extern void _ZNK15QGuiApplication16devicePixelRatioEv(void* qthis);
  // proto: static QFunctionPointer QGuiApplication::platformFunction(const QByteArray & function);
extern void _ZN15QGuiApplication16platformFunctionERK10QByteArray(void* arg0);
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

// class sizeof(QGuiApplication)=1
type QGuiApplication struct {
  /*qbase*/ QCoreApplication;
  qclsinst uint64 /* *mut c_void*/;
//  _saveStateRequest QGuiApplication_saveStateRequest_signal;
//  _applicationStateChanged QGuiApplication_applicationStateChanged_signal;
//  _screenAdded QGuiApplication_screenAdded_signal;
//  _screenRemoved QGuiApplication_screenRemoved_signal;
//  _focusObjectChanged QGuiApplication_focusObjectChanged_signal;
//  _paletteChanged QGuiApplication_paletteChanged_signal;
//  _focusWindowChanged QGuiApplication_focusWindowChanged_signal;
//  _lastWindowClosed QGuiApplication_lastWindowClosed_signal;
//  _fontDatabaseChanged QGuiApplication_fontDatabaseChanged_signal;
//  _layoutDirectionChanged QGuiApplication_layoutDirectionChanged_signal;
//  _commitDataRequest QGuiApplication_commitDataRequest_signal;
}

  // proto:  void QGuiApplication::~QGuiApplication();
func (this *QGuiApplication) FreeQGuiApplication(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "~QGuiApplication", args)
  }

}

  // proto: static void QGuiApplication::setFont(const QFont & );
func (this *QGuiApplication) setFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setFont", args)
  }

}

  // proto: static QString QGuiApplication::platformName();
func (this *QGuiApplication) platformName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformName", args)
  }

}

  // proto: static QList<QScreen *> QGuiApplication::screens();
func (this *QGuiApplication) screens_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "screens", args)
  }

}

  // proto: static void QGuiApplication::setPalette(const QPalette & pal);
func (this *QGuiApplication) setPalette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setPalette", args)
  }

}

  // proto: static QInputMethod * QGuiApplication::inputMethod();
func (this *QGuiApplication) inputMethod_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "inputMethod", args)
  }

}

  // proto:  void QGuiApplication::QGuiApplication(const QGuiApplication & );
func NewQGuiApplication(args ...interface{}) QGuiApplication {
  return QGuiApplication{}
}

  // proto:  bool QGuiApplication::isSavingSession();
func (this *QGuiApplication) isSavingSession(args ...interface{}) () {
  // isSavingSession()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication15isSavingSessionEv
    // invoke: bool isSavingSession()
    C._ZNK15QGuiApplication15isSavingSessionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSavingSession", args)
  }

}

  // proto: static QFont QGuiApplication::font();
func (this *QGuiApplication) font_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "font", args)
  }

}

  // proto:  bool QGuiApplication::isSessionRestored();
func (this *QGuiApplication) isSessionRestored(args ...interface{}) () {
  // isSessionRestored()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication17isSessionRestoredEv
    // invoke: bool isSessionRestored()
    C._ZNK15QGuiApplication17isSessionRestoredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSessionRestored", args)
  }

}

  // proto:  QString QGuiApplication::sessionKey();
func (this *QGuiApplication) sessionKey(args ...interface{}) () {
  // sessionKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication10sessionKeyEv
    // invoke: QString sessionKey()
    C._ZNK15QGuiApplication10sessionKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionKey", args)
  }

}

  // proto: static bool QGuiApplication::desktopSettingsAware();
func (this *QGuiApplication) desktopSettingsAware_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "desktopSettingsAware", args)
  }

}

  // proto: static void QGuiApplication::sync();
func (this *QGuiApplication) sync_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "sync", args)
  }

}

  // proto: static void QGuiApplication::setQuitOnLastWindowClosed(bool quit);
func (this *QGuiApplication) setQuitOnLastWindowClosed_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setQuitOnLastWindowClosed", args)
  }

}

  // proto: static QScreen * QGuiApplication::primaryScreen();
func (this *QGuiApplication) primaryScreen_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "primaryScreen", args)
  }

}

  // proto: static QCursor * QGuiApplication::overrideCursor();
func (this *QGuiApplication) overrideCursor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "overrideCursor", args)
  }

}

  // proto: static QIcon QGuiApplication::windowIcon();
func (this *QGuiApplication) windowIcon_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "windowIcon", args)
  }

}

  // proto: static QStyleHints * QGuiApplication::styleHints();
func (this *QGuiApplication) styleHints_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "styleHints", args)
  }

}

  // proto: static QClipboard * QGuiApplication::clipboard();
func (this *QGuiApplication) clipboard_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "clipboard", args)
  }

}

  // proto: static QPalette QGuiApplication::palette();
func (this *QGuiApplication) palette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "palette", args)
  }

}

  // proto:  bool QGuiApplication::notify(QObject * , QEvent * );
func (this *QGuiApplication) notify(args ...interface{}) () {
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
    // invoke: _ZN15QGuiApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN15QGuiApplication6notifyEP7QObjectP6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGuiApplication", "notify", args)
  }

}

  // proto: static QWindowList QGuiApplication::topLevelWindows();
func (this *QGuiApplication) topLevelWindows_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelWindows", args)
  }

}

  // proto: static bool QGuiApplication::isRightToLeft();
func (this *QGuiApplication) isRightToLeft_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "isRightToLeft", args)
  }

}

  // proto: static void QGuiApplication::changeOverrideCursor(const QCursor & );
func (this *QGuiApplication) changeOverrideCursor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "changeOverrideCursor", args)
  }

}

  // proto: static QWindowList QGuiApplication::allWindows();
func (this *QGuiApplication) allWindows_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "allWindows", args)
  }

}

  // proto: static void QGuiApplication::setOverrideCursor(const QCursor & );
func (this *QGuiApplication) setOverrideCursor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setOverrideCursor", args)
  }

}

  // proto: static void QGuiApplication::setWindowIcon(const QIcon & icon);
func (this *QGuiApplication) setWindowIcon_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setWindowIcon", args)
  }

}

  // proto:  QString QGuiApplication::sessionId();
func (this *QGuiApplication) sessionId(args ...interface{}) () {
  // sessionId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication9sessionIdEv
    // invoke: QString sessionId()
    C._ZNK15QGuiApplication9sessionIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionId", args)
  }

}

  // proto: static void QGuiApplication::setApplicationDisplayName(const QString & name);
func (this *QGuiApplication) setApplicationDisplayName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setApplicationDisplayName", args)
  }

}

  // proto: static bool QGuiApplication::isLeftToRight();
func (this *QGuiApplication) isLeftToRight_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "isLeftToRight", args)
  }

}

  // proto: static QWindow * QGuiApplication::topLevelAt(const QPoint & pos);
func (this *QGuiApplication) topLevelAt_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelAt", args)
  }

}

  // proto: static void QGuiApplication::setDesktopSettingsAware(bool on);
func (this *QGuiApplication) setDesktopSettingsAware_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "setDesktopSettingsAware", args)
  }

}

  // proto: static QWindow * QGuiApplication::modalWindow();
func (this *QGuiApplication) modalWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "modalWindow", args)
  }

}

  // proto: static QString QGuiApplication::applicationDisplayName();
func (this *QGuiApplication) applicationDisplayName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "applicationDisplayName", args)
  }

}

  // proto: static int QGuiApplication::exec();
func (this *QGuiApplication) exec_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "exec", args)
  }

}

  // proto: static bool QGuiApplication::quitOnLastWindowClosed();
func (this *QGuiApplication) quitOnLastWindowClosed_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "quitOnLastWindowClosed", args)
  }

}

  // proto: static void QGuiApplication::restoreOverrideCursor();
func (this *QGuiApplication) restoreOverrideCursor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "restoreOverrideCursor", args)
  }

}

  // proto: static QPlatformNativeInterface * QGuiApplication::platformNativeInterface();
func (this *QGuiApplication) platformNativeInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformNativeInterface", args)
  }

}

  // proto:  const QMetaObject * QGuiApplication::metaObject();
func (this *QGuiApplication) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QGuiApplication10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "metaObject", args)
  }

}

  // proto: static QObject * QGuiApplication::focusObject();
func (this *QGuiApplication) focusObject_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusObject", args)
  }

}

  // proto: static QWindow * QGuiApplication::focusWindow();
func (this *QGuiApplication) focusWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusWindow", args)
  }

}

  // proto:  qreal QGuiApplication::devicePixelRatio();
func (this *QGuiApplication) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    C._ZNK15QGuiApplication16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "devicePixelRatio", args)
  }

}

  // proto: static QFunctionPointer QGuiApplication::platformFunction(const QByteArray & function);
func (this *QGuiApplication) platformFunction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformFunction", args)
  }

}

// <= body block end

