// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldUpdateBy holds the string denoting the update_by field in the database.
	FieldUpdateBy = "update_by"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldRealname holds the string denoting the realname field in the database.
	FieldRealname = "realname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldTelephone holds the string denoting the telephone field in the database.
	FieldTelephone = "telephone"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAuthority holds the string denoting the authority field in the database.
	FieldAuthority = "authority"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldPositionID holds the string denoting the position_id field in the database.
	FieldPositionID = "position_id"
	// FieldWorkID holds the string denoting the work_id field in the database.
	FieldWorkID = "work_id"
	// FieldRoles holds the string denoting the roles field in the database.
	FieldRoles = "roles"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateBy,
	FieldUpdateBy,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldRemark,
	FieldStatus,
	FieldTenantID,
	FieldUsername,
	FieldNickname,
	FieldRealname,
	FieldEmail,
	FieldMobile,
	FieldTelephone,
	FieldAvatar,
	FieldAddress,
	FieldRegion,
	FieldDescription,
	FieldGender,
	FieldAuthority,
	FieldLastLoginTime,
	FieldLastLoginIP,
	FieldOrgID,
	FieldPositionID,
	FieldWorkID,
	FieldRoles,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// TenantIDValidator is a validator for the "tenant_id" field. It is called by the builders before save.
	TenantIDValidator func(uint32) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// RealnameValidator is a validator for the "realname" field. It is called by the builders before save.
	RealnameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultMobile holds the default value on creation for the "mobile" field.
	DefaultMobile string
	// MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	MobileValidator func(string) error
	// DefaultTelephone holds the default value on creation for the "telephone" field.
	DefaultTelephone string
	// TelephoneValidator is a validator for the "telephone" field. It is called by the builders before save.
	TelephoneValidator func(string) error
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// DefaultRegion holds the default value on creation for the "region" field.
	DefaultRegion string
	// RegionValidator is a validator for the "region" field. It is called by the builders before save.
	RegionValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultLastLoginIP holds the default value on creation for the "last_login_ip" field.
	DefaultLastLoginIP string
	// LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	LastLoginIPValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusON is the default value of the Status enum.
const DefaultStatus = StatusON

// Status values.
const (
	StatusOFF Status = "OFF"
	StatusON  Status = "ON"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOFF, StatusON:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderSecret Gender = "SECRET"
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderSecret, GenderMale, GenderFemale:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for gender field: %q", ge)
	}
}

// Authority defines the type for the "authority" enum field.
type Authority string

// AuthorityCustomerUser is the default value of the Authority enum.
const DefaultAuthority = AuthorityCustomerUser

// Authority values.
const (
	AuthoritySysAdmin     Authority = "SYS_ADMIN"
	AuthorityTenantAdmin  Authority = "TENANT_ADMIN"
	AuthorityCustomerUser Authority = "CUSTOMER_USER"
	AuthorityGuest        Authority = "GUEST"
)

func (a Authority) String() string {
	return string(a)
}

// AuthorityValidator is a validator for the "authority" field enum values. It is called by the builders before save.
func AuthorityValidator(a Authority) error {
	switch a {
	case AuthoritySysAdmin, AuthorityTenantAdmin, AuthorityCustomerUser, AuthorityGuest:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for authority field: %q", a)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateBy orders the results by the create_by field.
func ByCreateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateBy, opts...).ToFunc()
}

// ByUpdateBy orders the results by the update_by field.
func ByUpdateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateBy, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByRealname orders the results by the realname field.
func ByRealname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealname, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByTelephone orders the results by the telephone field.
func ByTelephone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTelephone, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAuthority orders the results by the authority field.
func ByAuthority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthority, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByPositionID orders the results by the position_id field.
func ByPositionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPositionID, opts...).ToFunc()
}

// ByWorkID orders the results by the work_id field.
func ByWorkID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkID, opts...).ToFunc()
}
