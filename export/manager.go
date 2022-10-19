package export

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"plugin-demo-go/db"
	"plugin-demo-go/db/gen/item"
	"plugin-demo-go/plugins/monitor"
	"plugin-demo-go/tools"
	"strings"
	"time"
)

var target_path = "/home/amelia/Desktop/py-proj/plugin-demo-py/impl"

func (c *Client) Install(source_path string) error {
	// 把 zip 解压到指定目录,同时返回文件名
	file_name := tools.Unzip(source_path, target_path)
	// 依照 文件名 将数据存储到数据库，同时返回服务名
	srv_name := c.doStorage(file_name)
	// 依照服务名向插件服务器更新服务(批量)
	c.doNotify("install", srv_name)
	return nil
}

func (c *Client) Update(plugin string) {
	// plugin 是更新的插件名
	// conf 每次更新都会刷新配置文件
	// 如果仅更新配置文件就将 plugin 设置为默认值 ""
	list := []string{"conf"}
	if plugin != "conf" {
		list = append(list, plugin)
	}
	c.doNotify("update", list)
}

func (c *Client) doStorage(dir []string) (service_name []string) {
	dbs := db.Connection()
	for _, v := range dir {
		// 打开文件
		f, err := os.Open(filepath.Join(target_path, v))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		defer f.Close()
		// 读出数据
		b, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		// 存入数据库
		s := strings.Split(v, ".")[0]
		effect := dbs.Item.Update().
			Where(item.NameEQ(s)).
			SetFile(b).
			SaveX(context.Background())
		if effect == 0 {
			// 如果不存在就创建，一般都是默认存在的
			dbs.Item.Create().
				SetName(s).
				SetFile(b).
				SaveX(context.Background())
		}
		service_name = append(service_name, strings.Split(s, "_")[0])
	}
	return service_name
}

func (c *Client) doNotify(command string, srv []string) {
	m := c.GetSrv("mon").(monitor.MT)
	m.Commands(command, srv)
}

func (c *Client) UnInstall(src_name string) {
	// 应该只能修改配置文件
	// 我们不允许将插件服务器中的插件代码直接删掉，因为可能会导致问题
	// 我们只能通过配置去禁掉该功能
	m := c.GetSrv("mon").(monitor.MT)
	m.Command("uninstall", src_name)
}
// 启用/禁用
func (c *Client) Banned(src_name string) {
	m := c.GetSrv("mon").(monitor.MT)
	m.Command("banned", src_name)
}

// func (c *Client) UpdateAll() {
// 	// 所有 plugin 包括 配置文件 全部更新
// }

// new 插件客户端后
// 自动启动该自动监控功能，监控远端数据库的信息变化
func (c *Client) AutoManage() {
	// 准备工作
	ctx := context.Background()
	dbs := db.Connection()
	ticker := time.NewTicker(5 * time.Second)
	var old_info []pluginInfo
	var new_info []pluginInfo
	dbs.Item.Query().
		Select(item.FieldID,item.FieldName, item.FieldUpdatedAt).
		Scan(ctx, &old_info)
	// 开始巡回监控
	for {
		select {
		case <-ticker.C:
			dbs.Item.Query().
				Select(item.FieldID,item.FieldName, item.FieldUpdatedAt).
				Scan(ctx, &new_info)
			// fmt.Printf("old_info: %v\n", old_info)
			// fmt.Printf("new_info: %v\n", new_info)
			new_len := len(new_info)
			old_len := len(old_info)
			
			for i:=0;i<old_len;i++{
				// 如果更新时间不同，则拉取下来，做更新操作
				if new_info[i].Update != old_info[i].Update{
					item := dbs.Item.Query().
						Where(item.NameEQ(new_info[i].Name)).
						OnlyX(ctx)
					if item.Name != "conf" {
						// 如果不是配置文件
						f, err := os.OpenFile(filepath.Join(target_path, item.Name+".py"),os.O_WRONLY,0755)
						if err != nil {
							fmt.Printf("err: %v\n", err)
						}
						defer f.Close()
						// 覆盖写
						f.Write(item.File)
						// 更新
						c.Update(strings.Split(item.Name, "_")[0])
					}else {
						// 如果是配置文件
						f, err := os.OpenFile(filepath.Join(target_path, item.Name+".json"),os.O_WRONLY,0755)
						if err != nil {
							fmt.Printf("err: %v\n", err)
						}
						defer f.Close()
						// 覆盖写
						f.Write(item.File)
						// 更新
						c.Update(item.Name)
					}
				}
			}

			// 如果确实存在新增的插件话
			for i:=old_len;i<new_len;i++{
				
			}

			old_info = new_info
			new_info = nil
		case <-c.done:
			println("监视系统退出")
			return
		}
	}
}

type pluginInfo struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Update time.Time `json:"updated_at"`
}
