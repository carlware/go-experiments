package main

import (
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func main() {
	sql, _, _ := goqu.From("test").Where(goqu.Ex{
		"d": []string{"a", "b", "c"},
	}).ToSQL()
	fmt.Println(sql)

	// filter := []goqu.Expression{
	// 	goqu.C("").Gt("2020-05-27 12:47:50.183925-05"),
	// 	goqu.C("created").Lt("2020-05-27 16:47:49.107218-05"),
	// }

	now := time.Now()
	filter := []exp.Expression{
		goqu.C("created").Gt(now),
		goqu.C("created").Lt(now),
	}
	ordering := []exp.OrderedExpression{
		goqu.I("created").Desc(),
		goqu.I("name").Desc(),
	}

	sql, p, e := goqu.From("providers").
		Limit(2).
		// Select(
		// 	goqu.COUNT("*").As("total"),
		// ).
		Where(
			filter...,
		).
		Order(
			ordering...,
		).
		// Prepared(true).
		ToSQL()
	fmt.Println(sql)
	fmt.Println(p)
	fmt.Println(e)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(now)

}
