package requests

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func TestValidate(t *testing.T) {
	type args struct {
		c       *gin.Context
		obj     interface{}
		handler ValidatorFunc
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.c, tt.args.obj, tt.args.handler); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		data     interface{}
		rules    govalidator.MapData
		messages govalidator.MapData
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.data, tt.args.rules, tt.args.messages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
