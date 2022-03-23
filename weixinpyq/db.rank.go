package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"gorm.io/gorm"
	"learn.go/pkg/apis"
	"log"
)

var _ ServerInterface = &dbRank{}
var _ RankInitInterface = &dbRank{}

func NewDbRank(conn *gorm.DB, embedRank ServerInterface) ServerInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{
		conn:      conn,
		embedRank: embedRank,
	}
}

type dbRank struct {
	conn      *gorm.DB
	embedRank ServerInterface
}

func (d dbRank) Init() error {
	output := make([]*apis.PersonInformation, 0)
	resp := d.conn.Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return resp.Error
	}
	for _, item := range output {
		if _, err := d.embedRank.UpdatePersonInformation(item); err != nil {
			log.Fatalf("初始化%s时失败:%v", item.PersonalName, err)
		}
	}
	return nil
}

func (d dbRank) RegisterPersonInformation(pi *apis.PersonInformation) error {
	bmi, err := gobmi.BMI(float64(pi.ByTimeWeight), float64(pi.ByTimeHeight))
	if err != nil {
		fmt.Println("bmi计算错误:", err)
		return err
	}
	ft := gobmi.CalcFatRate(int(pi.ByTimeAge), bmi, pi.PersonalSex)
	pi.ByTimeFatrate = float32(ft)
	resp := d.conn.Create(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("创建%s时失败%v:", pi.PersonalName, err)
		return err
	}

	fmt.Printf("创建%s成功\n", pi.PersonalName)
	_ = d.embedRank.RegisterPersonInformation(pi) //handle error
	return nil
}

func (d dbRank) UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败%v:", pi.PersonalName, err)
		return nil, err
	}
	fmt.Printf("更新%s成功\n", pi.PersonalName)
	return d.embedRank.UpdatePersonInformation(pi)

}

func (d dbRank) GetFateRate(PersonalName string) (*apis.PersonRank, error) {
	return d.embedRank.GetFateRate(PersonalName)
}

func (d dbRank) GetTop() ([]*apis.PersonRank, error) {
	return d.embedRank.GetTop()
}

//删除数据时，标记要删除的用户数据Visiable为1   表示数据不可见
func (d dbRank) DeletePersonPyq(id int) error {
	sqlStr := "update pyq set visiable =? where id = ?"
	ret := d.conn.Exec(sqlStr, 1, id)
	err := ret.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	n := ret.RowsAffected // 操作影响的行数

	fmt.Printf("update success, affected rows:%d\n", n)
	return nil
}

func (d dbRank) ShowPersonPyq() ([]*apis.PersonInformation, error) {
	//true = 1   代表该pyq状态不可见
	resp := d.conn.Where("visiable = 0  ").Find(&apis.PersonInformation{})
	//resp := d.conn.Delete(id)
	if err := resp.Error; err != nil {
		fmt.Printf("显示所有pyq状态失败:%v", err)
		return nil, err
	}
	output := make([]*apis.PersonInformation, 0)
	resp1 := resp.Find(&output)
	if resp1.Error != nil {
		fmt.Println("查找失败")
		return nil, resp.Error
	}
	fmt.Println("显示所有pyq状态成功")
	return output, nil
}
