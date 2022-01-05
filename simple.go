package main

import (
	"fmt"

	"github.com/fatih/color"
	_ "github.com/gin-gonic/gin"
)

type smartLock struct {
	id        string `json:"id"`
	name      string `json:"name"`
	door_pass string `json:"door_pass"`
	prog_pass string `json:"prog_pass"`
	desc      string `json:"desc"`
}

func newSmartLock(id string, name string, prog_pass string, door_pass string, desc string) *smartLock {

	p := smartLock{name: name}
	p.id = id
	p.door_pass = door_pass
	p.prog_pass = prog_pass
	p.desc = desc
	return &p
}

var smartLocks = []smartLock{}

func main() {

	perkotek := smartLock{
		id:        "1",
		name:      "Perkotek",
		prog_pass: "666888",
		door_pass: "Admin Şifresi : *#1234 8'e bas 0 'a bas",
		desc:      "- Alarmları Kapatma *#1234 0'a bas 1'e bas 0'a bas \n- Bütün Kullanıcı Şifrelerini Silme 4444 \n- Kapının Açık Kalma Süresinin Ayarlanması 5XX (XX sayısı 00-99 arasında saniye olarak zamanı gösterir.)",
	}

	erd_1120 := smartLock{
		id:        "2",
		name:      "ERD-1120",
		prog_pass: "#4567#",
		door_pass: "Undefined",
		desc:      "- Sistem şifresini değiştirmek için *#SISTEM_PASS 8 NEW_PASS",
	}

	efes_dig_panel := smartLock{
		id:        "3",
		name:      "Efes Digital Panel",
		prog_pass: "0000",
		door_pass: "1234",
		desc:      "- Sistem şifresini değiştirmek için CALL -> 1234 -> CALL -> NEW_PASS",
	}

	mas_door := smartLock{
		id:        "4",
		name:      "Mas",
		door_pass: "1234",
		prog_pass: "12345",
		desc:      "- 0 tuşuna 10 saniye basılı tutunuz. Üç led yanıp sönecektir. Yönetici şifresini giriniz. (Yönetici şifresi fabrika çıkışı: 12345'tir.) Yönetici şifresi doğru ise üç led bir kez yanıp sönecektir. 1 tuşuna basınız. Üç led yanıp sönecektir. Belirlediğiniz 4 rakamlı yeni kapı şifresini giriniz. \n Yönetici Şifresi Değişme: 0 tuşuna 10 saniye basılı tutunuz. Üç led yanıp sönecektir. Yönetici şifresini giriniz. (Yönetici şifresi fabrika çıkışı: 12345'tir.) Yönetici şifresi doğru ise üç led bir kez yanıp sönecektir. 2 tuşuna basınız. Üç led yanıp sönecektir. Belirlediğiniz 5 rakamlı yeni yönetici şifresini giriniz.",
	}

	ac_13px := smartLock{
		id:        "5",
		name:      "AC 13PX",
		door_pass: "Undefined",
		prog_pass: "*1234",
		desc:      "- Kullanıcı Ekleme \n *1234 -> 3",
	}

	burg_wachter := smartLock{
		id:        "6",
		name:      "Burg Wachter",
		door_pass: "Undefined",
		prog_pass: "123456",
		desc:      "Undefined",
	}

	dip40 := smartLock{
		id:        "7",
		name:      "DIP40",
		door_pass: "Undefined",
		prog_pass: "0000",
		desc:      "Undefined",
	}

	lorex_lrdph2 := smartLock{
		id:        "8",
		name:      "Lorex LR-DPH2",
		door_pass: "Undefined",
		prog_pass: "999999",
		desc:      "Undefined",
	}

	m100 := smartLock{
		id:        "9",
		name:      "M100",
		door_pass: "#8888#",
		prog_pass: "*9999#",
		desc:      "- Kapı Şifresi Değişme \n 92# Yeni Şifre",
	}

	mb05_03 := smartLock{
		id:        "10",
		name:      "MB05-03",
		door_pass: "Undefined",
		prog_pass: "7447",
		desc:      "- Şifre Değiştirme \n CALL -> 99 -> 7447",
	}

	mb_dyf40 := smartLock{
		id:        "11",
		name:      "MB DYF40",
		door_pass: "Undefined",
		prog_pass: "7447",
		desc:      "- Şifre Değiştirme \n CALL -> 99 -> 7447",
	}

	mls_1470 := smartLock{
		id:        "12",
		name:      "MLŞ 14-70",
		door_pass: "0000",
		prog_pass: "Undefined",
		desc:      "Undefined",
	}

	mls_14107 := smartLock{
		id:        "13",
		name:      "MLŞ 14-107",
		door_pass: "Undefined",
		prog_pass: "123456",
		desc:      "Undefined",
	}

	mra_101 := smartLock{
		id:        "14",
		name:      "MRA 101",
		door_pass: "Undefined",
		prog_pass: "990101",
		desc:      "Undefined",
	}

	netalsan_obsidian := smartLock{
		id:        "15",
		name:      "Netalsan Obsidian",
		door_pass: "Undefined",
		prog_pass: "654321",
		desc:      "• Yönetici menüsünde iken 3 tuşuna basınız.\n • Kapı açma şifresini gireceğiniz/değiştireceğiniz ya da sileceğiniz\n daire numarasını tuşlayınız ve onaylayınız.\n • ŞiFRE GiRME / DEĞiŞTiRME için 1 , ŞiFRE SiLME için 2 tuşuna\n basınız.\n • Şifre girme/değiştirme işlemi seçildiyse ŞİFRE : mesajına 4 haneli\n kapı açma şifresini, ardından ŞİFRE TEKRARI : mesajına 4 haneli\n kapı açma şifresini tekrar giriniz. ŞiFRE DEĞİŞTİRİLDİ mesajı\n görülmesi şifrenin tanımlandığı anlamında gelir.",
	}

	ondrive_ed07 := smartLock{
		id:        "16",
		name:      "ONDRIVE ED07",
		door_pass: "Undefined",
		prog_pass: "1234",
		desc:      "Undefined",
	}

	op705 := smartLock{
		id:        "17",
		name:      "OP705",
		door_pass: "#8888#",
		prog_pass: "*9999#",
		desc:      "Undefined",
	}

	op_m400 := smartLock{
		id:        "18",
		name:      "OP M400",
		door_pass: "1234",
		prog_pass: "0000",
		desc:      "Undefined",
	}

	op_m500 := smartLock{
		id:        "19",
		name:      "OP M500",
		door_pass: "1234",
		prog_pass: "0000",
		desc:      "Undefined",
	}

	pratik_kart := smartLock{
		id:        "20",
		name:      "Pratik Kart",
		door_pass: "1234#",
		prog_pass: "0000",
		desc:      "Undefined",
	}

	desi_steely := smartLock{
		id:        "21",
		name:      "Desi Steely",
		door_pass: "Undefined",
		prog_pass: "1111",
		desc:      "Undefined",
	}

	audio := smartLock{
		id:        "22",
		name:      "Audio",
		door_pass: "**1234",
		prog_pass: "*#9999 Eski panellerde **9999",
		desc:      "Undefined",
	}

	teknoline := smartLock{
		id:        "23",
		name:      "Teknoline",
		door_pass: "Undefined",
		prog_pass: "#9008",
		desc:      "- Şifre değiştirmek \n 1# 66666666 -> 2 -> 24 -> NEW_PASS",
	}

	teknoline_imr18 := smartLock{
		id:        "24",
		name:      "Teknoline IMR18",
		door_pass: "1111",
		prog_pass: "1 # 66666666",
		desc:      "#9008 66666666",
	}

	desi_utopic := smartLock{
		id:        "25",
		name:      "Desi UTOPIC",
		door_pass: "Undefined",
		prog_pass: "Undefined",
		desc:      "- Kapı Şifresi değişme \n M A 12345",
	}

	wl02 := smartLock{
		id:        "26",
		name:      "WL02",
		door_pass: "1234",
		prog_pass: "0000",
		desc:      "- Tüm şifreleri silmek \n 253  8888#8888",
	}

	d45 := smartLock{
		id:        "27",
		name:      "D45",
		door_pass: "2000",
		prog_pass: "686868",
		desc:      "Undefined",
	}
	a20_door := smartLock{
		id:        "28",
		name:      "A20 Kapı Kilidi",
		prog_pass: "12345678 ve **12345678",
		door_pass: "88888888",
		desc:      "- Kapı Şifresi Değişme : \n Menüye gir \n 2->2 bas \n Mevcut şifreyi gir. \n Yeni şifreyi gir. \n Alarm Kapatma 2->6->1->0 \n USB ile Firmware \n 3->2",
	}

	smartLocks = append(smartLocks, perkotek)
	smartLocks = append(smartLocks, erd_1120)
	smartLocks = append(smartLocks, efes_dig_panel)
	smartLocks = append(smartLocks, mas_door)
	smartLocks = append(smartLocks, ac_13px)
	smartLocks = append(smartLocks, burg_wachter)
	smartLocks = append(smartLocks, dip40)
	smartLocks = append(smartLocks, lorex_lrdph2)
	smartLocks = append(smartLocks, m100)
	smartLocks = append(smartLocks, mb05_03)
	smartLocks = append(smartLocks, mb_dyf40)
	smartLocks = append(smartLocks, mls_1470)
	smartLocks = append(smartLocks, mls_14107)
	smartLocks = append(smartLocks, mra_101)
	smartLocks = append(smartLocks, netalsan_obsidian)
	smartLocks = append(smartLocks, ondrive_ed07)
	smartLocks = append(smartLocks, op705)
	smartLocks = append(smartLocks, op_m400)
	smartLocks = append(smartLocks, op_m500)
	smartLocks = append(smartLocks, pratik_kart)
	smartLocks = append(smartLocks, desi_steely)
	smartLocks = append(smartLocks, audio)
	smartLocks = append(smartLocks, teknoline)
	smartLocks = append(smartLocks, teknoline_imr18)
	smartLocks = append(smartLocks, desi_utopic)
	smartLocks = append(smartLocks, wl02)
	smartLocks = append(smartLocks, d45)
	smartLocks = append(smartLocks, a20_door)

	//fmt.Print(smartLocks[0])
	for i := range smartLocks {
		color.HiYellow(smartLocks[i].id + " - " + smartLocks[i].name)
	}

	color.HiRed("Choose : ")
	// var then variable name then variable type
	var first int

	// Taking input from user
	fmt.Scanln(&first)

	for i := range smartLocks {
		//fmt.Println(smartLocks[i].door_pass)
		if i == first {
			color.HiRed("ID : " + smartLocks[i-1].id)
			color.HiRed("Name : " + smartLocks[i-1].name)
			color.HiRed("Door Password : " + smartLocks[i-1].door_pass)
			color.HiRed("Programming Password : " + smartLocks[i-1].prog_pass)
			color.HiRed("Description : " + "\n" + smartLocks[i-1].desc)
		}

	}
}
