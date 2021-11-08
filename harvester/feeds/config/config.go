//Import a feed config file for use by the scheduler and feeds.
//Files live in feeds.d and are in YAML.
package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//As Go uses marshalling this should be capable of representing any feed.
type FeedConfig struct {
	Format      string
	Frequency   string          //Google style duration or "Once" to run fetch immediately and return
	Credentials map[string]string   `yaml:",omitempty"`
	Source      string              `yaml:"feed source"`
	Subs        map[string][]string `yaml:"sub feeds,omitempty"`
	Paths       []map[string]string
	Params      map[string]interface{} `yaml:",omitempty"` //extra params for specific fetchers
}

//Split out for ease of testing
func (fc *FeedConfig) Populate(config []byte) {
	//log.Println(string(config))

	err := yaml.Unmarshal([]byte(config), &fc)
	if err != nil {
		log.Panicf("error: %v", err)
	}
	//log.Println(fc)
	//panic(err)
}

func (fc *FeedConfig) ToBytes() ([]byte, error) {
	if fc.Format == "" {
		return nil, errors.New("Config file must have a Format")
	}
	if fc.Frequency == "" {
		return nil, errors.New("Config file must have a Frequency")
	}
	if fc.Source == "" {
		return nil, errors.New("Config file must have a Source")
	}
	if fc.Paths == nil {
		return nil, errors.New("Config file must have at least one Path")
	}
	confYAML, err := yaml.Marshal(fc)
	if err != nil {
		log.Println("Config Format Error : ", err)
		return nil, err
	}
	return confYAML, nil
}

//Inserts one credential value into the feed source path, if the key is in the path
func (fc *FeedConfig) ApplyCredentials() {
	if fc.Credentials == nil {
		return
	}
	sourcePath := strings.Split(fc.Source, "|")
  if len(sourcePath) < 3 {
    return
  }

	//Very minimal could put this in a for loop for multiple creds
	fc.Source = fmt.Sprintf("%s%s%s", sourcePath[0], fc.Credentials[sourcePath[1]], sourcePath[2])
}

/*
 * Allow moving dates based on today
 *
 * Expandable Date string is TODAY-15D or TODAY+03M
 * It must start with TODAY followed by an optional offest
 * an offset comprises +/- then a 2 digit number
 * and units of either D (days) or M (months)
 *
 * There can be more than one expandable date in a URL
*/
func (fc *FeedConfig) ExpandDates() {
  r, err := regexp.Compile(`TODAY([+-]?\w*)`)
  if err != nil {
    log.Println("ExpandDates Regex  Error : ", err)
    return
  }
  substrs := r.FindAllStringIndex(fc.Source, -1)
  if len(substrs) == 0 {
    return
  }

  src := fc.Source
  newUrl := ""
  subStrStart := 0
  for _, slice := range substrs {
    datStr1 := src[slice[0]:slice[1]]
    delta_str, err := DateDelta(datStr1)
    if err != nil {
      return
    }
    newUrl = newUrl + src[subStrStart:slice[0]] + delta_str
    subStrStart = slice[1]
  }
  fc.Source = newUrl + src[subStrStart:]
}

func DateDelta(deltaStr string) (string, error) {
  t := time.Now()

  if deltaStr != "TODAY" {
    num, err := strconv.Atoi(deltaStr[6:8])
    if err != nil {
      log.Println("ExpandDates Invalid date in source url  Error : ", err)
      return "", err
    }
    switch deltaStr[8] {
    case 'D':
      if deltaStr[5] == '+' {
        t = t.AddDate(0, 0, num)
      } else {
        t = t.AddDate(0, 0, -num)
      }
    case 'M':
      if deltaStr[5] == '+' {
        t = t.AddDate(0, num, 0)
      } else {
        t = t.AddDate(0, -num, 0)
      }
    default:
      log.Println("ExpandDates Invalid date in source url  Error : ", err)
      return "", err
    }
  }

  tStr := t.Format("2006/01/02")
  return tStr, nil
}
/*
func (fc *FeedConfig)ExpandSource() {
  //expandedPaths []map[string]string



    for key, val := range fc.Subs {
      for _, param := range val {
        source := fmt.Sprintf("%s%s=%s", fc.Source, key, param)
      }
    }

}
*/

//Top Level
func (fc *FeedConfig) Read(config_file string) {
	config, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.Panicf("Failed to read config file : error: %v", err)
	}
	fc.Populate(config)
	fc.ApplyCredentials()
  fc.ExpandDates()
}
