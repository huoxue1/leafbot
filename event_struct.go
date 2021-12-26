package leafbot

import "github.com/huoxue1/leafBot/message"

type (
	anonymous struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	}

	Files struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Size    int64  `json:"size"`
		Busid   int64  `json:"busid"`
		FileUrl string `json:"url"`
	}

	Status struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
	}

	MessageIds struct {
		MessageID int32 `json:"message_id"`
	}

	Senders struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		NickName string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserId   int    `json:"user_id"`
	}

	Event struct {
		Anonymous     anonymous       `json:"anonymous"`
		Font          int             `json:"font"`
		GroupId       int             `json:"group_id"`
		Message       message.Message `json:"message"`
		MessageType   string          `json:"message_type"`
		PostType      string          `json:"post_type"`
		RawMessage    string          `json:"raw_message"`
		SelfId        int             `json:"self_id"`
		Sender        Senders         `json:"sender"`
		SubType       string          `json:"sub_type"`
		UserId        int             `json:"user_id"`
		Time          int             `json:"time"`
		NoticeType    string          `json:"notice_type"`
		RequestType   string          `json:"request_type"`
		Comment       string          `json:"comment"`
		Flag          string          `json:"flag"`
		OperatorId    int             `json:"operator_id"`
		File          Files           `json:"file"`
		Duration      int64           `json:"duration"`
		TargetId      int64           `json:"target_id"` //运气王id
		HonorType     string          `json:"honor_type"`
		MetaEventType string          `json:"meta_event_type"`
		Status        Status          `json:"status"`
		Interval      int             `json:"interval"`
		CardNew       string          `json:"card_new"` //新名片
		CardOld       string          `json:"card_old"` //旧名片
		MessageIds
	}

	MiraiEvent struct {
		Type   string `json:"type"`
		QQ     int64  `json:"qq"`
		Friend struct {
			Id       int64  `json:"id"`
			NickName string `json:"nick_name"`
			Remark   string `json:"remark"`
		} `json:"friend"`
	}
)

//GetMsg
/**
 * @Description:
 * @receiver e
 * @return message.Message
 */
func (e Event) GetMsg() message.Message {
	return e.Message
}

func (e Event) GetPlainText() string {
	content := ""
	for _, mes := range e.Message {
		if mes.Type == "text" {
			content += mes.Data["text"]
		}
	}
	return content
}

func (e Event) GetImages() []message.MessageSegment {
	var images []message.MessageSegment
	for _, mes := range e.Message {
		if mes.Type == "image" {
			images = append(images, mes)
		}
	}
	return images
}
