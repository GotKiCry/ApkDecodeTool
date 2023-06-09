package sdktool

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	Assets []string `yaml:"assets"`
	Libs   []string `yaml:"libs"`
	Res    []string `yaml:"res"`
	Smali  []string `yaml:"smali"`
}

// 全局变量
var rule *Rule

// loadFilterByLocal 从本地初始化规则
func loadFilterByLocal() error {
	localPath, _ := os.Getwd()
	rulePath := localPath + "\\assets\\skipRule.yaml"
	log.Println(rulePath)
	data, readError := os.ReadFile(rulePath)
	if readError != nil {
		return readError
	}
	//log.Println(data)
	if err := yaml.Unmarshal(data, &rule); err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("规则内容 ：%s", rule)
	return nil
}

func isSkip(ruleName string, targetValue string) bool {

	if rule == nil {
		err := loadFilterByLocal()
		if err != nil {
			log.Fatalf("无法初始化规则 %e \r\n", err)
			return false
		}
	}
	var ruleList []string
	switch ruleName {
	case "assets":
		ruleList = rule.Assets
	case "libs":
		ruleList = rule.Libs
	case "res":
		ruleList = rule.Res
	case "smali":
		ruleList = rule.Smali
	}

	for _, content := range ruleList {
		matchRule := strings.Replace(content, "*", ".+", -1)
		matchRule = strings.Replace(matchRule, "\\", "\\\\", -1)
		matchRule = strings.Replace(matchRule, "$", "\\$", -1)
		result, err := regexp.MatchString(matchRule, targetValue)
		if err != nil {
			log.Printf("Error = %s \r\n", err)
			log.Printf("targetValue = %s,content = %s", targetValue, content)
		}

		if result {
			log.Printf("跳过文件 %s", targetValue)
			return true
		}
		continue
	}
	return false
}
