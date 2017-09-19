// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: redis_meta.proto

/*
Package redis is a generated protocol buffer package.

It is generated from these files:
	redis_meta.proto

It has these top-level messages:
	IPAddr
	Slave
	Instance
	RawInstance
*/
package redis

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_IPAddr_IP = regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")

func (this *IPAddr) Validate() error {
	if !_regex_IPAddr_IP.MatchString(this.IP) {
		return github_com_mwitkow_go_proto_validators.FieldError("IP", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$"`, this.IP))
	}
	if !(this.Port > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '1'`, this.Port))
	}
	if !(this.Port < 65535) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be less than '65535'`, this.Port))
	}
	return nil
}
func (this *Slave) Validate() error {
	if this.Addr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Addr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Addr", err)
		}
	}
	return nil
}
func (this *Instance) Validate() error {
	if this.Master != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Master); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Master", err)
		}
	}
	for _, item := range this.Slaves {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Slaves", err)
			}
		}
	}
	return nil
}
func (this *RawInstance) Validate() error {
	if nil == this.Addr {
		return github_com_mwitkow_go_proto_validators.FieldError("Addr", fmt.Errorf("message must exist"))
	}
	if this.Addr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Addr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Addr", err)
		}
	}
	if !(this.Epoch > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Epoch", fmt.Errorf(`value '%v' must be greater than '0'`, this.Epoch))
	}
	if !(this.Sdowntime > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sdowntime", fmt.Errorf(`value '%v' must be greater than '0'`, this.Sdowntime))
	}
	if !(this.FailoverTimeout > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("FailoverTimeout", fmt.Errorf(`value '%v' must be greater than '0'`, this.FailoverTimeout))
	}
	return nil
}