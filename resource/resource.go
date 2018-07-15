package resource

import (
	"strings"
	"strconv"
)

type resource struct {
	url    string
	target string
	start  int
	end    int
}

/**
url 生成规则
 */
func ReuleResource() (result []resource ){
	result = []resource{}

	//首页
	result = append(result, resource{
		url:    "http://blog.test/",
		target: "",
		start:  0,
		end:    0,
	})

	//分类
	result = append(result,resource{
		url:    "http://blog.test/categories/{$id}",
		target: "{$id}",
		start:  1,
		end:    4,
	})

	//话题
	result = append(result,resource{
		url:    "http://blog.test/topics/{$id}",
		target: "{$id}",
		start:  1,
		end:    87,
	})

	//标签
	result = append(result,resource{
		url:    "http://blog.test/tags/{$id}",
		target: "{$id}",
		start:  1,
		end:    41,
	})

	return result
}


func BuildUlr(result [] resource)(list []string){
	for _,item := range result{
		if item.target == ""{
			list = append(list,item.url)
		}else{
			for i:= item.start; i<=item.end;i++{
				urlStr := strings.Replace(item.url,item.target,strconv.Itoa(i),-1)
				list = append(list,urlStr)
			}
		}
	}

	return
}