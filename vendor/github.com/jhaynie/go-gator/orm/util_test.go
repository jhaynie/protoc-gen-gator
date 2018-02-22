package orm

import (
	"database/sql"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type tmp struct {
	msg string
}

func TestToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(ToString(nil), "<nil>", "should have been <nil> string")
	assert.Equal(ToString(""), "", "should have been empty string")
	assert.Equal(ToString("123"), "123", "should have been 123")
	assert.Equal(ToString(sql.NullString{String: "", Valid: false}), "", "should have been empty string")
	assert.Equal(ToString(&sql.NullString{String: "", Valid: false}), "", "should have been empty string")
	assert.Equal(ToString(sql.NullString{String: "yes", Valid: true}), "yes", "should have been yes")
	assert.Equal(ToString(&sql.NullString{String: "yes", Valid: true}), "yes", "should have been yes")
	assert.Equal(ToString(123), "123", "should have been 123")
	assert.Equal(ToString(int32(123)), "123", "should have been 123")
	assert.Equal(ToString(int64(123)), "123", "should have been 123")
	assert.Equal(ToString(float32(123)), "123.000000", "should have been 123.000000")
	assert.Equal(ToString(float64(123)), "123.000000", "should have been 123.000000")
	assert.Equal(ToString(true), "true", "should have been true")
	assert.Equal(ToString(false), "false", "should have been false")
	assert.Equal(ToString(sql.NullInt64{Int64: 123, Valid: true}), "123", "should have been 123")
	assert.Equal(ToString(&sql.NullInt64{Int64: 123, Valid: true}), "123", "should have been 123")
	assert.Equal(ToString(&sql.NullInt64{Int64: 0, Valid: false}), "0", "should have been 0")
	assert.Equal(ToString(&sql.NullFloat64{Float64: 0, Valid: false}), "0.000000", "should have been 0.000000")
	assert.Equal(ToString(&sql.NullFloat64{Float64: 123, Valid: true}), "123.000000", "should have been 123.000000")
	assert.Equal(ToString(sql.NullFloat64{Float64: 123, Valid: true}), "123.000000", "should have been 123.000000")
	assert.Equal(ToString(sql.NullBool{Bool: true, Valid: true}), "true", "should have been true")
	assert.Equal(ToString(&sql.NullBool{Bool: true, Valid: true}), "true", "should have been true")
	assert.Equal(ToString(sql.NullBool{Bool: false, Valid: true}), "false", "should have been false")
	assert.Equal(ToString(&sql.NullBool{Bool: false, Valid: true}), "false", "should have been false")
	tv := time.Now()
	tvs := tv.String()
	assert.Equal(ToString(tv), tvs, "should have been "+tvs)
	assert.Equal(ToString(mysql.NullTime{Time: tv, Valid: true}), tvs, "should have been "+tvs)
	s := "abc"
	assert.Equal(ToString(&s), "abc", "should have been abc")
	sd := ToSQLDate("2017-03-17T21:35:27Z")
	tm := ToTimestamp(sd)
	assert.Equal(ToString(tm), "1489786527.0", "should have been 1489786527.0")
}

func TestSQLString(t *testing.T) {
	assert := assert.New(t)
	s := "select * from foo"
	str := ToSQLString(s)
	assert.Equal(str.String, s, "should have been "+s)
	str = ToSQLString(&s)
	assert.Equal(str.String, s, "should have been "+s)
}

func TestSQLDate(t *testing.T) {
	assert := assert.New(t)
	v := ToSQLDate(nil)
	assert.Equal(v.Valid, false, "should have been invalid")
	v = ToSQLDate("now")
	assert.Equal(v.Valid, true, "should have been valid")
	tv := time.Now()
	v = ToSQLDate(tv)
	assert.Equal(v.Valid, true, "should have been valid")
	v = ToSQLDate("2017-03-17T21:35:27Z")
	assert.Equal(v.Valid, true, "should have been valid")
	v = ToSQLDate(ISODate())
	assert.Equal(v.Valid, true, "should have been valid")
}

func TestInt64(t *testing.T) {
	assert := assert.New(t)
	v := toInt64("123")
	assert.Equal(v, int64(123), "should have been 123")
	v = toInt64("")
	assert.Equal(v, int64(0), "should have been 0")
}

func TestInt32(t *testing.T) {
	assert := assert.New(t)
	v := toInt32("123")
	assert.Equal(v, int32(123), "should have been 123")
	v = toInt32("")
	assert.Equal(v, int32(0), "should have been 0")
}

func TestFloat64(t *testing.T) {
	assert := assert.New(t)
	v := toFloat64("123.0")
	assert.Equal(v, float64(123), "should have been 123.0000")
	v = toFloat64("")
	assert.Equal(v, float64(0), "should have been 0.0000")
}

func TestSQLInt64(t *testing.T) {
	assert := assert.New(t)
	v := ToSQLInt64("123")
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Int64, int64(123), "should have been 123")
	i := 123
	v = ToSQLInt64(&i)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Int64, int64(123), "should have been 123")
	v = ToSQLInt64(i)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Int64, int64(123), "should have been 123")
	v = ToSQLInt64("")
	assert.Equal(v.Valid, false, "should have been false")
	assert.Equal(v.Int64, int64(0), "should have been 0")
}

func TestSQLFloat64(t *testing.T) {
	assert := assert.New(t)
	v := ToSQLFloat64("123.0")
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Float64, float64(123.0), "should have been 123.0000")

	f := 123.0
	v = ToSQLFloat64(&f)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Float64, float64(123.0), "should have been 123.0000")

	v = ToSQLFloat64(f)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Float64, float64(123.0), "should have been 123.0000")

	v = ToSQLFloat64("")
	assert.Equal(v.Valid, false, "should have been false")
	assert.Equal(v.Float64, float64(0), "should have been 0")
}

func TestSQLBool(t *testing.T) {
	assert := assert.New(t)
	v := ToSQLBool("true")
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, true, "should have been true")
	v = ToSQLBool("false")
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, false, "should have been false")
	v = ToSQLBool(false)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, false, "should have been false")
	f := false
	v = ToSQLBool(&f)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, false, "should have been false")
	v = ToSQLBool("")
	assert.Equal(v.Valid, false, "should have been false")
	assert.Equal(v.Bool, false, "should have been false")
	v = ToSQLBool(1)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, true, "should have been true")
	v = ToSQLBool(0)
	assert.Equal(v.Valid, true, "should have been true")
	assert.Equal(v.Bool, false, "should have been false")
}

func TestHash(t *testing.T) {
	assert := assert.New(t)
	v := HashStrings("1")
	assert.Equal("b7b41276360564d4", v, "should have been b7b41276360564d4")
	v = HashStrings("1", "2")
	assert.Equal("5460f49adbe7aba2", v, "should have been 5460f49adbe7aba2")
}

func TestGeometry(t *testing.T) {
	assert := assert.New(t)
	g := ToGeometry("POINT(-122.3890954 37.6145378)")
	assert.Equal(g.String(), "latitude:37.614536 longitude:-122.3891 ")
	assert.Equal(g.Latitude, float32(37.614536))
	assert.Equal(g.Longitude, float32(-122.3891))
}

func TestTimestamp(t *testing.T) {
	assert := assert.New(t)
	tv := time.Now()
	ts := ToTimestamp(ToSQLDate(tv))
	assert.Equal(ts.Nanos, int32(tv.Nanosecond()))
	assert.Equal(ts.Seconds, tv.Unix())

	dt := ToTimestamp(mysql.NullTime{Time: time.Now(), Valid: true})
	assert.NotNil(dt)
	sdt := ToSQLDate(dt)
	assert.Equal(sdt.Valid, true)
}

func TestNullInt32(t *testing.T) {
	v := NullInt32
	assert := assert.New(t)
	assert.Equal(v, NullInt32)
	assert.False(IsNullInt(123))
	assert.True(IsNullInt(int32(v)))
}

func TestJSON(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(`{"a":"b"}`, Stringify(map[string]string{"a": "b"}))

	assert.Equal(`{
	"a": "b"
}`, Stringify(map[string]string{"a": "b"}, true))
}

func TestHashValues(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("b7b41276360564d4", HashValues("1"))
	assert.Equal("c990c404884a29c1", HashValues(1, true))

	s := "1"
	assert.Equal("b7b41276360564d4", HashValues(&s))

	n := 1
	assert.Equal("b7b41276360564d4", HashValues(&n))

	n2 := 1
	assert.Equal("b7b41276360564d4", HashValues(n2))

	assert.Equal("b7b41276360564d4", HashValues(int32(1)))

	assert.Equal("b7b41276360564d4", HashValues(int64(1)))

	var intV = 12
	assert.Equal("5460f49adbe7aba2", HashValues(intV))
	assert.Equal("5460f49adbe7aba2", HashValues(&intV))

	var int32V int32 = 12
	assert.Equal("5460f49adbe7aba2", HashValues(&int32V))

	var int64V int64 = 45
	assert.Equal("fd6817dc570d1402", HashValues(&int64V))

	var float32V float32 = 45
	assert.Equal("d66eb8eb7ef9c413", HashValues(&float32V))

	var float64V float64 = 455
	assert.Equal("bbd50ab530eafec2", HashValues(&float64V))

	var float32P *float32
	assert.Equal("ef46db3751d8e999", HashValues(float32P))

	var float64P *float64
	assert.Equal("ef46db3751d8e999", HashValues(float64P))

	var intP *int
	assert.Equal("ef46db3751d8e999", HashValues(intP))
	intP = &intV
	assert.Equal("5460f49adbe7aba2", HashValues(intP))

	var int32P *int32
	assert.Equal("ef46db3751d8e999", HashValues(int32P))

	var int64P *int64
	assert.Equal("ef46db3751d8e999", HashValues(int64P))
	assert.Equal("7c5b4e400f80bf7c", HashValues(nil))

	var uintV uint = 1
	assert.NotEmpty(HashValues(&uintV))
	assert.Equal("b7b41276360564d4", HashValues(uintV))
	var uintP = &uintV
	assert.NotEmpty(HashValues(uintP))
	uintP = nil
	assert.Equal("7c5b4e400f80bf7c", HashValues(uintP))

	boolV := false
	var boolP *bool
	assert.NotEmpty(HashValues(&boolV))
	assert.Equal("d7c9b97948142e4a", HashValues(true))
	assert.Equal("6d3f99ccc0c03a7a", HashValues(false))
	boolP = nil
	assert.NotEmpty(HashValues(&boolP))

	obj := &tmp{msg: "mg"}
	assert.Equal("2151b3b520c20d4f", HashValues(obj))
	assert.Equal("393845fb4cad9040", HashValues(float32(23)))
	assert.Equal("45abfe9e1ba43375", HashValues(float64(33)))

	assert.Equal("ea8842e9ea2638fa", HashValues([]byte("hi")))
}

var result string

// go test -run=Bench -bench=.  ./pkg/util
func BenchmarkHashValuesString(b *testing.B) {
	var r string
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = HashValues("1")
	}
	// prevent compiler optimizations
	result = r
}

func BenchmarkHashValuesStringPointer(b *testing.B) {
	var r string
	p := "1"
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = HashValues(&p)
	}
	// prevent compiler optimizations
	result = r
}

func BenchmarkHashValuesInt(b *testing.B) {
	var r string
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = HashValues(n)
	}
	// prevent compiler optimizations
	result = r
}

func BenchmarkHashValuesInt32(b *testing.B) {
	var r string
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = HashValues(int32(n))
	}
	// prevent compiler optimizations
	result = r
}

func BenchmarkHashValuesIntPtr(b *testing.B) {
	var r string
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = HashValues(&n)
	}
	// prevent compiler optimizations
	result = r
}
