package core

import "encoding/xml"

const (
	MsgTypeText string = "text"
)
//接收普通消息
//当普通微信用户向公众账号发消息时，微信服务器将POST消息的XML数据包到开发者填写的URL上。
type Message struct {
	XMLName xml.Name `xml:"xml"` // 指定最外层的标签为xml
	// 读取ToUserName配置项，并将结果保存到ToUserName变量中
	//开发者微信号
	ToUserName string `xml:"ToUserName"`
	//发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName"`

	//消息创建时间 （整型）
	CreateTime uint64 `xml:"CreateTime"`

	//消息类型，文本为text
	MsgType string `xml:"MsgType"`

	//消息id，64位整型
	MsgId int64 `xml:"MsgId"`

	//------------------------------------------------------------
	//文本消息内容(消息类型为text时)
	Content string `xml:"Content"`

	//------------------------------------------------------------
	//语音消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId int64 `xml:"MediaId"`
    //语音格式
	Format string `xml:"Format"`
	//------------------------------------------------------------
}

