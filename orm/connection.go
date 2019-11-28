package orm

import (
	"context"
	"database/sql"
	"log"
)

type Connection struct {
	ctx context.Context
	conn *sql.Conn
	tx *sql.Tx
	name string
}

func newConnection() *Connection {
	c := new(Connection)
	c.ctx = context.Background()
	return c
}

func (c *Connection) WithDB(name string) *Connection {
	c.name = name
	return c
}

func (c *Connection) WithContext(ctx context.Context) *Connection {
	c.ctx = ctx
	return c
}

func (c *Connection) Select(query string, bindings []interface{}) *Rows {
	rows, err := c.queryRows(query, bindings)
	if err != nil {
		return &Rows{rs:nil, lastError:err}
	}

	return &Rows{rs:rows, lastError:err}
}

func (c *Connection) queryRows(query string, bindings []interface{}) (rows sql.Rows, err error) {
	log.Println("query:", query, "| bindings:", bindings)
	if c.tx != nil {
		rows, err = c.tx.QueryContext(c.ctx, query, bindings...)
		return
	}

	var conn *sql.Conn
	conn, err = c.getConn()
	if err != nil {
		return nil, err
	}

	rows, err = conn.QueryContext(c.ctx, query, bindings...)

	return
}
