package leafBot

import (
	"github.com/tidwall/gjson"
)

// Api
// @Description:
//
type Api interface {
	SendGroupMsg(groupId int, message interface{}) int32
	SendPrivateMsg(userId int, message interface{}) int32
	DeleteMsg(messageId int32)
	GetMsg(messageId int32) gjson.Result
	SetGroupBan(groupId int, userId int, duration int)
	SetGroupCard(groupId int, userId int, card string)
	SendMsg(messageType string, userId int, groupId int, message interface{}) int32
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
	GetLoginInfo() gjson.Result
	GetStrangerInfo(userId int, noCache bool) gjson.Result
	GetFriendList() gjson.Result
	GetGroupInfo(groupId int, noCache bool) gjson.Result
	GetGroupList() gjson.Result
	GetGroupMemberInfo(groupId int, UserId int, noCache bool) gjson.Result
	GetGroupMemberList(groupId int) gjson.Result
	GetGroupHonorInfo(groupId int, honorType string) gjson.Result
	GetCookies(domain string) gjson.Result
	GetCsrfToken() gjson.Result
	GetCredentials(domain string) gjson.Result
	GetRecord(file, outFormat string) gjson.Result
	GetImage(file string) gjson.Result
	CanSendImage() bool
	CanSendRecord() bool
	GetStatus() gjson.Result
	SetRestart(delay int)
	CleanCache()

	DownloadFile(url string, threadCount int, headers []string) gjson.Result
	GetGroupMsgHistory(messageSeq int64, groupId int) gjson.Result
	GetOnlineClients(noCache bool) gjson.Result
	GetVipInfoTest(UserId int) gjson.Result
	SendGroupNotice(groupId int, content string)
	ReloadEventFilter()
	SetEssenceMsg(messageId int)
	DeleteEssenceMsg(messageId int)
	GetEssenceMsgList(groupId int)
	CheckUrlSafely(url string) int
	UploadGroupFile(groupId int, file string, name string, folder string)

	SetGroupNameSpecial(groupId int, groupName string)
	SetGroupPortrait(groupId int, file string, cache int)
	GetMsgSpecial(messageId int) gjson.Result
	GetForwardMsg(messageId int) gjson.Result
	SendGroupForwardMsg(groupId int, messages interface{})
	GetWordSlices(content string) gjson.Result
	OcrImage(image string) gjson.Result
	GetGroupSystemMsg() gjson.Result
	GetGroupFileSystemInfo(groupId int) gjson.Result
	GetGroupRootFiles(groupId int) gjson.Result
	GetGroupFilesByFolder(groupId int, folderId string) gjson.Result
	GetGroupFileUrl(groupId int, fileId string, busid int) gjson.Result
	GetGroupAtAllRemain(groupId int) gjson.Result

	CallApi(action string, params interface{}) interface{}
}

func (e Event) Send(message interface{}) int32 {
	bot := GetBotById(e.SelfId)
	if e.MessageType == "private" {
		return bot.SendPrivateMsg(e.UserId, message)
	} else if e.MessageType == "group" {
		return bot.SendGroupMsg(e.GroupId, message)
	}
	return 0
}

// type (
//	FileUrl struct {
//		Url string `json:"url"`
//	}
//	MsgData struct {
//		MessageId int     `json:"message_id"`
//		RealId    int     `json:"real_id"`
//		Sender    Senders `json:"sender"`
//		Time      int     `json:"time"`
//		Message   string  `json:"message"`
//	}
//	ForwardMsg struct {
//		Content string  `json:"content"`
//		Sender  Senders `json:"sender"`
//		Time    int     `json:"time"`
//	}
//	TextDetection struct {
//		Text        string      `json:"text"`
//		Confidence  int         `json:"confidence"`
//		Coordinates interface{} `json:"coordinates"`
//	}
//	OcrImage struct {
//		Texts    []TextDetection `json:"texts"`
//		Language string          `json:"language"`
//	}
//	InvitedRequest struct {
//		RequestId   int    `json:"request_id"`   //请求ID
//		InvitorUin  int    `json:"invitor_uin"`  //邀请者
//		InvitorNick string `json:"invitor_nick"` //邀请者昵称
//		GroupId     int    `json:"group_id"`     //群号
//		GroupName   string `json:"group_name"`   //群名
//		Checked     bool   `json:"checked"`      //是否已被处理
//		Actor       int64  `json:"actor"`        //处理者, 未处理为0
//	}
//	JoinRequest struct {
//		RequestId     int    `json:"request_id"`     //请求ID
//		RequesterUin  int    `json:"requester_uin"`  //请求者ID
//		RequesterNick string `json:"requester_nick"` //请求者昵称
//		Message       string `json:"message"`        //验证消息
//		GroupId       int    `json:"group_id"`       //群号
//		GroupName     string `json:"group_name"`     //群名
//		Checked       bool   `json:"checked"`        //是否已被处理
//		Actor         int    `json:"actor"`          //处理者, 未处理为0
//	}
//	GroupSystemMsg struct {
//		InvitedRequests []InvitedRequest `json:"invited_requests"` //邀请消息列表
//		JoinRequests    []JoinRequest    `json:"join_requests"`    //进群消息列表
//	}
//	GroupFileSystemInfo struct {
//		FileCount  int `json:"file_count"`  //文件总数
//		LimitCount int `json:"limit_count"` //文件上限
//		UsedSpace  int `json:"used_space"`  //已使用空间
//		TotalSpace int `json:"total_space"` //空间上限
//	}
//	File struct {
//		FileId        string `json:"file_id"`        //文件ID
//		FileName      string `json:"file_name"`      //文件名
//		Busid         int    `json:"busid"`          //文件类型
//		FileSize      int64  `json:"file_size"`      //文件大小
//		UploadTime    int64  `json:"upload_time"`    //上传时间
//		DeadTime      int64  `json:"dead_time"`      //过期时间,永久文件恒为0
//		ModifyTime    int64  `json:"modify_time"`    //最后修改时间
//		DownloadTimes int32  `json:"download_times"` //下载次数
//		Uploader      int64  `json:"uploader"`       //上传者ID
//		UploaderName  string `json:"uploader_name"`  //上传者名字
//	}
//	Folder struct {
//		FolderId       string `json:"folder_id"`        //文件夹ID
//		FolderName     string `json:"folder_name"`      //文件名
//		CreateTime     int    `json:"create_time"`      //创建时间
//		Creator        int    `json:"creator"`          //创建者
//		CreatorName    string `json:"creator_name"`     //创建者名字
//		TotalFileCount int32  `json:"total_file_count"` //子文件数量
//	}
//	GroupRootFiles struct {
//		Files   []File   `json:"files"`
//		Folders []Folder `json:"folders"`
//	}
//	GroupFilesByFolder struct {
//		Files   []File   `json:"files"`
//		Folders []Folder `json:"folders"`
//	}
//	GroupAtAllRemain struct {
//		CanAtAll                 bool `json:"can_at_all"`                    //是否可以@全体成员
//		RemainAtAllCountForGroup int  `json:"remain_at_all_count_for_group"` //群内所有管理当天剩余@全体成员次数
//		RemainAtAllCountForUin   int  `json:"remain_at_all_count_for_uin"`   //BOT当天剩余@全体成员次数
//	}
//	DownloadFilePath struct {
//		File string `json:"file"`
//	}
//	MessageHistory struct {
//		Messages []string `json:"messages"`
//	}
//	Clients struct {
//		Clients []Device `json:"clients"` //在线客户端列表
//	}
//	Device struct {
//		AppId      int64  `json:"app_id"`      //客户端ID
//		DeviceName string `json:"device_name"` //设备名称
//		DeviceKind string `json:"device_kind"` //设备类型
//	}
//	VipInfo struct {
//		UserId         int64   `json:"user_id"`          //QQ 号
//		Nickname       string  `json:"nickname"`         //用户昵称
//		Level          int64   `json:"level"`            //QQ 等级
//		LevelSpeed     float64 `json:"level_speed"`      //等级加速度
//		VipLevel       string  `json:"vip_level"`        //会员等级
//		VipGrowthSpeed int64   `json:"vip_growth_speed"` //会员成长速度
//		VipGrowthTotal int64   `json:"vip_growth_total"` //会员成长总值
//	}
//)
//
//
//type (
//	ConstNoticeType struct {
//		GroupUpload   string //群文件上传
//		GroupAdmin    string //群管理员变动
//		GroupDecrease string //群成员减少
//		GroupIncrease string //群成员增加
//		GroupBan      string //群禁言
//		FriendAdd     string //好友添加
//		GroupRecall   string //群消息撤回
//		FriendRecall  string //好友消息撤回
//		Notify        string //群内戳一戳,群红包运气王,群成员荣誉变更,好友戳一戳
//		GroupCard     string //群成员名片更新
//		OfflineFile   string //接收到离线文件
//	}
//	ConstMessageType struct {
//		Group   string
//		Private string
//	}
//	//ConstMessageSubType struct {
//	//	Friend    string //好友消息
//	//	Group     string //临时会话
//	//	Other     string //其他消息
//	//	Normal    string //正常群消息
//	//	Anonymous string //匿名消息
//	//	Notice    string //群通知消息
//	//}
//)
//
//var (
//	NoticeTypeApi = ConstNoticeType{
//		GroupUpload:   "group_upload",
//		GroupAdmin:    "group_admin",
//		GroupDecrease: "group_decrease",
//		GroupIncrease: "group_increase",
//		GroupBan:      "group_ban",
//		FriendAdd:     "friend_add",
//		GroupRecall:   "group_recall",
//		FriendRecall:  "friend_recall",
//		Notify:        "notify",
//		OfflineFile:   "offline_file",
//		GroupCard:     "group_card"}
//)
//
//
//
//type LoginInfo struct {
//	UserId   int    `json:"user_id"`
//	NickName string `json:"nick_name"`
//}
//
//
//
//
//
//
//type responseJson struct {
//	Status  string `json:"status"`
//	RetCode int    `json:"retcode"`
//}
//
//type (
//
//
//	GroupMemberInfo struct {
//		GroupId         int  `json:"group_id"`
//		JoinTime        int  `json:"join_time"`
//		LastSentTime    int  `json:"last_sent_time"`
//		Unfriendly      bool `json:"unfriendly"`
//		TitleExpireTime int  `json:"title_expire_time"`
//		CardChangeable  bool `json:"card_changeable"`
//		Senders
//	}
//
//	getMsgJson struct {
//		responseJson
//		Data GetMessage `json:"data"`
//	}
//
//	GetMessage struct {
//		Time      int32            `json:"time"`
//		Group     bool             `json:"group"`
//		MessageId int32            `json:"message_id"`
//		RealId    int32            `json:"real_id"`
//		Sender    Senders          `json:"sender"`
//		Message   message2.Message `json:"message"`
//	}
//
//	FriendList struct {
//		UserId   int    `json:"user_id"`
//		NickName string `json:"nick_name"`
//		Remark   string `json:"remark"`
//	}
//
//	GroupInfo struct {
//		GroupId        int    `json:"group_id"`
//		GroupName      string `json:"group_name"`
//		MemberCount    int    `json:"member_count"`
//		MaxMemberCount int    `json:"max_member_count"`
//	}
//
//	GroupHonorInfo struct {
//		GroupId          int                  `json:"group_id"`
//		CurrentTalkative CurrentTalkativeS    `json:"current_talkative"`
//		TalkativeList    []GroupHonorInfoList `json:"talkative_list"`
//		PerformerList    []GroupHonorInfoList `json:"performer_list"`
//		LegendList       []GroupHonorInfoList `json:"legend_list"`
//		StrongNewbieList []GroupHonorInfoList `json:"strong_newbie_list"`
//		EmotionList      []GroupHonorInfoList `json:"emotion_list"`
//	}
//
//	CurrentTalkativeS struct {
//		UserId   int    `json:"user_id"`
//		NickName string `json:"nick_name"`
//		Avatar   string `json:"avatar"`
//		DayCount int    `json:"day_count"`
//	}
//
//	GroupHonorInfoList struct {
//		UserId      int    `json:"user_id"`
//		NickName    string `json:"nick_name"`
//		Avatar      string `json:"avatar"`
//		Description string `json:"description"`
//	}
//
//	Record struct {
//		File      string `json:"file"`
//		OutFormat string `json:"out_format"`
//	}
//
//	Cookie struct {
//		Cookies string `json:"cookies"`
//	}
//
//	CsrfToken struct {
//		Token string `json:"token"`
//	}
//
//	OnlineStatus struct {
//		Online bool `json:"online"`
//		Good   bool `json:"good"`
//	}
//
//	Bool struct {
//		Yes bool `json:"yes"`
//	}
//
//	Image struct {
//		File string `json:"file"`
//	}
//
//	Credentials struct {
//		Cookie
//		CsrfToken
//	}
//	Content struct {
//		Type string `json:"type"`
//		Data string `json:"data"`
//	}
//	Node struct {
//		Id      int     `json:"id"`
//		Name    string  `json:"name"`
//		Uin     int     `json:"uin"`
//		Content Content `json:"content"`
//	}
//)
