package util

import (
	"database/sql"
	"github.com/streadway/amqp"
	"text/template"

	"github.com/gorilla/sessions"
)

var DB *sql.DB
var Tpl *template.Template
var Store *sessions.CookieStore
var MessageBroker *amqp.Connection