package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QWidgetSet struct {
	*qtrt.CObject
}

// void swap(QSet<T> &)
func (this *QWidgetSet) Swap_0() {
	// QWidgetSet_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// bool operator==(const QSet<T> &)
func (this *QWidgetSet) Operator_equal_equal_0() bool {
	// QWidgetSet_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QSet<T> &)
func (this *QWidgetSet) Operator_not_equal_0() bool {
	// QWidgetSet_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QWidgetSet) Size_0() int {
	// QWidgetSet_size_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0
}

// bool isEmpty()
func (this *QWidgetSet) IsEmpty_0() bool {
	// QWidgetSet_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// int capacity()
func (this *QWidgetSet) Capacity_0() int {
	// QWidgetSet_capacity_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_capacity_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0
}

// void reserve(int)
func (this *QWidgetSet) Reserve_0() {
	// QWidgetSet_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// void squeeze()
func (this *QWidgetSet) Squeeze_0() {
	// QWidgetSet_squeeze_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_squeeze_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// void detach()
func (this *QWidgetSet) Detach_0() {
	// QWidgetSet_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QWidgetSet) IsDetached_0() bool {
	// QWidgetSet_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QWidgetSet) SetSharable_0() {
	// QWidgetSet_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// void clear()
func (this *QWidgetSet) Clear_0() {
	// QWidgetSet_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// bool remove(const T &)
func (this *QWidgetSet) Remove_0() bool {
	// QWidgetSet_remove_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_remove_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains(const T &)
func (this *QWidgetSet) Contains_0() bool {
	// QWidgetSet_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains(const QSet<T> &)
func (this *QWidgetSet) Contains_1() bool {
	// QWidgetSet_contains_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_contains_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// QSet::iterator begin()
func (this *QWidgetSet) Begin_0() {
	// QWidgetSet_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator begin()
func (this *QWidgetSet) Begin_1() {
	// QWidgetSet_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator cbegin()
func (this *QWidgetSet) Cbegin_0() {
	// QWidgetSet_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator constBegin()
func (this *QWidgetSet) ConstBegin_0() {
	// QWidgetSet_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::iterator end()
func (this *QWidgetSet) End_0() {
	// QWidgetSet_end_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator end()
func (this *QWidgetSet) End_1() {
	// QWidgetSet_end_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator cend()
func (this *QWidgetSet) Cend_0() {
	// QWidgetSet_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator constEnd()
func (this *QWidgetSet) ConstEnd_0() {
	// QWidgetSet_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin_0() {
	// QWidgetSet_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::reverse_iterator rend()
func (this *QWidgetSet) Rend_0() {
	// QWidgetSet_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin_1() {
	// QWidgetSet_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator rend()
func (this *QWidgetSet) Rend_1() {
	// QWidgetSet_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator crbegin()
func (this *QWidgetSet) Crbegin_0() {
	// QWidgetSet_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator crend()
func (this *QWidgetSet) Crend_0() {
	// QWidgetSet_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::iterator erase(class QSet::iterator)
func (this *QWidgetSet) Erase_0() {
	// QWidgetSet_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::iterator erase(class QSet::const_iterator)
func (this *QWidgetSet) Erase_1() {
	// QWidgetSet_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// int count()
func (this *QWidgetSet) Count_0() int {
	// QWidgetSet_count_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0
}

// QSet::iterator insert(const T &)
func (this *QWidgetSet) Insert_0() {
	// QWidgetSet_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::iterator find(const T &)
func (this *QWidgetSet) Find_0() {
	// QWidgetSet_find_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_find_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator find(const T &)
func (this *QWidgetSet) Find_1() {
	// QWidgetSet_find_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_find_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet::const_iterator constFind(const T &)
func (this *QWidgetSet) ConstFind_0() {
	// QWidgetSet_constFind_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_constFind_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet<T> & unite(const QSet<T> &)
func (this *QWidgetSet) Unite_0() *QWidgetSet {
	// QWidgetSet_unite_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_unite_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & intersect(const QSet<T> &)
func (this *QWidgetSet) Intersect_0() *QWidgetSet {
	// QWidgetSet_intersect_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_intersect_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// bool intersects(const QSet<T> &)
func (this *QWidgetSet) Intersects_0() bool {
	// QWidgetSet_intersects_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_intersects_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// QSet<T> & subtract(const QSet<T> &)
func (this *QWidgetSet) Subtract_0() *QWidgetSet {
	// QWidgetSet_subtract_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_subtract_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// bool empty()
func (this *QWidgetSet) Empty_0() bool {
	// QWidgetSet_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// QSet<T> & operator<<(const T &)
func (this *QWidgetSet) Operator_left_shift_0() *QWidgetSet {
	// QWidgetSet_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator|=(const QSet<T> &)
func (this *QWidgetSet) Operator_or_equal_0() *QWidgetSet {
	// QWidgetSet_operator_or_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_or_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator|=(const T &)
func (this *QWidgetSet) Operator_or_equal_1() *QWidgetSet {
	// QWidgetSet_operator_or_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_or_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator&=(const QSet<T> &)
func (this *QWidgetSet) Operator_and_equal_0() *QWidgetSet {
	// QWidgetSet_operator_and_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_and_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator&=(const T &)
func (this *QWidgetSet) Operator_and_equal_1() *QWidgetSet {
	// QWidgetSet_operator_and_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_and_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator+=(const QSet<T> &)
func (this *QWidgetSet) Operator_add_equal_0() *QWidgetSet {
	// QWidgetSet_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator+=(const T &)
func (this *QWidgetSet) Operator_add_equal_1() *QWidgetSet {
	// QWidgetSet_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator-=(const QSet<T> &)
func (this *QWidgetSet) Operator_minus_equal_0() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_minus_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator-=(const T &)
func (this *QWidgetSet) Operator_minus_equal_1() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_minus_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> operator|(const QSet<T> &)
func (this *QWidgetSet) Operator_or_0() *QWidgetSet {
	// QWidgetSet_operator_or_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_or_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> operator&(const QSet<T> &)
func (this *QWidgetSet) Operator_and_0() *QWidgetSet {
	// QWidgetSet_operator_and_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_and_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> operator+(const QSet<T> &)
func (this *QWidgetSet) Operator_add_0() *QWidgetSet {
	// QWidgetSet_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet<T> operator-(const QSet<T> &)
func (this *QWidgetSet) Operator_minus_0() *QWidgetSet {
	// QWidgetSet_operator_minus_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_operator_minus_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QList<T> toList()
func (this *QWidgetSet) ToList_0() {
	// QWidgetSet_toList_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_toList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QList<T> values()
func (this *QWidgetSet) Values_0() {
	// QWidgetSet_values_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_values_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// QSet<T> fromList(const QList<T> &)
func (this *QWidgetSet) FromList_0() *QWidgetSet {
	// QWidgetSet_fromList_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_fromList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return this
}

// QSet::const_iterator m2c(class QSet::iterator)
func (this *QWidgetSet) M2c_0() {
	// QWidgetSet_m2c_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_m2c_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// bool isValidIterator(const class QSet::iterator &)
func (this *QWidgetSet) IsValidIterator_0() bool {
	// QWidgetSet_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

// bool isValidIterator(const class QSet::const_iterator &)
func (this *QWidgetSet) IsValidIterator_1() bool {
	// QWidgetSet_isValidIterator_1()
	rv, err := qtrt.InvokeQtFunc6("QWidgetSet_isValidIterator_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	gopp.ErrPrint(err, rv)
	return 0 == 0
}

//  body block end
