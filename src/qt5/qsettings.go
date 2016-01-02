package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qsettings.h
// dst-file: /src/core/qsettings.go
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
  // proto:  void QSettings::QSettings(QObject * parent);
extern void* dector_ZN9QSettingsC1EP7QObject(void* arg0);
extern void _ZN9QSettingsC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QSettings::isWritable();
extern void _ZNK9QSettings10isWritableEv(void* qthis);
  // proto:  QString QSettings::fileName();
extern void _ZNK9QSettings8fileNameEv(void* qthis);
  // proto:  bool QSettings::fallbacksEnabled();
extern void _ZNK9QSettings16fallbacksEnabledEv(void* qthis);
  // proto:  QString QSettings::applicationName();
extern void _ZNK9QSettings15applicationNameEv(void* qthis);
  // proto:  void QSettings::sync();
extern void _ZN9QSettings4syncEv(void* qthis);
  // proto:  void QSettings::setValue(const QString & key, const QVariant & value);
extern void _ZN9QSettings8setValueERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QSettings::setArrayIndex(int i);
extern void _ZN9QSettings13setArrayIndexEi(void* qthis, int arg0);
  // proto:  void QSettings::QSettings(const QString & organization, const QString & application, QObject * parent);
extern void* dector_ZN9QSettingsC1ERK7QStringS2_P7QObject(void* arg0, void* arg1, void* arg2);
extern void _ZN9QSettingsC1ERK7QStringS2_P7QObject(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QSettings::setIniCodec(QTextCodec * codec);
extern void _ZN9QSettings11setIniCodecEP10QTextCodec(void* qthis, void* arg0);
  // proto:  void QSettings::setIniCodec(const char * codecName);
extern void _ZN9QSettings11setIniCodecEPKc(void* qthis, char* arg0);
  // proto:  int QSettings::beginReadArray(const QString & prefix);
extern void _ZN9QSettings14beginReadArrayERK7QString(void* qthis, void* arg0);
  // proto:  void QSettings::clear();
extern void _ZN9QSettings5clearEv(void* qthis);
  // proto:  void QSettings::~QSettings();
extern void _ZN9QSettingsD0Ev(void* qthis);
  // proto:  QTextCodec * QSettings::iniCodec();
extern void _ZNK9QSettings8iniCodecEv(void* qthis);
  // proto: static void QSettings::setUserIniPath(const QString & dir);
extern void _ZN9QSettings14setUserIniPathERK7QString(void* arg0);
  // proto:  QStringList QSettings::childGroups();
extern void _ZNK9QSettings11childGroupsEv(void* qthis);
  // proto:  QVariant QSettings::value(const QString & key, const QVariant & defaultValue);
extern void _ZNK9QSettings5valueERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  QString QSettings::organizationName();
extern void _ZNK9QSettings16organizationNameEv(void* qthis);
  // proto:  const QMetaObject * QSettings::metaObject();
extern void _ZNK9QSettings10metaObjectEv(void* qthis);
  // proto:  void QSettings::setFallbacksEnabled(bool b);
extern void _ZN9QSettings19setFallbacksEnabledEb(void* qthis, bool arg0);
  // proto:  bool QSettings::contains(const QString & key);
extern void _ZNK9QSettings8containsERK7QString(void* qthis, void* arg0);
  // proto:  void QSettings::remove(const QString & key);
extern void _ZN9QSettings6removeERK7QString(void* qthis, void* arg0);
  // proto:  void QSettings::endGroup();
extern void _ZN9QSettings8endGroupEv(void* qthis);
  // proto:  void QSettings::beginWriteArray(const QString & prefix, int size);
extern void _ZN9QSettings15beginWriteArrayERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  void QSettings::beginGroup(const QString & prefix);
extern void _ZN9QSettings10beginGroupERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QSettings::childKeys();
extern void _ZNK9QSettings9childKeysEv(void* qthis);
  // proto:  void QSettings::QSettings(const QSettings & );
extern void* dector_ZN9QSettingsC1ERKS_(void* arg0);
extern void _ZN9QSettingsC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSettings::endArray();
extern void _ZN9QSettings8endArrayEv(void* qthis);
  // proto: static void QSettings::setSystemIniPath(const QString & dir);
extern void _ZN9QSettings16setSystemIniPathERK7QString(void* arg0);
  // proto:  QStringList QSettings::allKeys();
extern void _ZNK9QSettings7allKeysEv(void* qthis);
  // proto:  QString QSettings::group();
extern void _ZNK9QSettings5groupEv(void* qthis);
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

// class sizeof(QSettings)=1
type QSettings struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSettings::QSettings(QObject * parent);
func NewQSettings(args ...interface{}) QSettings {
  return QSettings{}
}

  // proto:  bool QSettings::isWritable();
func (this *QSettings) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10isWritableEv
  default:
    qtrt.ErrorResolve("QSettings", "isWritable", args)
  }

}

  // proto:  QString QSettings::fileName();
func (this *QSettings) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8fileNameEv
  default:
    qtrt.ErrorResolve("QSettings", "fileName", args)
  }

}

  // proto:  bool QSettings::fallbacksEnabled();
func (this *QSettings) fallbacksEnabled(args ...interface{}) () {
  // fallbacksEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16fallbacksEnabledEv
  default:
    qtrt.ErrorResolve("QSettings", "fallbacksEnabled", args)
  }

}

  // proto:  QString QSettings::applicationName();
func (this *QSettings) applicationName(args ...interface{}) () {
  // applicationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings15applicationNameEv
  default:
    qtrt.ErrorResolve("QSettings", "applicationName", args)
  }

}

  // proto:  void QSettings::sync();
func (this *QSettings) sync(args ...interface{}) () {
  // sync()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings4syncEv
  default:
    qtrt.ErrorResolve("QSettings", "sync", args)
  }

}

  // proto:  void QSettings::setValue(const QString & key, const QVariant & value);
func (this *QSettings) setValue(args ...interface{}) () {
  // setValue(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8setValueERK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QSettings", "setValue", args)
  }

}

  // proto:  void QSettings::setArrayIndex(int i);
func (this *QSettings) setArrayIndex(args ...interface{}) () {
  // setArrayIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings13setArrayIndexEi
  default:
    qtrt.ErrorResolve("QSettings", "setArrayIndex", args)
  }

}

  // proto:  void QSettings::setIniCodec(QTextCodec * codec);
func (this *QSettings) setIniCodec(args ...interface{}) () {
  // setIniCodec(class QTextCodec *)
  // setIniCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings11setIniCodecEP10QTextCodec
  case 1:
    // invoke: _ZN9QSettings11setIniCodecEPKc
  default:
    qtrt.ErrorResolve("QSettings", "setIniCodec", args)
  }

}

  // proto:  int QSettings::beginReadArray(const QString & prefix);
func (this *QSettings) beginReadArray(args ...interface{}) () {
  // beginReadArray(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14beginReadArrayERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "beginReadArray", args)
  }

}

  // proto:  void QSettings::clear();
func (this *QSettings) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings5clearEv
  default:
    qtrt.ErrorResolve("QSettings", "clear", args)
  }

}

  // proto:  void QSettings::~QSettings();
func (this *QSettings) FreeQSettings(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "~QSettings", args)
  }

}

  // proto:  QTextCodec * QSettings::iniCodec();
func (this *QSettings) iniCodec(args ...interface{}) () {
  // iniCodec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8iniCodecEv
  default:
    qtrt.ErrorResolve("QSettings", "iniCodec", args)
  }

}

  // proto: static void QSettings::setUserIniPath(const QString & dir);
func (this *QSettings) setUserIniPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "setUserIniPath", args)
  }

}

  // proto:  QStringList QSettings::childGroups();
func (this *QSettings) childGroups(args ...interface{}) () {
  // childGroups()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings11childGroupsEv
  default:
    qtrt.ErrorResolve("QSettings", "childGroups", args)
  }

}

  // proto:  QVariant QSettings::value(const QString & key, const QVariant & defaultValue);
func (this *QSettings) value(args ...interface{}) () {
  // value(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5valueERK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QSettings", "value", args)
  }

}

  // proto:  QString QSettings::organizationName();
func (this *QSettings) organizationName(args ...interface{}) () {
  // organizationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16organizationNameEv
  default:
    qtrt.ErrorResolve("QSettings", "organizationName", args)
  }

}

  // proto:  const QMetaObject * QSettings::metaObject();
func (this *QSettings) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10metaObjectEv
  default:
    qtrt.ErrorResolve("QSettings", "metaObject", args)
  }

}

  // proto:  void QSettings::setFallbacksEnabled(bool b);
func (this *QSettings) setFallbacksEnabled(args ...interface{}) () {
  // setFallbacksEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings19setFallbacksEnabledEb
  default:
    qtrt.ErrorResolve("QSettings", "setFallbacksEnabled", args)
  }

}

  // proto:  bool QSettings::contains(const QString & key);
func (this *QSettings) contains(args ...interface{}) () {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8containsERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "contains", args)
  }

}

  // proto:  void QSettings::remove(const QString & key);
func (this *QSettings) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings6removeERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "remove", args)
  }

}

  // proto:  void QSettings::endGroup();
func (this *QSettings) endGroup(args ...interface{}) () {
  // endGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endGroupEv
  default:
    qtrt.ErrorResolve("QSettings", "endGroup", args)
  }

}

  // proto:  void QSettings::beginWriteArray(const QString & prefix, int size);
func (this *QSettings) beginWriteArray(args ...interface{}) () {
  // beginWriteArray(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings15beginWriteArrayERK7QStringi
  default:
    qtrt.ErrorResolve("QSettings", "beginWriteArray", args)
  }

}

  // proto:  void QSettings::beginGroup(const QString & prefix);
func (this *QSettings) beginGroup(args ...interface{}) () {
  // beginGroup(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings10beginGroupERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "beginGroup", args)
  }

}

  // proto:  QStringList QSettings::childKeys();
func (this *QSettings) childKeys(args ...interface{}) () {
  // childKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings9childKeysEv
  default:
    qtrt.ErrorResolve("QSettings", "childKeys", args)
  }

}

  // proto:  void QSettings::endArray();
func (this *QSettings) endArray(args ...interface{}) () {
  // endArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endArrayEv
  default:
    qtrt.ErrorResolve("QSettings", "endArray", args)
  }

}

  // proto: static void QSettings::setSystemIniPath(const QString & dir);
func (this *QSettings) setSystemIniPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "setSystemIniPath", args)
  }

}

  // proto:  QStringList QSettings::allKeys();
func (this *QSettings) allKeys(args ...interface{}) () {
  // allKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings7allKeysEv
  default:
    qtrt.ErrorResolve("QSettings", "allKeys", args)
  }

}

  // proto:  QString QSettings::group();
func (this *QSettings) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5groupEv
  default:
    qtrt.ErrorResolve("QSettings", "group", args)
  }

}

// <= body block end

