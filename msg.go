package workers

import (
	"reflect"

	"github.com/bitly/go-simplejson"
	"github.com/go-kit/kit/log/level"
)

type Args struct {
	*data
}

func NewMsg(content string) (*Msg, error) {
	if d, err := NewData(content); err != nil {
		return nil, err
	} else {
		return &Msg{d, content}, nil
	}
}

type Msg struct {
	*data
	original string
}

func (m *Msg) Jid() string {
	return m.Get("jid").MustString()
}

func (m *Msg) Args() *Args {
	if args, ok := m.CheckGet("args"); ok {
		return &Args{&data{args}}
	} else {
		d, _ := NewData("[]")
		return &Args{d}
	}
}

func (m *Msg) OriginalJson() string {
	return m.original
}

func NewData(content string) (*data, error) {
	if json, err := simplejson.NewJson([]byte(content)); err != nil {
		return nil, err
	} else {
		return &data{json}, nil
	}
}

type data struct {
	*simplejson.Json
}

func (d *data) ToJson() string {
	json, err := d.Encode()

	if err != nil {
		level.Error(Logger).Log(
			"msg", "failed to generate json",
			"err", err,
		)
	}

	return string(json)
}

func (d *data) Equals(other interface{}) bool {
	otherJson := reflect.ValueOf(other).MethodByName("ToJson").Call([]reflect.Value{})
	return d.ToJson() == otherJson[0].String()
}
