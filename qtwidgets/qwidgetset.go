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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
}

// bool operator==(const QSet<T> &)
func (this *QWidgetSet) Operator_equal_equal_0() bool {
	// QWidgetSet_operator_equal_equal_0()
	return 0 == 0
}

// bool operator!=(const QSet<T> &)
func (this *QWidgetSet) Operator_not_equal_0() bool {
	// QWidgetSet_operator_not_equal_0()
	return 0 == 0
}

// int size()
func (this *QWidgetSet) Size_0() int {
	// QWidgetSet_size_0()
	return 0
}

// bool isEmpty()
func (this *QWidgetSet) IsEmpty_0() bool {
	// QWidgetSet_isEmpty_0()
	return 0 == 0
}

// int capacity()
func (this *QWidgetSet) Capacity_0() int {
	// QWidgetSet_capacity_0()
	return 0
}

// void reserve(int)
func (this *QWidgetSet) Reserve_0() {
	// QWidgetSet_reserve_0()
}

// void squeeze()
func (this *QWidgetSet) Squeeze_0() {
	// QWidgetSet_squeeze_0()
}

// void detach()
func (this *QWidgetSet) Detach_0() {
	// QWidgetSet_detach_0()
}

// bool isDetached()
func (this *QWidgetSet) IsDetached_0() bool {
	// QWidgetSet_isDetached_0()
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QWidgetSet) SetSharable_0() {
	// QWidgetSet_setSharable_0()
}

// void clear()
func (this *QWidgetSet) Clear_0() {
	// QWidgetSet_clear_0()
}

// bool remove(const T &)
func (this *QWidgetSet) Remove_0() bool {
	// QWidgetSet_remove_0()
	return 0 == 0
}

// bool contains(const T &)
func (this *QWidgetSet) Contains_0() bool {
	// QWidgetSet_contains_0()
	return 0 == 0
}

// bool contains(const QSet<T> &)
func (this *QWidgetSet) Contains_1() bool {
	// QWidgetSet_contains_1()
	return 0 == 0
}

// QSet::iterator begin()
func (this *QWidgetSet) Begin_0() {
	// QWidgetSet_begin_0()
}

// QSet::const_iterator begin()
func (this *QWidgetSet) Begin_1() {
	// QWidgetSet_begin_1()
}

// QSet::const_iterator cbegin()
func (this *QWidgetSet) Cbegin_0() {
	// QWidgetSet_cbegin_0()
}

// QSet::const_iterator constBegin()
func (this *QWidgetSet) ConstBegin_0() {
	// QWidgetSet_constBegin_0()
}

// QSet::iterator end()
func (this *QWidgetSet) End_0() {
	// QWidgetSet_end_0()
}

// QSet::const_iterator end()
func (this *QWidgetSet) End_1() {
	// QWidgetSet_end_1()
}

// QSet::const_iterator cend()
func (this *QWidgetSet) Cend_0() {
	// QWidgetSet_cend_0()
}

// QSet::const_iterator constEnd()
func (this *QWidgetSet) ConstEnd_0() {
	// QWidgetSet_constEnd_0()
}

// QSet::reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin_0() {
	// QWidgetSet_rbegin_0()
}

// QSet::reverse_iterator rend()
func (this *QWidgetSet) Rend_0() {
	// QWidgetSet_rend_0()
}

// QSet::const_reverse_iterator rbegin()
func (this *QWidgetSet) Rbegin_1() {
	// QWidgetSet_rbegin_1()
}

// QSet::const_reverse_iterator rend()
func (this *QWidgetSet) Rend_1() {
	// QWidgetSet_rend_1()
}

// QSet::const_reverse_iterator crbegin()
func (this *QWidgetSet) Crbegin_0() {
	// QWidgetSet_crbegin_0()
}

// QSet::const_reverse_iterator crend()
func (this *QWidgetSet) Crend_0() {
	// QWidgetSet_crend_0()
}

// QSet::iterator erase(class QSet::iterator)
func (this *QWidgetSet) Erase_0() {
	// QWidgetSet_erase_0()
}

// QSet::iterator erase(class QSet::const_iterator)
func (this *QWidgetSet) Erase_1() {
	// QWidgetSet_erase_1()
}

// int count()
func (this *QWidgetSet) Count_0() int {
	// QWidgetSet_count_0()
	return 0
}

// QSet::iterator insert(const T &)
func (this *QWidgetSet) Insert_0() {
	// QWidgetSet_insert_0()
}

// QSet::iterator find(const T &)
func (this *QWidgetSet) Find_0() {
	// QWidgetSet_find_0()
}

// QSet::const_iterator find(const T &)
func (this *QWidgetSet) Find_1() {
	// QWidgetSet_find_1()
}

// QSet::const_iterator constFind(const T &)
func (this *QWidgetSet) ConstFind_0() {
	// QWidgetSet_constFind_0()
}

// QSet<T> & unite(const QSet<T> &)
func (this *QWidgetSet) Unite_0() *QWidgetSet {
	// QWidgetSet_unite_0()
	return this
}

// QSet<T> & intersect(const QSet<T> &)
func (this *QWidgetSet) Intersect_0() *QWidgetSet {
	// QWidgetSet_intersect_0()
	return this
}

// bool intersects(const QSet<T> &)
func (this *QWidgetSet) Intersects_0() bool {
	// QWidgetSet_intersects_0()
	return 0 == 0
}

// QSet<T> & subtract(const QSet<T> &)
func (this *QWidgetSet) Subtract_0() *QWidgetSet {
	// QWidgetSet_subtract_0()
	return this
}

// bool empty()
func (this *QWidgetSet) Empty_0() bool {
	// QWidgetSet_empty_0()
	return 0 == 0
}

// QSet<T> & operator<<(const T &)
func (this *QWidgetSet) Operator_left_shift_0() *QWidgetSet {
	// QWidgetSet_operator_left_shift_0()
	return this
}

// QSet<T> & operator|=(const QSet<T> &)
func (this *QWidgetSet) Operator_or_equal_0() *QWidgetSet {
	// QWidgetSet_operator_or_equal_0()
	return this
}

// QSet<T> & operator|=(const T &)
func (this *QWidgetSet) Operator_or_equal_1() *QWidgetSet {
	// QWidgetSet_operator_or_equal_1()
	return this
}

// QSet<T> & operator&=(const QSet<T> &)
func (this *QWidgetSet) Operator_and_equal_0() *QWidgetSet {
	// QWidgetSet_operator_and_equal_0()
	return this
}

// QSet<T> & operator&=(const T &)
func (this *QWidgetSet) Operator_and_equal_1() *QWidgetSet {
	// QWidgetSet_operator_and_equal_1()
	return this
}

// QSet<T> & operator+=(const QSet<T> &)
func (this *QWidgetSet) Operator_add_equal_0() *QWidgetSet {
	// QWidgetSet_operator_add_equal_0()
	return this
}

// QSet<T> & operator+=(const T &)
func (this *QWidgetSet) Operator_add_equal_1() *QWidgetSet {
	// QWidgetSet_operator_add_equal_1()
	return this
}

// QSet<T> & operator-=(const QSet<T> &)
func (this *QWidgetSet) Operator_minus_equal_0() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_0()
	return this
}

// QSet<T> & operator-=(const T &)
func (this *QWidgetSet) Operator_minus_equal_1() *QWidgetSet {
	// QWidgetSet_operator_minus_equal_1()
	return this
}

// QSet<T> operator|(const QSet<T> &)
func (this *QWidgetSet) Operator_or_0() *QWidgetSet {
	// QWidgetSet_operator_or_0()
	return this
}

// QSet<T> operator&(const QSet<T> &)
func (this *QWidgetSet) Operator_and_0() *QWidgetSet {
	// QWidgetSet_operator_and_0()
	return this
}

// QSet<T> operator+(const QSet<T> &)
func (this *QWidgetSet) Operator_add_0() *QWidgetSet {
	// QWidgetSet_operator_add_0()
	return this
}

// QSet<T> operator-(const QSet<T> &)
func (this *QWidgetSet) Operator_minus_0() *QWidgetSet {
	// QWidgetSet_operator_minus_0()
	return this
}

// QList<T> toList()
func (this *QWidgetSet) ToList_0() {
	// QWidgetSet_toList_0()
}

// QList<T> values()
func (this *QWidgetSet) Values_0() {
	// QWidgetSet_values_0()
}

// QSet<T> fromList(const QList<T> &)
func (this *QWidgetSet) FromList_0() *QWidgetSet {
	// QWidgetSet_fromList_0()
	return this
}

// QSet::const_iterator m2c(class QSet::iterator)
func (this *QWidgetSet) M2c_0() {
	// QWidgetSet_m2c_0()
}

// bool isValidIterator(const class QSet::iterator &)
func (this *QWidgetSet) IsValidIterator_0() bool {
	// QWidgetSet_isValidIterator_0()
	return 0 == 0
}

// bool isValidIterator(const class QSet::const_iterator &)
func (this *QWidgetSet) IsValidIterator_1() bool {
	// QWidgetSet_isValidIterator_1()
	return 0 == 0
}

//  body block end
