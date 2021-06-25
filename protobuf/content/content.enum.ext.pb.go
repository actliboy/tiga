// Code generated by protoc-gen-enum. DO NOT EDIT.
// source: content/content.enum.proto

package content

import (
	"errors"
	"github.com/liov/tiga/utils/strings"
	"io"
)

func (x ContentType) String() string {
	switch x {
	case ContentPlaceholder:
		return "占位"
	case ContentMoment:
		return "瞬间"
	case ContentNote:
		return "笔记"
	case ContentDairy:
		return "日记"
	case ContentDairyBook:
		return "日记本"
	case ContentFavorites:
		return "收藏夹"
	case ContentCollection:
		return "收藏"
	case ContentComment:
		return "评论"
	}
	return ""
}

func (x ContentType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *ContentType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = ContentType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x AttrType) String() string {
	switch x {
	case AttrPlaceholder:
		return "占位"
	case AttrImage:
		return "图片"
	case AttrVideo:
		return "视频"
	case AttrAudio:
		return "音频"
	case AttrCover:
		return "封面"
	case AttrCategory:
		return "分类"
	case AttrTitle:
		return "标题"
	case AttrAbstract:
		return "摘要"
	case AttrIntro:
		return "介绍"
	case AttrContentType:
		return "文本类型"
	case AttrModifyTimes:
		return "修改次数"
	}
	return ""
}

func (x AttrType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *AttrType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = AttrType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x TagType) String() string {
	switch x {
	case TagPlaceholder:
		return "占位"
	case TagContent:
		return "内容"
	case TagMood:
		return "心情"
	case TagWeather:
		return "天气"
	case TagLocation:
		return "地点"
	case TagCategory:
		return "分类"
	case TagNotice:
		return "提示"
	}
	return ""
}

func (x TagType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *TagType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = TagType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x ViewPermission) String() string {
	switch x {
	case ViewPermissionPlaceholder:
		return "占位"
	case ViewPermissionAll:
		return "无限制"
	case ViewPermissionSelf:
		return "仅自己"
	case ViewPermissionHomePage:
		return "主页"
	case ViewPermissionStranger:
		return "陌生人"
	case ViewPermissionShield:
		return "屏蔽部分人"
	case ViewPermissionOpen:
		return "开放部分人"
	}
	return ""
}

func (x ViewPermission) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *ViewPermission) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = ViewPermission(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x MomentType) String() string {
	switch x {
	case MomentTypePlaceholder:
		return "占位"
	case MomentTypeImage:
		return "图片"
	case MomentTypeVideo:
		return "视频"
	case MomentTypeAudio:
		return "音频"
	}
	return ""
}

func (x MomentType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *MomentType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = MomentType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x ContainerType) String() string {
	switch x {
	case ContainerTypePlaceholder:
		return "占位"
	case ContainerTypeFavorites:
		return "收藏夹"
	case ContainerTypeDiaryBook:
		return "日记本"
	case ContainerTypeAlbum:
		return "专辑"
	case ContainerTypeCollection:
		return "合集"
	}
	return ""
}

func (x ContainerType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *ContainerType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = ContainerType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}