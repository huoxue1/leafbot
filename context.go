package leafbot

import (
	"strconv"

	"github.com/tidwall/gjson"

	message2 "github.com/huoxue1/leafbot/message"
)

// Context
// @Description: 上下文管理对象
//
type Context struct {
	Event *Event
	Bot   API

	State    *State
	RawEvent gjson.Result

	UserID  int64
	GroupID int64
	SelfID  int64
}

// Send
/**
 * @Description: 使用上下文对象方便的回复当前会话
 * @receiver ctx
 * @param message
 * @return int32
 */
func (ctx *Context) Send(message interface{}) int32 {
	if ctx.Event.MessageType == "group" {
		return ctx.Bot.SendGroupMsg(ctx.Event.GroupId, message)
	} else if ctx.Event.MessageType == "guild" && ctx.Event.SubType == "channel" {
		result := ctx.SendGuildChannelMsg(ctx.Event.GuildID, ctx.Event.ChannelID, message)
		if result.Get("message_id").Exists() {
			return 0
		} else {
			return -1
		}
	}
	return ctx.Bot.SendPrivateMsg(ctx.Event.UserId, message)
}

// GetAtUsers
/**
 * @Description: 获取消息中At的所有人的qq
 * @receiver ctx
 * @return users
 */
func (ctx *Context) GetAtUsers() (users []int64) {
	for _, segment := range ctx.Event.Message {
		if segment.Type == "at" {
			qq, err := strconv.ParseInt(segment.Data["qq"], 10, 64)
			if err != nil {
				continue
			}
			users = append(users, qq)
		}
	}
	return
}

// GetImages
/**
 * @Description: 获取消息中所有的图片
 * @receiver ctx
 * @return images
 */
func (ctx *Context) GetImages() (images []message2.MessageSegment) {
	for _, segment := range ctx.Event.Message {
		if segment.Type == "image" {
			images = append(images, segment)
		}
	}
	return
}

// SendGroupMsg
/**
 * @Description:
 * @receiver ctx
 * @param groupId
 * @param message
 * @return int32
 */

func (ctx *Context) SendGroupMsg(groupId int64, message interface{}) int32 {
	if _, ok := message.(string); ok {
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	type param struct {
		GroupId int64       `json:"group_id"`
		Message interface{} `json:"message"`
	}
	result := ctx.CallApi("send_group_msg", param{
		GroupId: groupId,
		Message: message,
	})

	return int32(result.Get("message_id").Int())
}

// SendPrivateMsg
/**
 * @Description:
 * @receiver ctx
 * @param userId
 * @param message
 * @return int32
 */
func (ctx *Context) SendPrivateMsg(userId int64, message interface{}) int32 {
	if _, ok := message.(string); ok {
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	type param struct {
		UserId  int64       `json:"user_id"`
		Message interface{} `json:"message"`
	}
	result := ctx.CallApi("send_private_msg", param{
		UserId:  userId,
		Message: message,
	})

	return int32(result.Get("message_id").Int())
}

// DeleteMsg
/**
 * @Description:
 * @receiver ctx
 * @param messageId
 */
func (ctx *Context) DeleteMsg(messageId int32) {
	ctx.CallApi("delete_msg", map[string]int32{"message_id": messageId})
}

// GetMsg
/**
 * @Description:
 * @receiver ctx
 * @param messageId
 * @return gjson.Result
 */
func (ctx *Context) GetMsg(messageId int32) gjson.Result {
	return ctx.CallApi("get_msg", map[string]int32{"message_id": messageId})
}

// SetGroupBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param duration int
*/
func (ctx *Context) SetGroupBan(groupId int64, userId int64, duration int64) {
	ctx.CallApi("set_group_ban", map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"duration": duration},
	)
}

// SetGroupCard
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param card string
*/
func (ctx *Context) SetGroupCard(groupID int64, userId int64, card string) {
	ctx.CallApi("set_group_card", map[string]interface{}{
		"group_id": groupID,
		"user_id":  userId,
		"card":     card},
	)
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
func (ctx *Context) SendMsg(messageType string, userId int64, groupId int64, message interface{}) int32 {
	if messageType == "group" {
		return ctx.SendGroupMsg(groupId, message)
	} else {
		return ctx.SendPrivateMsg(userId, message)
	}
}

// SendLike
/*
   @Description:
   @receiver b
   @param userId int
   @param times int
*/
func (ctx *Context) SendLike(userID int64, times int) {
	ctx.CallApi("send_like", map[string]interface{}{
		"user_id": userID,
		"times":   times},
	)
}

// SetGroupKick
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param rejectAddRequest bool
*/
func (ctx *Context) SetGroupKick(groupID int64, userId int64, rejectAddRequest bool) {
	ctx.CallApi("set_group_kick", map[string]interface{}{
		"group_id":           groupID,
		"user_id":            userId,
		"reject_add_request": rejectAddRequest},
	)
}

// SetGroupAnonymousBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param flag string
   @param duration int
*/
func (ctx *Context) SetGroupAnonymousBan(groupID int64, flag string, duration int) {
	ctx.CallApi("set_group_anonymous_ban", map[string]interface{}{
		"group_id": groupID,
		"flag":     flag,
		"duration": duration},
	)
}

// SetGroupWholeBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param enable bool
*/
func (ctx *Context) SetGroupWholeBan(groupId int64, enable bool) {
	ctx.CallApi("set_group_whole_ban", map[string]interface{}{"group_id": groupId, "enable": enable})
}

// SetGroupAdmin
/*
   @Description:
   @receiver b
   @param groupId int
   @param userId int
   @param enable bool
*/
func (ctx *Context) SetGroupAdmin(groupID int64, userId int64, enable bool) {
	ctx.CallApi("set_group_admin", map[string]interface{}{"group_id": groupID, "user_id": userId, "enable": enable})
}

// SetGroupAnonymous
/*
   @Description:
   @receiver b
   @param groupID int64
   @param enable bool
*/
func (ctx *Context) SetGroupAnonymous(groupID int64, enable bool) {
	ctx.CallApi("set_group_anonymous", map[string]interface{}{"group_id": groupID, "enable": enable})
}

// SetGroupName
/*
   @Description:
   @receiver b
   @param groupID int64
   @param groupName string
*/
func (ctx *Context) SetGroupName(groupID int64, groupName string) {
	ctx.CallApi("set_group_name", map[string]interface{}{"group_id": groupID, "group_name": groupName})
}

// SetGroupLeave
/*
   @Description:
   @receiver b
   @param groupID int64
   @param isDisMiss bool
*/
func (ctx *Context) SetGroupLeave(groupID int64, isDisMiss bool) {
	ctx.CallApi("set_group_leave", map[string]interface{}{"group_id": groupID, "is_dismiss": isDisMiss})
}

// SetGroupSpecialTitle
/*
   @Description:
   @receiver b
   @param groupID int64
   @param userId int
   @param specialTitle string
   @param duration int
*/
func (ctx *Context) SetGroupSpecialTitle(groupID int64, userId int64, specialTitle string, duration int) {
	ctx.CallApi("set_group_special_title", map[string]interface{}{"group_id": groupID, "user_id": userId, "special_title": specialTitle, "duration": duration})
}

// SetFriendAddRequest
/*
   @Description:
   @receiver b
   @param flag string
   @param approve bool
   @param remark string
*/
func (ctx *Context) SetFriendAddRequest(flag string, approve bool, remark string) {
	ctx.CallApi("set_friend_add_request", map[string]interface{}{"flag": flag, "approve": approve, "remark": remark})
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
func (ctx *Context) SetGroupAddRequest(flag string, subType string, approve bool, reason string) {
	ctx.CallApi("set_group_add_request", map[string]interface{}{
		"flag":     flag,
		"sub_type": subType,
		"approve":  approve,
		"reason":   reason})
}

// GetLoginInfo
/*
   @Description:
   @receiver b
   @return LoginInfo
*/
func (ctx *Context) GetLoginInfo() gjson.Result {
	return ctx.CallApi("get_login_info", nil)
}

// GetStrangerInfo
/*
   @Description:
   @receiver b
   @param userId int
   @param noCache bool
   @return Senders
*/
func (ctx *Context) GetStrangerInfo(userId int, noCache bool) gjson.Result {
	return ctx.CallApi("get_stranger_info", map[string]interface{}{"user_id": userId, "no_cache": noCache})
}

// GetFriendList
/**
 * @Description:
 * @receiver b
 * @return gjson.Result
 * example
 */
func (ctx *Context) GetFriendList() gjson.Result {
	return ctx.CallApi("get_friend_list", nil)
}

// GetGroupInfo
/**
 * @Description:
 * @receiver b
 * @param groupID
 * @param noCache
 * @return gjson.Result
 * example
 */
func (ctx *Context) GetGroupInfo(groupID int64, noCache bool) gjson.Result {
	return ctx.CallApi("get_group_info", map[string]interface{}{"group_id": groupID, "no_cache": noCache})
}

// GetGroupList
/*
   @Description:
   @receiver b
   @return []GroupInfo
*/
func (ctx *Context) GetGroupList() gjson.Result {
	return ctx.CallApi("get_group_list", nil)
}

// GetGroupMemberInfo
/*
   @Description:
   @receiver b
   @param groupID int64
   @param userId int
   @param noCache bool
   @return GroupMemberInfo
*/
func (ctx *Context) GetGroupMemberInfo(groupID int64, userId int64, noCache bool) gjson.Result {
	return ctx.CallApi("get_group_member_info", map[string]interface{}{"group_id": groupID, "user_id": userId, "no_cache": noCache})
}

// GetGroupMemberList
/*
   @Description:
   @receiver b
   @param groupID int64
   @return []GroupMemberInfo
*/
func (ctx *Context) GetGroupMemberList(groupID int64) gjson.Result {
	return ctx.CallApi("get_group_member_list", map[string]interface{}{"group_id": groupID})
}

// GetGroupHonorInfo
/*
   @Description:
   @receiver b
   @param groupID int64
   @param honorType string
   @return GroupHonorInfo
*/
func (ctx *Context) GetGroupHonorInfo(groupID int64, honorType string) gjson.Result {
	return ctx.CallApi("get_stranger_info", map[string]interface{}{"group_id": groupID, "honor_type": honorType})
}

// GetCookies
/*
   @Description:
   @receiver b
   @param domain string
   @return Cookie
*/
func (ctx *Context) GetCookies(domain string) gjson.Result {
	return ctx.CallApi("get_cookies", map[string]interface{}{"domain": domain})
}

// GetCsrfToken
/*
   @Description:
   @receiver b
   @return CsrfToken
*/
func (ctx *Context) GetCsrfToken() gjson.Result {
	return ctx.CallApi("get_csrf_token", nil)
}

// GetCredentials
/*
   @Description:
   @receiver b
   @param domain string
   @return Credentials
*/
func (ctx *Context) GetCredentials(domain string) gjson.Result {
	return ctx.CallApi("get_credentials", map[string]interface{}{"domain": domain})
}

// GetRecord
/*
   @Description:
   @receiver b
   @param file file
   @param outFormat string
   @return Record
*/
func (ctx *Context) GetRecord(file, outFormat string) gjson.Result {
	return ctx.CallApi("get_record", map[string]interface{}{"file": file, "out_format": outFormat})
}

// GetImage
/*
   @Description:
   @receiver b
   @param file string
   @return Image
*/
func (ctx *Context) GetImage(file string) gjson.Result {
	return ctx.CallApi("get_image", map[string]interface{}{"file": file})
}

// CanSendImage
/*
   @Description:
   @receiver b
   @return Bool
*/
func (ctx *Context) CanSendImage() bool {
	return ctx.CallApi("can_send_image", nil).Bool()
}

// CanSendRecord
/*
   @Description:
   @receiver b
   @return Bool
*/
func (ctx *Context) CanSendRecord() bool {
	return ctx.CallApi("can_send_record", nil).Bool()
}

// GetStatus
/*
   @Description:
   @receiver b
   @return OnlineStatus
*/
func (ctx *Context) GetStatus() gjson.Result {
	return ctx.CallApi("get_status", nil)
}

// SetRestart
/*
   @Description:
   @receiver b
   @param delay int
*/
func (ctx *Context) SetRestart(delay int) {
	ctx.CallApi("set_restart", map[string]interface{}{"delay": delay})
}

// CleanCache
/**
 * @Description:
 * @receiver b
 * example
 */
func (ctx *Context) CleanCache() {
	ctx.CallApi("clean_cache", nil)
}

// 新增

// DownloadFile
/*
   @Description:
   @receiver b
   @param url string
   @param threadCount int
   @param headers []string
   @return DownloadFilePath
*/
func (ctx *Context) DownloadFile(url string, threadCount int, headers []string) gjson.Result {
	return ctx.CallApi("download_file", map[string]interface{}{
		"url":          url,
		"thread_count": threadCount,
		"headers":      headers})
}

// GetGroupMsgHistory
/*
   @Description:
   @receiver b
   @param messageSeq int64
   @param groupID int64
   @return MessageHistory
*/
func (ctx *Context) GetGroupMsgHistory(messageSeq int64, groupID int64) gjson.Result {
	return ctx.CallApi("get_group_msh_history", map[string]interface{}{"message_seq": messageSeq, "group_id": groupID})
}

// GetOnlineClients
/**
 * @Description:
 * @receiver b
 * @param noCache
 * @return gjson.Result
 * example
 */
func (ctx *Context) GetOnlineClients(noCache bool) gjson.Result {
	return ctx.CallApi("get_online_clients", map[string]interface{}{"no_cache": noCache})
}

// GetVipInfoTest
/*
   @Description:
   @receiver b
   @param UserId int
   @return VipInfo
*/
func (ctx *Context) GetVipInfoTest(userId int64) gjson.Result {
	return ctx.CallApi("_get_vip_info", map[string]interface{}{"user_id": userId})
}

// SendGroupNotice
/*
   @Description:
   @receiver b
   @param groupID int64
   @param content string
*/
func (ctx *Context) SendGroupNotice(groupID int64, content string) {
	ctx.CallApi("send_group_notice", map[string]interface{}{"group_id": groupID, "content": content})
}

// ReloadEventFilter
/*
   @Description:
   @receiver b
*/
func (ctx *Context) ReloadEventFilter() {
	ctx.CallApi("reload_event_filter", nil)
}

// SetGroupNameSpecial
/*
   @Description:
   @receiver b
   @param groupID int64
   @param groupName string
*/
func (ctx *Context) SetGroupNameSpecial(groupID int64, groupName string) {
	ctx.CallApi("set_group_name", map[string]interface{}{"group_id": groupID, "group_name": groupName})
}

// SetGroupPortrait
/*
   @Description:
   @receiver b
   @param groupID int64
   @param file string
   @param cache int
*/
func (ctx *Context) SetGroupPortrait(groupID int64, file string, cache int) {
	ctx.CallApi("set_group+portrait", map[string]interface{}{"group_id": groupID, "file": file, "cache": cache})
}

// GetMsgSpecial
/*
   @Description:
   @receiver b
   @param messageId int
   @return MsgData
*/
func (ctx *Context) GetMsgSpecial(messageId int) gjson.Result {
	return ctx.CallApi("get_msg", map[string]interface{}{"message_id": messageId})
}

// GetForwardMsg
/*
   @Description:
   @receiver b
   @param messageId int
   @return []ForwardMsg
*/
func (ctx *Context) GetForwardMsg(messageId int) gjson.Result {
	return ctx.CallApi("get_forward_msg", map[string]interface{}{"message_id": messageId})
}

// SendGroupForwardMsg
/*
   @Description:
   @receiver b
   @param groupID int64
   @param messages []Node
*/
func (ctx *Context) SendGroupForwardMsg(groupID int64, messages interface{}) {
	ctx.CallApi("send_group_forward_msg", map[string]interface{}{"group_id": groupID, "message": messages})
}

// GetWordSlices
/*
   @Description:
   @receiver b
   @param content string
   @return []string
*/
func (ctx *Context) GetWordSlices(content string) gjson.Result {
	return ctx.CallApi(".get_word_slices", map[string]interface{}{"content": content})
}

// OcrImage
/*
   @Description:
   @receiver b
   @param image string
   @return OcrImage
*/
func (ctx *Context) OcrImage(image string) gjson.Result {
	return ctx.CallApi("ocr_image", map[string]interface{}{"image": image})
}

// GetGroupSystemMsg
/*
   @Description:
   @receiver b
   @return GroupSystemMsg
*/
func (ctx *Context) GetGroupSystemMsg() gjson.Result {
	return ctx.CallApi("get_group_system_msg", nil)
}

// GetGroupFileSystemInfo
/*
   @Description:
   @receiver b
   @param groupID int64
   @return GroupFileSystemInfo
*/
func (ctx *Context) GetGroupFileSystemInfo(groupID int64) gjson.Result {
	return ctx.CallApi("get_group_file_system_info", map[string]interface{}{"group_id": groupID})
}

// GetGroupRootFiles
/*
   @Description:
   @receiver b
   @param groupID int64
   @return GroupRootFiles
*/
func (ctx *Context) GetGroupRootFiles(groupID int64) gjson.Result {
	return ctx.CallApi("get_group_root_files", map[string]interface{}{"group_id": groupID})
}

// GetGroupFilesByFolder
/*
   @Description:
   @receiver b
   @param groupID int64
   @param folderId string
   @return GroupFilesByFolder
*/
func (ctx *Context) GetGroupFilesByFolder(groupID int64, folderId string) gjson.Result {
	return ctx.CallApi("get_group_files_by_folder", map[string]interface{}{"group_id": groupID, "folder_id": folderId})
}

// GetGroupFileUrl
/*
   @Description:
   @receiver b
   @param groupID int64
   @param fileId string
   @param busid int
   @return FileUrl
*/
func (ctx *Context) GetGroupFileUrl(groupID int64, fileId string, busid int) gjson.Result {
	return ctx.CallApi("get_group_file_url", map[string]interface{}{"group_id": groupID, "file_id": fileId, "busid": busid})
}

// GetGroupAtAllRemain
/*
   @Description:
   @receiver b
   @param groupID int64
   @return GroupAtAllRemain
*/
func (ctx *Context) GetGroupAtAllRemain(groupID int64) gjson.Result {
	return ctx.CallApi("get_group_at_all_remain", map[string]interface{}{"group_id": groupID})
}

// UploadGroupFile
/*
   @Description:
   @receiver b
   @param groupID int64
   @param file string
   @param name string
   @param folder string
*/
func (ctx *Context) UploadGroupFile(groupID int64, file string, name string, folder string) {
	ctx.CallApi("upload_group_file", map[string]interface{}{"group_id": groupID, "file": file, "name": name, "folder": folder})
}

func (ctx *Context) CallApi(action string, param interface{}) gjson.Result {
	return ctx.Bot.CallApi(action, param)
}

// SetEssenceMsg
/**
 * @Description:
 * @receiver ctx
 * @param messageId
 */
func (ctx *Context) SetEssenceMsg(messageId int) {
	ctx.CallApi("set_essence_msg", map[string]interface{}{"message_id": messageId})
}

// DeleteEssenceMsg
/**
 * @Description:
 * @receiver ctx
 * @param messageId
 */
func (ctx *Context) DeleteEssenceMsg(messageId int) {
	ctx.CallApi("delete_essence_msg", map[string]interface{}{"message_id": messageId})
}

// GetEssenceMsgList
/**
 * @Description:
 * @receiver ctx
 * @param groupID
 * @return gjson.Result
 */
func (ctx *Context) GetEssenceMsgList(groupID int64) gjson.Result {
	return ctx.CallApi("get_essence_msg_list", map[string]interface{}{"group_id": groupID})
}

// CheckUrlSafely
/**
 * @Description:
 * @receiver ctx
 * @param url
 * @return int
 */
func (ctx *Context) CheckUrlSafely(url string) int {
	return int(ctx.CallApi("check_url_safely", map[string]interface{}{"url": url}).Int())
}

func (ctx *Context) GetGuildServiceProfile() gjson.Result {
	return ctx.CallApi("get_guild_service_profile", nil)
}

func (ctx *Context) GetGuildList() gjson.Result {
	return ctx.CallApi("get_guild_list", nil)
}

func (ctx *Context) GetGuildMetaByQuest(guildID int64) gjson.Result {
	return ctx.CallApi("get_guild_meta_by_guest", map[string]interface{}{"guild_id": guildID})
}

func (ctx *Context) GetGuildChannelList(guildID int64, noCache bool) gjson.Result {
	return ctx.CallApi("get_guild_channel_list", map[string]interface{}{"guild_id": guildID, "no_cache": noCache})
}

func (ctx *Context) GetGuildMembers(guildID int64) gjson.Result {
	return ctx.CallApi("get_guild_members", map[string]interface{}{"guild_id": guildID})
}

func (ctx *Context) SendGuildChannelMsg(guildID, channelID int64, message interface{}) gjson.Result {
	return ctx.CallApi("send_guild_channel_msg", map[string]interface{}{"guild_id": guildID, "channel_id": channelID, "message": message})
}
