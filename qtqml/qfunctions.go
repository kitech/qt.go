package qtqml

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

func init_unused_10017() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtQml/qjsengine.h:137
// index:89
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QJSEngine::Extensions::enum_type, int)

/*

 */
func Operator_or89(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN9QJSEngine9ExtensionEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:80
// index:90
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QQmlImageProviderBase::Flags::enum_type, int)

/*

 */
func Operator_or90(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN21QQmlImageProviderBase4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQml/qqml.h:581
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QQmlEngine * qmlEngine(const QObject *)

/*

 */
func QmlEngine(arg0 qtcore.QObject_ITF /*777 const QObject **/) *QQmlEngine /*777 QQmlEngine **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml9qmlEngineEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlinfo.h:54
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlDebug(const QObject *, const QQmlError &)

/*

 */
func QmlDebug(me qtcore.QObject_ITF /*777 const QObject **/, error QQmlError_ITF) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if error != nil && error.QQmlError_PTR() != nil {
		convArg1 = error.QQmlError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml8qmlDebugEPK7QObjectRK9QQmlError", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:53
// index:1
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlDebug(const QObject *)

/*

 */
func QmlDebug1(me qtcore.QObject_ITF /*777 const QObject **/) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml8qmlDebugEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:58
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlInfo(const QObject *, const QQmlError &)

/*

 */
func QmlInfo(me qtcore.QObject_ITF /*777 const QObject **/, error QQmlError_ITF) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if error != nil && error.QQmlError_PTR() != nil {
		convArg1 = error.QQmlError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml7qmlInfoEPK7QObjectRK9QQmlError", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:57
// index:1
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlInfo(const QObject *)

/*

 */
func QmlInfo1(me qtcore.QObject_ITF /*777 const QObject **/) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml7qmlInfoEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqml.h:582
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QObject * qmlAttachedPropertiesObjectById(int, const QObject *, bool)

/*

 */
func QmlAttachedPropertiesObjectById(arg0 int, arg1 qtcore.QObject_ITF /*777 const QObject **/, create bool) *qtcore.QObject /*777 QObject **/ {
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QObject_PTR() != nil {
		convArg1 = arg1.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml31qmlAttachedPropertiesObjectByIdEiPK7QObjectb", qtrt.FFI_TYPE_POINTER, arg0, convArg1, create)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqml.h:583
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QObject * qmlAttachedPropertiesObject(int *, const QObject *, const QMetaObject *, bool)

/*

 */
func QmlAttachedPropertiesObject(arg0 unsafe.Pointer /*666*/, arg1 qtcore.QObject_ITF /*777 const QObject **/, arg2 qtcore.QMetaObject_ITF /*777 const QMetaObject **/, create bool) *qtcore.QObject /*777 QObject **/ {
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QObject_PTR() != nil {
		convArg1 = arg1.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QMetaObject_PTR() != nil {
		convArg2 = arg2.QMetaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml27qmlAttachedPropertiesObjectEPiPK7QObjectPK11QMetaObjectb", qtrt.FFI_TYPE_POINTER, arg0, convArg1, convArg2, create)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqml.h:579
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qmlExecuteDeferred(QObject *)

/*

 */
func QmlExecuteDeferred(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml18qmlExecuteDeferredEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlinfo.h:62
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlWarning(const QObject *, const QQmlError &)

/*

 */
func QmlWarning(me qtcore.QObject_ITF /*777 const QObject **/, error QQmlError_ITF) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if error != nil && error.QQmlError_PTR() != nil {
		convArg1 = error.QQmlError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml10qmlWarningEPK7QObjectRK9QQmlError", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:61
// index:1
// Invalid Visibility=Default Availability=Available
// [16] QQmlInfo qmlWarning(const QObject *)

/*

 */
func QmlWarning1(me qtcore.QObject_ITF /*777 const QObject **/) *QQmlInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if me != nil && me.QObject_PTR() != nil {
		convArg0 = me.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml10qmlWarningEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqml.h:580
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QQmlContext * qmlContext(const QObject *)

/*

 */
func QmlContext(arg0 qtcore.QObject_ITF /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtQml10qmlContextEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlprivate.h:318
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qmlregister(QQmlPrivate::RegistrationType, void *)

/*

 */
func Qmlregister(arg0 int, arg1 unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlPrivate11qmlregisterENS_16RegistrationTypeEPv", qtrt.FFI_TYPE_POINTER, arg0, arg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqml.h:680
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qmlTypeId(const char *, int, int, const char *)

/*

 */
func QmlTypeId(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	var convArg3 = qtrt.CString(qmlName)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_Z9qmlTypeIdPKciiS0_", qtrt.FFI_TYPE_POINTER, convArg0, versionMajor, versionMinor, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qjsengine.h:164
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QJSEngine * qjsEngine(const QObject *)

/*

 */
func QjsEngine(arg0 qtcore.QObject_ITF /*777 const QObject **/) *QJSEngine /*777 QJSEngine **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z9qjsEnginePK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlproperty.h:130
// index:52
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QQmlProperty &)

/*

 */
func QHash52(key QQmlProperty_ITF) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QQmlProperty_PTR() != nil {
		convArg0 = key.QQmlProperty_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK12QQmlProperty", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtQml/qqml.h:280
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qmlRegisterUncreatableMetaObject(const QMetaObject &, const char *, int, int, const char *, const QString &)

/*

 */
func QmlRegisterUncreatableMetaObject(staticMetaObject qtcore.QMetaObject_ITF, uri string, versionMajor int, versionMinor int, qmlName string, reason string) int {
	var convArg0 unsafe.Pointer
	if staticMetaObject != nil && staticMetaObject.QMetaObject_PTR() != nil {
		convArg0 = staticMetaObject.QMetaObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	var convArg4 = qtrt.CString(qmlName)
	defer qtrt.FreeMem(convArg4)
	var tmpArg5 = qtcore.NewQString5(reason)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z32qmlRegisterUncreatableMetaObjectRK11QMetaObjectPKciiS3_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, versionMajor, versionMinor, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqml.h:136
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qmlRegisterTypeNotAvailable(const char *, int, int, const char *, const QString &)

/*

 */
func QmlRegisterTypeNotAvailable(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	var convArg3 = qtrt.CString(qmlName)
	defer qtrt.FreeMem(convArg3)
	var tmpArg4 = qtcore.NewQString5(message)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z27qmlRegisterTypeNotAvailablePKciiS0_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, versionMajor, versionMinor, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqml.h:102
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qmlClearTypeRegistrations()

/*

 */
func QmlClearTypeRegistrations() {
	rv, err := qtrt.InvokeQtFunc6("_Z25qmlClearTypeRegistrationsv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqml.h:642
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qmlRegisterSingletonType(const QUrl &, const char *, int, int, const char *)

/*

 */
func QmlRegisterSingletonType(url qtcore.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	var convArg4 = qtrt.CString(qmlName)
	defer qtrt.FreeMem(convArg4)
	rv, err := qtrt.InvokeQtFunc6("_Z24qmlRegisterSingletonTypeRK4QUrlPKciiS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, versionMajor, versionMinor, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqml.h:610
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qmlRegisterSingletonType(const char *, int, int, const char *, QJSValue (*)(QQmlEngine *, QJSEngine *))

/*

 */
func QmlRegisterSingletonType1(uri string, versionMajor int, versionMinor int, typeName string, callback unsafe.Pointer /*666*/) int {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	var convArg3 = qtrt.CString(typeName)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_Z24qmlRegisterSingletonTypePKciiS0_PF8QJSValueP10QQmlEngineP9QJSEngineE", qtrt.FFI_TYPE_POINTER, convArg0, versionMajor, versionMinor, convArg3, callback)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqml.h:608
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qmlRegisterBaseTypes(const char *, int, int)

/*

 */
func QmlRegisterBaseTypes(uri string, versionMajor int, versionMinor int) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z20qmlRegisterBaseTypesPKcii", qtrt.FFI_TYPE_POINTER, convArg0, versionMajor, versionMinor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqml.h:599
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qmlRegisterModule(const char *, int, int)

/*

 */
func QmlRegisterModule(uri string, versionMajor int, versionMinor int) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z17qmlRegisterModulePKcii", qtrt.FFI_TYPE_POINTER, convArg0, versionMajor, versionMinor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqml.h:598
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qmlProtectModule(const char *, int)

/*

 */
func QmlProtectModule(uri string, majVersion int) bool {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z16qmlProtectModulePKci", qtrt.FFI_TYPE_POINTER, convArg0, majVersion)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqml.h:661
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qmlRegisterType(const QUrl &, const char *, int, int, const char *)

/*

 */
func QmlRegisterType(url qtcore.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	var convArg4 = qtrt.CString(qmlName)
	defer qtrt.FreeMem(convArg4)
	rv, err := qtrt.InvokeQtFunc6("_Z15qmlRegisterTypeRK4QUrlPKciiS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, versionMajor, versionMinor, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
