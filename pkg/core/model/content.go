package model

import "content-core/pkg/common/util"

type ContentType string

const (
	ContentTypeCOMMENT ContentType = "COMMENT"
	ContentTypeARITCLE ContentType = "QUESTION"
	ContentTypeUnknown ContentType = "UNKNOWN"
)

func ContentTypeFromString(name string) (ContentType, error) {
	switch name {
	case ContentTypeARITCLE.String():
		return ContentTypeARITCLE, nil
	case ContentTypeCOMMENT.String():
		return ContentTypeCOMMENT, nil
	}
	return ContentTypeUnknown, util.BadContentType
}

func (t ContentType) String() string{
	return string(t)
}