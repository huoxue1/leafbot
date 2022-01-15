package cqhttp_default_driver

import (
	"encoding/json" //nolint:gci

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	message2 "github.com/huoxue1/leafbot/message"
)

// SendGroupMsg
/**
 * @Description:
 * @receiver b
 * @param groupID
 * @param message
 * @return int32
 */
func (b *Bot) SendGroupMsg(groupID int, message interface{}) int32 {
	if _, ok := message.(string); ok {
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	result := b.CallApi("send_group_msg", map[string]interface{}{"group_id": groupID, "message": message})

	return int32(result.Get("message_id").Int())
}

// SendPrivateMsg
/*
   @Description:
   @receiver b
   @param userID int
   @param message string
   @param autoEscape bool
   @return int32
*/
func (b *Bot) SendPrivateMsg(userID int, message interface{}) int32 {
	if _, ok := message.(string); ok {
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	type param struct {
		UserId  int         `json:"user_id"`
		Message interface{} `json:"message"`
	}
	result := b.CallApi("send_private_msg", param{
		UserId:  userID,
		Message: message,
	})

	return int32(result.Get("message_id").Int())
}

// DeleteMsg
/*
   @Description:
   @receiver b
   @param messageID int32
*/
func (b *Bot) DeleteMsg(messageID int32) {
	b.CallApi("delete_msg", map[string]interface{}{"message_id": messageID})
}

// GetMsg
/*
   @Description:
   @receiver b
   @param messageID int32
   @return GetMessage
*/
func (b *Bot) GetMsg(messageID int32) gjson.Result {
	return b.CallApi("get_msg", map[string]interface{}{"message_id": messageID})
}

// SetGroupBan
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param duration int
*/
func (b *Bot) SetGroupBan(groupID int, userID int, duration int) {
	b.CallApi("set_group_ban", map[string]interface{}{
		"group_id": groupID,
		"user_id":  userID,
		"duration": duration},
	)
}

// SetGroupCard
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param card string
*/
func (b *Bot) SetGroupCard(groupID int, userID int, card string) {
	b.CallApi("set_group_card", map[string]interface{}{
		"group_id": groupID,
		"user_id":  userID,
		"card":     card},
	)
}

// SendMsg
/*
   @Description:
   @receiver b
   @param messageType string
   @param userID int
   @param groupID int
   @param message string
   @param autoEscape bool
   @return int32
*/
func (b *Bot) SendMsg(messageType string, userID int, groupID int, message interface{}) int32 {
	if messageType == "group" {
		return b.SendGroupMsg(groupID, message)
	}
	return b.SendPrivateMsg(userID, message)
}

// SendLike
/*
   @Description:
   @receiver b
   @param userID int
   @param times int
*/
func (b *Bot) SendLike(userID int, times int) {
	b.CallApi("send_like", map[string]interface{}{
		"user_id": userID,
		"times":   times},
	)
}

// SetGroupKick
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param rejectAddRequest bool
*/
func (b *Bot) SetGroupKick(groupID int, userID int, rejectAddRequest bool) {
	b.CallApi("set_group_kick", map[string]interface{}{
		"group_id":           groupID,
		"user_id":            userID,
		"reject_add_request": rejectAddRequest},
	)
}

// SetGroupAnonymousBan
/*
   @Description:
   @receiver b
   @param groupID int
   @param flag string
   @param duration int
*/
func (b *Bot) SetGroupAnonymousBan(groupID int, flag string, duration int) {
	b.CallApi("set_group_anonymous_ban", map[string]interface{}{
		"group_id": groupID,
		"flag":     flag,
		"duration": duration},
	)
}

// SetGroupWholeBan
/*
   @Description:
   @receiver b
   @param groupID int
   @param enable bool
*/
func (b *Bot) SetGroupWholeBan(groupID int, enable bool) {
	b.CallApi("set_group_whole_ban", map[string]interface{}{"group_id": groupID, "enable": enable})
}

// SetGroupAdmin
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param enable bool
*/
func (b *Bot) SetGroupAdmin(groupID int, userID int, enable bool) {
	b.CallApi("set_group_admin", map[string]interface{}{"group_id": groupID, "user_id": userID, "enable": enable})
}

// SetGroupAnonymous
/*
   @Description:
   @receiver b
   @param groupID int
   @param enable bool
*/
func (b *Bot) SetGroupAnonymous(groupID int, enable bool) {
	b.CallApi("set_group_anonymous", map[string]interface{}{"group_id": groupID, "enable": enable})
}

// SetGroupName
/*
   @Description:
   @receiver b
   @param groupID int
   @param groupName string
*/
func (b *Bot) SetGroupName(groupID int, groupName string) {
	b.CallApi("set_group_name", map[string]interface{}{"group_id": groupID, "group_name": groupName})
}

// SetGroupLeave
/*
   @Description:
   @receiver b
   @param groupID int
   @param isDisMiss bool
*/
func (b *Bot) SetGroupLeave(groupID int, isDisMiss bool) {
	b.CallApi("set_group_leave", map[string]interface{}{"group_id": groupID, "is_dismiss": isDisMiss})
}

// SetGroupSpecialTitle
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param specialTitle string
   @param duration int
*/
func (b *Bot) SetGroupSpecialTitle(groupID int, userID int, specialTitle string, duration int) {
	b.CallApi("set_group_special_title", map[string]interface{}{"group_id": groupID, "user_id": userID, "special_title": specialTitle, "duration": duration})
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
	b.CallApi("set_friend_add_request", map[string]interface{}{"flag": flag, "approve": approve, "remark": remark})
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
	b.CallApi("set_group_add_request", map[string]interface{}{
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
func (b *Bot) GetLoginInfo() gjson.Result {
	return b.CallApi("get_login_info", nil)
}

// GetStrangerInfo
/*
   @Description:
   @receiver b
   @param userID int
   @param noCache bool
   @return Senders
*/
func (b *Bot) GetStrangerInfo(userID int, noCache bool) gjson.Result {
	return b.CallApi("get_stranger_info", map[string]interface{}{"user_id": userID, "no_cache": noCache})
}

// GetFriendList
/**
 * @Description:
 * @receiver b
 * @return gjson.Result
 * example
 */
func (b *Bot) GetFriendList() gjson.Result {
	return b.CallApi("get_friend_list", nil)
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
func (b *Bot) GetGroupInfo(groupID int, noCache bool) gjson.Result {
	return b.CallApi("get_group_info", map[string]interface{}{"group_id": groupID, "no_cache": noCache})
}

// GetGroupList
/*
   @Description:
   @receiver b
   @return []GroupInfo
*/
func (b *Bot) GetGroupList() gjson.Result {
	return b.CallApi("get_group_list", nil)
}

// GetGroupMemberInfo
/*
   @Description:
   @receiver b
   @param groupID int
   @param userID int
   @param noCache bool
   @return GroupMemberInfo
*/
func (b *Bot) GetGroupMemberInfo(groupID int, userID int, noCache bool) gjson.Result {
	return b.CallApi("get_group_member_info", map[string]interface{}{"group_id": groupID, "user_id": userID, "no_cache": noCache})
}

// GetGroupMemberList
/*
   @Description:
   @receiver b
   @param groupID int
   @return []GroupMemberInfo
*/
func (b *Bot) GetGroupMemberList(groupID int) gjson.Result {
	return b.CallApi("get_group_member_list", map[string]interface{}{"group_id": groupID})
}

// GetGroupHonorInfo
/*
   @Description:
   @receiver b
   @param groupID int
   @param honorType string
   @return GroupHonorInfo
*/
func (b *Bot) GetGroupHonorInfo(groupID int, honorType string) gjson.Result {
	return b.CallApi("get_stranger_info", map[string]interface{}{"group_id": groupID, "honor_type": honorType})
}

// GetCookies
/*
   @Description:
   @receiver b
   @param domain string
   @return Cookie
*/
func (b *Bot) GetCookies(domain string) gjson.Result {
	return b.CallApi("get_cookies", map[string]interface{}{"domain": domain})
}

// GetCsrfToken
/*
   @Description:
   @receiver b
   @return CsrfToken
*/
func (b *Bot) GetCsrfToken() gjson.Result {
	return b.CallApi("get_csrf_token", nil)
}

// GetCredentials
/*
   @Description:
   @receiver b
   @param domain string
   @return Credentials
*/
func (b *Bot) GetCredentials(domain string) gjson.Result {
	return b.CallApi("get_credentials", map[string]interface{}{"domain": domain})
}

// GetRecord
/*
   @Description:
   @receiver b
   @param file file
   @param outFormat string
   @return Record
*/
func (b *Bot) GetRecord(file, outFormat string) gjson.Result {
	return b.CallApi("get_record", map[string]interface{}{"file": file, "out_format": outFormat})
}

// GetImage
/*
   @Description:
   @receiver b
   @param file string
   @return Image
*/
func (b *Bot) GetImage(file string) gjson.Result {
	return b.CallApi("get_image", map[string]interface{}{"file": file})
}

// CanSendImage
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendImage() bool {
	return b.CallApi("can_send_image", nil).Bool()
}

// CanSendRecord
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendRecord() bool {
	return b.CallApi("can_send_record", nil).Bool()
}

// GetStatus
/*
   @Description:
   @receiver b
   @return OnlineStatus
*/
func (b *Bot) GetStatus() gjson.Result {
	return b.CallApi("get_status", nil)
}

// SetRestart
/*
   @Description:
   @receiver b
   @param delay int
*/
func (b *Bot) SetRestart(delay int) {
	b.CallApi("set_restart", map[string]interface{}{"delay": delay})
}

// CleanCache
/**
 * @Description:
 * @receiver b
 * example
 */
func (b *Bot) CleanCache() {
	b.CallApi("clean_cache", nil)
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
func (b *Bot) DownloadFile(url string, threadCount int, headers []string) gjson.Result {
	return b.CallApi("download_file", map[string]interface{}{
		"url":          url,
		"thread_count": threadCount,
		"headers":      headers})
}

// GetGroupMsgHistory
/*
   @Description:
   @receiver b
   @param messageSeq int64
   @param groupID int
   @return MessageHistory
*/
func (b *Bot) GetGroupMsgHistory(messageSeq int64, groupID int) gjson.Result {
	return b.CallApi("get_group_msh_history", map[string]interface{}{"message_seq": messageSeq, "group_id": groupID})
}

// GetOnlineClients
/**
 * @Description:
 * @receiver b
 * @param noCache
 * @return gjson.Result
 * example
 */
func (b *Bot) GetOnlineClients(noCache bool) gjson.Result {
	return b.CallApi("get_online_clients", map[string]interface{}{"no_cache": noCache})
}

// GetVipInfoTest
/*
   @Description:
   @receiver b
   @param UserId int
   @return VipInfo
*/
func (b *Bot) GetVipInfoTest(userID int) gjson.Result {
	return b.CallApi("_get_vip_info", map[string]interface{}{"user_id": userID})
}

// SendGroupNotice
/*
   @Description:
   @receiver b
   @param groupID int
   @param content string
*/
func (b *Bot) SendGroupNotice(groupID int, content string) {
	b.CallApi("send_group_notice", map[string]interface{}{"group_id": groupID, "content": content})
}

// ReloadEventFilter
/*
   @Description:
   @receiver b
*/
func (b *Bot) ReloadEventFilter() {
	b.CallApi("reload_event_filter", nil)
}

// SetGroupNameSpecial
/*
   @Description:
   @receiver b
   @param groupID int
   @param groupName string
*/
func (b *Bot) SetGroupNameSpecial(groupID int, groupName string) {
	b.CallApi("set_group_name", map[string]interface{}{"group_id": groupID, "group_name": groupName})
}

// SetGroupPortrait
/*
   @Description:
   @receiver b
   @param groupID int
   @param file string
   @param cache int
*/
func (b *Bot) SetGroupPortrait(groupID int, file string, cache int) {
	b.CallApi("set_group_portrait", map[string]interface{}{"group_id": groupID, "file": file, "cache": cache})
}

// GetMsgSpecial
/*
   @Description:
   @receiver b
   @param messageID int
   @return MsgData
*/
func (b *Bot) GetMsgSpecial(messageID int) gjson.Result {
	return b.CallApi("get_msg", map[string]interface{}{"message_id": messageID})
}

// GetForwardMsg
/*
   @Description:
   @receiver b
   @param messageID int
   @return []ForwardMsg
*/
func (b *Bot) GetForwardMsg(messageID int) gjson.Result {
	return b.CallApi("get_forward_msg", map[string]interface{}{"message_id": messageID})
}

// SendGroupForwardMsg
/*
   @Description:
   @receiver b
   @param groupID int
   @param messages []Node
*/
func (b *Bot) SendGroupForwardMsg(groupID int, messages interface{}) {
	b.CallApi("send_group_forward_msg", map[string]interface{}{"group_id": groupID, "message": messages})
}

// GetWordSlices
/*
   @Description:
   @receiver b
   @param content string
   @return []string
*/
func (b *Bot) GetWordSlices(content string) gjson.Result {
	return b.CallApi(".get_word_slices", map[string]interface{}{"content": content})
}

// OcrImage
/*
   @Description:
   @receiver b
   @param image string
   @return OcrImage
*/
func (b *Bot) OcrImage(image string) gjson.Result {
	return b.CallApi("ocr_image", map[string]interface{}{"image": image})
}

// GetGroupSystemMsg
/*
   @Description:
   @receiver b
   @return GroupSystemMsg
*/
func (b *Bot) GetGroupSystemMsg() gjson.Result {
	return b.CallApi("get_group_system_msg", nil)
}

// GetGroupFileSystemInfo
/*
   @Description:
   @receiver b
   @param groupID int
   @return GroupFileSystemInfo
*/
func (b *Bot) GetGroupFileSystemInfo(groupID int) gjson.Result {
	return b.CallApi("get_group_file_system_info", map[string]interface{}{"group_id": groupID})
}

// GetGroupRootFiles
/*
   @Description:
   @receiver b
   @param groupID int
   @return GroupRootFiles
*/
func (b *Bot) GetGroupRootFiles(groupID int) gjson.Result {
	return b.CallApi("get_group_root_files", map[string]interface{}{"group_id": groupID})
}

// GetGroupFilesByFolder
/*
   @Description:
   @receiver b
   @param groupID int
   @param folderId string
   @return GroupFilesByFolder
*/
func (b *Bot) GetGroupFilesByFolder(groupID int, folderId string) gjson.Result {
	return b.CallApi("get_group_files_by_folder", map[string]interface{}{"group_id": groupID, "folder_id": folderId})
}

// GetGroupFileUrl
/*
   @Description:
   @receiver b
   @param groupID int
   @param fileId string
   @param busid int
   @return FileUrl
*/
func (b *Bot) GetGroupFileUrl(groupID int, fileId string, busid int) gjson.Result {
	return b.CallApi("get_group_file_url", map[string]interface{}{"group_id": groupID, "file_id": fileId, "busid": busid})
}

// GetGroupAtAllRemain
/*
   @Description:
   @receiver b
   @param groupID int
   @return GroupAtAllRemain
*/
func (b *Bot) GetGroupAtAllRemain(groupID int) gjson.Result {
	return b.CallApi("get_group_at_all_remain", map[string]interface{}{"group_id": groupID})
}

// UploadGroupFile
/*
   @Description:
   @receiver b
   @param groupID int
   @param file string
   @param name string
   @param folder string
*/
func (b *Bot) UploadGroupFile(groupID int, file string, name string, folder string) {
	b.CallApi("upload_group_file", map[string]interface{}{"group_id": groupID, "file": file, "name": name, "folder": folder})
}

func (b *Bot) CallApi(action string, param interface{}) gjson.Result {
	echo := uuid.NewV4().String()
	var d = userAPi{
		Action: action,
		Params: param,
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	content, _ := json.Marshal(d)
	log.Infoln(string(content) + "\n\t\t\t\t\t" + string(data))
	if gjson.GetBytes(data, "status").String() != "ok" {
		log.Errorln("调用API出现了错误")
		log.Panicln(gjson.GetBytes(data, "msg"), ",", gjson.GetBytes(data, "wording"))
	}
	return gjson.GetBytes(content, "data")
}

func (b *Bot) SetEssenceMsg(messageID int) {
	b.CallApi("set_essence_msg", map[string]interface{}{"message_id": messageID})
}

func (b *Bot) DeleteEssenceMsg(messageID int) {
	b.CallApi("delete_essence_msg", map[string]interface{}{"message_id": messageID})
}

func (b *Bot) GetEssenceMsgList(groupID int) gjson.Result {
	return b.CallApi("get_essence_msg_list", map[string]interface{}{"group_id": groupID})
}

func (b *Bot) CheckUrlSafely(url string) int {
	return int(b.CallApi("check_url_safely", map[string]interface{}{"url": url}).Int())
}
