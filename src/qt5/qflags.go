package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qflags.h
// dst-file: /src/core/qflags.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QIncompatibleFlag)=4
type QIncompatibleFlag struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFlag)=4
type QFlag struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQIncompatibleFlag(args ...interface{})() {
}


func NewQFlag(args ...interface{})() {
}

// <= body block end

