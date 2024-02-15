package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJSONFile(filePath string) (interface{}, error) {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("文件 %s 不存在", filePath)
	}

	// 检查文件是否为 JSON 文件
	ext := filepath.Ext(filePath)
	if ext != ".json" {
		return nil, fmt.Errorf("文件 %s 不是 JSON 文件", filePath)
	}

	// 读取文件内容
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取文件内容：%v", err)
	}

	// 根据 JSON 类型解析文件内容
	var data interface{}
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, fmt.Errorf("无法解析 JSON 内容：%v", err)
	}

	return data, nil
}

// 读取Json数组
func ReadJsonArray(filePath string) []map[string]string {
	// 读取文件中的JSON数据
	jsonFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		Err("Cannot opening file ! ", err)
		return nil
	}
	defer jsonFile.Close() // 确保文件句柄在函数执行结束时被关闭

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		Err("Error reading file:", err)
		return nil
	}

	// 解析JSON数据
	var data []map[string]string
	if err := json.Unmarshal(byteValue, &data); err != nil {
		Err("Error unmarshalling JSON:", err)
		return nil
	}
	return data
}

// 写入Json数组
func WriteJsonArray(filePath string, newData map[string]string) {
	data := ReadJsonArray(filePath)

	// 追加新的对象
	data = append(data, newData)

	// println(data[0])

	// 将修改后的JSON数据重新编码为JSON格式
	encodedData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		Err("Error encoding JSON:", err)
		return
	}

	// 将修改后的JSON数据写回到文件中
	if err := ioutil.WriteFile(filePath, encodedData, 0644); err != nil {
		Err("Error writing to file:", err)
		return
	}

}
