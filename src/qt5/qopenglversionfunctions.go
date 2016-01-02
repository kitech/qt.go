package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qopenglversionfunctions.h
// dst-file: /src/gui/qopenglversionfunctions.go
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
  // proto: static void QAbstractOpenGLFunctionsPrivate::insertFunctionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v, QOpenGLVersionFunctionsBackend * backend);
extern void _ZN31QAbstractOpenGLFunctionsPrivate22insertFunctionsBackendEP14QOpenGLContextRK20QOpenGLVersionStatusP30QOpenGLVersionFunctionsBackend(void* arg0, void* arg1, void* arg2);
  // proto: static void QAbstractOpenGLFunctionsPrivate::insertExternalFunctions(QOpenGLContext * context, QAbstractOpenGLFunctions * f);
extern void _ZN31QAbstractOpenGLFunctionsPrivate23insertExternalFunctionsEP14QOpenGLContextP24QAbstractOpenGLFunctions(void* arg0, void* arg1);
  // proto: static QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctionsPrivate::get(QAbstractOpenGLFunctions * q);
extern void _ZN31QAbstractOpenGLFunctionsPrivate3getEP24QAbstractOpenGLFunctions(void* arg0);
  // proto: static QOpenGLVersionFunctionsBackend * QAbstractOpenGLFunctionsPrivate::functionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v);
extern void _ZN31QAbstractOpenGLFunctionsPrivate16functionsBackendEP14QOpenGLContextRK20QOpenGLVersionStatus(void* arg0, void* arg1);
  // proto: static void QAbstractOpenGLFunctionsPrivate::removeFunctionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v);
extern void _ZN31QAbstractOpenGLFunctionsPrivate22removeFunctionsBackendEP14QOpenGLContextRK20QOpenGLVersionStatus(void* arg0, void* arg1);
  // proto: static void QAbstractOpenGLFunctionsPrivate::removeExternalFunctions(QOpenGLContext * context, QAbstractOpenGLFunctions * f);
extern void _ZN31QAbstractOpenGLFunctionsPrivate23removeExternalFunctionsEP14QOpenGLContextP24QAbstractOpenGLFunctions(void* arg0, void* arg1);
  // proto:  void QAbstractOpenGLFunctionsPrivate::QAbstractOpenGLFunctionsPrivate();
extern void* dector_ZN31QAbstractOpenGLFunctionsPrivateC1Ev();
extern void _ZN31QAbstractOpenGLFunctionsPrivateC1Ev(void* qthis);
  // proto:  void QOpenGLFunctions_4_5_DeprecatedBackend::QOpenGLFunctions_4_5_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_4_5_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_4_5_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_5_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_4_5_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_2_DeprecatedBackend::QOpenGLFunctions_1_2_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_1_2_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_1_2_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_2_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_1_2_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_4_1_CoreBackend::QOpenGLFunctions_4_1_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_1_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_1_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_1_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_1_CoreBackend13versionStatusEv();
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_3_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_3_3_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_3_3_CoreBackend::QOpenGLFunctions_3_3_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_3_3_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_3_3_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_5_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_5_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_5_CoreBackend::QOpenGLFunctions_1_5_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_5_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_5_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_5_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_5_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_4_5_CoreBackend::QOpenGLFunctions_4_5_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_5_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_5_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions_4_4_CoreBackend::QOpenGLFunctions_4_4_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_4_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_4_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_4_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_4_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_4_3_CoreBackend::QOpenGLFunctions_4_3_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_3_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_3_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_3_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_3_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_3_0_DeprecatedBackend::QOpenGLFunctions_3_0_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_3_0_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_3_0_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_0_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_3_0_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_2_1_CoreBackend::QOpenGLFunctions_2_1_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_2_1_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_2_1_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_1_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_2_1_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_0_DeprecatedBackend::QOpenGLFunctions_1_0_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_1_0_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_1_0_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_0_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_1_0_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_3_0_CoreBackend::QOpenGLFunctions_3_0_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_3_0_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_3_0_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_0_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_3_0_CoreBackend13versionStatusEv();
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_2_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_2_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_2_CoreBackend::QOpenGLFunctions_1_2_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_2_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_2_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_1_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_1_1_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_1_DeprecatedBackend::QOpenGLFunctions_1_1_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_1_1_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_1_1_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_2_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_2_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_4_2_CoreBackend::QOpenGLFunctions_4_2_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_2_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_2_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_0_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_2_0_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_2_0_CoreBackend::QOpenGLFunctions_2_0_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_2_0_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_2_0_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions_3_2_CoreBackend::QOpenGLFunctions_3_2_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_3_2_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_3_2_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_2_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_3_2_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLVersionFunctionsBackend::QOpenGLVersionFunctionsBackend(QOpenGLContext * ctx);
extern void* dector_ZN30QOpenGLVersionFunctionsBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN30QOpenGLVersionFunctionsBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  bool QAbstractOpenGLFunctions::initializeOpenGLFunctions();
extern void _ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv(void* qthis);
  // proto:  void QAbstractOpenGLFunctions::QAbstractOpenGLFunctions();
extern void* dector_ZN24QAbstractOpenGLFunctionsC1Ev();
extern void _ZN24QAbstractOpenGLFunctionsC1Ev(void* qthis);
  // proto:  QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctions::d_func();
extern void _ZN24QAbstractOpenGLFunctions6d_funcEv(void* qthis);
  // proto:  void QAbstractOpenGLFunctions::~QAbstractOpenGLFunctions();
extern void _ZN24QAbstractOpenGLFunctionsD0Ev(void* qthis);
  // proto:  void QOpenGLFunctions_2_0_DeprecatedBackend::QOpenGLFunctions_2_0_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_2_0_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_2_0_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_0_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_2_0_DeprecatedBackend13versionStatusEv();
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_3_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_1_3_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_3_DeprecatedBackend::QOpenGLFunctions_1_3_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_1_3_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_1_3_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions_1_4_DeprecatedBackend::QOpenGLFunctions_1_4_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_1_4_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_1_4_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_4_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_1_4_DeprecatedBackend13versionStatusEv();
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_3_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_3_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_3_CoreBackend::QOpenGLFunctions_1_3_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_3_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_3_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLVersionStatus::QOpenGLVersionStatus();
extern void* dector_ZN20QOpenGLVersionStatusC1Ev();
extern void _ZN20QOpenGLVersionStatusC1Ev(void* qthis);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_0_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_0_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_0_CoreBackend::QOpenGLFunctions_1_0_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_0_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_0_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_1_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_3_1_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_3_1_CoreBackend::QOpenGLFunctions_3_1_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_3_1_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_3_1_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions_1_1_CoreBackend::QOpenGLFunctions_1_1_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_1_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_1_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_1_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_1_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_1_4_CoreBackend::QOpenGLFunctions_1_4_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_1_4_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_1_4_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_4_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_1_4_CoreBackend13versionStatusEv();
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_0_CoreBackend::versionStatus();
extern void _ZN32QOpenGLFunctions_4_0_CoreBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_4_0_CoreBackend::QOpenGLFunctions_4_0_CoreBackend(QOpenGLContext * context);
extern void* dector_ZN32QOpenGLFunctions_4_0_CoreBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN32QOpenGLFunctions_4_0_CoreBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_3_DeprecatedBackend::versionStatus();
extern void _ZN38QOpenGLFunctions_3_3_DeprecatedBackend13versionStatusEv();
  // proto:  void QOpenGLFunctions_3_3_DeprecatedBackend::QOpenGLFunctions_3_3_DeprecatedBackend(QOpenGLContext * context);
extern void* dector_ZN38QOpenGLFunctions_3_3_DeprecatedBackendC1EP14QOpenGLContext(void* arg0);
extern void _ZN38QOpenGLFunctions_3_3_DeprecatedBackendC1EP14QOpenGLContext(void* qthis, void* arg0);
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

// class sizeof(QAbstractOpenGLFunctionsPrivate)=16
type QAbstractOpenGLFunctionsPrivate struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_5_DeprecatedBackend)=1
type QOpenGLFunctions_4_5_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_2_DeprecatedBackend)=1
type QOpenGLFunctions_1_2_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_1_CoreBackend)=1
type QOpenGLFunctions_4_1_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_3_CoreBackend)=1
type QOpenGLFunctions_3_3_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_5_CoreBackend)=1
type QOpenGLFunctions_1_5_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_5_CoreBackend)=1
type QOpenGLFunctions_4_5_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_4_CoreBackend)=1
type QOpenGLFunctions_4_4_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_3_CoreBackend)=1
type QOpenGLFunctions_4_3_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_0_DeprecatedBackend)=1
type QOpenGLFunctions_3_0_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_2_1_CoreBackend)=1
type QOpenGLFunctions_2_1_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_0_DeprecatedBackend)=1
type QOpenGLFunctions_1_0_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_0_CoreBackend)=1
type QOpenGLFunctions_3_0_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_2_CoreBackend)=1
type QOpenGLFunctions_1_2_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_1_DeprecatedBackend)=1
type QOpenGLFunctions_1_1_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_2_CoreBackend)=1
type QOpenGLFunctions_4_2_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_2_0_CoreBackend)=1
type QOpenGLFunctions_2_0_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_2_CoreBackend)=1
type QOpenGLFunctions_3_2_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLVersionFunctionsBackend)=1
type QOpenGLVersionFunctionsBackend struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractOpenGLFunctions)=16
type QAbstractOpenGLFunctions struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_2_0_DeprecatedBackend)=1
type QOpenGLFunctions_2_0_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_3_DeprecatedBackend)=1
type QOpenGLFunctions_1_3_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_4_DeprecatedBackend)=1
type QOpenGLFunctions_1_4_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_3_CoreBackend)=1
type QOpenGLFunctions_1_3_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLVersionStatus)=1
type QOpenGLVersionStatus struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_0_CoreBackend)=1
type QOpenGLFunctions_1_0_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_1_CoreBackend)=1
type QOpenGLFunctions_3_1_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_1_CoreBackend)=1
type QOpenGLFunctions_1_1_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_1_4_CoreBackend)=1
type QOpenGLFunctions_1_4_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_4_0_CoreBackend)=1
type QOpenGLFunctions_4_0_CoreBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions_3_3_DeprecatedBackend)=1
type QOpenGLFunctions_3_3_DeprecatedBackend struct {
  /*qbase*/ QOpenGLVersionFunctionsBackend;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static void QAbstractOpenGLFunctionsPrivate::insertFunctionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v, QOpenGLVersionFunctionsBackend * backend);
func (this *QAbstractOpenGLFunctionsPrivate) insertFunctionsBackend_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "insertFunctionsBackend", args)
  }

}

  // proto: static void QAbstractOpenGLFunctionsPrivate::insertExternalFunctions(QOpenGLContext * context, QAbstractOpenGLFunctions * f);
func (this *QAbstractOpenGLFunctionsPrivate) insertExternalFunctions_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "insertExternalFunctions", args)
  }

}

  // proto: static QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctionsPrivate::get(QAbstractOpenGLFunctions * q);
func (this *QAbstractOpenGLFunctionsPrivate) get_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "get", args)
  }

}

  // proto: static QOpenGLVersionFunctionsBackend * QAbstractOpenGLFunctionsPrivate::functionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v);
func (this *QAbstractOpenGLFunctionsPrivate) functionsBackend_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "functionsBackend", args)
  }

}

  // proto: static void QAbstractOpenGLFunctionsPrivate::removeFunctionsBackend(QOpenGLContext * context, const QOpenGLVersionStatus & v);
func (this *QAbstractOpenGLFunctionsPrivate) removeFunctionsBackend_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "removeFunctionsBackend", args)
  }

}

  // proto: static void QAbstractOpenGLFunctionsPrivate::removeExternalFunctions(QOpenGLContext * context, QAbstractOpenGLFunctions * f);
func (this *QAbstractOpenGLFunctionsPrivate) removeExternalFunctions_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctionsPrivate", "removeExternalFunctions", args)
  }

}

  // proto:  void QAbstractOpenGLFunctionsPrivate::QAbstractOpenGLFunctionsPrivate();
func NewQAbstractOpenGLFunctionsPrivate(args ...interface{}) QAbstractOpenGLFunctionsPrivate {
  return QAbstractOpenGLFunctionsPrivate{}
}

  // proto:  void QOpenGLFunctions_4_5_DeprecatedBackend::QOpenGLFunctions_4_5_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_5_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_4_5_DeprecatedBackend {
  return QOpenGLFunctions_4_5_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_5_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_4_5_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_5_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_2_DeprecatedBackend::QOpenGLFunctions_1_2_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_2_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_1_2_DeprecatedBackend {
  return QOpenGLFunctions_1_2_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_2_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_1_2_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_2_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_4_1_CoreBackend::QOpenGLFunctions_4_1_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_1_CoreBackend(args ...interface{}) QOpenGLFunctions_4_1_CoreBackend {
  return QOpenGLFunctions_4_1_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_1_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_1_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_1_CoreBackend", "versionStatus", args)
  }

}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_3_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_3_3_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_3_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_3_3_CoreBackend::QOpenGLFunctions_3_3_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_3_CoreBackend(args ...interface{}) QOpenGLFunctions_3_3_CoreBackend {
  return QOpenGLFunctions_3_3_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_5_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_5_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_5_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_5_CoreBackend::QOpenGLFunctions_1_5_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_5_CoreBackend(args ...interface{}) QOpenGLFunctions_1_5_CoreBackend {
  return QOpenGLFunctions_1_5_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_5_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_5_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_5_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_4_5_CoreBackend::QOpenGLFunctions_4_5_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_5_CoreBackend(args ...interface{}) QOpenGLFunctions_4_5_CoreBackend {
  return QOpenGLFunctions_4_5_CoreBackend{}
}

  // proto:  void QOpenGLFunctions_4_4_CoreBackend::QOpenGLFunctions_4_4_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_4_CoreBackend(args ...interface{}) QOpenGLFunctions_4_4_CoreBackend {
  return QOpenGLFunctions_4_4_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_4_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_4_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_4_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_4_3_CoreBackend::QOpenGLFunctions_4_3_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_3_CoreBackend(args ...interface{}) QOpenGLFunctions_4_3_CoreBackend {
  return QOpenGLFunctions_4_3_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_3_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_3_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_3_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_3_0_DeprecatedBackend::QOpenGLFunctions_3_0_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_0_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_3_0_DeprecatedBackend {
  return QOpenGLFunctions_3_0_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_0_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_3_0_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_0_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_2_1_CoreBackend::QOpenGLFunctions_2_1_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_2_1_CoreBackend(args ...interface{}) QOpenGLFunctions_2_1_CoreBackend {
  return QOpenGLFunctions_2_1_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_1_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_2_1_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_2_1_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_0_DeprecatedBackend::QOpenGLFunctions_1_0_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_0_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_1_0_DeprecatedBackend {
  return QOpenGLFunctions_1_0_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_0_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_1_0_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_0_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_3_0_CoreBackend::QOpenGLFunctions_3_0_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_0_CoreBackend(args ...interface{}) QOpenGLFunctions_3_0_CoreBackend {
  return QOpenGLFunctions_3_0_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_0_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_3_0_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_0_CoreBackend", "versionStatus", args)
  }

}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_2_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_2_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_2_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_2_CoreBackend::QOpenGLFunctions_1_2_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_2_CoreBackend(args ...interface{}) QOpenGLFunctions_1_2_CoreBackend {
  return QOpenGLFunctions_1_2_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_1_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_1_1_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_1_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_1_DeprecatedBackend::QOpenGLFunctions_1_1_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_1_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_1_1_DeprecatedBackend {
  return QOpenGLFunctions_1_1_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_2_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_2_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_2_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_4_2_CoreBackend::QOpenGLFunctions_4_2_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_2_CoreBackend(args ...interface{}) QOpenGLFunctions_4_2_CoreBackend {
  return QOpenGLFunctions_4_2_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_0_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_2_0_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_2_0_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_2_0_CoreBackend::QOpenGLFunctions_2_0_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_2_0_CoreBackend(args ...interface{}) QOpenGLFunctions_2_0_CoreBackend {
  return QOpenGLFunctions_2_0_CoreBackend{}
}

  // proto:  void QOpenGLFunctions_3_2_CoreBackend::QOpenGLFunctions_3_2_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_2_CoreBackend(args ...interface{}) QOpenGLFunctions_3_2_CoreBackend {
  return QOpenGLFunctions_3_2_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_2_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_3_2_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_2_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLVersionFunctionsBackend::QOpenGLVersionFunctionsBackend(QOpenGLContext * ctx);
func NewQOpenGLVersionFunctionsBackend(args ...interface{}) QOpenGLVersionFunctionsBackend {
  return QOpenGLVersionFunctionsBackend{}
}

  // proto:  bool QAbstractOpenGLFunctions::initializeOpenGLFunctions();
func (this *QAbstractOpenGLFunctions) initializeOpenGLFunctions(args ...interface{}) () {
  // initializeOpenGLFunctions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "initializeOpenGLFunctions", args)
  }

}

  // proto:  void QAbstractOpenGLFunctions::QAbstractOpenGLFunctions();
func NewQAbstractOpenGLFunctions(args ...interface{}) QAbstractOpenGLFunctions {
  return QAbstractOpenGLFunctions{}
}

  // proto:  QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctions::d_func();
func (this *QAbstractOpenGLFunctions) d_func(args ...interface{}) () {
  // d_func()
  // d_func()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions6d_funcEv
  case 1:
    // invoke: _ZNK24QAbstractOpenGLFunctions6d_funcEv
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "d_func", args)
  }

}

  // proto:  void QAbstractOpenGLFunctions::~QAbstractOpenGLFunctions();
func (this *QAbstractOpenGLFunctions) FreeQAbstractOpenGLFunctions(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "~QAbstractOpenGLFunctions", args)
  }

}

  // proto:  void QOpenGLFunctions_2_0_DeprecatedBackend::QOpenGLFunctions_2_0_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_2_0_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_2_0_DeprecatedBackend {
  return QOpenGLFunctions_2_0_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_2_0_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_2_0_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_2_0_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_3_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_1_3_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_3_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_3_DeprecatedBackend::QOpenGLFunctions_1_3_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_3_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_1_3_DeprecatedBackend {
  return QOpenGLFunctions_1_3_DeprecatedBackend{}
}

  // proto:  void QOpenGLFunctions_1_4_DeprecatedBackend::QOpenGLFunctions_1_4_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_4_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_1_4_DeprecatedBackend {
  return QOpenGLFunctions_1_4_DeprecatedBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_4_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_1_4_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_4_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_3_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_3_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_3_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_3_CoreBackend::QOpenGLFunctions_1_3_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_3_CoreBackend(args ...interface{}) QOpenGLFunctions_1_3_CoreBackend {
  return QOpenGLFunctions_1_3_CoreBackend{}
}

  // proto:  void QOpenGLVersionStatus::QOpenGLVersionStatus();
func NewQOpenGLVersionStatus(args ...interface{}) QOpenGLVersionStatus {
  return QOpenGLVersionStatus{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_0_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_0_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_0_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_0_CoreBackend::QOpenGLFunctions_1_0_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_0_CoreBackend(args ...interface{}) QOpenGLFunctions_1_0_CoreBackend {
  return QOpenGLFunctions_1_0_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_1_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_3_1_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_1_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_3_1_CoreBackend::QOpenGLFunctions_3_1_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_1_CoreBackend(args ...interface{}) QOpenGLFunctions_3_1_CoreBackend {
  return QOpenGLFunctions_3_1_CoreBackend{}
}

  // proto:  void QOpenGLFunctions_1_1_CoreBackend::QOpenGLFunctions_1_1_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_1_CoreBackend(args ...interface{}) QOpenGLFunctions_1_1_CoreBackend {
  return QOpenGLFunctions_1_1_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_1_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_1_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_1_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_1_4_CoreBackend::QOpenGLFunctions_1_4_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_1_4_CoreBackend(args ...interface{}) QOpenGLFunctions_1_4_CoreBackend {
  return QOpenGLFunctions_1_4_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_1_4_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_1_4_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_1_4_CoreBackend", "versionStatus", args)
  }

}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_4_0_CoreBackend::versionStatus();
func (this *QOpenGLFunctions_4_0_CoreBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_4_0_CoreBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_4_0_CoreBackend::QOpenGLFunctions_4_0_CoreBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_4_0_CoreBackend(args ...interface{}) QOpenGLFunctions_4_0_CoreBackend {
  return QOpenGLFunctions_4_0_CoreBackend{}
}

  // proto: static QOpenGLVersionStatus QOpenGLFunctions_3_3_DeprecatedBackend::versionStatus();
func (this *QOpenGLFunctions_3_3_DeprecatedBackend) versionStatus_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions_3_3_DeprecatedBackend", "versionStatus", args)
  }

}

  // proto:  void QOpenGLFunctions_3_3_DeprecatedBackend::QOpenGLFunctions_3_3_DeprecatedBackend(QOpenGLContext * context);
func NewQOpenGLFunctions_3_3_DeprecatedBackend(args ...interface{}) QOpenGLFunctions_3_3_DeprecatedBackend {
  return QOpenGLFunctions_3_3_DeprecatedBackend{}
}

// <= body block end

