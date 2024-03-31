package example

import (
	"fmt"
	"strconv"
)

type role struct {
	Name       string
	Profession string
	Attack     int
	Skill      []string
	Money      int
	Goods      []map[string]int
	Wear       map[string]string
}

func ExecRescue() {
	var r = role{
		Money:  0,
		Skill:  []string{"普通攻击"},
		Attack: 10,
	}
	var name, profession string
	fmt.Println("欢迎来到西西村！请输入并创建你的名字：")
	fmt.Scanln(&name)
	r.Name = name

for1:
	for {
		fmt.Println("请选择你的职业：1.剑士，2.法师，3.弓箭手")
		fmt.Scanln(&profession)

		switch profession {
		case "1":
			r.Profession = "剑士"
			r.Skill = append(r.Skill, "基础剑术")
			break for1
		case "2":
			r.Profession = "法师"
			r.Skill = append(r.Skill, "基础法术")
			break for1
		case "3":
			r.Profession = "弓箭手"
			r.Skill = append(r.Skill, "基础箭法")
			break for1
		default:
			fmt.Println("输入有误，请重新选择")
		}

	}

	fmt.Println("人物名称：", r.Name)
	fmt.Println("人物职业：", r.Profession)
	fmt.Println("人物技能：", r.Skill)
	fmt.Println("人物装备：", r.Wear)

	var skill, ways, wear string
	if r.Profession == "剑士" {
		skill = "剑刃冲击"
	} else if r.Profession == "法师" {
		skill = "火球术"
	} else if r.Profession == "弓箭手" {
		skill = "心神凝聚"
	}

	// 学习技能
	fmt.Printf("你好，%v，我是西西村村长，最近野兽森林里出现了山贼，拐走了采药的村民，我这里有一本%v，希望你能营救村民。\n",
		name, skill)
	fmt.Println("系统提示：习得", skill)
	r.Skill = append(r.Skill, skill)

	// 获得金钱和物品
	fmt.Printf("你好，%v，我是药铺老板，我的员工在采药时被山贼桌去，这里有10瓶金创药喝100两银子，希望你能解救我的员工\n", name)
	r.Goods = append(r.Goods, map[string]int{"金疮药": 10})
	r.Money += 100
	fmt.Println("系统提示，获得10瓶金疮药喝100两银子")

	for {
		fmt.Println("选择你的路线：1.兵器铺，2.野兽森林，3.退出")
		fmt.Scanln(&ways)

		if ways == "1" { // 兵器铺路线
			fmt.Println("本店出售兵器，总有一款适合你")
			fmt.Println("木杖-60两 +18攻击：选择1")
			fmt.Println("木剑-60两 +18攻击：选择2")
			fmt.Println("木弓-60两 +18攻击：选择3")
			fmt.Println("布衣-30两 +10防御：选择4")
			fmt.Scanln(&wear)
			r.Wear = map[string]string{}
			var arr = []string{"木杖", "木剑", "木弓", "布衣"}
			var i, _ = strconv.Atoi(wear)

			if wear == "4" {
				// 将衣服的防御力乘以 0.2 转化为攻击力
				r.Attack += 10 * 0.2
				r.Money -= 30
				r.Wear["clothes"] = arr[i-1]
			} else {
				r.Attack += 18
				r.Money -= 60
				r.Wear["arms"] = arr[i-1]
			}

		} else if ways == "3" {
			break
		}
	}

	fmt.Println("人物名称：", r.Name)
	fmt.Println("人物职业：", r.Profession)
	fmt.Println("人物技能：", r.Skill)
	fmt.Println("人物攻击力：", r.Attack)
	fmt.Println("人物金钱：", r.Money)
	fmt.Println("人物装备：", r.Wear)
	fmt.Println("人物物品：", r.Goods)

}
