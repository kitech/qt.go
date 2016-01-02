package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qfileiconprovider.h
// dst-file: /src/widgets/qfileiconprovider.go
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
  // proto:  QString QFileIconProvider::type(const QFileInfo & info);
extern void _ZNK17QFileIconProvider4typeERK9QFileInfo(void* qthis, void* arg0);
  // proto:  QIcon QFileIconProvider::icon(const QFileInfo & info);
extern void _ZNK17QFileIconProvider4iconERK9QFileInfo(void* qthis, void* arg0);
  // proto:  void QFileIconProvider::QFileIconProvider(const QFileIconProvider & );
extern void* dector_ZN17QFileIconProviderC1ERKS_(void* arg0);
extern void _ZN17QFileIconProviderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFileIconProvider::QFileIconProvider();
extern void* dector_ZN17QFileIconProviderC1Ev();
extern void _ZN17QFileIconProviderC1Ev(void* qthis);
  // proto:  void QFileIconProvider::~QFileIconProvider();
extern void _ZN17QFileIconProviderD0Ev(void* qthis);
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

// class sizeof(QFileIconProvider)=1
type QFileIconProvider struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QString QFileIconProvider::type(const QFileInfo & info);
func (this *QFileIconProvider) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileIconProvider", "type", args)
  }

}

  // proto:  QIcon QFileIconProvider::icon(const QFileInfo & info);
func (this *QFileIconProvider) icon(args ...interface{}) () {
  // icon(const class QFileInfo &)
  // icon(enum QFileIconProvider::IconType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "const QFileInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QFileIconProvider::IconType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4iconERK9QFileInfo
    // invoke: QIcon icon(const class QFileInfo &)
    var arg0 = args[0].(QFileInfo).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QFileIconProvider4iconERK9QFileInfo(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "icon", args)
  }

}

  // proto:  void QFileIconProvider::QFileIconProvider(const QFileIconProvider & );
func NewQFileIconProvider(args ...interface{}) QFileIconProvider {
  return QFileIconProvider{}
}

  // proto:  void QFileIconProvider::~QFileIconProvider();
func (this *QFileIconProvider) FreeQFileIconProvider(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileIconProvider", "~QFileIconProvider", args)
  }

}

// <= body block end

