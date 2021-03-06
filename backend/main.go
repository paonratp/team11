package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team11/app/controllers"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/role"
	"github.com/team11/app/ent/status"
)

//User struct
type User struct {
	Name     string
	Email    string
	Password string
	Role     int
}

//Location struct
type Location struct {
	Name string
}

//ServicePoint struct
type ServicePoint struct {
	CounterNumber string
}

//ClientEntity struct
type ClientEntity struct {
	Name  string
	State int
}

//Book struct
type Book struct {
	BookName string
	Category int
	Author   int
	User     int
	Status   int
}

// Author struct
type Author struct {
	NameAuthor string
}

// Researchtype struct
type Researchtype struct {
	NameResearchtype string
}

//Bookborrow struct
type Bookborrow struct {
	UserID         int
	BookID         int
	ServicePointID int
	Status         int
}

// Roominfos for struct
type Roominfos struct {
	Roominfo []Roominfo
}

// Roominfo for struct
type Roominfo struct {
	RoomID     string
	RoomNo     string
	RoomType   string
	RoomTime   string
	RoomStatus string
}

// Purposes for struct
type Purposes struct {
	Purpose []Purpose
}

// Purpose for struct
type Purpose struct {
	PurposeName string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewAuthorController(v1, client)
	controllers.NewBookController(v1, client)
	controllers.NewBookborrowController(v1, client)
	controllers.NewBookingController(v1, client)
	controllers.NewBookreturnController(v1, client)

	controllers.NewCategoryController(v1, client)
	controllers.NewClientEntityController(v1, client)
	controllers.NewLocationController(v1, client)
	controllers.NewPreemptionController(v1, client)
	controllers.NewPurposeController(v1, client)

	controllers.NewResearchController(v1, client)
	controllers.NewResearchTypeController(v1, client)
	controllers.NewRoleController(v1, client)
	controllers.NewRoominfoController(v1, client)
	controllers.NewServicePointController(v1, client)

	controllers.NewStatusController(v1, client)
	controllers.NewUserController(v1, client)

	Role := []string{"Library Member", "Librarian"}
	for _, r := range Role {
		client.Role.
			Create().
			SetROLENAME(r).
			Save(context.Background())
	}
	// Set Users Data
	Users := []User{
		{"Suphachai Phetthamrong", "B6111427@gmail.com", "B6111427", 1},
		{"Name Surname", "me@example.com", "pass", 1},
		{"root", "root@gmail.com", "root", 2},
	}
	for _, u := range Users {

		r, err := client.Role.
			Query().
			Where(role.IDEQ(int(u.Role))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.User.
			Create().
			SetUSEREMAIL(u.Email).
			SetUSERNAME(u.Name).
			SetPASSWORD(u.Password).
			SetPosition(r).
			Save(context.Background())
	}

	Status := []string{"Available", "In Use", "No Service", "Borrowed", "Returned"}
	for _, st := range Status {
		client.Status.
			Create().
			SetSTATUSNAME(st).
			Save(context.Background())
	}

	ClientEntity := []ClientEntity{
		{"Video On Demand - 01", 1},
		{"Video On Demand - 02", 1},
		{"Video On Demand - 02", 1},
		{"Video On Demand - 03", 1},
		{"Video On Demand - 04", 1},
		{"Video On Demand - 05", 1},
		{"Video On Demand - 06", 1},
		{"Video On Demand - 07", 1},
		{"Video On Demand - 08", 1},
		{"Video On Demand - 09", 1}}
	for _, cl := range ClientEntity {

		s, err := client.Status.
			Query().
			Where(status.IDEQ(int(cl.State))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.ClientEntity.
			Create().
			SetCLIENTNAME(cl.Name).
			SetState(s).
			Save(context.Background())
	}

	ServicePoint := []ServicePoint{
		{"เค้าเตอร์ 1"},
		{"เค้าเตอร์ 2"}}
	for _, se := range ServicePoint {
		client.ServicePoint.
			Create().
			SetCOUNTERNUMBER(se.CounterNumber).
			Save(context.Background())
	}

	Author := []string{"โยชิฮิโระ โทงาชิ", "เออิจิโร โอดะ", "เจ.อาร์.อาร์ โทลคีน"}
	for _, a := range Author {
		client.Author.
			Create().
			SetName(a).
			Save(context.Background())
	}

	Researchtype := []string{"การวิจัยเชิงประวัติศาสตร์", "การวิจัยเชิงบรรยาย", "การวิจัยเชิงทดลอง", "การวิจัยเชิงปริมาณ", "การวิจัยพื้นฐาน", "การวิจัยประยุกต์", ""}
	for _, rt := range Researchtype {
		client.Researchtype.
			Create().
			SetTYPENAME(rt).
			Save(context.Background())
	}

	Category := []string{"นิยาย", "เรื่องสั้น", "การ์ตูน", "ศิลปกรรม", "ความรู้ทั้วไป", "ปรัชญา จิตวิทยา ศาสนา", "ประวัติศาสตร์", "ภูมิศาสตร์", "เทคโนโลยี", "การศึกษา"}
	for _, ca := range Category {
		client.Category.
			Create().
			SetCategoryName(ca).
			Save(context.Background())
	}

	Location := []string{"Building A", "Building B", "Building C"}
	for _, l := range Location {
		client.Location.
			Create().
			SetLOCATIONNAME(l).
			Save(context.Background())
	}

	// Set Purposes Data
	purposes := Purposes{
		Purpose: []Purpose{
			Purpose{"อ่านทบทวน"},
			Purpose{"ติวหนังสือ"},
			Purpose{"ทำวิจัย"},
			Purpose{"เพื่อความบันเทิง"},
		},
	}

	for _, p := range purposes.Purpose {
		client.Purpose.
			Create().
			SetPurposeName(p.PurposeName).
			Save(context.Background())
	}

	// Set Roominfos Data
	roominfos := Roominfos{
		Roominfo: []Roominfo{
			Roominfo{"RS108", "RS01", "ห้องเดี่ยว", "08.00 AM - 08.50 AM", "ว่าง"},
			Roominfo{"RS109", "RS01", "ห้องเดี่ยว", "09.00 AM - 09.50 AM", "ว่าง"},
			Roominfo{"RS110", "RS01", "ห้องเดี่ยว", "10.00 AM - 10.50 AM", "ว่าง"},
			Roominfo{"RS208", "RS02", "ห้องเดี่ยว", "08.00 AM - 08.50 AM", "ว่าง"},
			Roominfo{"RS209", "RS02", "ห้องเดี่ยว", "09.00 AM - 09.50 AM", "ว่าง"},
			Roominfo{"RS210", "RS02", "ห้องเดี่ยว", "10.00 AM - 10.50 AM", "ว่าง"},
			Roominfo{"RG208", "RS02", "ห้องกลุ่ม", "08.00 AM - 08.50 AM", "ว่าง"},
			Roominfo{"RG209", "RS02", "ห้องกลุ่ม", "09.00 AM - 09.50 AM", "ว่าง"},
			Roominfo{"RG210", "RS02", "ห้องกลุ่ม", "10.00 AM - 10.50 AM", "ว่าง"},
		},
	}

	for _, r := range roominfos.Roominfo {
		client.Roominfo.
			Create().
			SetRoomID(r.RoomID).
			SetRoomNo(r.RoomNo).
			SetRoomType(r.RoomType).
			SetRoomTime(r.RoomTime).
			SetRoomStatus(r.RoomStatus).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
