/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/05/16        Tong Yuehong
 */

package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tongyuehong1/golang-training/log"
)


type Student struct {
	//Id        uint32 `orm:"column(id)"        json:"id"`
	Name      string `orm:"column(name)"      json:"name"`
	Class     string `orm:"column(class)"     json:"class"`
	StudentId string `orm:"column(studentid)" json:"studentid"`
	Avatar    string `orm:"column(avatar)"    json:"avatar"`
	Sex       string `orm:"column(sex)"       json:"sex"`
	Age       string `orm:"column(age)"       json:"age"`
	Phone     string `orm:"column(phone)"     json:"phone"`
	Address   string `orm:"column(address)"   json:"address"`
	Duty      string `orm:"column(duty)"      json:"duty"`
	Isonly    string `orm:"column(isonly)"    json:"isonly"`
	Status    int8   `orm:"column(status)"    json:"status"`
}

func main() {
	var stu Student
	xlsx, err := excelize.OpenFile("/Users/a5/Documents/Book.xlsx")
	if err != nil {
		//fmt.Println(err)
		for i := 0; i< 5; i ++{
			log.Error(fmt.Sprintf("example: %v", err))
		}
		return
	}

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			stu.Name = row[0]
			stu.Class = row[1]
			stu.StudentId = row[2]
			stu.Sex = row[3]
			stu.Age = row[4]
			stu.Phone = row[5]
			stu.Duty = row[6]
			stu.Isonly = row[7]
			stu.Address = row[8]
			fmt.Println(stu)
		}
		fmt.Println(stu)
	}
}