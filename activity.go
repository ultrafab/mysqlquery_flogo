package mysqlquery_flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"database/sql"
	"fmt"
	"strings"
	"encoding/json"
)
import _ "github.com/go-sql-driver/mysql"
var log = logger.GetLogger("mysqlquery_log")
// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	mm := context.GetInput("message").(string)
	connection_url := context.GetInput("username").(string)+":"+context.GetInput("password").(string)+"@/"+context.GetInput("db").(string)
	db, err := sql.Open("mysql", connection_url)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var dat map[string]interface{}

	if err2 := json.Unmarshal([]byte(mm), &dat); err2 != nil {
        panic(err2)
    }

	var query = context.GetInput("query").(string)
	i := strings.Index(query, "$MESSAGE_KEYS")
	if i>-1 {
		a := query[:i]
		j := strings.Index(query, "$MESSAGE_VALUES")
		b := query[i+13:j]
		c := query[j+15:]
		for k, v := range dat {
			ss := fmt.Sprintf("%f", v)
				a = a +k+","
				b = b +ss+","
		}
		a = a[:len(a)-1]
		b = b[:len(b)-1]
		query = a+b+c
	}

	fmt.Println(query)
	var r = true
	 stmtIns, err := db.Query(query) // ? = placeholder
	 if err != nil {
		 r = false
		 panic(err.Error()) // proper error handling instead of panic in your app
	 }
	 defer stmtIns.Close()

	context.SetOutput("result", r)
	return true, nil
}
