package datastore

/*
Basically a path based cache

if you pass true to Init(true) it will use the
datawarehouse to back the cache for straight Sets, Gets and Dels
*/

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/redis.v5"

	dw "github.com/derbexuk/wurzel/harvester/datawarehouse"
)

var UpdateChannel = "feedUpdates"
var redisClient *redis.Client

var dbh *dw.Datawarehouse

var DB = "duff"
const COL = "cache"

//Sets up a singleton client
//Will keep retrying
func Init(params ...bool) {
	tries := 1
	redisAddr := "redis:6379"
	ok := os.Getenv("REDIS")
	if ok != "" {
		redisAddr = ok
	}
	for redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:        redisAddr,
			Password:    "", // no password set
			DB:          0,  // use default DB
			MaxRetries:  5,
			DialTimeout: time.Second * 10,
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			log.Println(err)
			if tries > 4 {
				log.Panic("Redis connection failied")
				log.Panic(err)
			}
			time.Sleep(time.Second * time.Duration(tries))
			tries++

			redisClient = nil
		}
	}
	if len(params) > 0 {
		//Don't spawn
		if params[0] == true {
			log.Println("Use DB ", DB)
			if s := os.Getenv("SERVICE"); s != "" {
				DB = s
			}
			if dbh == nil {
				dbh = &dw.Datawarehouse{}
				dbh.Open()
			}
		}
	}
}

func Set(path, data string) error {
	return SetExp(path, data, 0)
}

func SetExp(path, data string, ttl int) error {
	expiry := time.Duration(ttl) * time.Second
	err := redisClient.Set(path, data, expiry).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	if dbh != nil {
		//dw stores objects
		redisData := map[string]interface{}{"stuff": data}
		if ttl != 0 {
			redisData["ttl"] = ttl
		}
		dbh.Upsert(DB, COL, path, redisData)
	}
	return err
}

func Del(path string) error {
	i, err := redisClient.Del(path).Result()
	if err != nil {
		log.Println(err)
	}
	if i == 0 {
		log.Printf("DS failed to delete from redis : path %s\n", path)
	}
	if dbh != nil && i > 0 {
		i := dbh.DelPath(DB, COL, path)
		if i == 0 {
			log.Printf("DS failed to delete : db %s col %s path %s\n", DB, COL, path)
		}
	}
	//We seem to suffer a race condition between db Del and redis get
	//This is an attempt to avoid implementing locking
	i, _ = redisClient.Del(path).Result()
	if i != 0 {
		log.Printf("DS secondary delete from redis : path %s\n", path)
	}
	return err
}

func Get(path string) (string, error) {
	val, err := redisClient.Get(path).Result()
	if err == redis.Nil && dbh != nil {
		//Go and look for it in Mongo
		result := dbh.Get(DB, COL, path)
		if result != nil {
			if result["ttl"] != nil {
				ttl := result["ttl"].(int32)
				exp := time.Duration(ttl) * time.Second
				err = redisClient.Set(path, result["stuff"].(string), exp).Err()
			} else {
				err = redisClient.Set(path, result["stuff"].(string), 0).Err()
			}
			if err != nil {
				log.Println(err)
			}
			return result["stuff"].(string), nil
		}
		err = errors.New("Path Not Found : " + path)
	} else if err == redis.Nil {
		err = errors.New("Path Not Found : " + path)
	} else if err != nil {
		log.Println(err)
	} else {
		return val, nil
	}
	return "", err
}

//This is OK for small dbs, but slow for large ones
func GetPaths(path string) ([]string, error) {
	log.Println("Path : " + path)
	val, err := redisClient.Keys(path).Result()
	log.Printf("Val %v, err %v\n", val, err)
	if len(val) == 0 && dbh != nil {
		//Go and look for it in Mongo
		log.Println("MONGO PATH : " + strings.Split(path, "*")[0])
		ress := dbh.GetLike(DB, COL, strings.Split(path, "*")[0])
		if len(ress) == 0 {
			err = errors.New("Path Not Found : " + path)
		} else {
			for _, res := range ress {
				log.Println(res.Path)
				err = redisClient.Set(res.Path, res.Data["stuff"].(string), 0).Err()
				if err != nil {
					log.Println(err)
				}
			}
			log.Println("Path : " + path)
			val, err = redisClient.Keys(path).Result()
			log.Printf("Val %v, err %v\n", val, err)
		}
	} else if len(val) == 0 {
		err = errors.New("Path Not Found : " + path)
	} else if err != nil {
		log.Println(err)
	}
	return val, err
}

//A stack that can test for membership
func AddToSeriesIndex(path string, v string) error {
	var err error
	_, err = redisClient.SAdd("/index/members/"+path, v).Result()
	if err != nil {
		log.Println(err)
	}
	_, err = redisClient.LPush("/index/list/"+path, v).Result()
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetSeriesRange(path string, start, stop int64) ([]string, error) {
	r, err := redisClient.LRange("/index/list/"+path, start, stop).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("RANGE : ", r)
	return r, err
}

func InIndex(path string, value string) (bool, error) {
	r, err := redisClient.SIsMember("/index/members/"+path, value).Result()
	return r, err
}

//Hashes
//According to the docs data ashould be map[string]interface{} but this errors
func MapSet(path string, data map[string]string) error {
	err := redisClient.HMSet(path, data).Err()
	if err != nil {
		log.Println(err)
	}
	return err
}

func MapElSet(path, key string, val interface{}) error {
	err := redisClient.HSet(path, key, val).Err()
	if err != nil {
		log.Println(err)
	}
	return err
}

func MapElGet(path, key string) (string, error) {
	val, err := redisClient.HGet(path, key).Result()
	if err != nil {
		if err == redis.Nil {
			err = errors.New("Not Found : " + path + " or " + key)
			return "", err
		}
		return "", err
	}
	return val, nil
}

func MapElDel(path, key string) error {
	_, err := redisClient.HDel(path, key).Result()
	return err
}

func Subscribe(path string) *redis.PubSub {
	ps, err := redisClient.Subscribe(path)
	if err != nil {
		log.Println(err)
	}
	return ps
}
func SubscribeUpdates() *redis.PubSub {
	return Subscribe(UpdateChannel)
}

func Publish(path string, stuff string) {
	err := redisClient.Publish(path, stuff).Err()
	if err != nil {
		log.Println(err)
	}
}

func PublishUpdate(path string) {
	Publish(UpdateChannel, path)
}
