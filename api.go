package leafBot

import (
	"encoding/json"
	message2 "github.com/3343780376/leafBot/message"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var (
	apiResChan = make(chan []byte, 1)
	eventChan  = make(chan []byte)
)

type response struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Echo    string `json:"echo"`
}

type (
	ConstNoticeType struct {
		GroupUpload   string //群文件上传
		GroupAdmin    string //群管理员变动
		GroupDecrease string //群成员减少
		GroupIncrease string //群成员增加
		GroupBan      string //群禁言
		FriendAdd     string //好友添加
		GroupRecall   string //群消息撤回
		FriendRecall  string //好友消息撤回
		Notify        string //群内戳一戳,群红包运气王,群成员荣誉变更,好友戳一戳
		GroupCard     string //群成员名片更新
		OfflineFile   string //接收到离线文件
	}
	ConstMessageType struct {
		Group   string
		Private string
	}
	//ConstMessageSubType struct {
	//	Friend    string //好友消息
	//	Group     string //临时会话
	//	Other     string //其他消息
	//	Normal    string //正常群消息
	//	Anonymous string //匿名消息
	//	Notice    string //群通知消息
	//}
)

var (

	//MessageSubTypeApi = ConstMessageSubType{
	//	Friend: "friend",
	//	Group: "group",
	//	Other: "other",
	//	Normal: "normal",
	//	Anonymous: "anonymous",
	//	Notice: "notice"}
	NoticeTypeApi = ConstNoticeType{
		GroupUpload:   "group_upload",
		GroupAdmin:    "group_admin",
		GroupDecrease: "group_decrease",
		GroupIncrease: "group_increase",
		GroupBan:      "group_ban",
		FriendAdd:     "friend_add",
		GroupRecall:   "group_recall",
		FriendRecall:  "friend_recall",
		Notify:        "notify",
		OfflineFile:   "offline_file",
		GroupCard:     "group_card"}

	MessageTypeApi = ConstMessageType{
		Group:   "group",
		Private: "private"}
)

type MessageIds struct {
	MessageId int32 `json:"message_id"`
}

type LoginInfo struct {
	UserId   int    `json:"user_id"`
	NickName string `json:"nick_name"`
}

type Status struct {
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

type Event struct {
	Anonymous     anonymous `json:"anonymous"`
	Font          int       `json:"font"`
	GroupId       int       `json:"group_id"`
	Message       string    `json:"message"`
	MessageType   string    `json:"message_type"`
	PostType      string    `json:"post_type"`
	RawMessage    string    `json:"raw_message"`
	SelfId        int       `json:"self_id"`
	Sender        Senders   `json:"sender"`
	SubType       string    `json:"sub_type"`
	UserId        int       `json:"user_id"`
	Time          int       `json:"time"`
	NoticeType    string    `json:"notice_type"`
	RequestType   string    `json:"request_type"`
	Comment       string    `json:"comment"`
	Flag          string    `json:"flag"`
	OperatorId    int       `json:"operator_id"`
	File          Files     `json:"file"`
	Duration      int64     `json:"duration"`
	TargetId      int64     `json:"target_id"` //运气王id
	HonorType     string    `json:"honor_type"`
	MetaEventType string    `json:"meta_event_type"`
	Status        Status    `json:"status"`
	Interval      int       `json:"interval"`
	CardNew       string    `json:"card_new"` //新名片
	CardOld       string    `json:"card_old"` //旧名片
	MessageIds
}

type Senders struct {
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

type responseJson struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
}

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

	GroupMemberInfo struct {
		GroupId         int  `json:"group_id"`
		JoinTime        int  `json:"join_time"`
		LastSentTime    int  `json:"last_sent_time"`
		Unfriendly      bool `json:"unfriendly"`
		TitleExpireTime int  `json:"title_expire_time"`
		CardChangeable  bool `json:"card_changeable"`
		Senders
	}

	getMsgJson struct {
		responseJson
		Data GetMessage `json:"data"`
	}

	GetMessage struct {
		Time      int32   `json:"time"`
		Group     bool    `json:"group"`
		MessageId int32   `json:"message_id"`
		RealId    int32   `json:"real_id"`
		Sender    Senders `json:"sender"`
		Message   string  `json:"message"`
	}

	FriendList struct {
		UserId   int    `json:"user_id"`
		NickName string `json:"nick_name"`
		Remark   string `json:"remark"`
	}

	GroupInfo struct {
		GroupId        int    `json:"group_id"`
		GroupName      string `json:"group_name"`
		MemberCount    int    `json:"member_count"`
		MaxMemberCount int    `json:"max_member_count"`
	}

	GroupHonorInfo struct {
		GroupId          int                  `json:"group_id"`
		CurrentTalkative CurrentTalkativeS    `json:"current_talkative"`
		TalkativeList    []GroupHonorInfoList `json:"talkative_list"`
		PerformerList    []GroupHonorInfoList `json:"performer_list"`
		LegendList       []GroupHonorInfoList `json:"legend_list"`
		StrongNewbieList []GroupHonorInfoList `json:"strong_newbie_list"`
		EmotionList      []GroupHonorInfoList `json:"emotion_list"`
	}

	CurrentTalkativeS struct {
		UserId   int    `json:"user_id"`
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
		DayCount int    `json:"day_count"`
	}

	GroupHonorInfoList struct {
		UserId      int    `json:"user_id"`
		NickName    string `json:"nick_name"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	}

	Record struct {
		File      string `json:"file"`
		OutFormat string `json:"out_format"`
	}

	Cookie struct {
		Cookies string `json:"cookies"`
	}

	CsrfToken struct {
		Token string `json:"token"`
	}

	OnlineStatus struct {
		Online bool `json:"online"`
		Good   bool `json:"good"`
	}

	Bool struct {
		Yes bool `json:"yes"`
	}

	Image struct {
		File string `json:"file"`
	}

	Credentials struct {
		Cookie
		CsrfToken
	}
	Content struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}
	Node struct {
		Id      int     `json:"id"`
		Name    string  `json:"name"`
		Uin     int     `json:"uin"`
		Content Content `json:"content"`
	}
)

type (
	responseGroupListJson struct {
		response
		Data []GroupInfo `json:"data"`
	}

	responseStrangerInfoJson struct {
		response
		Data Senders `json:"data"`
	}

	responseGroupInfoJson struct {
		response
		Data GroupInfo `json:"data"`
	}

	responseFriendListJson struct {
		response
		Data []FriendList `json:"data"`
	}

	responseMessageJson struct {
		response
		Data MessageIds `json:"data"`
	}

	responseLogoInfoJson struct {
		response
		Data LoginInfo `json:"data"`
	}

	responseGroupMemberInfoJson struct {
		response
		Data GroupMemberInfo `json:"data"`
	}
	responseGroupMemberListJson struct {
		response
		Data []GroupMemberInfo `json:"data"`
	}
	responseGroupHonorInfoJson struct {
		response
		Data GroupHonorInfo `json:"data"`
	}
	responseCookiesJson struct {
		response
		Data Cookie `json:"data"`
	}
	responseCsrfTokenJson struct {
		response
		Data CsrfToken `json:"data"`
	}
	responseCredentialsJson struct {
		response
		Data Credentials `json:"data"`
	}
	responseRecordJson struct {
		response
		Data Record `json:"data"`
	}
	responseImageJson struct {
		response
		Data Image `json:"data"`
	}
	responseCanSendJson struct {
		response
		Data Bool `json:"data"`
	}
	responseOnlineStatus struct {
		response
		Data OnlineStatus `json:"data"`
	}
	responseMsgDataJson struct {
		response
		Data MsgData `json:"data"`
	}
	responseForwardMsgJson struct {
		response
		Data []ForwardMsg `json:"data"`
	}
	responseWordSliceJson struct {
		response
		Data []string `json:"data"`
	}
	responseOcrImageJson struct {
		response
		Data OcrImage `json:"data"`
	}
	responseGroupSystemMsgJson struct {
		response
		Data GroupSystemMsg `json:"data"`
	}
	responseGroupFileSystemInfoJson struct {
		response
		Data GroupFileSystemInfo `json:"data"`
	}
	responseGroupRootFilesJson struct {
		response
		Data GroupRootFiles `json:"data"`
	}
	responseGroupFilesByFolderJson struct {
		response
		Data GroupFilesByFolder `json:"data"`
	}
	responseGroupFileUrlJson struct {
		response
		Data FileUrl `json:"data"`
	}
	responseGroupAtAllRemainJson struct {
		response
		Data GroupAtAllRemain `json:"data"`
	}
	responseDownloadFilePathJson struct {
		response
		Data DownloadFilePath `json:"data"`
	}
	responseMessageHistoryJson struct {
		response
		Data MessageHistory `json:"data"`
	}
	responseOnlineClientsJson struct {
		response
		Data Clients `json:"data"`
	}
	responseVipInfoJson struct {
		response
		Data VipInfo `json:"data"`
	}
)

type (
	FileUrl struct {
		Url string `json:"url"`
	}
	MsgData struct {
		MessageId int     `json:"message_id"`
		RealId    int     `json:"real_id"`
		Sender    Senders `json:"sender"`
		Time      int     `json:"time"`
		Message   string  `json:"message"`
	}
	ForwardMsg struct {
		Content string  `json:"content"`
		Sender  Senders `json:"sender"`
		Time    int     `json:"time"`
	}
	TextDetection struct {
		Text        string      `json:"text"`
		Confidence  int         `json:"confidence"`
		Coordinates interface{} `json:"coordinates"`
	}
	OcrImage struct {
		Texts    []TextDetection `json:"texts"`
		Language string          `json:"language"`
	}
	InvitedRequest struct {
		RequestId   int    `json:"request_id"`   //请求ID
		InvitorUin  int    `json:"invitor_uin"`  //邀请者
		InvitorNick string `json:"invitor_nick"` //邀请者昵称
		GroupId     int    `json:"group_id"`     //群号
		GroupName   string `json:"group_name"`   //群名
		Checked     bool   `json:"checked"`      //是否已被处理
		Actor       int64  `json:"actor"`        //处理者, 未处理为0
	}
	JoinRequest struct {
		RequestId     int    `json:"request_id"`     //请求ID
		RequesterUin  int    `json:"requester_uin"`  //请求者ID
		RequesterNick string `json:"requester_nick"` //请求者昵称
		Message       string `json:"message"`        //验证消息
		GroupId       int    `json:"group_id"`       //群号
		GroupName     string `json:"group_name"`     //群名
		Checked       bool   `json:"checked"`        //是否已被处理
		Actor         int    `json:"actor"`          //处理者, 未处理为0
	}
	GroupSystemMsg struct {
		InvitedRequests []InvitedRequest `json:"invited_requests"` //邀请消息列表
		JoinRequests    []JoinRequest    `json:"join_requests"`    //进群消息列表
	}
	GroupFileSystemInfo struct {
		FileCount  int `json:"file_count"`  //文件总数
		LimitCount int `json:"limit_count"` //文件上限
		UsedSpace  int `json:"used_space"`  //已使用空间
		TotalSpace int `json:"total_space"` //空间上限
	}
	File struct {
		FileId        string `json:"file_id"`        //文件ID
		FileName      string `json:"file_name"`      //文件名
		Busid         int    `json:"busid"`          //文件类型
		FileSize      int64  `json:"file_size"`      //文件大小
		UploadTime    int64  `json:"upload_time"`    //上传时间
		DeadTime      int64  `json:"dead_time"`      //过期时间,永久文件恒为0
		ModifyTime    int64  `json:"modify_time"`    //最后修改时间
		DownloadTimes int32  `json:"download_times"` //下载次数
		Uploader      int64  `json:"uploader"`       //上传者ID
		UploaderName  string `json:"uploader_name"`  //上传者名字
	}
	Folder struct {
		FolderId       string `json:"folder_id"`        //文件夹ID
		FolderName     string `json:"folder_name"`      //文件名
		CreateTime     int    `json:"create_time"`      //创建时间
		Creator        int    `json:"creator"`          //创建者
		CreatorName    string `json:"creator_name"`     //创建者名字
		TotalFileCount int32  `json:"total_file_count"` //子文件数量
	}
	GroupRootFiles struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	}
	GroupFilesByFolder struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	}
	GroupAtAllRemain struct {
		CanAtAll                 bool `json:"can_at_all"`                    //是否可以@全体成员
		RemainAtAllCountForGroup int  `json:"remain_at_all_count_for_group"` //群内所有管理当天剩余@全体成员次数
		RemainAtAllCountForUin   int  `json:"remain_at_all_count_for_uin"`   //BOT当天剩余@全体成员次数
	}
	DownloadFilePath struct {
		File string `json:"file"`
	}
	MessageHistory struct {
		Messages []string `json:"messages"`
	}
	Clients struct {
		Clients []Device `json:"clients"` //在线客户端列表
	}
	Device struct {
		AppId      int64  `json:"app_id"`      //客户端ID
		DeviceName string `json:"device_name"` //设备名称
		DeviceKind string `json:"device_kind"` //设备类型
	}
	VipInfo struct {
		UserId         int64   `json:"user_id"`          //QQ 号
		Nickname       string  `json:"nickname"`         //用户昵称
		Level          int64   `json:"level"`            //QQ 等级
		LevelSpeed     float64 `json:"level_speed"`      //等级加速度
		VipLevel       string  `json:"vip_level"`        //会员等级
		VipGrowthSpeed int64   `json:"vip_growth_speed"` //会员成长速度
		VipGrowthTotal int64   `json:"vip_growth_total"` //会员成长总值
	}
)

type QuickUseApi interface {
	Send(event Event, message string) int
	SendAt(event Event, message string) int
}

//go-cqhttp新增api
type IncreaseApi interface {
	DownloadFile(url string, threadCount int, headers []string) DownloadFilePath
	GetGroupMsgHistory(messageSeq int64, groupId int) MessageHistory
	GetOnlineClients(noCache bool) Clients
	GetVipInfoTest(UserId int) VipInfo
	SendGroupNotice(groupId int, content string)
	ReloadEventFilter()
	UploadGroupFile(groupId int, file string, name string, folder string) error
}

type SpecialApi interface {
	SetGroupNameSpecial(groupId int, groupName string)
	SetGroupPortrait(groupId int, file string, cache int)
	GetMsgSpecial(messageId int) MsgData
	GetForwardMsg(messageId int) []ForwardMsg
	SendGroupForwardMsg(groupId int, messages []Node)
	GetWordSlices(content string) []string
	OcrImage(image string) OcrImage
	GetGroupSystemMsg() GroupSystemMsg
	GetGroupFileSystemInfo(groupId int) GroupFileSystemInfo
	GetGroupRootFiles(groupId int) GroupRootFiles
	GetGroupFilesByFolder(groupId int, folderId string) GroupFilesByFolder
	GetGroupFileUrl(groupId int, fileId string, busid int) FileUrl
	GetGroupAtAllRemain(groupId int) GroupAtAllRemain
}

type Api interface {
	SendGroupMsg(groupId int, message string, autoEscape bool) int32
	SendPrivateMsg(userId int, message string, autoEscape bool) int32
	DeleteMsg(messageId int32)
	GetMsg(messageId int32) GetMessage
	SetGroupBan(groupId int, userId int, duration int)
	SetGroupCard(groupId int, userId int, card string)
	SendMsg(messageType string, userId int, groupId int, message string, autoEscape bool) int32
	SendLike(userId int, times int)
	SetGroupKick(groupId int, userId int, rejectAddRequest bool)
	SetGroupAnonymousBan(groupId int, flag string, duration int)
	SetGroupWholeBan(groupId int, enable bool)
	SetGroupAdmin(groupId int, UserId int, enable bool)
	SetGroupAnonymous(groupId int, enable bool)
	SetGroupName(groupId int, groupName string)
	SetGroupLeave(groupId int, isDisMiss bool)
	SetGroupSpecialTitle(groupId int, userId int, specialTitle string, duration int)
	SetFriendAddRequest(flag string, approve bool, remark string)
	SetGroupAddRequest(flag string, subType string, approve bool, reason string)
	GetLoginInfo() LoginInfo
	GetStrangerInfo(userId int, noCache bool) Senders
	GetFriendList() []FriendList
	GetGroupInfo(groupId int, noCache bool) GroupInfo
	GetGroupList() []GroupInfo
	GetGroupMemberInfo(groupId int, UserId int, noCache bool) GroupMemberInfo
	GetGroupMemberList(groupId int) []GroupMemberInfo
	GetGroupHonorInfo(groupId int, honorType string) GroupHonorInfo
	GetCookies(domain string) Cookie
	GetCsrfToken() CsrfToken
	GetCredentials(domain string) Credentials
	GetRecord(file, outFormat string) Record
	GetImage(file string) Image
	CanSendImage() Bool
	CanSendRecord() Bool
	GetStatus() OnlineStatus
	SetRestart(delay int)
	CleanCache()
}

type UseApi struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

// Send
/*
   @Description:
   @receiver b
   @param event Event
   @param message string
   @return int
*/
func (b *Bot) Send(event Event, message string) int {
	msgId := b.SendMsg(event.MessageType, event.UserId, event.GroupId, message, false)
	return int(msgId)
}

// SendAt
/*
   @Description:
   @receiver b
   @param event Event
   @param message string
   @return int
*/
func (b *Bot) SendAt(event Event, message string) int {
	msgId := b.SendMsg(event.MessageType, event.UserId, event.GroupId, message+message2.At(event.UserId), false)
	return int(msgId)
}

// SendGroupMsg
/*
   @Description:
   @receiver b
   @param groupId int
   @param message string
   @param autoEscape bool
   @return int32
*/
func (b *Bot) SendGroupMsg(groupId int, message string, autoEscape bool) int32 {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId    int    `json:"group_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	}
	var d = UseApi{
		Action: "send_group_msg",
		Params: param{
			GroupId:    groupId,
			Message:    message,
			AutoEscape: autoEscape,
		},
		Echo: echo}

	b.Client.OutChan <- d
	var responseMsgJson responseMessageJson
	data := getResponse(&response{}, echo)
	_ = json.Unmarshal(data, &responseMsgJson)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return responseMsgJson.Data.MessageId
}

// SendPrivateMsg
/*
   @Description:
   @receiver b
   @param userId int
   @param message string
   @param autoEscape bool
   @return int32
*/
func (b *Bot) SendPrivateMsg(userId int, message string, autoEscape bool) int32 {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		UserId     int    `json:"user_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	}
	var d = UseApi{
		Action: "send_private_msg",
		Params: param{
			UserId:     userId,
			Message:    message,
			AutoEscape: autoEscape,
		},
		Echo: echo}

	b.Client.OutChan <- d
	var responseMsgJson responseMessageJson
	data := getResponse(&response{}, echo)
	_ = json.Unmarshal(data, &responseMsgJson)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return responseMsgJson.Data.MessageId
}

// DeleteMsg
/*
   @Description:
   @receiver b
   @param messageId int32
*/
func (b *Bot) DeleteMsg(messageId int32) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		MessageId int32 `json:"message_id"`
	}
	var d = UseApi{
		Action: "delete_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// GetMsg
/*
   @Description:
   @receiver b
   @param messageId int32
   @return GetMessage
*/
func (b *Bot) GetMsg(messageId int32) GetMessage {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		MessageId int32 `json:"message_id"`
	}
	var d = UseApi{
		Action: "get_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo}

	b.Client.OutChan <- d
	var defaultJson getMsgJson
	data := getResponse(&response{}, echo)
	_ = json.Unmarshal(data, &defaultJson)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return defaultJson.Data
}

// SetGroupBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param duration int
*/
func (b *Bot) SetGroupBan(groupId int, userId int, duration int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId  int `json:"group_id"`
		UserId   int `json:"user_id"`
		Duration int `json:"duration"`
	}
	var d = UseApi{
		Action: "set_group_ban",
		Params: param{
			GroupId:  groupId,
			UserId:   userId,
			Duration: duration,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupCard
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param card string
*/
func (b *Bot) SetGroupCard(groupId int, userId int, card string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int    `json:"group_id"`
		UserId  int    `json:"user_id"`
		Card    string `json:"card"`
	}
	var d = UseApi{
		Action: "set_group_card",
		Params: param{
			GroupId: groupId,
			UserId:  userId,
			Card:    card,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SendMsg
/*
   @Description:
   @receiver b
   @param messageType string
   @param userId int
   @param groupId int
   @param message string
   @param autoEscape bool
   @return int32
*/
func (b *Bot) SendMsg(messageType string, userId int, groupId int, message string, autoEscape bool) int32 {
	if messageType == "group" {
		return b.SendGroupMsg(groupId, message, autoEscape)
	} else {
		return b.SendPrivateMsg(userId, message, autoEscape)
	}
}

// SendLike
/*
   @Description:
   @receiver b
   @param userId int
   @param times int
*/
func (b *Bot) SendLike(userId int, times int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		UserId int `json:"user_id"`
		Times  int `json:"times"`
	}
	var d = UseApi{
		Action: "send_like",
		Params: param{
			UserId: userId,
			Times:  times,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupKick
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param rejectAddRequest bool
*/
func (b *Bot) SetGroupKick(groupId int, userId int, rejectAddRequest bool) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId          int  `json:"group_id"`
		UserId           int  `json:"user_id"`
		RejectAddRequest bool `json:"reject_add_request"`
	}
	var d = UseApi{
		Action: "set_group_kick",
		Params: param{
			GroupId:          groupId,
			UserId:           userId,
			RejectAddRequest: rejectAddRequest,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupAnonymousBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param flag string
   @param duration int
*/
func (b *Bot) SetGroupAnonymousBan(groupId int, flag string, duration int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId  int    `json:"group_id"`
		Flag     string `json:"flag"`
		Duration int    `json:"duration"`
	}
	var d = UseApi{
		Action: "set_group_anonymous_ban",
		Params: param{
			GroupId:  groupId,
			Flag:     flag,
			Duration: duration,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupWholeBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param enable bool
*/
func (b *Bot) SetGroupWholeBan(groupId int, enable bool) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int  `json:"group_id"`
		Enable  bool `json:"enable"`
	}
	var d = UseApi{
		Action: "set_group_whole_ban",
		Params: param{
			GroupId: groupId,

			Enable: enable,
		},
		Echo: echo}

	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupAdmin
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param enable bool
*/
func (b *Bot) SetGroupAdmin(groupId int, userId int, enable bool) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int  `json:"group_id"`
		UserId  int  `json:"user_id"`
		Enable  bool `json:"enable"`
	}
	var d = UseApi{
		Action: "set_group_admin",
		Params: param{
			GroupId: groupId,
			UserId:  userId,
			Enable:  enable,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupAnonymous
/*
   @Description:
   @receiver b
   @param groupId int
   @param enable bool
*/
func (b *Bot) SetGroupAnonymous(groupId int, enable bool) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int  `json:"group_id"`
		Enable  bool `json:"enable"`
	}
	var d = UseApi{
		Action: "set_group_anonymous",
		Params: param{
			GroupId: groupId,
			Enable:  enable,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupName
/*
   @Description:
   @receiver b
   @param groupId int
   @param groupName string
*/
func (b *Bot) SetGroupName(groupId int, groupName string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId   int    `json:"group_id"`
		GroupName string `json:"group_name"`
	}
	var d = UseApi{
		Action: "set_group_name",
		Params: param{
			GroupId:   groupId,
			GroupName: groupName,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupLeave
/*
   @Description:
   @receiver b
   @param groupId int
   @param isDisMiss bool
*/
func (b *Bot) SetGroupLeave(groupId int, isDisMiss bool) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId   int  `json:"group_id"`
		IsDisMiss bool `json:"is_dismiss"`
	}
	var d = UseApi{
		Action: "set_group_leave",
		Params: param{
			GroupId:   groupId,
			IsDisMiss: isDisMiss,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupSpecialTitle
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param specialTitle string
   @param duration int
*/
func (b *Bot) SetGroupSpecialTitle(groupId int, userId int, specialTitle string, duration int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId      int    `json:"group_id"`
		UserId       int    `json:"user_id"`
		SpecialTitle string `json:"special_title"`
		Duration     int    `json:"duration"`
	}
	var d = UseApi{
		Action: "set_group_special_title",
		Params: param{
			GroupId:      groupId,
			UserId:       userId,
			SpecialTitle: specialTitle,
			Duration:     duration,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetFriendAddRequest
/*
   @Description:
   @receiver b
   @param flag string
   @param approve bool
   @param remark string
*/
func (b *Bot) SetFriendAddRequest(flag string, approve bool, remark string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Flag    string `json:"flag"`
		Approve bool   `json:"approve"`
		Remark  string `json:"remark"`
	}
	var d = UseApi{
		Action: "set_group_add_request",
		Params: param{
			Flag:    flag,
			Approve: approve,
			Remark:  remark,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupAddRequest
/*
   @Description:
   @receiver b
   @param flag string
   @param subType string
   @param approve bool
   @param reason string
*/
func (b *Bot) SetGroupAddRequest(flag string, subType string, approve bool, reason string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Flag    string `json:"flag"`
		SubType string `json:"sub_type"`
		Approve bool   `json:"approve"`
		Reason  string `json:"reason"`
	}
	var d = UseApi{
		Action: "set_group_add_request",
		Params: param{
			Flag:    flag,
			SubType: subType,
			Approve: approve,
			Reason:  reason,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// GetLoginInfo
/*
   @Description:
   @receiver b
   @return LoginInfo
*/
func (b *Bot) GetLoginInfo() LoginInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_login_info",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseLogoInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetStrangerInfo
/*
   @Description:
   @receiver b
   @param userId int
   @param noCache bool
   @return Senders
*/
func (b *Bot) GetStrangerInfo(userId int, noCache bool) Senders {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		UserId  int  `json:"user_id"`
		NoCache bool `json:"no_cache"`
	}
	var d = UseApi{
		Action: "get_stranger_info",
		Params: param{
			UserId:  userId,
			NoCache: noCache,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseStrangerInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetFriendList
/*
   @Description:
   @receiver b
   @return []FriendList
*/
func (b *Bot) GetFriendList() []FriendList {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_friend_list",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseFriendListJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @param noCache bool
   @return GroupInfo
*/
func (b *Bot) GetGroupInfo(groupId int, noCache bool) GroupInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int  `json:"group_id"`
		NoCache bool `json:"no_cache"`
	}
	var d = UseApi{
		Action: "get_group_info",
		Params: param{
			GroupId: groupId,
			NoCache: noCache,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupList
/*
   @Description:
   @receiver b
   @return []GroupInfo
*/
func (b *Bot) GetGroupList() []GroupInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_group_list",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupListJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupMemberInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param noCache bool
   @return GroupMemberInfo
*/
func (b *Bot) GetGroupMemberInfo(groupId int, userId int, noCache bool) GroupMemberInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int  `json:"group_id"`
		UserId  int  `json:"user_id"`
		NoCache bool `json:"no_cache"`
	}
	var d = UseApi{
		Action: "get_group_member_info",
		Params: param{
			GroupId: groupId,
			UserId:  userId,
			NoCache: noCache,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupMemberInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupMemberList
/*
   @Description:
   @receiver b
   @param groupId int
   @return []GroupMemberInfo
*/
func (b *Bot) GetGroupMemberList(groupId int) []GroupMemberInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int `json:"group_id"`
	}
	var d = UseApi{
		Action: "get_group_member_list",
		Params: param{
			GroupId: groupId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupMemberListJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupHonorInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @param honorType string
   @return GroupHonorInfo
*/
func (b *Bot) GetGroupHonorInfo(groupId int, honorType string) GroupHonorInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId   int    `json:"group_id"`
		HonorType string `json:"honor_type"`
	}
	var d = UseApi{
		Action: "get_stranger_info",
		Params: param{
			GroupId:   groupId,
			HonorType: honorType,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupHonorInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetCookies
/*
   @Description:
   @receiver b
   @param domain string
   @return Cookie
*/
func (b *Bot) GetCookies(domain string) Cookie {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Domain string `json:"domain"`
	}
	var d = UseApi{
		Action: "get_cookies",
		Params: param{
			Domain: domain,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseCookiesJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetCsrfToken
/*
   @Description:
   @receiver b
   @return CsrfToken
*/
func (b *Bot) GetCsrfToken() CsrfToken {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_csrf_token",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseCsrfTokenJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetCredentials
/*
   @Description:
   @receiver b
   @param domain string
   @return Credentials
*/
func (b *Bot) GetCredentials(domain string) Credentials {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Domain string `json:"domain"`
	}
	var d = UseApi{
		Action: "get_credentials",
		Params: param{
			Domain: domain,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseCredentialsJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetRecord
/*
   @Description:
   @receiver b
   @param file file
   @param outFormat string
   @return Record
*/
func (b *Bot) GetRecord(file, outFormat string) Record {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		File      string `json:"file"`
		OutFormat string `json:"out_format"`
	}
	var d = UseApi{
		Action: "get_Record",
		Params: param{
			File:      file,
			OutFormat: outFormat,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseRecordJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetImage
/*
   @Description:
   @receiver b
   @param file string
   @return Image
*/
func (b *Bot) GetImage(file string) Image {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		File string `json:"file"`
	}
	var d = UseApi{
		Action: "get_Record",
		Params: param{
			File: file,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseImageJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// CanSendImage
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendImage() Bool {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "can_send_image",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseCanSendJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// CanSendRecord
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendRecord() Bool {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "can_send_record",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseCanSendJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetStatus
/*
   @Description:
   @receiver b
   @return OnlineStatus
*/
func (b *Bot) GetStatus() OnlineStatus {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_status",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseOnlineStatus
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// SetRestart
/*
   @Description:
   @receiver b
   @param delay int
*/
func (b *Bot) SetRestart(delay int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Delay int `json:"delay"`
	}
	var d = UseApi{
		Action: "set_restart",
		Params: param{
			Delay: delay,
		},
		Echo: echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// CleanCache
/*
   @Description:
   @receiver b
*/
func (b *Bot) CleanCache() {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "clean_cache",
		Echo:   echo}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

//新增

// DownloadFile
/*
   @Description:
   @receiver b
   @param url string
   @param threadCount int
   @param headers []string
   @return DownloadFilePath
*/
func (b *Bot) DownloadFile(url string, threadCount int, headers []string) DownloadFilePath {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Url         string   `json:"url"`
		ThreadCount int      `json:"thread_count"`
		Headers     []string `json:"headers"`
	}
	var d = UseApi{
		Action: "download_file",
		Params: param{
			Url:         url,
			ThreadCount: threadCount,
			Headers:     headers,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseDownloadFilePathJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupMsgHistory
/*
   @Description:
   @receiver b
   @param messageSeq int64
   @param groupId int
   @return MessageHistory
*/
func (b *Bot) GetGroupMsgHistory(messageSeq int64, groupId int) MessageHistory {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		MessageSeq int64 `json:"message_seq"`
		GroupId    int   `json:"group_id"`
	}
	var d = UseApi{
		Action: "get_group_msg_history",
		Params: param{
			MessageSeq: messageSeq,
			GroupId:    groupId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseMessageHistoryJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetOnlineClients
/*
   @Description:
   @receiver b
   @param noCache bool
   @return Clients
*/
func (b *Bot) GetOnlineClients(noCache bool) Clients {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		NoCache bool `json:"no_cache"`
	}
	var d = UseApi{
		Action: "get_online_clients",
		Params: param{
			NoCache: noCache,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseOnlineClientsJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetVipInfoTest
/*
   @Description:
   @receiver b
   @param UserId int
   @return VipInfo
*/
func (b *Bot) GetVipInfoTest(UserId int) VipInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		UserId int `json:"user_id"`
	}
	var d = UseApi{
		Action: "_get_vip_info",
		Params: param{
			UserId: UserId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseVipInfoJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// SendGroupNotice
/*
   @Description:
   @receiver b
   @param groupId int
   @param content string
*/
func (b *Bot) SendGroupNotice(groupId int, content string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int    `json:"group_id"`
		Content string `json:"content"`
	}
	var d = UseApi{
		Action: "send_group_notice",
		Params: param{
			GroupId: groupId,
			Content: content,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
}

// ReloadEventFilter
/*
   @Description:
   @receiver b
*/
func (b *Bot) ReloadEventFilter() {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "reload_event_filter",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupNameSpecial
/*
   @Description:
   @receiver b
   @param groupId int
   @param groupName string
*/
func (b *Bot) SetGroupNameSpecial(groupId int, groupName string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId   int    `json:"group_id"`
		GroupName string `json:"group_name"`
	}
	var d = UseApi{
		Action: "send_group_name",
		Params: param{
			GroupId:   groupId,
			GroupName: groupName,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
}

// SetGroupPortrait
/*
   @Description:
   @receiver b
   @param groupId int
   @param file string
   @param cache int
*/
func (b *Bot) SetGroupPortrait(groupId int, file string, cache int) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int    `json:"group_id"`
		File    string `json:"file"`
		Cache   int    `json:"cache"`
	}
	var d = UseApi{
		Action: "send_group_portrait",
		Params: param{
			GroupId: groupId,
			File:    file,
			Cache:   cache,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
}

// GetMsgSpecial
/*
   @Description:
   @receiver b
   @param messageId int
   @return MsgData
*/
func (b *Bot) GetMsgSpecial(messageId int) MsgData {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		MessageId int `json:"message_id"`
	}
	var d = UseApi{
		Action: "get_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseMsgDataJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetForwardMsg
/*
   @Description:
   @receiver b
   @param messageId int
   @return []ForwardMsg
*/
func (b *Bot) GetForwardMsg(messageId int) []ForwardMsg {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		MessageId int `json:"message_id"`
	}
	var d = UseApi{
		Action: "get_forward_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseForwardMsgJson
	_ = json.Unmarshal(data, &response)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// SendGroupForwardMsg
/*
   @Description:
   @receiver b
   @param groupId int
   @param messages []Node
*/
func (b *Bot) SendGroupForwardMsg(groupId int, messages []Node) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId  int    `json:"group_id"`
		Messages []Node `json:"messages"`
	}
	var d = UseApi{
		Action: "send_group_forward_msg",
		Params: param{
			GroupId:  groupId,
			Messages: messages,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}

// GetWordSlices
/*
   @Description:
   @receiver b
   @param content string
   @return []string
*/
func (b *Bot) GetWordSlices(content string) []string {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Content string `json:"content"`
	}
	var d = UseApi{
		Action: ".get_word_slices",
		Params: param{
			Content: content,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseWordSliceJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// OcrImage
/*
   @Description:
   @receiver b
   @param image string
   @return OcrImage
*/
func (b *Bot) OcrImage(image string) OcrImage {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		Image string `json:"image"`
	}
	var d = UseApi{
		Action: ".ocr_image",
		Params: param{
			Image: image,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseOcrImageJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupSystemMsg
/*
   @Description:
   @receiver b
   @return GroupSystemMsg
*/
func (b *Bot) GetGroupSystemMsg() GroupSystemMsg {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	var d = UseApi{
		Action: "get_group_system_msg",
		Echo:   echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupSystemMsgJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupFileSystemInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupFileSystemInfo
*/
func (b *Bot) GetGroupFileSystemInfo(groupId int) GroupFileSystemInfo {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int `json:"group_id"`
	}
	var d = UseApi{
		Action: "get_group_file_system_info",
		Params: param{
			GroupId: groupId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupFileSystemInfoJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupRootFiles
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupRootFiles
*/
func (b *Bot) GetGroupRootFiles(groupId int) GroupRootFiles {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int `json:"group_id"`
	}
	var d = UseApi{
		Action: "get_group_root_files",
		Params: param{
			GroupId: groupId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupRootFilesJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupFilesByFolder
/*
   @Description:
   @receiver b
   @param groupId int
   @param folderId string
   @return GroupFilesByFolder
*/
func (b *Bot) GetGroupFilesByFolder(groupId int, folderId string) GroupFilesByFolder {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId  int `json:"group_id"`
		FolderId string
	}
	var d = UseApi{
		Action: "get_group_files_by_folder",
		Params: param{
			GroupId:  groupId,
			FolderId: folderId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupFilesByFolderJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupFileUrl
/*
   @Description:
   @receiver b
   @param groupId int
   @param fileId string
   @param busid int
   @return FileUrl
*/
func (b *Bot) GetGroupFileUrl(groupId int, fileId string, busid int) FileUrl {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int    `json:"group_id"`
		FileId  string `json:"file_id"`
		Busid   int    `json:"busid"`
	}
	var d = UseApi{
		Action: "get_group_file_url",
		Params: param{
			GroupId: groupId,
			FileId:  fileId,
			Busid:   busid,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupFileUrlJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// GetGroupAtAllRemain
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupAtAllRemain
*/
func (b *Bot) GetGroupAtAllRemain(groupId int) GroupAtAllRemain {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int `json:"group_id"`
	}
	var d = UseApi{
		Action: "get_group_at_all_remain",
		Params: param{
			GroupId: groupId,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	var response responseGroupAtAllRemainJson
	_ = json.Unmarshal(data, &response)
	contents, _ := json.Marshal(d)
	log.Println(string(contents) + "\n\t\t\t\t\t" + string(data))
	return response.Data
}

// UploadGroupFile
/*
   @Description:
   @receiver b
   @param groupId int
   @param file string
   @param name string
   @param folder string
*/
func (b *Bot) UploadGroupFile(groupId int, file string, name string, folder string) {
	rand.Seed(time.Now().Unix())
	echo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63n(1000), 10)
	type param struct {
		GroupId int    `json:"group_id"`
		File    string `json:"file"`
		Name    string `json:"name"`
		Folder  string `json:"folder"`
	}
	var d = UseApi{
		Action: "upload_group_file",
		Params: param{
			GroupId: groupId,
			File:    file,
			Name:    name,
			Folder:  folder,
		},
		Echo: echo,
	}
	b.Client.OutChan <- d
	data := getResponse(&response{}, echo)
	content, _ := json.Marshal(d)
	log.Println(string(content) + "\n\t\t\t\t\t" + string(data))
}
