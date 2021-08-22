// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ocp-instruction-api.proto

package ocp_instruction_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Instruction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Instruction) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return InstructionValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() < 0 {
		return InstructionValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than or equal to 0",
		}
	}

	if utf8.RuneCountInString(m.GetText()) < 0 {
		return InstructionValidationError{
			field:  "Text",
			reason: "value length must be at least 0 runes",
		}
	}

	if m.GetPrevId() < 0 {
		return InstructionValidationError{
			field:  "PrevId",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// InstructionValidationError is the validation error returned by
// Instruction.Validate if the designated constraints aren't met.
type InstructionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InstructionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InstructionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InstructionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InstructionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InstructionValidationError) ErrorName() string { return "InstructionValidationError" }

// Error satisfies the builtin error interface
func (e InstructionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInstruction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InstructionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InstructionValidationError{}

// Validate checks the field values on CreateV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInstruction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateV1RequestValidationError{
				field:  "Instruction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateV1RequestValidationError is the validation error returned by
// CreateV1Request.Validate if the designated constraints aren't met.
type CreateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateV1RequestValidationError) ErrorName() string { return "CreateV1RequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateV1RequestValidationError{}

// Validate checks the field values on CreateV1Response with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateV1ResponseValidationError is the validation error returned by
// CreateV1Response.Validate if the designated constraints aren't met.
type CreateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateV1ResponseValidationError) ErrorName() string { return "CreateV1ResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateV1ResponseValidationError{}

// Validate checks the field values on DescribeV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DescribeV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return DescribeV1RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeV1RequestValidationError is the validation error returned by
// DescribeV1Request.Validate if the designated constraints aren't met.
type DescribeV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeV1RequestValidationError) ErrorName() string {
	return "DescribeV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeV1RequestValidationError{}

// Validate checks the field values on DescribeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInstruction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeV1ResponseValidationError{
				field:  "Instruction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeV1ResponseValidationError is the validation error returned by
// DescribeV1Response.Validate if the designated constraints aren't met.
type DescribeV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeV1ResponseValidationError) ErrorName() string {
	return "DescribeV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeV1ResponseValidationError{}

// Validate checks the field values on ListV1Request with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() < 0 {
		return ListV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetOffset() < 0 {
		return ListV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListV1RequestValidationError is the validation error returned by
// ListV1Request.Validate if the designated constraints aren't met.
type ListV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListV1RequestValidationError) ErrorName() string { return "ListV1RequestValidationError" }

// Error satisfies the builtin error interface
func (e ListV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListV1RequestValidationError{}

// Validate checks the field values on ListV1Response with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetInstruction() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListV1ResponseValidationError{
					field:  fmt.Sprintf("Instruction[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListV1ResponseValidationError is the validation error returned by
// ListV1Response.Validate if the designated constraints aren't met.
type ListV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListV1ResponseValidationError) ErrorName() string { return "ListV1ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListV1ResponseValidationError{}

// Validate checks the field values on RemoveV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RemoveV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return RemoveV1RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveV1RequestValidationError is the validation error returned by
// RemoveV1Request.Validate if the designated constraints aren't met.
type RemoveV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveV1RequestValidationError) ErrorName() string { return "RemoveV1RequestValidationError" }

// Error satisfies the builtin error interface
func (e RemoveV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveV1RequestValidationError{}

// Validate checks the field values on RemoveV1Response with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RemoveV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveV1ResponseValidationError is the validation error returned by
// RemoveV1Response.Validate if the designated constraints aren't met.
type RemoveV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveV1ResponseValidationError) ErrorName() string { return "RemoveV1ResponseValidationError" }

// Error satisfies the builtin error interface
func (e RemoveV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveV1ResponseValidationError{}
