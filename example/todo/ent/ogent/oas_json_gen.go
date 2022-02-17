// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Encode implements json.Marshaler.
func (s CreateTodoReq) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"title\"" + ":")
		e.Str(s.Title)
	}
	{
		if s.Done.Set {
			e.Comma()
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfCreateTodoReq = [2]string{
	0: "title",
	1: "done",
}

// Decode decodes CreateTodoReq from json.
func (s *CreateTodoReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateTodoReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateTodoReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfCreateTodoReq) {
					name = jsonFieldsNameOfCreateTodoReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode encodes ListTodoOKApplicationJSON as json.
func (s ListTodoOKApplicationJSON) Encode(e *jx.Writer) {
	unwrapped := []TodoList(s)
	e.ArrStart()
	if len(unwrapped) >= 1 {
		// Encode first element without comma.
		{
			elem := unwrapped[0]
			elem.Encode(e)
		}
		for _, elem := range unwrapped[1:] {
			e.Comma()
			elem.Encode(e)
		}
	}
	e.ArrEnd()
}

// Decode decodes ListTodoOKApplicationJSON from json.
func (s *ListTodoOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ListTodoOKApplicationJSON to nil")
	}
	var unwrapped []TodoList
	if err := func() error {
		unwrapped = nil
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem TodoList
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ListTodoOKApplicationJSON(unwrapped)
	return nil
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	switch d.Next() {
	case jx.Bool:
		o.Set = true
		v, err := d.Bool()
		if err != nil {
			return err
		}
		o.Value = bool(v)
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptBool", d.Next())
	}
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptString", d.Next())
	}
}

// Encode implements json.Marshaler.
func (s R400) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"code\"" + ":")
		e.Int(s.Code)
	}
	{
		e.Comma()

		e.RawStr("\"status\"" + ":")
		e.Str(s.Status)
	}
	{
		if len(s.Errors) != 0 {
			e.Comma()
		}

		if len(s.Errors) != 0 {
			e.RawStr("\"errors\"" + ":")
			e.Raw(s.Errors)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfR400 = [3]string{
	0: "code",
	1: "status",
	2: "errors",
}

// Decode decodes R400 from json.
func (s *R400) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode R400 to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "errors":

			if err := func() error {
				v, err := d.Raw()
				s.Errors = jx.Raw(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode R400")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfR400) {
					name = jsonFieldsNameOfR400[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s R404) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"code\"" + ":")
		e.Int(s.Code)
	}
	{
		e.Comma()

		e.RawStr("\"status\"" + ":")
		e.Str(s.Status)
	}
	{
		if len(s.Errors) != 0 {
			e.Comma()
		}

		if len(s.Errors) != 0 {
			e.RawStr("\"errors\"" + ":")
			e.Raw(s.Errors)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfR404 = [3]string{
	0: "code",
	1: "status",
	2: "errors",
}

// Decode decodes R404 from json.
func (s *R404) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode R404 to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "errors":

			if err := func() error {
				v, err := d.Raw()
				s.Errors = jx.Raw(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode R404")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfR404) {
					name = jsonFieldsNameOfR404[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s R409) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"code\"" + ":")
		e.Int(s.Code)
	}
	{
		e.Comma()

		e.RawStr("\"status\"" + ":")
		e.Str(s.Status)
	}
	{
		if len(s.Errors) != 0 {
			e.Comma()
		}

		if len(s.Errors) != 0 {
			e.RawStr("\"errors\"" + ":")
			e.Raw(s.Errors)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfR409 = [3]string{
	0: "code",
	1: "status",
	2: "errors",
}

// Decode decodes R409 from json.
func (s *R409) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode R409 to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "errors":

			if err := func() error {
				v, err := d.Raw()
				s.Errors = jx.Raw(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode R409")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfR409) {
					name = jsonFieldsNameOfR409[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s R500) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"code\"" + ":")
		e.Int(s.Code)
	}
	{
		e.Comma()

		e.RawStr("\"status\"" + ":")
		e.Str(s.Status)
	}
	{
		if len(s.Errors) != 0 {
			e.Comma()
		}

		if len(s.Errors) != 0 {
			e.RawStr("\"errors\"" + ":")
			e.Raw(s.Errors)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfR500 = [3]string{
	0: "code",
	1: "status",
	2: "errors",
}

// Decode decodes R500 from json.
func (s *R500) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode R500 to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "errors":

			if err := func() error {
				v, err := d.Raw()
				s.Errors = jx.Raw(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode R500")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfR500) {
					name = jsonFieldsNameOfR500[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s TodoCreate) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"id\"" + ":")
		e.Int(s.ID)
	}
	{
		e.Comma()

		e.RawStr("\"title\"" + ":")
		e.Str(s.Title)
	}
	{
		if s.Done.Set {
			e.Comma()
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTodoCreate = [3]string{
	0: "id",
	1: "title",
	2: "done",
}

// Decode decodes TodoCreate from json.
func (s *TodoCreate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodoCreate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodoCreate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodoCreate) {
					name = jsonFieldsNameOfTodoCreate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s TodoList) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"id\"" + ":")
		e.Int(s.ID)
	}
	{
		e.Comma()

		e.RawStr("\"title\"" + ":")
		e.Str(s.Title)
	}
	{
		if s.Done.Set {
			e.Comma()
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTodoList = [3]string{
	0: "id",
	1: "title",
	2: "done",
}

// Decode decodes TodoList from json.
func (s *TodoList) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodoList to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodoList")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodoList) {
					name = jsonFieldsNameOfTodoList[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s TodoRead) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"id\"" + ":")
		e.Int(s.ID)
	}
	{
		e.Comma()

		e.RawStr("\"title\"" + ":")
		e.Str(s.Title)
	}
	{
		if s.Done.Set {
			e.Comma()
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTodoRead = [3]string{
	0: "id",
	1: "title",
	2: "done",
}

// Decode decodes TodoRead from json.
func (s *TodoRead) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodoRead to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodoRead")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodoRead) {
					name = jsonFieldsNameOfTodoRead[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s TodoUpdate) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"id\"" + ":")
		e.Int(s.ID)
	}
	{
		e.Comma()

		e.RawStr("\"title\"" + ":")
		e.Str(s.Title)
	}
	{
		if s.Done.Set {
			e.Comma()
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTodoUpdate = [3]string{
	0: "id",
	1: "title",
	2: "done",
}

// Decode decodes TodoUpdate from json.
func (s *TodoUpdate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodoUpdate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodoUpdate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodoUpdate) {
					name = jsonFieldsNameOfTodoUpdate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s UpdateTodoReq) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.Title.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Title.Set {
			e.RawStr("\"title\"" + ":")
			s.Title.Encode(e)
		}
	}
	{
		if s.Done.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Done.Set {
			e.RawStr("\"done\"" + ":")
			s.Done.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfUpdateTodoReq = [2]string{
	0: "title",
	1: "done",
}

// Decode decodes UpdateTodoReq from json.
func (s *UpdateTodoReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateTodoReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":

			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "done":

			if err := func() error {
				s.Done.Reset()
				if err := s.Done.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"done\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateTodoReq")
	}

	return nil
}
