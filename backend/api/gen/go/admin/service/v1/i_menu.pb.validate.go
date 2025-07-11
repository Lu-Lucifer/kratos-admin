// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/service/v1/i_menu.proto

package servicev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Menu with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Menu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Menu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MenuMultiError, or nil if none found.
func (m *Menu) ValidateAll() error {
	return m.validate(true)
}

func (m *Menu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if m.Redirect != nil {
		// no validation rules for Redirect
	}

	if m.Alias != nil {
		// no validation rules for Alias
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Component != nil {
		// no validation rules for Component
	}

	if m.Meta != nil {

		if all {
			switch v := interface{}(m.GetMeta()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "Meta",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "Meta",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.CreateBy != nil {
		// no validation rules for CreateBy
	}

	if m.UpdateBy != nil {
		// no validation rules for UpdateBy
	}

	if m.CreateTime != nil {

		if all {
			switch v := interface{}(m.GetCreateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UpdateTime != nil {

		if all {
			switch v := interface{}(m.GetUpdateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DeleteTime != nil {

		if all {
			switch v := interface{}(m.GetDeleteTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeleteTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MenuMultiError(errors)
	}

	return nil
}

// MenuMultiError is an error wrapping multiple validation errors returned by
// Menu.ValidateAll() if the designated constraints aren't met.
type MenuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuMultiError) AllErrors() []error { return m }

// MenuValidationError is the validation error returned by Menu.Validate if the
// designated constraints aren't met.
type MenuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuValidationError) ErrorName() string { return "MenuValidationError" }

// Error satisfies the builtin error interface
func (e MenuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuValidationError{}

// Validate checks the field values on ListMenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMenuResponseMultiError, or nil if none found.
func (m *ListMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMenuResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListMenuResponseMultiError(errors)
	}

	return nil
}

// ListMenuResponseMultiError is an error wrapping multiple validation errors
// returned by ListMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type ListMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMenuResponseMultiError) AllErrors() []error { return m }

// ListMenuResponseValidationError is the validation error returned by
// ListMenuResponse.Validate if the designated constraints aren't met.
type ListMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMenuResponseValidationError) ErrorName() string { return "ListMenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMenuResponseValidationError{}

// Validate checks the field values on GetMenuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetMenuRequestMultiError,
// or nil if none found.
func (m *GetMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetMenuRequestMultiError(errors)
	}

	return nil
}

// GetMenuRequestMultiError is an error wrapping multiple validation errors
// returned by GetMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuRequestMultiError) AllErrors() []error { return m }

// GetMenuRequestValidationError is the validation error returned by
// GetMenuRequest.Validate if the designated constraints aren't met.
type GetMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuRequestValidationError) ErrorName() string { return "GetMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuRequestValidationError{}

// Validate checks the field values on CreateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuRequestMultiError, or nil if none found.
func (m *CreateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMenuRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateMenuRequestMultiError(errors)
	}

	return nil
}

// CreateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by CreateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuRequestMultiError) AllErrors() []error { return m }

// CreateMenuRequestValidationError is the validation error returned by
// CreateMenuRequest.Validate if the designated constraints aren't met.
type CreateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuRequestValidationError) ErrorName() string {
	return "CreateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuRequestValidationError{}

// Validate checks the field values on UpdateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMenuRequestMultiError, or nil if none found.
func (m *UpdateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateMenuRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateMenuRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.AllowMissing != nil {
		// no validation rules for AllowMissing
	}

	if len(errors) > 0 {
		return UpdateMenuRequestMultiError(errors)
	}

	return nil
}

// UpdateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMenuRequestMultiError) AllErrors() []error { return m }

// UpdateMenuRequestValidationError is the validation error returned by
// UpdateMenuRequest.Validate if the designated constraints aren't met.
type UpdateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMenuRequestValidationError) ErrorName() string {
	return "UpdateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMenuRequestValidationError{}

// Validate checks the field values on DeleteMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMenuRequestMultiError, or nil if none found.
func (m *DeleteMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteMenuRequestMultiError(errors)
	}

	return nil
}

// DeleteMenuRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMenuRequestMultiError) AllErrors() []error { return m }

// DeleteMenuRequestValidationError is the validation error returned by
// DeleteMenuRequest.Validate if the designated constraints aren't met.
type DeleteMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMenuRequestValidationError) ErrorName() string {
	return "DeleteMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMenuRequestValidationError{}
