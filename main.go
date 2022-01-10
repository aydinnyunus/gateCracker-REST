package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveIndex(s []smartLock, index int) []smartLock {
	return append(s[:index], s[index+1:]...)
}

// smartLock represents data about a record smartLock.
type smartLock struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Door_Pass string `json:"door_pass"`
	Prog_Pass string `json:"prog_pass"`
	Desc      string `json:"desc"`
}

// smartLocks slice to seed record smartLock data.
var smartLocks = []smartLock{
    {ID: "1", Name: "Perkotek", Prog_Pass: "666888", Door_Pass: "Admin Password : *#1234 8'e bas 0 'a bas", Desc: "- Turn off Alarms : *#1234 Press 0 Press 1 Press 0 \\n- Deleting All User Passwords 4444 \\n- Setting the Door Open Time 5XX (XX number shows the time in seconds between 00-99.)"},
    {ID: "2", Name: "ERD-1120", Prog_Pass: "#4567#", Door_Pass: "Undefined", Desc: "- *#SYSTEM_PASS 8 NEW_PASS to change system password"},

    {ID: "3", Name: "Efes Digital Panel", Prog_Pass: "0000", Door_Pass: "1234", Desc: "- CALL -> 1234 -> CALL -> NEW_PASS to change system password"},

    {ID: "4", Name: "Mas", Prog_Pass: "12345", Door_Pass: "1234", Desc: "- Press and hold the 0 button for 10 seconds. Three LEDs will flash. Enter the administrator password. (The administrator password is factory default: 12345.) If the administrator password is correct, the three leds will flash once. Press button 1. Three LEDs will flash. Enter the new 4-digit door password you have determined. \\n Change Administrator Password: Press and hold the 0 key for 10 seconds. Three LEDs will flash. Enter the administrator password. (The administrator password is factory default: 12345.) If the administrator password is correct, the three leds will flash once. Press button 2. Three LEDs will flash. Enter the new 5-digit administrator password you have specified."},

    {ID: "5", Name: "AC 13PX", Prog_Pass: "*1234", Door_Pass: "Undefined", Desc: "- Add User \\n *1234 -> 3"},

    {ID: "6", Name: "Burg Wachter", Prog_Pass: "123456", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "7", Name: "DIP40", Prog_Pass: "0000", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "8", Name: "Lorex LR-DPH2", Prog_Pass: "999999", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "9", Name: "M100", Prog_Pass: "*9999#", Door_Pass: "#8888#", Desc: "- Changing Door Password \\n 92# New Password"},
    {ID: "10", Name: "MB05-03", Prog_Pass: "7447", Door_Pass: "Undefined", Desc: "- Change Password \\n CALL -> 99 -> 7447"},
    {ID: "11", Name: "MB DYF40", Prog_Pass: "7447", Door_Pass: "Undefined", Desc: "- Change Password \\n CALL -> 99 -> 7447"},
    {ID: "12", Name: "MLŞ 14-70", Prog_Pass: "Undefined", Door_Pass: "0000", Desc: "Undefined"},
    {ID: "13", Name: "MLŞ 14-107", Prog_Pass: "123456", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "14", Name: "MRA 101", Prog_Pass: "990101", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "15", Name: "Netalsan Obsidian", Prog_Pass: "654321", Door_Pass: "Undefined", Desc: "• Press the 3 key while in the Admin menu.\\n • Key in and confirm the apartment number where you will enter/change or delete the door unlock password\\n and confirm.\\n • Press 1 key for PASSWORD ENTER/CHANGE, 2 key for DELETE PASSWORD\\n .\\n • If the password entry/change operation is selected, re-enter the 4-digit\\n door unlocking password in the PASSWORD: message, and then re-enter the 4-digit\\n door unlocking password in the PASSWORD REPEAT: message. Seeing PASSWORD CHANGED message\\n means that the password has been defined."},
    {ID: "16", Name: "ONDRIVE ED07", Prog_Pass: "1234", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "17", Name: "OP705", Prog_Pass: "*9999#", Door_Pass: "#8888#", Desc: "Undefined"},
    {ID: "18", Name: "OP M400", Prog_Pass: "0000", Door_Pass: "1234", Desc: "Undefined"},
    {ID: "19", Name: "OP M500", Prog_Pass: "0000", Door_Pass: "1234", Desc: "Undefined"},
    {ID: "20", Name: "Pratik Kart", Prog_Pass: "0000", Door_Pass: "1234#", Desc: "Undefined"},
    {ID: "21", Name: "Desi Steely", Prog_Pass: "1111", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "22", Name: "Audio", Prog_Pass: "*#9999 Eski panellerde **9999", Door_Pass: "**1234", Desc: "Undefined"},
    {ID: "23", Name: "Teknoline", Prog_Pass: "#9008", Door_Pass: "Undefined", Desc: "- Change password \\n 1# 66666666 -> 2 -> 24 -> NEW_PASS"},
    {ID: "24", Name: "Teknoline IMR18", Prog_Pass: "1 # 66666666", Door_Pass: "1111", Desc: "#9008 66666666"},
    {ID: "25", Name: "Desi UTOPIC", Prog_Pass: "Undefined", Door_Pass: "Undefined", Desc: "- Changing the Door Password \\n M A 12345"},
    {ID: "26", Name: "WL02", Prog_Pass: "0000", Door_Pass: "1234", Desc: "- Deleting all passwords \\n 253 8888#8888"},
    {ID: "27", Name: "D45", Prog_Pass: "686868", Door_Pass: "2000", Desc: "Undefined"},
    {ID: "28", Name: "A20 Kapı Kiliti", Prog_Pass: "12345678 ve **12345678", Door_Pass: "88888888", Desc: "- Changing the Door Password: \\n Enter the menu \\n Press 2->2 \\n Enter the current password. \\n Enter the new password. \\n Alarm Off 2->6->1->0 \\n Firmware via USB \\n 3->2"},
    {ID: "29", Name: "Kwikset 275 Contemporary Deadbolt", Prog_Pass: "0000", Door_Pass: "Undefined", Desc: "1. Make sure the door is open and unlocked. Press and hold the Reset button for 5 seconds, until you hear 3 long beeps.\\n2. Enter the default Programming Code (O-O-O-0).\\n3. Press enter arrow. You will hear one beep.\\n4. Press 0.\\n5. Press enter arrow. The latch bolt will extend and retract to learn the\\norientation of the door. If successful, you will hear 2 beeps. If unsuccessful, you will hear 3 beeps (make sure the door is open and unlocked and that you are using a fresh set of batteries).\\n6. Test the lock: Activate the screen. With the door open and unlocked, press the Lock symbol. Make sure it locks the door.\\n7. Test the default User Code: Activate the screen. Touch the random digits that appear. Enter the default User Code (1-2-3-4), then press the Lock symbol. Make sure it unlocks the door.\\nThe default Programming Code is 0-0-0-0. It is recommended that you change it to a code of your own.\\nThe lock is pre-programmed with a default User Code of 1-2-3-4. It is recommended that you delete this code."},
}

func main() {

	router := gin.Default()

	router.GET("/smartLocks", getSmartLocks)
	router.GET("/smartLocks/:ID", getSmartLockByID)
	router.POST("/smartLocks", postSmartLocks)
	router.PUT("/smartLocks/:ID", changeSmartLock)
	router.DELETE("/smartLocks/:ID", deleteSmartLocks)

	router.Run()
}

// getSmartLocks responds with the list of all smartLocks as JSON.
func getSmartLocks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, smartLocks)
}

// postSmartLocks adds an smartLock from JSON received in the request body.
func postSmartLocks(c *gin.Context) {
	var newsmartLock smartLock

	// Call BindJSON to bind the received JSON to
	// newsmartLock.
	if err := c.BindJSON(&newsmartLock); err != nil {
		return
	}

	// Add the new smartLock to the slice.
	smartLocks = append(smartLocks, newsmartLock)
	c.IndentedJSON(http.StatusCreated, newsmartLock)
}

// getSmartLockByID locates the smartLock whose ID value matches the ID
// parameter sent by the client, then returns that smartLock as a response.
func getSmartLockByID(c *gin.Context) {
	ID := c.Param("ID")

	// Loop through the list of smartLocks, looking for
	// an smartLock whose ID value matches the parameter.
	for _, a := range smartLocks {
		if a.ID == ID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "smartLock not found"})
}

func changeSmartLock(c *gin.Context) {
	ID := c.Param("ID")
	var newsmartLock smartLock
	intVar, _ := strconv.Atoi(ID)
	if err := c.BindJSON(&newsmartLock); err != nil {
		return
	} else {
		if newsmartLock.ID != "" {
			smartLocks[intVar-1].ID = newsmartLock.ID
		}
		if newsmartLock.Name != "" {
			smartLocks[intVar-1].Name = newsmartLock.Name
		}
		if newsmartLock.Prog_Pass != "" {
			smartLocks[intVar-1].Prog_Pass = newsmartLock.Prog_Pass
		}
		if newsmartLock.Door_Pass != "" {
			smartLocks[intVar-1].Door_Pass = newsmartLock.Door_Pass
		}
		if newsmartLock.Desc != "" {
			smartLocks[intVar-1].Desc = newsmartLock.Desc
		}
		c.IndentedJSON(http.StatusCreated, smartLocks[intVar-1])
		return
	}

	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "smartLock not found"})
}

func deleteSmartLocks(c *gin.Context) {
	ID := c.Param("ID")
	intVar, _ := strconv.Atoi(ID)
	//fmt.Println(smartLocks)
	smartLocks = RemoveIndex(smartLocks, intVar-1)
	//fmt.Print(smartLocks)
	c.IndentedJSON(http.StatusOK, smartLocks[intVar-1])
}
