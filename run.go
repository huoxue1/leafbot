package leafBot

import (
	log "github.com/sirupsen/logrus"
)

// run
/*
   @Description:
   @param bot *Bot
*/
func run(bot *Bot) {
	log.Infoln("已加载bot：" + bot.Name)

}
