package typ

import (
	"errors"
	"fmt"
	"reflect"
)

var matrixSuite = newMatrix()

type Convert func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool)

type Comparator func(a interface{}, b interface{}) (bool)

func newMatrix() *matrix {
	return &matrix{
		data:        make(map[string][]dataItem),
		converters:  make(map[converterKey]converterValue),
		comparators: make(map[reflect.Type]Comparator),
	}
}

type converterKey struct {
	from string
	to   string
}

type converterValue struct {
	converter Convert
	from      reflect.Type
	to        reflect.Type
}

type matrix struct {
	data        map[string][]dataItem
	converters  map[converterKey]converterValue
	comparators map[reflect.Type]Comparator
}

type dataItem struct {
	value reflect.Value
	opts  []interface{}
}

// Typ Value used as default value
func (m *matrix) Register(typ reflect.Type, data []dataItem) {
	cdata := make([]dataItem, len(data))
	copy(cdata, data)
	nameType := typ.String()
	for _, i := range cdata {
		ri := i.value
		if nameType != ri.Type().String() {
			panic(fmt.Sprintf("data contains item with different type %s, expected %s", ri.Type(), typ))
		}
		if ri.Interface() == nil {
			continue
		}
		m.data[nameType] = append(m.data[nameType], i)
	}
}

func (m *matrix) SetConverters(from []reflect.Type, to []reflect.Type, converter Convert) {
	for _, cf := range from {
		for _, ct := range to {
			m.SetConverter(cf, ct, converter)
		}
	}
}

func (m *matrix) SetConverter(from reflect.Type, to reflect.Type, converter Convert) {
	var ft, tt string
	if from == nil {
		ft = "nil"
	} else {
		ft = from.String()
	}
	if to == nil {
		panic("to reflect type cannot be nil")
	}
	tt = to.String()
	if converter == nil {
		panic("converter cannot be nil")
	}
	ci := converterKey{ft, tt}
	if _, ok := m.converters[ci]; ok {
		panic(fmt.Sprintf("converter from %v to %v already registered", ft, tt))
	}
	if ci.from == ci.to {
		panic("from & to type must be different")
	}
	m.converters[ci] = converterValue{converter, from, to}
}

func (m *matrix) SetComparators(comparator Comparator, typ []reflect.Type) {
	for _, t := range typ {
		m.SetComparator(t, comparator)
	}
}

func (m *matrix) SetComparator(typ reflect.Type, comparator Comparator) {
	if comparator == nil {
		panic("comparator cannot be nil")
	}
	if typ.Kind() == reflect.Invalid {
		panic("from & to type cannot be reflect.Invalid")
	}
	m.comparators[typ] = comparator
}

func (m *matrix) Test(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool, error) {
	return m.Convert(from, to, opts...)
}

func (m *matrix) Convert(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool, error) {
	if to == nil {
		panic("to reflect type cannot be nil")
	}
	ft := "<nil>"
	rf := reflect.TypeOf(from)
	if rf != nil {
		ft = rf.String()
	}
	tt := to.String()
	if ft == tt {
		return from, true, nil
	}
	defaultConverter := false
	c, ok := m.converters[converterKey{ft, tt}]
	if !ok {
		c, ok = m.converters[converterKey{ft, "nil"}]
		defaultConverter = ok
	}
	if !ok {
		c, ok = m.converters[converterKey{"nil", tt}]
		defaultConverter = ok
	}
	if !ok {
		panic(errors.New(fmt.Sprintf("no converter from %s to %s type", ft, to)))
	}
	iv, ok := c.converter(from, to, opts...)
	if !defaultConverter {
		tiv := "<nil>"
		dt := reflect.TypeOf(iv)
		if dt != nil {
			tiv = dt.String()
		}
		if tiv != tt && ok {
			panic(fmt.Sprintf("converter from %s to %s returned unexpected type %s", ft, to, tiv))
		}
	}
	var err error
	if iv == nil {
		err = errors.New(fmt.Sprintf("no converter from %s to %s type", ft, to))
	}
	return iv, ok, err
}

func (m *matrix) Generate() []dataItem {
	var out []dataItem
	for _, di := range m.data {
		out = append(out, m.GenerateToTyp(di, nil)...)
	}
	return out
}

func (m *matrix) GenerateToTyp(di []dataItem, to reflect.Type) []dataItem {
	var out []dataItem
	for _, i := range di {
		ivi := reflect.Indirect(i.value)
		if to != nil {
			if i.value.Type() == to {
				out = append(out, i)
			}
		} else {
			out = append(out, i)
		}
		if ivi.Interface() == nil {
			continue
		}
		ivt := ivi.Type().String()
		if to != nil {
			if iv, ok, _ := m.Convert(ivi.Interface(), to); ok {
				out = append(out, dataItem{reflect.ValueOf(iv), nil})
			}
		} else {
			for ci, c := range m.converters {
				if (ci.from == ivt || ci.from == "nil") && ivt != ci.to {
					if iv, ok := c.converter(ivi.Interface(), c.to); ok {
						out = append(out, dataItem{reflect.ValueOf(iv), nil})
					}
				}
			}
		}
	}
	return out
}

func (m *matrix) GenerateComparable() []dataItem {
	var out []dataItem
	for _, di := range m.Generate() {
		if di.value.Type().Comparable() {
			out = append(out, di)
		}
	}
	return out
}

func (m *matrix) GenerateImplements(typ reflect.Type) []dataItem {
	var out []dataItem
	for _, di := range m.Generate() {
		if di.value.Type().Implements(typ) {
			out = append(out, di)
		}
	}
	return out
}

func (m *matrix) Compare(a interface{}, b interface{}) (bool) {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if a == nil && b == nil {
		return true
	}
	ra := reflect.TypeOf(a)
	rb := reflect.TypeOf(b)
	if ra.String() != rb.String() {
		return false
	}
	if c, ok := m.comparators[ra]; ok {
		return c(a, b)
	}
	if !ra.Comparable() || !rb.Comparable() {
		return reflect.DeepEqual(a, b)
	}
	return a == b
}

func (m *matrix) ListConverters() []string {
	var out []string
	for ci, _ := range m.converters {
		out = append(out, fmt.Sprintf("%s ==> %s", ci.from, ci.to))
	}
	return out
}

func (m *matrix) ListTypes() []string {
	var out []string
	for ci, _ := range m.data {
		out = append(out, ci)
	}
	return out
}

func (m *matrix) GetOptByType(opts []interface{}, typ reflect.Type) interface{} {
	for _, v := range opts {
		if reflect.TypeOf(v) == typ {
			return v
		}
	}
	return nil
}
