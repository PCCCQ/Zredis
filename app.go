package main

import (
	"changeme/control/data_structure/data_hash"
	"changeme/control/data_structure/data_list"
	"changeme/control/data_structure/data_public"
	"changeme/control/data_structure/data_set"
	"changeme/control/data_structure/data_string"
	"changeme/control/data_structure/data_zset"
	"changeme/control/redis_init"
	"changeme/model"
	"context"
	"fmt"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test(t any) any {
	//type People struct {
	//	Name   string `json:"name"`
	//	Age    int    `json:"age"`
	//	Adress string `json:"address"`
	//}
	//People1 := People{"张志轩", 22, "河北省邯郸市峰峰矿区"}
	//marshal, err := json.Marshal(People1)
	return t
}

// RedisInit 初始化redis，即连接redis
func (a *App) RedisInit(r redis_init.RedisClient) model.SimpleResponse {
	fmt.Println("执行RedisInit", r)
	rdb, err := r.RedisInit()
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: "连接失败，请重试"}
	}
	RedisDB = rdb
	return model.SimpleResponse{Code: 200, Data: "连接成功"}
}

// GetDBList 获取全部数据库
func (a *App) GetDBList() model.SimpleResponse {
	r, err := data_public.GetDBList(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: r}
}

func (a *App) ChangeDB(DB int) model.SimpleResponse {
	err := data_public.ChangeDB(a.ctx, RedisDB, DB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: "切换成功"}
}

// AllKeyTypes 获取全部数据，key和type
func (a *App) AllKeyTypes() model.SimpleResponse {

	allKeyTypes, err := redis_init.AllKeyTypes(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: allKeyTypes}
}

// Del 删除数据key
func (a *App) Del(key string) model.SimpleResponse {
	if data_public.Del(a.ctx, RedisDB, key) {
		return model.SimpleResponse{Code: 200, Data: "删除成功"}
	}
	return model.SimpleResponse{Code: 500, Data: "删除失败"}

}

// SetKey 修改key
func (a *App) SetKey(oldKey string, newKey string) model.SimpleResponse {

	key, err := data_public.SetKey(a.ctx, RedisDB, oldKey, newKey)
	if err != nil {
		return model.SimpleResponse{500, err}
	}

	return model.SimpleResponse{Code: 200, Data: key}
}

// GetTTL 获取超时时间
func (a *App) GetTTL(key string) model.SimpleResponse {

	TTL := data_public.GetTTL(a.ctx, RedisDB, key)

	return model.SimpleResponse{Code: 200, Data: TTL}
}

// SetTTL 修改超时时间超时时间
func (a *App) SetTTL(key string, TTL string) model.SimpleResponse {
	tmp, _ := strconv.ParseFloat(TTL, 10)
	if data_public.SetTTL(a.ctx, RedisDB, key, float32(tmp)) {
		return model.SimpleResponse{Code: 200, Data: "修改成功"}
	}
	return model.SimpleResponse{Code: 500, Data: "修改失败"}

}

/*以下为set数据结构的操作函数*/

// SetSet 给set类型设置值
func (a *App) SetSet(s data_set.SetData) model.SimpleResponse {
	fmt.Println("执行1")
	err := s.SetSet(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: "成功"}
}

// GetSet 获取set类型值
func (a *App) GetSet(s data_set.SetData) model.SimpleResponse {
	data, err := s.GetSet(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{500, data}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

// DelSetOne 删除set某一个值
func (a *App) DelSetOne(s data_set.SetData) model.SimpleResponse {
	data, err := s.DelSetOne(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{500, data}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

/*以下为String数据结构的操作函数*/

func (a *App) SetString(s data_string.StringData) model.SimpleResponse {
	data, err := s.SetString(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

func (a *App) GetString(s data_string.StringData) model.SimpleResponse {
	data, err := s.GetString(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

func (a *App) DelString(s data_string.StringData) model.SimpleResponse {
	data, err := s.DelString(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

/*以下为list数据结构的操作函数*/

// SetList 给set类型设置值
func (a *App) SetList(l data_list.ListData) model.SimpleResponse {
	if l.AddTab == "" {
		l.AddTab = "after"
	}
	err := l.SetList(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: "成功"}
}

func (a *App) SetListSet(l data_list.ListData) model.SimpleResponse {
	err := l.SetListSet(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: "成功"}
}

// GetList 获取set类型值
func (a *App) GetList(l data_list.ListData) model.SimpleResponse {
	data, err := l.GetList(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{500, err}
	}
	return model.SimpleResponse{Code: 200, Data: data}
}

// DelListOne 删除set某一个值
func (a *App) DelListOne(l data_list.ListData) model.SimpleResponse {
	err := l.DelListOne(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{500, err}
	}
	return model.SimpleResponse{Code: 200, Data: err}
}

func (a *App) GetZset(z data_zset.ZsetData) model.SimpleResponse {
	err := z.GetZset(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: z}
	}
	return model.SimpleResponse{Code: 200, Data: z}
}
func (a *App) AddZset(z data_zset.ZsetData) model.SimpleResponse {
	err := z.AddZset(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: z}
	}
	return model.SimpleResponse{Code: 200, Data: z}
}
func (a *App) DelZsetOne(z data_zset.ZsetData) model.SimpleResponse {
	err := z.DelZsetOne(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: z}
	}
	return model.SimpleResponse{Code: 200, Data: z}
}

func (a *App) SetZsetScore(z data_zset.ZsetData) model.SimpleResponse {
	// Set the score of the zset in Redis
	err := z.SetZsetScore(a.ctx, RedisDB)

	// If there is an error, return an error response
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: z}
	}
	// Otherwise, return a successful response
	return model.SimpleResponse{Code: 200, Data: z}
}

// GetHash  获取hash值
func (a *App) GetHash(h data_hash.HashData) model.SimpleResponse {
	err := h.GetHash(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: h}
}

func (a *App) DelHashOne(h data_hash.HashData) model.SimpleResponse {
	err := h.DelHashOne(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: h}
}

func (a *App) SetHash(h data_hash.HashData) model.SimpleResponse {
	err := h.SetHash(a.ctx, RedisDB)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: h}
}

func (a *App) AnyCmd(cmd string) model.SimpleResponse {
	r, err := data_public.AnyCmd(a.ctx, RedisDB, cmd)
	if err != nil {
		return model.SimpleResponse{Code: 500, Data: err}
	}
	return model.SimpleResponse{Code: 200, Data: r}
}
