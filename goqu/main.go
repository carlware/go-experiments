package main

import (
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
)

func main() {
	sql, _, _ := goqu.From("test").Where(goqu.Ex{
		"d": []string{"a", "b", "c"},
	}).ToSQL()
	fmt.Println(sql)

	now := time.Now()
	sql, _, _ = goqu.From("providers").
		Limit(2).
		// Select(
		// 	goqu.COUNT("*").As("total"),
		// ).
		Where(
			goqu.C("created").Gt("2020-05-27 12:47:50.183925-05"),
			goqu.C("created").Lt("2020-05-27 16:47:49.107218-05"),
		).
		Order(goqu.I("created").Desc()).
		Prepared(true).
		ToSQL()
	fmt.Println(sql)
	// fmt.Println(a)
	// fmt.Println(b)
	fmt.Println(now)
}
