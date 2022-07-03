package cqhttp_default_driver

import (
	"encoding/json" //nolint:gci

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// SendGroupMsg
/* @Description:
 * @receiver b
 * @param groupID
 * @param message
 * @return int32
 */
func (b *Bot) SendGroupMsg(groupID int64, message interface{}) int32 {
	return int32(b.CallApi("send_group_msg", map[string]interface{}{"group_id": groupID, "message": message}).Int())
}

// SendPrivateMsg
/* @Description:
 * @receiver b
 * @param userID
 * @param message
 * @return int32
 */
func (b *Bot) SendPrivateMsg(userID int64, message interface{}) int32 {
	return int32(b.CallApi("send_private_msg", map[string]interface{}{"user_id": userID, "message": message}).Int())
}

// CallApi
/* @Description:
 * @receiver b
 * @param action
 * @param param
 * @return gjson.Result
 */
func (b *Bot) CallApi(action string, param interface{}) gjson.Result {
	echo := uuid.NewV4().String()
	var d = userAPI{
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
