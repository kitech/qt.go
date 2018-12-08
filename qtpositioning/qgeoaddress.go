package qtpositioning

// /usr/include/qt/QtPositioning/qgeoaddress.h
// #include <qgeoaddress.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGeoAddress struct {
	*qtrt.CObject
}
type QGeoAddress_ITF interface {
	QGeoAddress_PTR() *QGeoAddress
}

func (ptr *QGeoAddress) QGeoAddress_PTR() *QGeoAddress { return ptr }

func (this *QGeoAddress) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoAddress) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoAddressFromPointer(cthis unsafe.Pointer) *QGeoAddress {
	return &QGeoAddress{&qtrt.CObject{cthis}}
}
func (*QGeoAddress) NewFromPointer(cthis unsafe.Pointer) *QGeoAddress {
	return NewQGeoAddressFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoAddress()

/*
Default constructor.
*/
func (*QGeoAddress) NewForInherit() *QGeoAddress {
	return NewQGeoAddress()
}
func NewQGeoAddress() *QGeoAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddressC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoAddress)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoAddress()

/*

 */
func DeleteQGeoAddress(this *QGeoAddress) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddressD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:58
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoAddress & operator=(const QGeoAddress &)

/*

 */
func (this *QGeoAddress) Operator_equal(other QGeoAddress_ITF) *QGeoAddress {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAddress_PTR() != nil {
		convArg0 = other.QGeoAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddressaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoAddress)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoAddress &) const

/*

 */
func (this *QGeoAddress) Operator_equal_equal(other QGeoAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAddress_PTR() != nil {
		convArg0 = other.QGeoAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddresseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoAddress &) const

/*

 */
func (this *QGeoAddress) Operator_not_equal(other QGeoAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAddress_PTR() != nil {
		convArg0 = other.QGeoAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddressneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*
Returns the address as a single formatted string. It is the recommended string to use to display the address to the user. It typically takes the format of an address as found on an envelope, but this is not always necessarily the case.

The address text is either automatically generated or explicitly assigned. This can be determined by checking isTextGenerated.

If an empty string is provided to setText(), then isTextGenerated() will be set to true and text() will return a string which is locally formatted according to countryCode() and based on the elements of the address such as street, city and so on. Because the text string is generated from the address elements, a sequence of calls such as text(), setStreet(), text() may return different strings for each invocation of text().

If a non-empty string is provided to setText(), then isTextGenerated() will be set to false and text() will always return the explicitly assigned string. Calls to modify other elements such as setStreet(), setCity() and so on will not affect the resultant string from text().

See also setText().
*/
func (this *QGeoAddress) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
If text is not empty, explicitly assigns text as the string to be returned by text(). isTextGenerated() will return false.

If text is empty, indicates that text() should be automatically generated from the address elements. isTextGenerated() will return true.

See also text().
*/
func (this *QGeoAddress) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString country() const

/*
Returns the country name.

See also setCountry().
*/
func (this *QGeoAddress) Country() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress7countryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCountry(const QString &)

/*
Sets the country name.

See also country().
*/
func (this *QGeoAddress) SetCountry(country string) {
	var tmpArg0 = qtcore.NewQString5(country)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress10setCountryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString countryCode() const

/*
Returns the country code according to ISO 3166-1 alpha-3

See also setCountryCode().
*/
func (this *QGeoAddress) CountryCode() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress11countryCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCountryCode(const QString &)

/*
Sets the countryCode according to ISO 3166-1 alpha-3

See also countryCode().
*/
func (this *QGeoAddress) SetCountryCode(countryCode string) {
	var tmpArg0 = qtcore.NewQString5(countryCode)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress14setCountryCodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QString state() const

/*
Returns the state. The state is considered the first subdivision below country.

See also setState().
*/
func (this *QGeoAddress) State() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setState(const QString &)

/*
Sets the state.

See also state().
*/
func (this *QGeoAddress) SetState(state string) {
	var tmpArg0 = qtcore.NewQString5(state)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress8setStateERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString county() const

/*
Returns the county. The county is considered the second subdivision below country.

See also setCounty().
*/
func (this *QGeoAddress) County() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress6countyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCounty(const QString &)

/*
Sets the county.

See also county().
*/
func (this *QGeoAddress) SetCounty(county string) {
	var tmpArg0 = qtcore.NewQString5(county)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress9setCountyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString city() const

/*
Returns the city.

See also setCity().
*/
func (this *QGeoAddress) City() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress4cityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCity(const QString &)

/*
Sets the city.

See also city().
*/
func (this *QGeoAddress) SetCity(city string) {
	var tmpArg0 = qtcore.NewQString5(city)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress7setCityERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString district() const

/*
Returns the district. The district is considered the subdivison below city.

See also setDistrict().
*/
func (this *QGeoAddress) District() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress8districtEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDistrict(const QString &)

/*
Sets the district.

See also district().
*/
func (this *QGeoAddress) SetDistrict(district string) {
	var tmpArg0 = qtcore.NewQString5(district)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress11setDistrictERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString postalCode() const

/*
Returns the postal code.

See also setPostalCode().
*/
func (this *QGeoAddress) PostalCode() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress10postalCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPostalCode(const QString &)

/*
Sets the postalCode.

See also postalCode().
*/
func (this *QGeoAddress) SetPostalCode(postalCode string) {
	var tmpArg0 = qtcore.NewQString5(postalCode)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress13setPostalCodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString street() const

/*
Returns the street-level component of the address.

This typically includes a street number and street name but may also contain things like a unit number, a building name, or anything else that might be used to distinguish one address from another.

See also setStreet().
*/
func (this *QGeoAddress) Street() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress6streetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStreet(const QString &)

/*
Sets the street-level component of the address to street.

This typically includes a street number and street name but may also contain things like a unit number, a building name, or anything else that might be used to distinguish one address from another.

See also street().
*/
func (this *QGeoAddress) SetStreet(street string) {
	var tmpArg0 = qtcore.NewQString5(street)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress9setStreetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns whether this address is empty. An address is considered empty if all of its fields are empty.
*/
func (this *QGeoAddress) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears all of the address' data fields.
*/
func (this *QGeoAddress) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoAddress5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoaddress.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTextGenerated() const

/*
Returns true if QGeoAddress::text() is automatically generated from address elements, otherwise returns false if text() has been explicitly assigned.

See also text() and setText().
*/
func (this *QGeoAddress) IsTextGenerated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoAddress15isTextGeneratedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
