// Code generated by protoc-gen-enum. DO NOT EDIT.
// source: content/action.enum.proto

package content

import (
	"errors"
	"github.com/liov/tiga/utils/strings"
	"io"
)

func (x ActionType) String() string {
	switch x {
	case ActionPlaceholder:
		return "占位"
	case ActionBrowse:
		return "浏览"
	case ActionLike:
		return "点赞"
	case ActionUnlike:
		return "不喜欢"
	case ActionComment:
		return "评论"
	case ActionCollect:
		return "收藏"
	case ActionReport:
		return "举报"
	case ActionGiveAction:
		return "回馈"
	case ActionApprove:
		return "赞同"
	case ActionDelete:
		return "删除"
	}
	return ""
}

func (x ActionType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *ActionType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = ActionType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x DelReason) String() string {
	switch x {
	case DelReasonPlaceholder:
		return "占位"
	case DelReasonViolationOfLawsAndRegulations:
		return "违返法律法规"
	case DelReasonEroticViolence:
		return "色情暴力"
	case DelReasonOther:
		return "其他原因"
	}
	return ""
}

func (x DelReason) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *DelReason) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = DelReason(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}

func (x CommentType) String() string {
	switch x {
	case CommentPlaceholder:
		return "占位"
	case CommentMoment:
		return "瞬间"
	case CommentDiary:
		return "日记"
	case CommentDiaryBook:
		return "日记本"
	case CommentArticle:
		return "文章"
	}
	return ""
}

func (x CommentType) MarshalGQL(w io.Writer) {
	w.Write(stringsi.QuoteToBytes(x.String()))
}

func (x *CommentType) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(uint32); ok {
		*x = CommentType(i)
		return nil
	}
	return errors.New("枚举值需要数字类型")
}
