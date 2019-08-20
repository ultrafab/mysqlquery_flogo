package mysqlquery_flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"database/sql"
	"fmt"
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
	db, err := sql.Open("mysql", "ukeau:kaplan@/rub")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var dat map[string]interface{}

	if err2 := json.Unmarshal([]byte(mm), &dat); err2 != nil {
        panic(err2)
    }

	var pi = dat["pressioneIngresso"]
	var pp = dat["pressionePerdita"]
	var spi = "NULL"
	var spp = "NULL"
	if pi != nil {
		spi = fmt.Sprintf("%f", pi.(float64))
	}
	if pp != nil {
		spp = fmt.Sprintf("%f", pp.(float64))
	}
	var s = "INSERT INTO data VALUES (UTC_TIMESTAMP(),"+spi+","+spp+")"
	var r = true
	 stmtIns, err := db.Query(s) // ? = placeholder
	 if err != nil {
		 r = false
		 panic(err.Error()) // proper error handling instead of panic in your app
	 }
	 defer stmtIns.Close()

	context.SetOutput("result", r)
	return true, nil
}
