# Typ

[![Build Status](https://travis-ci.org/gurukami/typ.svg "Travis CI status")](https://travis-ci.org/gurukami/typ)
[![GoDoc](https://godoc.org/github.com/gurukami/typ?status.svg)](https://godoc.org/github.com/gurukami/typ)

Typ is a library providing a powerful interface to impressive user experience with conversion and fetching data from built-in types in Golang

## Features

* Safe conversion along built-in types like as `bool`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128`, `string`
* Null types for all primitive types with supported interfaces: ```json.Unmarshaler```, ```json.Marshaler```, ```sql.Scanner```, ```driver.Valuer```
* Value retriever for multidimensional unstructured data from interface
* Conversion functions present via interface (reflection) and native types for better performance
* Some humanize string conversion functions

## Installation

Use ```go get``` to install the latest version of the library. 

```go
    go get -u github.com/gurukami/typ
```

Then include package in your application and enjoy.

```go
    import "github.com/gurukami/typ"
```
## Usage

**Of(interface{})** conversion from interface value to built-in type

```go
// typ.Of(v interface{}, options ...Option).{Type}(defaultValue ...{Type})
//
// Where {Type} any of 
//      Bool, 
//      Int, Int8, Int16, Int32, Int64, 
//      Uint, Uint8, Uint16, Uint32, Uint64, 
//      Float32, Float, 
//      Complex64, Complex, 
//      String
//
// All methods for conversion returns Null{Type} struct with helpful methods & fields
//
//  fields:
//      P - pointer to value
//      Error - conversion error
//
//  methods:
//      V() - value of type
//      Present() - determines whether a value has been set
//      Valid() - determines whether a value has been valid (without error)
//      
//      Scan(value interface{})         | sql.Scanner
//      Value() (driver.Value, error)   | driver.Valuer
//
//      UnmarshalJSON(b []byte) error   | json.Unmarshaler
//      MarshalJSON() ([]byte, error)   | json.Marshaler

// Valid
nv := typ.Of(3.1415926535, typ.FmtByte('g'), typ.Precision(4)).String()
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: 3.142, Valid: true, Present: true, Error: <nil>

// Not valid
nv = typ.Of(3.1415926535).Int()
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: 3, Valid: false, Present: true, Error: value can't safely convert
```

**Native conversion without `reflection` when type is know** 

```go
// For the best performance always use this way if you know about exact type
//
// typ.{FromType}{ToType}(value , [options ...{FromType}{ToType}Option]).V()
// 
// Where {FromType}, {ToType} any of 
//      Bool, 
//      Int, Int8, Int16, Int32, Int64, 
//      Uint, Uint8, Uint16, Uint32, Uint64, 
//      Float32, Float, 
//      Complex64, Complex, 
//      String
//
// All methods for conversion returns Null{Type} struct with helpful methods & fields, additional info you can read in example above

// Valid
nv := typ.FloatString(3.1415926535, typ.FloatStringFmtByte('g'), typ.FloatStringPrecision(4))
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: 3.142, Valid: true, Present: true, Error: <nil>

// Not valid
nv = typ.FloatInt(3.1415926535)
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: 3, Valid: false, Present: true, Error: value can't safely convert
```

**Retrieve multidimensional unstructured data from interface** 

```go
data := map[int]interface{}{
   0: []interface{}{
      0: map[string]int{
         "0": 42,
      },
   },
}

// Instead of do something like this 
//  data[0].([]interface{})[0].(map[string]int)["0”]
//      and not caught a panic
// use this

// Value exists
nv := typ.Of(data).Get(0, 0, "0").Interface()
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: 42, Valid: true, Present: true, Error: <nil>

// Value not exists
nv = typ.Of(data).Get(3, 7, "5").Interface()
fmt.Printf("Value: %v, Valid: %v, Present: %v, Error: %v\n", nv.V(), nv.Valid(), nv.Present(), nv.Error)
// Output: Value: <nil>, Valid: false, Present: false, Error: out of bounds on given data
```

**Rules of safely type conversion along types**

| From / to   | Bool | Int* |  String |  Uint*  |  Float* | Complex*  |
|----------|:-------------:|:------:|:---:|:---:|:---:|:---:|
| Bool      | + | + | + | +  | +  |  + |
| Int* |   +    | +   | `formatting`  | `>= 0`  | `24bit or 53bit`  | `real`, `24bit or 53bit`  |
| String | `parsing` | `parsing`  | +  | `parsing`  | `parsing`  | `parsing`  |
| Uint*  | + | `63bit` | `formatting`  | +  | `24bit or 53bit`  |  `24bit or 53bit`  |
| Float*   |       +        |   `24bit or 53bit`    |  `formatting` | `24bit or 53bit`  |  + | +  |
| Complex*   |       +        |   `real`, `24bit or 53bit`    | +  | `>= 0`, `real`, `24bit or 53bit`  | `real` | +  |

\* based on bit size capacity, `8,16,32,64` for `Int`,`Uint`; `32,64` for `Float`,`Complex`

## Donation for amazing goal

I like airplanes and i want to get private pilot licence, and i believe you can help me to make my dream come true :)  

[ >>>>>>>>>> **Make a dream come true** <<<<<<<<<< ](https://gist.github.com/Nerufa/0d868899d628b1b105f74b6da501bc1f)


## License

The MIT license  
Copyright (c) 2019 Gurukami