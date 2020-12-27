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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  keep block begin

func init_unused_10347() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QWidgetSet struct {
	*qtrt.CObject
}

// void swap(QSet<T> &)
func (this *QWidgetSet) Swap0() {
	// QWidgetSet_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QSet<T> &)
func (this *QWidgetSet) Operator_equal_equal0() bool {
	// QWidgetSet_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QSet<T> &)
func (this *QWidgetSet) Operator_not_equal0() bool {
	// QWidgetSet_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QWidgetSet) Size0() int {
	// QWidgetSet_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool isEmpty()
func (this *QWidgetSet) IsEmpty0() bool {
	// QWidgetSet_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int capacity()
func (this *QWidgetSet) Capacity0() int {
	// QWidgetSet_capacity_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_capacity_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void reserve(int)
func (this *QWidgetSet) Reserve0() {
	// QWidgetSet_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void squeeze()
func (this *QWidgetSet) Squeeze0() {
	// QWidgetSet_squeeze_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_squeeze_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach()
func (this *QWidgetSet) Detach0() {
	// QWidgetSet_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QWidgetSet) IsDetached0() bool {
	// QWidgetSet_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QWidgetSet) SetSharable0() {
	// QWidgetSet_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void clear()
func (this *QWidgetSet) Clear0() {
	// QWidgetSet_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool remove(const T &)
func (this *QWidgetSet) Remove0() bool {
	// QWidgetSet_remove_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_remove_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains(const T &)
func (this *QWidgetSet) Contains0() bool {
	// QWidgetSet_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains(const QSet<T> &)
func (this *QWidgetSet) Contains1() bool {
	// QWidgetSet_contains_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_contains_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QSet::iterator begin()
func (this *QWidgetSet) Begin0() {
	// QWidgetSet_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator begin()
func (this *QWidgetSet) Begin1() {
	// QWidgetSet_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator cbegin()
func (this *QWidgetSet) Cbegin0() {
	// QWidgetSet_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator constBegin()
func (this *QWidgetSet) ConstBegin0() {
	// QWidgetSet_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::iterator end()
func (this *QWidgetSet) End0() {
	// QWidgetSet_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator end()
func (this *QWidgetSet) End1() {
	// QWidgetSet_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator cend()
func (this *QWidgetSet) Cend0() {
	// QWidgetSet_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator constEnd()
func (this *QWidgetSet) ConstEnd0() {
	// QWidgetSet_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin0() {
	// QWidgetSet_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::reverse_iterator rend()
func (this *QWidgetSet) Rend0() {
	// QWidgetSet_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin1() {
	// QWidgetSet_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator rend()
func (this *QWidgetSet) Rend1() {
	// QWidgetSet_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator crbegin()
func (this *QWidgetSet) Crbegin0() {
	// QWidgetSet_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_reverse_iterator crend()
func (this *QWidgetSet) Crend0() {
	// QWidgetSet_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::iterator erase(QSet::iterator)
func (this *QWidgetSet) Erase0() {
	// QWidgetSet_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::iterator erase(QSet::const_iterator)
func (this *QWidgetSet) Erase1() {
	// QWidgetSet_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QWidgetSet) Count0() int {
	// QWidgetSet_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QSet::iterator insert(const T &)
func (this *QWidgetSet) Insert0() {
	// QWidgetSet_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::iterator find(const T &)
func (this *QWidgetSet) Find0() {
	// QWidgetSet_find_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_find_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator find(const T &)
func (this *QWidgetSet) Find1() {
	// QWidgetSet_find_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_find_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet::const_iterator constFind(const T &)
func (this *QWidgetSet) ConstFind0() {
	// QWidgetSet_constFind_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_constFind_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> & unite(const QSet<T> &)
func (this *QWidgetSet) Unite0() *QWidgetSet {
	// QWidgetSet_unite_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_unite_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & intersect(const QSet<T> &)
func (this *QWidgetSet) Intersect0() *QWidgetSet {
	// QWidgetSet_intersect_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_intersect_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// bool intersects(const QSet<T> &)
func (this *QWidgetSet) Intersects0() bool {
	// QWidgetSet_intersects_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_intersects_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QSet<T> & subtract(const QSet<T> &)
func (this *QWidgetSet) Subtract0() *QWidgetSet {
	// QWidgetSet_subtract_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_subtract_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// bool empty()
func (this *QWidgetSet) Empty0() bool {
	// QWidgetSet_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QSet<T> & operator<<(const T &)
func (this *QWidgetSet) Operator_left_shift0() *QWidgetSet {
	// QWidgetSet_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator|=(const QSet<T> &)
func (this *QWidgetSet) Operator_or_equal0() *QWidgetSet {
	// QWidgetSet_operator_or_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_or_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator|=(const T &)
func (this *QWidgetSet) Operator_or_equal1() *QWidgetSet {
	// QWidgetSet_operator_or_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_or_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator&=(const QSet<T> &)
func (this *QWidgetSet) Operator_and_equal0() *QWidgetSet {
	// QWidgetSet_operator_and_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_and_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator&=(const T &)
func (this *QWidgetSet) Operator_and_equal1() *QWidgetSet {
	// QWidgetSet_operator_and_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_and_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator+=(const QSet<T> &)
func (this *QWidgetSet) Operator_add_equal0() *QWidgetSet {
	// QWidgetSet_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator+=(const T &)
func (this *QWidgetSet) Operator_add_equal1() *QWidgetSet {
	// QWidgetSet_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator-=(const QSet<T> &)
func (this *QWidgetSet) Operator_minus_equal0() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_minus_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> & operator-=(const T &)
func (this *QWidgetSet) Operator_minus_equal1() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_minus_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> operator|(const QSet<T> &)
func (this *QWidgetSet) Operator_or0() *QWidgetSet {
	// QWidgetSet_operator_or_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_or_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> operator&(const QSet<T> &)
func (this *QWidgetSet) Operator_and0() *QWidgetSet {
	// QWidgetSet_operator_and_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_and_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> operator+(const QSet<T> &)
func (this *QWidgetSet) Operator_add0() *QWidgetSet {
	// QWidgetSet_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet<T> operator-(const QSet<T> &)
func (this *QWidgetSet) Operator_minus0() *QWidgetSet {
	// QWidgetSet_operator_minus_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_operator_minus_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> values()
func (this *QWidgetSet) Values0() {
	// QWidgetSet_values_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_values_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> toList()
func (this *QWidgetSet) ToList0() {
	// QWidgetSet_toList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_toList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> fromList(const QList<T> &)
func (this *QWidgetSet) FromList0() *QWidgetSet {
	// QWidgetSet_fromList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_fromList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QSet::const_iterator m2c(QSet::iterator)
func (this *QWidgetSet) M2c0() {
	// QWidgetSet_m2c_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_m2c_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QSet::iterator &)
func (this *QWidgetSet) IsValidIterator0() bool {
	// QWidgetSet_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isValidIterator(const QSet::const_iterator &)
func (this *QWidgetSet) IsValidIterator1() bool {
	// QWidgetSet_isValidIterator_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_isValidIterator_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

//  body block end
