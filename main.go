package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "encoding/json"
    "io"
)


func PrettyEncode(data interface{}, out io.Writer) error {
    enc := json.NewEncoder(out)
    enc.SetIndent("", "    ")
    if err := enc.Encode(data); err != nil {
        return err
    }
    return nil
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
    {ID: "1", Name: "Perkotek", Prog_Pass: "666888", Door_Pass: "Admin Şifresi : *#1234 8'e bas 0 'a bas", Desc: "- Alarmları Kapatma *#1234 0'a bas 1'e bas 0'a bas \\n- Bütün Kullanıcı Şifrelerini Silme 4444 \\n- Kapının Açık Kalma Süresinin Ayarlanması 5XX (XX sayısı 00-99 arasında saniye olarak zamanı gösterir.)"},
    {ID: "2", Name: "ERD-1120", Prog_Pass: "#4567#", Door_Pass: "Undefined", Desc: "- Sistem şifresini değiştirmek için *#SISTEM_PASS 8 NEW_PASS"},

    {ID: "3", Name: "Efes Digital Panel", Prog_Pass: "0000", Door_Pass: "1234", Desc: "- Sistem şifresini değiştirmek için CALL -> 1234 -> CALL -> NEW_PASS"},

    {ID: "4", Name: "Mas", Prog_Pass: "12345", Door_Pass: "1234", Desc: "- 0 tuşuna 10 saniye basılı tutunuz. Üç led yanıp sönecektir. Yönetici şifresini giriniz. (Yönetici şifresi fabrika çıkışı: 12345'tir.) Yönetici şifresi doğru ise üç led bir kez yanıp sönecektir. 1 tuşuna basınız. Üç led yanıp sönecektir. Belirlediğiniz 4 rakamlı yeni kapı şifresini giriniz. \\n Yönetici Şifresi Değişme: 0 tuşuna 10 saniye basılı tutunuz. Üç led yanıp sönecektir. Yönetici şifresini giriniz. (Yönetici şifresi fabrika çıkışı: 12345'tir.) Yönetici şifresi doğru ise üç led bir kez yanıp sönecektir. 2 tuşuna basınız. Üç led yanıp sönecektir. Belirlediğiniz 5 rakamlı yeni yönetici şifresini giriniz."},

    {ID: "5", Name: "AC 13PX", Prog_Pass: "*1234", Door_Pass: "Undefined", Desc: "- Kullanıcı Ekleme \\n *1234 -> 3"},

    {ID: "6", Name: "Burg Wachter", Prog_Pass: "123456", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "7", Name: "DIP40", Prog_Pass: "0000", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "8", Name: "Lorex LR-DPH2", Prog_Pass: "999999", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "9", Name: "M100", Prog_Pass: "*9999#", Door_Pass: "#8888#", Desc: "- Kapı Şifresi Değişme \\n 92# Yeni Şifre"},
    {ID: "10", Name: "MB05-03", Prog_Pass: "7447", Door_Pass: "Undefined", Desc: "- Şifre Değiştirme \\n CALL -> 99 -> 7447"},
    {ID: "11", Name: "MB DYF40", Prog_Pass: "7447", Door_Pass: "Undefined", Desc: "- Şifre Değiştirme \\n CALL -> 99 -> 7447"},
    {ID: "12", Name: "MLŞ 14-70", Prog_Pass: "Undefined", Door_Pass: "0000", Desc: "Undefined"},
    {ID: "13", Name: "MLŞ 14-107", Prog_Pass: "123456", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "14", Name: "MRA 101", Prog_Pass: "990101", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "15", Name: "Netalsan Obsidian", Prog_Pass: "654321", Door_Pass: "Undefined", Desc: "• Yönetici menüsünde iken 3 tuşuna basınız.\\n • Kapı açma şifresini gireceğiniz/değiştireceğiniz ya da sileceğiniz\\n daire numarasını tuşlayınız ve onaylayınız.\\n • ŞiFRE GiRME / DEĞiŞTiRME için 1 , ŞiFRE SiLME için 2 tuşuna\\n basınız.\\n • Şifre girme/değiştirme işlemi seçildiyse ŞİFRE : mesajına 4 haneli\\n kapı açma şifresini, ardından ŞİFRE TEKRARI : mesajına 4 haneli\\n kapı açma şifresini tekrar giriniz. ŞiFRE DEĞİŞTİRİLDİ mesajı\\n görülmesi şifrenin tanımlandığı anlamında gelir."},
    {ID: "16", Name: "ONDRIVE ED07", Prog_Pass: "1234", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "17", Name: "OP705", Prog_Pass: "*9999#", Door_Pass: "#8888#", Desc: "Undefined"},
    {ID: "18", Name: "OP M400", Prog_Pass: "0000", Door_Pass: "1234", Desc: "Undefined"},
    {ID: "19", Name: "OP M500", Prog_Pass: "0000", Door_Pass: "1234", Desc: "Undefined"},
    {ID: "20", Name: "Pratik Kart", Prog_Pass: "0000", Door_Pass: "1234#", Desc: "Undefined"},
    {ID: "21", Name: "Desi Steely", Prog_Pass: "1111", Door_Pass: "Undefined", Desc: "Undefined"},
    {ID: "22", Name: "Audio", Prog_Pass: "*#9999 Eski panellerde **9999", Door_Pass: "**1234", Desc: "Undefined"},
    {ID: "23", Name: "Teknoline", Prog_Pass: "#9008", Door_Pass: "Undefined", Desc: "- Şifre değiştirmek \\n 1# 66666666 -> 2 -> 24 -> NEW_PASS"},
    {ID: "24", Name: "Teknoline IMR18", Prog_Pass: "1 # 66666666", Door_Pass: "1111", Desc: "#9008 66666666"},
    {ID: "25", Name: "Desi UTOPIC", Prog_Pass: "Undefined", Door_Pass: "Undefined", Desc: "- Kapı Şifresi değişme \\n M A 12345"},
    {ID: "26", Name: "WL02", Prog_Pass: "0000", Door_Pass: "1234", Desc: "- Tüm şifreleri silmek \\n 253  8888#8888"},
    {ID: "27", Name: "D45", Prog_Pass: "686868", Door_Pass: "2000", Desc: "Undefined"},
    {ID: "28", Name: "A20 Kapı Kiliti", Prog_Pass: "12345678 ve **12345678", Door_Pass: "88888888", Desc: "- Kapı Şifresi Değişme : \\n Menüye gir \\n 2->2 bas \\n Mevcut şifreyi gir. \\n Yeni şifreyi gir. \\n Alarm Kapatma 2->6->1->0 \\n USB ile Firmware \\n 3->2"},
}



func main() {

    router := gin.Default()

    router.GET("/smartLocks", getSmartLocks)
    router.GET("/smartLocks/:ID", getSmartLockByID)
    router.POST("/smartLocks", postSmartLocks)

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
