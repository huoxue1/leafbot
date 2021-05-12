package message

import (
	"github.com/tidwall/gjson"
	"reflect"
	"unsafe"
)

// Modified from https://github.com/catsworld/qq-bot-api

// ParseMessage parses msg, which might have 2 types, string or array,
// depending on the configuration of cqhttp, to a Message.
// msg is the value of key "message" of the data unmarshalled from the
// API response JSON.
func ParseMessage(msg []byte) Message {
	x := gjson.Parse(BytesToString(msg))
	if x.IsArray() {
		return ParseMessageFromArray(x)
	} else {
		return ParseMessageFromString(x.String())
	}
}

// ParseMessageFromArray parses msg as type array to a Message.
// msg is the value of key "message" of the data unmarshalled from the
// API response JSON.
// ParseMessageFromArray cq字符串转化为json对象
func ParseMessageFromArray(msgs gjson.Result) Message {
	message := Message{}
	parse2map := func(val gjson.Result) map[string]string {
		m := map[string]string{}
		val.ForEach(func(key, value gjson.Result) bool {
			m[key.String()] = value.String()
			return true
		})
		return m
	}
	msgs.ForEach(func(_, item gjson.Result) bool {
		message = append(message, MessageSegment{
			Type: item.Get("type").String(),
			Data: parse2map(item.Get("data")),
		})
		return true
	})
	return message
}

// CQString returns the CQEncoded string. All media in the message will be converted
// to its CQCode.
// CQString 解码cq字符串
func (m Message) CQString() string {
	str := ""
	for _, media := range m {
		if media.Type != "text" {
			str += media.CQCode()
		} else {
			str += EscapeCQText(media.Data["text"])
		}
	}
	return str
}

// ExtractPlainText 提取消息中的纯文本
func (m Message) ExtractPlainText() string {
	msg := ""
	for _, val := range m {
		if val.Type == "text" {
			msg += val.Data["text"]
		}
	}
	return msg
}

// BytesToString 没有内存开销的转换
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 没有内存开销的转换
func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s)) // nolint
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
