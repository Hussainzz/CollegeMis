package models

import (
	
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
	//"strconv"
	"reflect"
	"errors"
	"strings"
	"log"
)

type Notice struct {
	Nid int `orm:"pk;auto" json:"Nid"`
	NoticeTitle string `json:"NoticeTitle"`
	NoticeType *NoticeType `orm:"rel(fk)" json:"NoticeType"`
	NoticeContent string `orm:"size(700)" json:"NoticeContent"`
	NoticeCreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"NoticeCreatedAt"`
}



type NoticeType struct{
	Ntid int `orm:"pk;auto" json:"Ntid"`
	NoticeTypeName string `json:"NoticeTypeName"`
}


type Keyword struct{
	
	Key string `json:"Key"`

}


func init(){
	orm.RegisterModel(new(NoticeType), new(Notice))
}

func AddingNotice(nDetails *Notice) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	notice := &Notice{}

	notice.NoticeTitle = nDetails.NoticeTitle
	notice.NoticeType = nDetails.NoticeType
	notice.NoticeContent = nDetails.NoticeContent

	o.Insert(notice)
	success = "New Notice Created Successfully"
	
	var notices []*Notice
	num, err := o.QueryTable("notice").Filter("notice_type_id", 2).RelatedSel().All(&notices)
	if err == nil {
    fmt.Printf("%d posts read\n", num)
    for _, post := range notices {
        fmt.Printf(post.NoticeType.NoticeTypeName)
    }
		}
   	fmt.Println(notice.NoticeType.NoticeTypeName)
	return success, nil
}
	

func AddingNoticeType(nTypeDetails *NoticeType) (string , error){
	o := orm.NewOrm()
	o.Using("default")
	noticeType := &NoticeType{}

	noticeType.NoticeTypeName = nTypeDetails.NoticeTypeName

	o.Insert(noticeType)



	success := "New Notice Type Created Successfully"
	return success, nil
}

//THIS IS FOR NOTICE TABLE
func (notice *Notice) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(notice); err != nil {
		beego.Info(err)
	}
	return
}

//THIS IS FOR NOTICE TABLE
func (notice Notice) GetAllNotice() (ret []Notice) {
	o := orm.NewOrm()
	o.QueryTable("notice").RelatedSel().OrderBy("-nid").All(&ret)
	
	return

}

//-------------------------------------------------------------
//THIS IS FOR NOTICE_TYPE TABLE
func (noticetype *NoticeType) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(noticetype); err != nil {
		beego.Info(err)
	}
	return
}

//THIS IS OF NOTICE_TYPE
func (noticetype NoticeType) GetAllNoticType() (ret []NoticeType) {
	o := orm.NewOrm()
	o.QueryTable("notice_type").RelatedSel().All(&ret)
	
	return
}

func (nty *NoticeType) checkNoticeType(nt int) orm.QuerySeter{
	o := orm.NewOrm()

	return o.QueryTable("notice_type").Filter("ntid__contains", nt)
}

func (n *Notice) cNotice() (ret []Notice){
	o := orm.NewOrm()
	o.QueryTable("notice").RelatedSel("notice_type_id")

	return
}


func Search(key *Keyword) ( []Notice, string, error){
	
	o := orm.NewOrm()
	
	query := o.QueryTable("notice")

	var ret []Notice
	query.Filter("notice_type_id__icontains", key.Key).RelatedSel().All(&ret)
	query.Filter("notice_title__icontains", key.Key).RelatedSel().All(&ret)
	
	for _, n := range ret {
        fmt.Printf(n.NoticeType.NoticeTypeName)
    }

	success := "Queryed Notice"

	return  ret, success, nil
}



func GetAllNotice(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notice))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Notice
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}