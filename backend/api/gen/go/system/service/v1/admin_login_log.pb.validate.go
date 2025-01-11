// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/service/v1/admin_login_log.proto

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

// Validate checks the field values on AdminLoginLog with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AdminLoginLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminLoginLog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AdminLoginLogMultiError, or
// nil if none found.
func (m *AdminLoginLog) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminLoginLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.LoginIp != nil {
		// no validation rules for LoginIp
	}

	if m.LoginMac != nil {
		// no validation rules for LoginMac
	}

	if m.LoginTime != nil {

		if all {
			switch v := interface{}(m.GetLoginTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "LoginTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "LoginTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLoginTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminLoginLogValidationError{
					field:  "LoginTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.StatusCode != nil {
		// no validation rules for StatusCode
	}

	if m.Success != nil {
		// no validation rules for Success
	}

	if m.Reason != nil {
		// no validation rules for Reason
	}

	if m.Location != nil {
		// no validation rules for Location
	}

	if m.UserAgent != nil {
		// no validation rules for UserAgent
	}

	if m.BrowserName != nil {
		// no validation rules for BrowserName
	}

	if m.BrowserVersion != nil {
		// no validation rules for BrowserVersion
	}

	if m.ClientId != nil {
		// no validation rules for ClientId
	}

	if m.ClientName != nil {
		// no validation rules for ClientName
	}

	if m.OsName != nil {
		// no validation rules for OsName
	}

	if m.OsVersion != nil {
		// no validation rules for OsVersion
	}

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if m.UserName != nil {
		// no validation rules for UserName
	}

	if m.CreateTime != nil {

		if all {
			switch v := interface{}(m.GetCreateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminLoginLogValidationError{
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
					errors = append(errors, AdminLoginLogValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminLoginLogValidationError{
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
					errors = append(errors, AdminLoginLogValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminLoginLogValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeleteTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminLoginLogValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AdminLoginLogMultiError(errors)
	}

	return nil
}

// AdminLoginLogMultiError is an error wrapping multiple validation errors
// returned by AdminLoginLog.ValidateAll() if the designated constraints
// aren't met.
type AdminLoginLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminLoginLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminLoginLogMultiError) AllErrors() []error { return m }

// AdminLoginLogValidationError is the validation error returned by
// AdminLoginLog.Validate if the designated constraints aren't met.
type AdminLoginLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminLoginLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminLoginLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminLoginLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminLoginLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminLoginLogValidationError) ErrorName() string { return "AdminLoginLogValidationError" }

// Error satisfies the builtin error interface
func (e AdminLoginLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminLoginLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminLoginLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminLoginLogValidationError{}

// Validate checks the field values on ListAdminLoginLogResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAdminLoginLogResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAdminLoginLogResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAdminLoginLogResponseMultiError, or nil if none found.
func (m *ListAdminLoginLogResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAdminLoginLogResponse) validate(all bool) error {
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
					errors = append(errors, ListAdminLoginLogResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAdminLoginLogResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAdminLoginLogResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListAdminLoginLogResponseMultiError(errors)
	}

	return nil
}

// ListAdminLoginLogResponseMultiError is an error wrapping multiple validation
// errors returned by ListAdminLoginLogResponse.ValidateAll() if the
// designated constraints aren't met.
type ListAdminLoginLogResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAdminLoginLogResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAdminLoginLogResponseMultiError) AllErrors() []error { return m }

// ListAdminLoginLogResponseValidationError is the validation error returned by
// ListAdminLoginLogResponse.Validate if the designated constraints aren't met.
type ListAdminLoginLogResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAdminLoginLogResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAdminLoginLogResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAdminLoginLogResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAdminLoginLogResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAdminLoginLogResponseValidationError) ErrorName() string {
	return "ListAdminLoginLogResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAdminLoginLogResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAdminLoginLogResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAdminLoginLogResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAdminLoginLogResponseValidationError{}

// Validate checks the field values on GetAdminLoginLogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAdminLoginLogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAdminLoginLogRequestMultiError, or nil if none found.
func (m *GetAdminLoginLogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAdminLoginLogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetAdminLoginLogRequestMultiError(errors)
	}

	return nil
}

// GetAdminLoginLogRequestMultiError is an error wrapping multiple validation
// errors returned by GetAdminLoginLogRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAdminLoginLogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAdminLoginLogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAdminLoginLogRequestMultiError) AllErrors() []error { return m }

// GetAdminLoginLogRequestValidationError is the validation error returned by
// GetAdminLoginLogRequest.Validate if the designated constraints aren't met.
type GetAdminLoginLogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAdminLoginLogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAdminLoginLogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAdminLoginLogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAdminLoginLogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAdminLoginLogRequestValidationError) ErrorName() string {
	return "GetAdminLoginLogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAdminLoginLogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAdminLoginLogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAdminLoginLogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAdminLoginLogRequestValidationError{}

// Validate checks the field values on CreateAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAdminLoginLogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAdminLoginLogRequestMultiError, or nil if none found.
func (m *CreateAdminLoginLogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAdminLoginLogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAdminLoginLogRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAdminLoginLogRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAdminLoginLogRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreateAdminLoginLogRequestMultiError(errors)
	}

	return nil
}

// CreateAdminLoginLogRequestMultiError is an error wrapping multiple
// validation errors returned by CreateAdminLoginLogRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateAdminLoginLogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAdminLoginLogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAdminLoginLogRequestMultiError) AllErrors() []error { return m }

// CreateAdminLoginLogRequestValidationError is the validation error returned
// by CreateAdminLoginLogRequest.Validate if the designated constraints aren't met.
type CreateAdminLoginLogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAdminLoginLogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAdminLoginLogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAdminLoginLogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAdminLoginLogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAdminLoginLogRequestValidationError) ErrorName() string {
	return "CreateAdminLoginLogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAdminLoginLogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAdminLoginLogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAdminLoginLogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAdminLoginLogRequestValidationError{}

// Validate checks the field values on UpdateAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAdminLoginLogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAdminLoginLogRequestMultiError, or nil if none found.
func (m *UpdateAdminLoginLogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAdminLoginLogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateAdminLoginLogRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateAdminLoginLogRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateAdminLoginLogRequestValidationError{
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
				errors = append(errors, UpdateAdminLoginLogRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateAdminLoginLogRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateAdminLoginLogRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if m.AllowMissing != nil {
		// no validation rules for AllowMissing
	}

	if len(errors) > 0 {
		return UpdateAdminLoginLogRequestMultiError(errors)
	}

	return nil
}

// UpdateAdminLoginLogRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateAdminLoginLogRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateAdminLoginLogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAdminLoginLogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAdminLoginLogRequestMultiError) AllErrors() []error { return m }

// UpdateAdminLoginLogRequestValidationError is the validation error returned
// by UpdateAdminLoginLogRequest.Validate if the designated constraints aren't met.
type UpdateAdminLoginLogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAdminLoginLogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAdminLoginLogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAdminLoginLogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAdminLoginLogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAdminLoginLogRequestValidationError) ErrorName() string {
	return "UpdateAdminLoginLogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAdminLoginLogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAdminLoginLogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAdminLoginLogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAdminLoginLogRequestValidationError{}

// Validate checks the field values on DeleteAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAdminLoginLogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAdminLoginLogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAdminLoginLogRequestMultiError, or nil if none found.
func (m *DeleteAdminLoginLogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAdminLoginLogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteAdminLoginLogRequestMultiError(errors)
	}

	return nil
}

// DeleteAdminLoginLogRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteAdminLoginLogRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteAdminLoginLogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAdminLoginLogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAdminLoginLogRequestMultiError) AllErrors() []error { return m }

// DeleteAdminLoginLogRequestValidationError is the validation error returned
// by DeleteAdminLoginLogRequest.Validate if the designated constraints aren't met.
type DeleteAdminLoginLogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAdminLoginLogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAdminLoginLogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAdminLoginLogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAdminLoginLogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAdminLoginLogRequestValidationError) ErrorName() string {
	return "DeleteAdminLoginLogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAdminLoginLogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAdminLoginLogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAdminLoginLogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAdminLoginLogRequestValidationError{}
