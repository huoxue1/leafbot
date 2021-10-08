package cqhttp_ws_driver

import (
	"encoding/json" //nolint:gci
	"fmt"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	message2 "github.com/huoxue1/leafBot/message"
)

type UseApi struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

// SendGroupMsg
/**
 * @Description:
 * @receiver b
 * @param groupId
 * @param message
 * @return int32
 */
func (b *Bot) SendGroupMsg(groupId int, message interface{}) int32 {
	echo := uuid.NewV4().String()
	if _, ok := message.(string); ok {
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	type param struct {
		GroupId int         `json:"group_id"`
		Message interface{} `json:"message"`
	}
	var d = UseApi{
		Action: "send_group_msg",
		Params: param{
			GroupId: groupId,
			Message: message,
		},
		Echo: echo}

	b.Do(d)

	data, _ := b.GetResponse(echo)

	log.Infoln("call api ==> " + d.Action)
	log.Debugln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return int32(gjson.GetBytes(data, "data.message_id").Int())
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
func (b *Bot) SendPrivateMsg(userId int, message interface{}) int32 {
	echo := uuid.NewV4().String()
	switch message.(type) {
	case string:
		{
			message = message2.ParseMessageFromString(message.(string))
		}
	}
	type param struct {
		UserId  int         `json:"user_id"`
		Message interface{} `json:"message"`
	}
	var d = UseApi{
		Action: "send_private_msg",
		Params: param{
			UserId:  userId,
			Message: message,
		},
		Echo: echo}

	b.Do(d)
	data, _ := b.GetResponse(echo)

	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return int32(gjson.GetBytes(data, "data.message_id").Int())
}

// DeleteMsg
/*
   @Description:
   @receiver b
   @param messageId int32
*/
func (b *Bot) DeleteMsg(messageId int32) {
	echo := uuid.NewV4().String()
	type param struct {
		MessageId int32 `json:"message_id"`
	}
	var d = UseApi{
		Action: "delete_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo}

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// GetMsg
/*
   @Description:
   @receiver b
   @param messageId int32
   @return GetMessage
*/
func (b *Bot) GetMsg(messageId int32) gjson.Result {
	echo := uuid.NewV4().String()
	type param struct {
		MessageId int32 `json:"message_id"`
	}
	var d = UseApi{
		Action: "get_msg",
		Params: param{
			MessageId: messageId,
		},
		Echo: echo}

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
	echo := uuid.NewV4().String()
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

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
func (b *Bot) SendMsg(messageType string, userId int, groupId int, message interface{}) int32 {
	if messageType == "group" {
		return b.SendGroupMsg(groupId, message)
	} else {
		return b.SendPrivateMsg(userId, message)
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
	echo := uuid.NewV4().String()
	type param struct {
		UserId int `json:"user_id"`
		Times  int `json:"times"`
	}
	var d = UseApi{
		Action: "Do()_like",
		Params: param{
			UserId: userId,
			Times:  times,
		},
		Echo: echo}

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// SetGroupWholeBan
/*
   @Description:
   @receiver b
   @param groupId int
   @param enable bool
*/
func (b *Bot) SetGroupWholeBan(groupId int, enable bool) {
	echo := uuid.NewV4().String()
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

	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// SetGroupAnonymous
/*
   @Description:
   @receiver b
   @param groupId int
   @param enable bool
*/
func (b *Bot) SetGroupAnonymous(groupId int, enable bool) {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// SetGroupName
/*
   @Description:
   @receiver b
   @param groupId int
   @param groupName string
*/
func (b *Bot) SetGroupName(groupId int, groupName string) {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// SetGroupLeave
/*
   @Description:
   @receiver b
   @param groupId int
   @param isDisMiss bool
*/
func (b *Bot) SetGroupLeave(groupId int, isDisMiss bool) {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
	type param struct {
		Flag    string `json:"flag"`
		Approve bool   `json:"approve"`
		Remark  string `json:"remark"`
	}
	var d = UseApi{
		Action: "set_friend_add_request",
		Params: param{
			Flag:    flag,
			Approve: approve,
			Remark:  remark,
		},
		Echo: echo}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// GetLoginInfo
/*
   @Description:
   @receiver b
   @return LoginInfo
*/
func (b *Bot) GetLoginInfo() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_login_info",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetStrangerInfo
/*
   @Description:
   @receiver b
   @param userId int
   @param noCache bool
   @return Senders
*/
func (b *Bot) GetStrangerInfo(userId int, noCache bool) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetFriendList
/**
 * @Description:
 * @receiver b
 * @return gjson.Result
 * example
 */
func (b *Bot) GetFriendList() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_friend_list",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupInfo
/**
 * @Description:
 * @receiver b
 * @param groupId
 * @param noCache
 * @return gjson.Result
 * example
 */
func (b *Bot) GetGroupInfo(groupId int, noCache bool) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupList
/*
   @Description:
   @receiver b
   @return []GroupInfo
*/
func (b *Bot) GetGroupList() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_group_list",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
func (b *Bot) GetGroupMemberInfo(groupId int, userId int, noCache bool) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupMemberList
/*
   @Description:
   @receiver b
   @param groupId int
   @return []GroupMemberInfo
*/
func (b *Bot) GetGroupMemberList(groupId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupHonorInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @param honorType string
   @return GroupHonorInfo
*/
func (b *Bot) GetGroupHonorInfo(groupId int, honorType string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetCookies
/*
   @Description:
   @receiver b
   @param domain string
   @return Cookie
*/
func (b *Bot) GetCookies(domain string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetCsrfToken
/*
   @Description:
   @receiver b
   @return CsrfToken
*/
func (b *Bot) GetCsrfToken() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_csrf_token",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetCredentials
/*
   @Description:
   @receiver b
   @param domain string
   @return Credentials
*/
func (b *Bot) GetCredentials(domain string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
	echo := uuid.NewV4().String()
	type param struct {
		File      string `json:"file"`
		OutFormat string `json:"out_format"`
	}
	var d = UseApi{
		Action: "get_record",
		Params: param{
			File:      file,
			OutFormat: outFormat,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetImage
/*
   @Description:
   @receiver b
   @param file string
   @return Image
*/
func (b *Bot) GetImage(file string) gjson.Result {
	echo := uuid.NewV4().String()
	type param struct {
		File string `json:"file"`
	}
	var d = UseApi{
		Action: "get_image",
		Params: param{
			File: file,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// CanSendImage
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendImage() bool {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "can_Do()_image",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data").Bool()
}

// CanSendRecord
/*
   @Description:
   @receiver b
   @return Bool
*/
func (b *Bot) CanSendRecord() bool {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "can_Do()_record",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data").Bool()
}

// GetStatus
/*
   @Description:
   @receiver b
   @return OnlineStatus
*/
func (b *Bot) GetStatus() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_status",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// SetRestart
/*
   @Description:
   @receiver b
   @param delay int
*/
func (b *Bot) SetRestart(delay int) {
	echo := uuid.NewV4().String()
	type param struct {
		Delay int `json:"delay"`
	}
	var d = UseApi{
		Action: "set_restart",
		Params: param{
			Delay: delay,
		},
		Echo: echo}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// CleanCache
/**
 * @Description:
 * @receiver b
 * example
 */
func (b *Bot) CleanCache() {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "clean_cache",
		Echo:   echo}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
func (b *Bot) DownloadFile(url string, threadCount int, headers []string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupMsgHistory
/*
   @Description:
   @receiver b
   @param messageSeq int64
   @param groupId int
   @return MessageHistory
*/
func (b *Bot) GetGroupMsgHistory(messageSeq int64, groupId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetVipInfoTest
/*
   @Description:
   @receiver b
   @param UserId int
   @return VipInfo
*/
func (b *Bot) GetVipInfoTest(UserId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// SendGroupNotice
/*
   @Description:
   @receiver b
   @param groupId int
   @param content string
*/
func (b *Bot) SendGroupNotice(groupId int, content string) {
	echo := uuid.NewV4().String()
	type param struct {
		GroupId int    `json:"group_id"`
		Content string `json:"content"`
	}
	var d = UseApi{
		Action: "Do()_group_notice",
		Params: param{
			GroupId: groupId,
			Content: content,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// ReloadEventFilter
/*
   @Description:
   @receiver b
*/
func (b *Bot) ReloadEventFilter() {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "reload_event_filter",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// SetGroupNameSpecial
/*
   @Description:
   @receiver b
   @param groupId int
   @param groupName string
*/
func (b *Bot) SetGroupNameSpecial(groupId int, groupName string) {
	echo := uuid.NewV4().String()
	type param struct {
		GroupId   int    `json:"group_id"`
		GroupName string `json:"group_name"`
	}
	var d = UseApi{
		Action: "Do()_group_name",
		Params: param{
			GroupId:   groupId,
			GroupName: groupName,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
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
	echo := uuid.NewV4().String()
	type param struct {
		GroupId int    `json:"group_id"`
		File    string `json:"file"`
		Cache   int    `json:"cache"`
	}
	var d = UseApi{
		Action: "Do()_group_portrait",
		Params: param{
			GroupId: groupId,
			File:    file,
			Cache:   cache,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// GetMsgSpecial
/*
   @Description:
   @receiver b
   @param messageId int
   @return MsgData
*/
func (b *Bot) GetMsgSpecial(messageId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetForwardMsg
/*
   @Description:
   @receiver b
   @param messageId int
   @return []ForwardMsg
*/
func (b *Bot) GetForwardMsg(messageId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// SendGroupForwardMsg
/*
   @Description:
   @receiver b
   @param groupId int
   @param messages []Node
*/
func (b *Bot) SendGroupForwardMsg(groupId int, messages interface{}) {
	echo := uuid.NewV4().String()
	type param struct {
		GroupId  int         `json:"group_id"`
		Messages interface{} `json:"messages"`
	}
	var d = UseApi{
		Action: "Do()_group_forward_msg",
		Params: param{
			GroupId:  groupId,
			Messages: messages,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

// GetWordSlices
/*
   @Description:
   @receiver b
   @param content string
   @return []string
*/
func (b *Bot) GetWordSlices(content string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// OcrImage
/*
   @Description:
   @receiver b
   @param image string
   @return OcrImage
*/
func (b *Bot) OcrImage(image string) gjson.Result {
	echo := uuid.NewV4().String()
	type param struct {
		Image string `json:"image"`
	}
	var d = UseApi{
		Action: "ocr_image",
		Params: param{
			Image: image,
		},
		Echo: echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupSystemMsg
/*
   @Description:
   @receiver b
   @return GroupSystemMsg
*/
func (b *Bot) GetGroupSystemMsg() gjson.Result {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: "get_group_system_msg",
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupFileSystemInfo
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupFileSystemInfo
*/
func (b *Bot) GetGroupFileSystemInfo(groupId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupRootFiles
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupRootFiles
*/
func (b *Bot) GetGroupRootFiles(groupId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupFilesByFolder
/*
   @Description:
   @receiver b
   @param groupId int
   @param folderId string
   @return GroupFilesByFolder
*/
func (b *Bot) GetGroupFilesByFolder(groupId int, folderId string) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
func (b *Bot) GetGroupFileUrl(groupId int, fileId string, busid int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
}

// GetGroupAtAllRemain
/*
   @Description:
   @receiver b
   @param groupId int
   @return GroupAtAllRemain
*/
func (b *Bot) GetGroupAtAllRemain(groupId int) gjson.Result {
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
	return gjson.GetBytes(data, "data")
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
	echo := uuid.NewV4().String()
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
	b.Do(d)
	data, _ := b.GetResponse(echo)
	log.Infoln("call api ==> " + d.Action)
	log.Infoln(fmt.Sprintf("the params ==> %v", d.Params))
	log.Infoln("the response ==> " + string(data))
}

func (b *Bot) CallApi(Action string, param interface{}) interface{} {
	echo := uuid.NewV4().String()
	var d = UseApi{
		Action: Action,
		Params: param,
		Echo:   echo,
	}
	b.Do(d)
	data, _ := b.GetResponse(echo)
	content, _ := json.Marshal(d)
	log.Infoln(string(content) + "\n\t\t\t\t\t" + string(data))
	return content
}

func (b *Bot) SetEssenceMsg(messageId int) {
	panic("implement me")
}

func (b *Bot) DeleteEssenceMsg(messageId int) {
	panic("implement me")
}

func (b *Bot) GetEssenceMsgList(groupId int) {
	panic("implement me")
}

func (b *Bot) CheckUrlSafely(url string) int {
	panic("implement me")
}
