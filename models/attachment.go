package models

import (
	"beegostudy/util/modelutil"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
*   pk      主键
*   auto        自增值（限数值）
*   column(N)   指定字段名N
*   null        非空
*   index       单个字段索引
*   unique      唯一键
*   auto_now_add    第一次插入数据时自动添加当前时间
*   auto_now    每一次保存时自动更新当前时间
*   type(T)     对应数据库的指定类型
*   size(S)     类型长度S
*   default(D)  默认值D（需要对应类型）
**/
type Attachment struct {
	Id          int       `orm:"pk;auto;column(id)"`
	FileName    string    `orm:"column(filaname);size(128)"`
	FileNewName string    `orm:"column(filenewfile);size(128)"`
	FilePath    string    `orm:"column(filepath);size(256)"`
	FileType    string    `orm:"column(filetype);size(64)"`
	FileSize    int64     `orm:"column(filesize)"`
	AddTime     time.Time `orm:"auto_now_add;type(datetime);column(addtime)"`
	AddUser     string    `orm:"column(adduser);size(64)"`
}

//自定义表名
func (m *Attachment) TableName() string {
	return "attachment"
}

func (m *Attachment) SetId(id interface{}) error {
	tmpId := fmt.Sprintf("%v", id)
	mid, err := strconv.Atoi(tmpId)
	if err == nil {
		m.Id = mid
	} else {
		beego.Error("Id字段必须为正整数型【%v】\n", id)
	}
	return err
}

func (m *Attachment) SetAddTime(t time.Time) {
	m.AddTime = t
}

func (m *Attachment) SetCurrentTime() {
	m.AddTime = time.Now()
}

func (m *Attachment) SetAddUser(uname string) {
	m.AddUser = uname
}

func (m *Attachment) SetValue(data map[string]interface{}) error {
	return modelutil.FillStruct(data, m)
}

//查询数据库
func (m *Attachment) Fill() error {
	o := orm.NewOrm()
	if m.Id > 0 {
		return o.Read(m, "Id")
	}

	return fmt.Errorf("请确认是否传递了Id", "")

}

//插入
func (m *Attachment) Insert() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

//修改
func (m *Attachment) Update(column ...string) (int64, error) {
	o := orm.NewOrm()
	return o.Update(m, column...)
}

func init() {
	orm.RegisterModel(new(Attachment))
}

func (m *Attachment) String() string {
	data, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		beego.Warn("JSON marshaling failed: %s", err)
	}
	return fmt.Sprintf("%s\n", data)
}

//相关函数
//新创建一个附件
func AddAttachment(filename, filenewname, filepath, filetype string, filesize int64, adduser string) *Attachment {
	m := new(Attachment)
	m.FileName = filename
	m.FileNewName = filenewname
	m.FilePath = filepath
	m.FileSize = filesize
	m.FileType = filetype
	m.SetCurrentTime()
	m.AddUser = adduser
	m.Insert()
	return m
}

//得到分页的菜单
/**
*   size    每页查询长度
*   index   查询的页码
*   ordercolumn 排序字段
*   orderby     升降序:desc\asc
**/
func GetAttchmentsPage(size, index int, ordercolumn, orderby string, data map[string]interface{}) (*DataGrid, error) {

	if ordercolumn == "" {
		ordercolumn = "addtime"
	} else if strings.EqualFold(orderby, "desc") {
		ordercolumn = "-" + ordercolumn
	}

	var atta []*Attachment
	o := orm.NewOrm()
	qt := o.QueryTable("attachment")
	if data["FileName"] != nil {
		qt = qt.Filter("FileName__icontains", data["FileName"])
	}
	_, err := qt.OrderBy(ordercolumn).Limit(size, (index-1)*size).All(&atta)

	if err == nil {
		cnt, err := qt.Count()

		pagetotal := cnt / int64(size)

		if cnt%int64(size) > 0 {
			pagetotal++
		}

		return GetDataGrid(atta, index, int(pagetotal), cnt), err
	}

	return nil, err
}
