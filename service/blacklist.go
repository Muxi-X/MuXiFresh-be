package service

//
// import (
// 	"time"
//
// 	"forum-gateway/log"
// 	"forum-gateway/model"
// 	m "forum/model"
//
// 	"github.com/spf13/viper"
// 	"go.uber.org/zap"
// )

// CheckInBlacklist checks whether a token is in blacklist
func CheckInBlacklist(token string) (bool, error) {
	// 从 redis 中查找是否有该 token 数据，若有，则在黑名单中
	// exist, err := m.HasExistedInRedis(token)
	// if err != nil {
	// log.Error("HasExistedInRedis error", zap.String("cause", err.Error()))
	// return false, err
	// }
	// return exist, nil
	return false, nil
}

//
// // AddToBlacklist adds a token into blacklist
// func AddToBlacklist(token string, expiresAt int64) error {
// 	// 加入 MySQL 黑名单
//
// 	record := &model.BlacklistModel{
// 		Token:     token,
// 		ExpiresAt: expiresAt,
// 	}
// 	if err := record.Create(); err != nil {
// 		return err
// 	}
//
// 	// 加入 redis 黑名单中
//
// 	// 过期时间(s)
// 	var expiration time.Duration = time.Duration(expiresAt-time.Now().Unix()) * time.Second
//
// 	if err := m.SetStringInRedis(token, 1, expiration); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // TidyBlacklist ... 定时清理过期的记录
// func TidyBlacklist() {
// 	tidyDay := viper.GetInt("blacklist.tidy_time")
// 	if tidyDay == 0 {
// 		log.Error("tidyDay failed to get")
// 		return
// 	}
//
// 	tidyDuration := time.Hour * time.Duration(tidyDay)
//
// 	for {
// 		if err := model.DeleteExpiredBlacklist(); err != nil {
// 			log.Error("TidyBlacklist error", zap.String("cause", err.Error()))
// 		}
//
// 		time.Sleep(tidyDuration)
// 	}
// }
//
// // SynchronizeBlacklistToRedis ... 同步 MySQL 黑名单到 redis 中，
// // 服务启动时调用
// func SynchronizeBlacklistToRedis() {
// 	// 从 MySQL 中读取所有数据
// 	records, err := model.GetAllBlacklist()
// 	if err != nil {
// 		log.Error("Getting blacklist records from DB failed", zap.String("cause", err.Error()))
// 		return
// 	}
//
// 	// 插入到 Redis 中
// 	// 因为数据不会很多，所以暂时还达不到 Redis 的性能瓶颈，若之后成为瓶颈，可以尝试使用 pipeline
//
// 	var expiration time.Duration
// 	for _, record := range records {
// 		// 过期时间(s)
// 		expiration = time.Duration((record.ExpiresAt - time.Now().Unix())) * time.Second
//
// 		if err := m.SetStringInRedis(record.Token, 1, expiration); err != nil {
// 			log.Error("SetStringInRedis function error", zap.String("cause", err.Error()))
// 			return
// 		}
// 	}
//
// 	log.Info("Synchronize blacklist successfully.")
// }
