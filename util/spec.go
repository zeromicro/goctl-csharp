package util

import (
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func IsOptionalOrOmitEmpty(m spec.Member) bool {
	tag := m.Tags()
	for _, item := range tag {
		if stringx.Contains(item.Options, "optional") || stringx.Contains(item.Options, "omitempty") {
			return true
		}
	}
	return false
}
