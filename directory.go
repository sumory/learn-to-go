package main

import (
	//"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Node struct {
	dir    int
	path   string
	name   string
	deepth int
}

type Nodes []Node

func (ms Nodes) Len() int {
	return len(ms)
}

func (ms Nodes) Less(i, j int) bool {
	if ms[i].deepth < ms[j].deepth { // 先按深度排序，由低到高
		return true
	} else if ms[i].deepth == ms[j].deepth { //深度相同，按类型排序，文件夹在前，文件在后
		if ms[i].dir > ms[j].dir {
			return true
		} else if ms[i].dir == ms[j].dir { // 类型相同，则再按名称排序
			return ms[i].path < ms[j].path
		} else {
			return false
		}
	} else {
		return false
	}
}

func (ms Nodes) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

/**
 * 根据不同系统获取路径分隔符
 */
func getSeparator() string {
	// OS := os.Getenv("GOOS")
	// if OS == "windows" {
	// 	return "\\"
	// } else {
	// 	return "/"
	// }

	return string(os.PathSeparator)
}

/**
 * 获取相对路径的深度
 */
func getDeepth(path string, prefix string) int {

	if strings.HasPrefix(path, prefix) {
		relative_path := strings.TrimPrefix(path, prefix)
		tmp_path_array := strings.Split(relative_path, getSeparator())
		//fmt.Println(path, tmp_path_array, len(tmp_path_array)-1)
		return len(tmp_path_array) - 1
	} else {
		return -1
	}
	return 0
}

/**
 * 遍历目录
 */
func walk(path string, limit_deepth int) Nodes {
	var nodes Nodes
	root := path

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		deepth := getDeepth(path, root)

		if deepth > limit_deepth {
			return nil
		} else {

			relative_path := strings.TrimPrefix(path, root)
			if f.IsDir() {
				nodes = append(nodes, Node{1, relative_path, f.Name(), deepth})
			} else {
				nodes = append(nodes, Node{0, relative_path, f.Name(), deepth})
			}

			return nil
		}
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return nil
	} else {
		// fmt.Println("总数量：", len(nodes))

		// for index, value := range nodes {
		// 	fmt.Println(index, value)
		// }

		return nodes
	}
}

func main() {
	//flag.Parse()
	//root := flag.Arg(0)
	//fmt.Println(root)
	directory := "/data/go/src/"
	nodes := walk(directory, 2)

	sort.Sort(nodes)

	for _, value := range nodes {
		fmt.Println(value)
	}

}
