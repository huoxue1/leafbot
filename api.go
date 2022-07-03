// Package leafbot
// @Description:
package leafbot

import (
	"github.com/tidwall/gjson"
)

// API
// @Description:
//
type API interface {
	CallApi(action string, params interface{}) gjson.Result
	SendGroupMsg(groupID int64, message interface{}) int32
	SendPrivateMsg(userID int64, message interface{}) int32
}

// GuildAPI
// @Description: 频道相关api
//
type GuildAPI interface {
	API
	// GetGuildServiceProfile
	// @Description: 获取频道系统内BOT的资料
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E7%B3%BB%E7%BB%9F%E5%86%85bot%E7%9A%84%E8%B5%84%E6%96%99
	GetGuildServiceProfile() gjson.Result

	// GetGuildList
	// @Description: 获取频道列表
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E5%88%97%E8%A1%A8
	//
	GetGuildList() gjson.Result

	// GetGuildMetaByQuest
	// @Description: 通过访客获取频道元数据
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E9%80%9A%E8%BF%87%E8%AE%BF%E5%AE%A2%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E5%85%83%E6%95%B0%E6%8D%AE
	//
	GetGuildMetaByQuest(guildID int64) gjson.Result

	// GetGuildChannelList
	// @Description: 获取子频道列表
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E8%8E%B7%E5%8F%96%E5%AD%90%E9%A2%91%E9%81%93%E5%88%97%E8%A1%A8
	//
	GetGuildChannelList(guildID int64, noCache bool) gjson.Result

	// GetGuildMembers
	// @Description: 获取频道成员列表
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8
	//
	GetGuildMembers(guildID int64) gjson.Result

	// SendGuildChannelMsg
	// @Description: 发送信息到子频道
	// @param guildID 频道ID
	// @param channelID 子频道ID
	// @param message 消息
	// https://github.com/Mrs4s/go-cqhttp/blob/master/docs/guild.md#%E5%8F%91%E9%80%81%E4%BF%A1%E6%81%AF%E5%88%B0%E5%AD%90%E9%A2%91%E9%81%93
	//
	SendGuildChannelMsg(guildID, channelID int64, message interface{}) gjson.Result
}

// OneBotAPI
// @Description:
//
type OneBotAPI interface {
	API
	DeleteMsg(messageID int32)
	GetMsg(messageID int32) gjson.Result
	SetGroupBan(groupID int64, userID int64, duration int64)
	SetGroupCard(groupID int64, userID int64, card string)
	SendMsg(messageType string, userID int64, groupID int64, message interface{}) int32
	SendLike(userID int64, times int)
	SetGroupKick(groupID int64, userID int64, rejectAddRequest bool)
	SetGroupAnonymousBan(groupID int64, flag string, duration int)
	SetGroupWholeBan(groupID int64, enable bool)
	SetGroupAdmin(groupID int64, UserID int64, enable bool)
	SetGroupAnonymous(groupID int64, enable bool)
	SetGroupName(groupID int64, groupName string)
	SetGroupLeave(groupID int64, isDisMiss bool)
	SetGroupSpecialTitle(groupID int64, userID int64, specialTitle string, duration int)
	SetFriendAddRequest(flag string, approve bool, remark string)
	SetGroupAddRequest(flag string, subType string, approve bool, reason string)
	GetLoginInfo() gjson.Result
	GetStrangerInfo(userID int, noCache bool) gjson.Result
	GetFriendList() gjson.Result
	GetGroupInfo(groupID int64, noCache bool) gjson.Result
	GetGroupList() gjson.Result
	GetGroupMemberInfo(groupID int64, UserID int64, noCache bool) gjson.Result
	GetGroupMemberList(groupID int64) gjson.Result
	GetGroupHonorInfo(groupID int64, honorType string) gjson.Result
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

	GetGroupFileSystemInfo(groupID int64) gjson.Result
	GetGroupRootFiles(groupID int64) gjson.Result
	GetGroupFilesByFolder(groupID int64, folderID string) gjson.Result
	GetGroupFileUrl(groupID int64, fileID string, busid int) gjson.Result
	DownloadFile(url string, threadCount int, headers []string) gjson.Result
	UploadGroupFile(groupID int64, file string, name string, folder string)
	UploadPrivateFile(userID int64, file, name string)

	GetGroupMsgHistory(messageSeq int64, groupID int64) gjson.Result
	GetOnlineClients(noCache bool) gjson.Result
	GetVipInfoTest(UserID int64) gjson.Result
	SendGroupNotice(groupID int64, content string)
	ReloadEventFilter()
	SetEssenceMsg(messageID int)
	DeleteEssenceMsg(messageID int)
	GetEssenceMsgList(groupID int64) gjson.Result
	CheckUrlSafely(url string) int

	SetGroupNameSpecial(groupID int64, groupName string)
	SetGroupPortrait(groupID int64, file string, cache int)
	GetMsgSpecial(messageID int) gjson.Result
	GetForwardMsg(messageID int) gjson.Result
	SendGroupForwardMsg(groupID int64, messages interface{})
	SendPrivateForwardMsg(userID int64, messages interface{})

	GetWordSlices(content string) gjson.Result
	OcrImage(image string) gjson.Result
	GetGroupSystemMsg() gjson.Result

	GetGroupAtAllRemain(groupID int64) gjson.Result
}

func (e Event) Send(message interface{}) int32 {
	bot := GetBotById(e.SelfId)
	if e.MessageType == "private" {
		return int32(bot.CallApi("send_private_msg", map[string]interface{}{
			"user_id": e.UserId,
			"message": message,
		}).Int())
	} else if e.MessageType == "group" {
		return int32(bot.CallApi("send_group_msg", map[string]interface{}{
			"group_id": e.GroupId,
			"message":  message,
		}).Int())
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
// type (
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
// var (
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
// type LoginInfo struct {
//	UserId   int    `json:"user_id"`
//	NickName string `json:"nick_name"`
//}
//
//
//
//
//
//
// type responseJson struct {
//	Status  string `json:"status"`
//	RetCode int    `json:"retcode"`
//}
//
// type (
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
