package sdktool

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// DecodeStyleableSmali 解析一个Styleable文件/**
func DecodeStyleableSmali(values string) (map[string][]string, error) {
	styleableSmali, fileNotExistError := os.Open(values)
	if fileNotExistError != nil {
		return map[string][]string{}, fmt.Errorf("文件不存在，请检查")
	}

	all, err := ioutil.ReadAll(styleableSmali)
	if err != nil {
		print(err)
	}
	builder := strings.Builder{}
	builder.Write(all)
	s := builder.String()
	split := strings.Split(s, ".method")
	for _, str := range split {
		if strings.Index(str, "<clinit>") != -1 {
			s = str
			break
		} else {
			s = ""
		}
	}
	styleableMap := make(map[string][]string)

	//切割文件
	smaliData := strings.Split(s, "return-void")
	styleableData := strings.Split(smaliData[0], ":[I")
	attrsData := strings.Split(smaliData[1], ".end array-data")
	//解析文件
	for _, styleableItem := range styleableData {
		index := strings.Index(styleableItem, "R$styleable;->")
		if index == -1 {
			continue
		}
		key := styleableItem[index+len("R$styleable;->"):]

		if strings.Index(styleableItem, ":array_") != -1 {

			s2 := styleableItem[strings.Index(styleableItem, ":array_"):]
			targetStr := s2[:strings.Index(s2, "\r")]
			item := findItem(targetStr, attrsData)
			styleableMap[key] = item
		} else {
			items := strings.Split(styleableItem, "0x")
			items = items[1:]
			for _, item := range items {
				index := strings.Index(item, "\r\n")
				if index < 5 {
					continue
				} else {
					styleableMap[key] = append([]string{}, "0x"+item[:index])
					break
				}
			}

		}
	}
	return styleableMap, nil
}

func findItem(targetStr string, strArray []string) []string {
	var i []string
	for _, str := range strArray {
		if strings.Index(str, targetStr) != -1 {
			resPrefix := "0x"
			str = strings.Replace(str, "\r\n", "", -1)
			str = strings.Replace(str, " ", "", -1)

			split := strings.Split(str, resPrefix)
			split = split[1:]
			for _, content := range split {
				i = append(i, resPrefix+content)
			}
			break
		}
	}
	return i
}

const androidFramework = "C:\\Users\\LZH\\AppData\\Local\\Android\\Sdk\\platforms\\android-32\\"

const frameworkPublicXml = androidFramework + "data\\res\\values\\public.xml"

// GenNewStyleableXml 根据系统与apk的Public.xml文件生成Styleable.xml文件 /**
func GenNewStyleableXml(appPath string) error {

	_, fileNotExist := os.Open(frameworkPublicXml)
	if fileNotExist != nil {
		return fmt.Errorf("frameworkPublicXml未找到,请指定androidFramework地址")
	}

	apkPublicXmlPath := appPath + "\\res\\values\\public.xml"
	attrXML, _ := DecodeResourcesXml(appPath + "\\res\\values\\attrs.xml")
	apkPublicXml, _ := DecodeResourcesXml(apkPublicXmlPath)
	frameworkPublicXml, _ := DecodeResourcesXml(frameworkPublicXml)

	styleMap, err := DecodeStyleableSmali(appPath + "\\smali\\com\\test\\supersdkdemo\\R$styleable.smali")
	if err != nil {
		styleMap, err = DecodeStyleableSmali(appPath + "\\smali_classes2\\com\\test\\supersdkdemo\\R$styleable.smali")
		if err != nil {
			return fmt.Errorf("smali与smali_classes2中未找到 com.test.supersdkdemo.R$styleable.smali文件")
		}
	}
	//设立一个cache，在结束前被添加过的attr都会被添加到这里面来
	attrsCache := Resources{}
	styleableXml := Resources{}
	//取一条styleable
	for key, values := range styleMap {
		//排除掉Font开头
		if strings.Index(key, "Font") != -1 {
			continue
		}
		styleable := Styleable{Name: key}
		//取styleable中的一条attr的ID
		for _, attrId := range values {
			attrs := Attrs{}
			//0x1对应framework中的0x01
			if strings.Index(attrId, "0x1") != -1 {
				//在public中查找attr,framework中统一用android:{attrName}
				for _, public := range frameworkPublicXml.Public {
					if strings.ReplaceAll(attrId, "0x1", "0x01") == public.Id {
						attrs.Name = "android:" + public.Name
						styleable.Attr = append(styleable.Attr, attrs)
						goto topLoop
					}
				}
			} else {
				for _, public := range apkPublicXml.Public {
					if public.Id == attrId {
						//通过public中查找到attr名称获取attrs文件中的具体值
						for index, attr := range attrXML.Attr {
							if attr.Name == public.Name {
								attrXML.Attr = append(attrXML.Attr[:index], attrXML.Attr[index+1:]...)
								attrsCache.Attr = append(attrsCache.Attr, attr)
								styleable.Attr = append(styleable.Attr, attr)
								goto attrLoop
							}
						}

						for _, attr := range attrsCache.Attr {
							if attr.Name == public.Name {
								if attr.Enum != nil {
									attr.Enum = nil
								}
								attr = Attrs{Name: attr.Name}
								styleable.Attr = append(styleable.Attr, attr)
								goto attrLoop
							}
						}
					attrLoop:
					}
				}
			}
		topLoop:
		}

		styleableXml.Styleable = append(styleableXml.Styleable, styleable)
	}

	sort.Slice(styleableXml.Styleable, func(i, j int) bool {
		strs := []string{styleableXml.Styleable[i].Name, styleableXml.Styleable[j].Name}
		sort.Strings(strs)
		return strs[0] == styleableXml.Styleable[i].Name
	})
	writeFile(appPath+"\\res\\values\\styleable.xml", styleableXml)
	writeFile(appPath+"\\res\\values\\attrs.xml", attrXML)
	return nil
}

func writeFile(path string, xmlRes Resources) {
	output, readErr := xml.MarshalIndent(&xmlRes, "", "\t")
	if readErr != nil {
		fmt.Println(readErr)
	}
	_ = ioutil.WriteFile(path, []byte(xml.Header+string(output)), 0666)

}

// DecodeResourcesXml 解析ResourcesXML文件/**
func DecodeResourcesXml(filePath string) (Resources, error) {
	attrsFile, err := os.Open(filePath)
	if err != nil {
		print("Error opening ", err)
	}
	attrsData, _ := ioutil.ReadAll(attrsFile)
	builder := strings.Builder{}
	builder.Write(attrsData)
	//print(builder.String())
	var v Resources
	err = xml.Unmarshal(attrsData, &v)
	if err != nil {
		print("Error opening ", err)
		return Resources{}, fmt.Errorf("please check file then decode")
	}
	return v, nil
}

type Resources struct {
	Attr      []Attrs     `xml:"attr"`
	Styleable []Styleable `xml:"declare-styleable"`
	Public    []Public    `xml:"public"`
}

type Public struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name,attr"`
	Id   string `xml:"id,attr"`
}

type Styleable struct {
	Name string  `xml:"name,attr"`
	Attr []Attrs `xml:"attr"`
}

type Attrs struct {
	Name   string `xml:"name,attr"`
	Format string `xml:"format,attr,omitempty"`
	Enum   []Enum `xml:"enum"`
	Flag   []Flag `xml:"flag"`
}

type Enum struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Flag struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
