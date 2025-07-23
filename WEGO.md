# WeGo debugging info

Both geocode and reverse geocode fail on a number of places, while the app search does succeed.

One possibility is that the data set is just not the same, which sucks.

Generally speaking, WeGo API is just awful in Japan, Finland, and probably other countries.

Also, geocode centered around lat and long is not fenced.
It will return places that are way far out from that lat and long.

## Geocoding failures

Some of these failures are "expected" (place does not exist anymore, or is not a place at all), but most are not.

```
10:49AM address="00192 Rome, Metropolitan City of Rome Capital, Italy" latitude=41.9076481 longitude=12.460114 name="Metropolitan City of Rome Capital"
10:49AM address="3-chōme-4-20 Hondamachi, Kanazawa, Ishikawa 920-0964, Japan" latitude=36.5576632 longitude=136.6609126 name="D.T. Suzuki Museum"
10:49AM address="Table Mountain (Nature Reserve), Cape Town, South Africa" latitude=-33.9575237 longitude=18.4029375 name="Upper Cable Car Station"
10:49AM address="2325 3rd St Floor 4R, San Francisco, CA 94107, USA" latitude=37.7601343 longitude=-122.3882609 name="Letterform Archive"
10:49AM address="Between 2810 &, 2812 Bayview Dr, Alameda, CA 94501, USA" latitude=37.7519531 longitude=-122.2437122 name="Public Walkway & Little Free Library with View of SF Bay"
10:49AM address="R. Mal. Saldanha 1, 1249-069 Lisboa, Portugal" latitude=38.7101732 longitude=-9.1471679 name="Pharmacy Museum"
10:49AM address="816 Folsom St, San Francisco, CA 94107, USA" latitude=37.781883 longitude=-122.4017023 name=Aphotic
10:49AM address="6-chōme-26-6 Tsukiji, Chuo City, Tokyo 104-0045, Japan" latitude=35.6646412 longitude=139.7713216 name="鮨聖"
10:49AM address="Japan, 〒150-0013 Tokyo, Shibuya, Ebisu, 3-chōme−28−２ SP15 EBISU１F" latitude=35.6437958 longitude=139.7156051 name="Toriyaki Ohana"
10:49AM address="Japan, 〒920-0849 Ishikawa, Kanazawa, Horikawashinmachi, 3−１ 6F" latitude=36.57934789999999 longitude=136.6498105 name="Morimori Sushi"
10:49AM address="11-1 Murasakino Higashifujinomorichō, Kita Ward, Kyoto, 603-8223, Japan" latitude=35.0375966 longitude=135.7467326 name="Teuichi Kanei Soba"
10:49AM ERR Giving up!!!!!!
10:49AM address="Piazza della Trinità dei Monti, 2, 00187 Roma RM, Italy" latitude=41.9077688 longitude=12.4824968 name="Ristorante Terrazza Ciampini di Marco Ciampini"
10:49AM address="130 1st Ave, New York, NY 10009, USA" latitude=40.7273481 longitude=-73.9851925 name="The Wild Son Lunch Counter"
10:49AM address="1 Rue de Cotte, 75012 Paris, France" latitude=48.8481646 longitude=2.3766494 name="Le Cotte Rôti"
10:49AM address="Japan, 〒107-0062 Tokyo, Minato City, Minamiaoyama, 5-chōme−4−４０ A FLAG 骨董通り" latitude=35.6621204 longitude=139.7131472 name="GENDY 南青山店"
10:49AM address="590 20th St, San Francisco, CA 94107, USA" latitude=37.7607716 longitude=-122.3873341 name="RH San Francisco | The Gallery at the Historic Bethlehem Steel Building"
10:49AM address="Calle del Scaleter, 2235, 30125 Venezia VE, Italy" latitude=45.4388571 longitude=12.3291232 name="Monica Daniele per Tabarro San Marco"
10:49AM address="Luolavuorentie 2, 20700 Turku, Finland" latitude=60.4405351 longitude=22.274051 name="Tyks Surgical Hospital"
10:49AM address="Table Mountain (Nature Reserve), Cape Town, South Africa" latitude=-33.9575237 longitude=18.4029375 name="Upper Cable Car Station"
10:50AM address="1 Rue de Cotte, 75012 Paris, France" latitude=48.8481646 longitude=2.3766494 name="Le Cotte Rôti"
10:50AM address="901 Divisadero St, San Francisco, CA 94115, USA" latitude=37.7779458 longitude=-122.4385695 name="Kava Lounge"
10:50AM address="1 Rue de Cotte, 75012 Paris, France" latitude=48.8481646 longitude=2.3766494 name="Le Cotte Rôti"
10:50AM address="C. 62 466-x 55-y 57, Parque Santa Lucia, Centro, 97000 Mérida, Yuc., Mexico" latitude=20.9705486 longitude=-89.62357469999999 name="Museum of Yucateca Gastronomy"
10:50AM address="P.º de Montejo 490, Zona Paseo Montejo, Centro, 97000 Mérida, Yuc., Mexico" latitude=20.9892552 longitude=-89.6171559 name="DECK Roof Lounge"
10:50AM address="130 1st Ave, New York, NY 10009, USA" latitude=40.7273481 longitude=-73.9851925 name="The Wild Son Lunch Counter"
10:50AM address="4166 24th St, San Francisco, CA 94114, USA" latitude=37.75131200000001 longitude=-122.4355664 name="Bon AppeTikka"
10:50AM address="901 Divisadero St, San Francisco, CA 94115, USA" latitude=37.7779458 longitude=-122.4385695 name="Kava Lounge"
10:50AM address="359 Divisadero St, San Francisco, CA 94117, USA" latitude=37.773 longitude=-122.43759 name="Vinyl Coffee & Wine Bar"
10:50AM address="NEWPORT HOUSE, Victoria Junction Complex, 123 Prestwich St, De Waterkant, Cape Town, 8001, South Africa" latitude=-33.9124343 longitude=18.4178815 name="No load shedding! 5 Star loft - private patio - Wifi 40Mbps - Pool - Waterkant Loft"
10:50AM address="Colmant rd, Franschhoek, 7690, South Africa" latitude=-33.9219841 longitude=19.1250069 name="Colmant Cap Classique and Champagne"
10:50AM address="Leeu Estates Dassenberg Road Franschhoek, Cape Town, 7690, South Africa" latitude=-33.9091014 longitude=19.1037483 name="Leeu Passant Winery"
10:50AM address="Vilhonkatu 2, 24100 Salo, Finland" latitude=60.3853006 longitude=23.1263324 name="Salo Main Library"
10:50AM address="26 Đ. Lê Anh Xuân, Phường Bến Thành, Quận 1, Hồ Chí Minh, Vietnam" latitude=10.7719108 longitude=106.6952932 name="Saigon Indian Restaurant – The Original Authentic South & North Indian Cuisine Since 1998 at District 1 in Saigon"
10:50AM address="142 Đường Trần Hưng Đạo, Khu Phố 7, Phú Quốc, Kiên Giang, Vietnam" latitude=10.1957411 longitude=103.967749 name="심커피"
10:50AM address="Japan, 〒160-0022 Tokyo, Shinjuku City, Shinjuku, 2-chōme−9−２０ Lions Mansion Shinjukugyoenmae, １Ｆ" latitude=35.6899066 longitude=139.7098321 name="CoCo Ichibanya Tokyo Metro Shinjuku Gyoenmae station shop"
10:50AM address="1 Chome-2 Nishishinjuku, Shinjuku City, Tokyo 160-0023, Japan" latitude=35.6929614 longitude=139.6995724 name="Omoide Yokocho Memory Lane"
10:50AM address="Japan, 〒150-0033 Tokyo, Shibuya, Sarugakuchō, 17−５ 代官山Ｔ－ＳＩＴＥ １号館～３号館 １階～２階" latitude=35.6488866 longitude=139.6997888 name="Tsutaya Books Daikanyama"
10:50AM address="2-chōme-8-1 Nishishinjuku, Shinjuku City, Tokyo 163-8001, Japan" latitude=35.6894807 longitude=139.6916863 name="Tokyo Metropolitan Government Building"
10:50AM address="Japan, 〒104-0061 Tokyo, Chuo City, Ginza, 3-chōme−2−１ マロニエゲート銀座２ １階～４階" latitude=35.67376369999999 longitude=139.7651281 name=UNIQLO
10:50AM address="Japan, 〒104-0061 Tokyo, Chuo City, Ginza, 6-chōme−9−５ ギンザコマツ東館 1－12F" latitude=35.6702448 longitude=139.7634686 name="Uniqlo Ginza Flagship Store"
10:50AM address="Japan, 〒160-0022 Tokyo, Shinjuku City, Shinjuku, 1-chōme−4−１３ 溝呂木第2ビル 2F" latitude=35.688218 longitude=139.711627 name="MONKEY GYM（モンキージム）"
10:50AM address="11 Naitōmachi, Shinjuku City, Tokyo 160-0014, Japan" latitude=35.68517629999999 longitude=139.7100517 name="Shinjuku Gyoen National Garden"
10:50AM address="21-4 Nishichōyabunouchidōri, Kanazawa, Ishikawa 920-0915, Japan" latitude=36.5697329 longitude=136.655615 name="エスコーラ焙煎考房"
10:50AM address="Japan, 〒150-0013 Tokyo, Shibuya, Ebisu, 3-chōme−28−２ SP15 EBISU１F" latitude=35.6437958 longitude=139.7156051 name="Toriyaki Ohana"
10:50AM address="2-chōme-1-1 Kōrinbō, Kanazawa, Ishikawa 920-0961, Japan" latitude=36.5623689 longitude=136.6531013 name="UNIQLO Tokyu Square Korinbo Store"
10:50AM address="1-chōme-2-2-12 Nomachi, Kanazawa, Ishikawa 921-8031, Japan" latitude=36.5553843 longitude=136.6489974 name="Myouryuji (Ninja Temple)"
10:50AM address="富士ビル 3f, 2-chōme-62 Senju, Adachi City, Tokyo 120-0034, Japan" latitude=35.7493452 longitude=139.8040087 name="貝と魚と炉ばたのバンビ"
10:50AM address="2-chōme-10-9 Kajichō, Chiyoda City, Tokyo 101-0044, Japan" latitude=35.6936985 longitude=139.7724063 name="Karashibi Miso Ramen Kikanbō Kanda Honten"
10:50AM address="Japan, 〒106-0031 Tokyo, Minato City, Nishiazabu, 1-chōme−4−１５ 寿司勇" latitude=35.6618348 longitude=139.725206 name="Sushi Yuu"
10:50AM address="329 Ebiyachō, Gokomachi-dori Sanjo sagaru, 329 海老屋町 中京区 京都市 京都府 604-8076, Japan" latitude=35.0076598 longitude=135.7663712 name="KIMONO TEA CEREMONY MAIKOYA NISHIKI KYOTO"
10:51AM ERR Giving up!!!!!!
10:51AM address="Japan, 〒543-0045 Osaka, Tennoji Ward, Teradachō, 2-chōme−5−１ 石村ビル 1F" latitude=34.6492539 longitude=135.5243882 name="Tetsudoukan Cafe and Rest Bar"
10:51AM address="Japan, 〒532-0011 Osaka, Yodogawa Ward, Nishinakajima, 5-chōme−1−４ モジュール新大阪 1F" latitude=34.7283309 longitude=135.5011788 name="Ramen Tokiya"
10:51AM address="Japan, 〒530-0003 Osaka, Kita Ward, Dōjima, 1-chōme−3−１９ 薬師堂ビル 3F" latitude=34.6969006 longitude=135.4974092 name="Sushi Dokoro Shinmon"
10:51AM address="2-chōme-3-22 Nishishinsaibashi, Chuo Ward, Osaka, 542-0086, Japan" latitude=34.6702053 longitude=135.4997748 name="BAR AQUAVIT"
10:51AM ERR Giving up!!!!!!
10:51AM address="Japan, 〒542-0076 Osaka, Chuo Ward, Namba, 1-chōme−7−１５ 江戸安ビル B1F" latitude=34.6680616 longitude=135.5007881 name="Okonomiyaki Sakaba O"
10:51AM address="3-chōme-2-4 Nakanoshima, Kita Ward, Osaka, 530-0005, Japan" latitude=34.6935117 longitude=135.4953805 name="40 Sky Bar & Lounge"
10:51AM ERR Giving up!!!!!!
10:51AM address="Đông Du/80 Ward, Bến Nghé, Quận 1, Hồ Chí Minh 700000, Vietnam" latitude=10.7757706 longitude=106.7042346 name="Level 23 Wine Bar - Sheraton Saigon Grand Opera Hotel"
10:51AM address="Đông Du/80 Ward, Bến Nghé, Quận 1, Hồ Chí Minh 700000, Vietnam" latitude=10.7757706 longitude=106.7042346 name="Level 23 Wine Bar - Sheraton Saigon Grand Opera Hotel"
10:51AM address="5-chōme-21-9 Shirokanedai, Minato City, Tokyo 108-0071, Japan" latitude=35.6369465 longitude=139.7190408 name="Tokyo Metropolitan Teien Art Museum"
10:51AM address="1-chōme-5-3 Yaesu, Chuo City, Tokyo 103-0028, Japan" latitude=35.6820486 longitude=139.7709832 name="Karaksa Hotel TOKYO STATION"
10:51AM address="1-chōme-1-2 Oshiage, Sumida City, Tokyo 131-0045, Japan" latitude=35.7100627 longitude=139.8107004 name="Tokyo Skytree"
10:51AM address="Japan, 〒104-0061 Tokyo, Chuo City, Ginza, 6-chōme−3−５ 第二 ソワレ・ド ビル" latitude=35.6716127 longitude=139.7614467 name="Hashigo Ginza Hon-ten"
10:51AM address="Japan, 〒160-0022 Tokyo, Shinjuku City, Shinjuku, 2-chōme−4−１ 第22宮庭マンション 1階105号室" latitude=35.6887843 longitude=139.7082917 name="Soba House Konjiki-Hototogisu"
10:51AM address="Japan, 〒150-0042 Tokyo, Shibuya, Udagawachō, 21−１ A Building, 8階 SEIBU Shibuya Store" latitude=35.6602148 longitude=139.7005026 name="Mawashizushi Katsu Seibu Shibuya Store"
10:51AM address="Via dei Banchi Nuovi, 39, 00186 Roma RM, Italy" latitude=41.8992581 longitude=12.4679841 name="Rituals Navona - Wellness & Spa Roma"
10:51AM address="2411 Alaskan Wy, Seattle, WA 98121, USA" latitude=47.61236599999999 longitude=-122.3522377 name="The Brim Coffee Shop at the Edgewater Hotel"
10:51AM address="424 Boren Ave N, Seattle, WA 98109, USA" latitude=47.6229396 longitude=-122.3354688 name="Tuk Tuk Mobile Feast"
10:51AM address="1506 Pike Pl #509, Seattle, WA 98101, USA" latitude=47.6088566 longitude=-122.3405796 name="Oriental Mart Filipino Restaurant"
10:51AM address="Luolavuorentie 2, 20700 Turku, Finland" latitude=60.4405351 longitude=22.274051 name="Tyks Surgical Hospital"
10:51AM address="1129 Sebastopol Rd, Santa Rosa, CA 95407, USA" latitude=38.4290857 longitude=-122.7329441 name="Sazón Authentic Peruvian Cuisine"
10:51AM address="53 Montgomery Dr, Santa Rosa, CA 95404, USA" latitude=38.4420097 longitude=-122.7035646 name="Rosso Pizzeria & Wine Bar"
10:51AM address="Calçada do Carmo 25 Sobreloja esquerda, 1200-090 Lisboa, Portugal" latitude=38.7134557 longitude=-9.1405263 name="Eu lavo seus cabelos com chás"
10:51AM address="Rua de O Século 158 A, 1200-437 Lisboa, Portugal" latitude=38.7148257 longitude=-9.1475204 name="Rhamus Hair Principe Real"
10:51AM address="R. São Boaventura 42, 1200-385 Lisboa, Portugal" latitude=38.7147079 longitude=-9.146201 name="Laundry Self Service - Cheir'a Lisbon"
10:51AM address="Bd de la Ligne, 84000 Avignon, France" latitude=43.9538522 longitude=4.8048781 name="The Bridge of Avignon"
10:52AM address="25 Rue Boulan, 33000 Bordeaux, France" latitude=44.8385785 longitude=-0.5812108 name="O-Zone Gym Avec Piscine"
10:52AM address="Marché des Capucins, Pl. des Capucins, 33800 Bordeaux, France" latitude=44.8307167 longitude=-0.5677801 name="Bistrot à huitres: \"Chez Jean-Mi\""
10:52AM address="Via Orolungo, 34, 13853 Lessona BI, Italy" latitude=45.5555703 longitude=8.219852099999999 name="Proprietà Sperino"
10:52AM address="Vaasankatu 7, 00500 Helsinki, Finland" latitude=60.18840299999999 longitude=24.9596758 name="Beer Restaurant Hilpeä Hauki"
10:52AM address="Fredrikinkatu 55, 00100 Helsinki, Finland" latitude=60.16724860000001 longitude=24.9332438 name="Helkatti Cat Cafe"
10:52AM address="Pohjoisranta 1, 00170 Helsinki, Finland" latitude=60.1693965 longitude=24.9588658 name="Restaurant Lightvessel Relandersgrund"
10:52AM address="Kiikunmäentie 28, 25660 Salo, Finland" latitude=60.2216742 longitude=22.9100949 name="Meri-Ruukin lomakylä/Meri-Ruukki Holiday Village INFO"
10:52AM address="C/ de Muntaner, 69, L'Eixample, 08011 Barcelona, Spain" latitude=41.38690709999999 longitude=2.1588377 name="kway jaquetes"
10:52AM address="1226 3rd St, Napa, CA 94559, USA" latitude=38.296956 longitude=-122.286327 name="Folklore Records+Drinks+Food"
10:52AM address="C. de Álvarez Gato, 4, Centro, 28012 Madrid, Spain" latitude=40.4150623 longitude=-3.7019135 name="Inclan Brutal Bar I Restaurante temático Madrid"
10:52AM address="816 Folsom St, San Francisco, CA 94107, USA" latitude=37.781883 longitude=-122.4017023 name=Aphotic
10:52AM address="Colmant rd, Franschhoek, 7690, South Africa" latitude=-33.9219841 longitude=19.1250069 name="Colmant Cap Classique and Champagne"
10:52AM address="Leeu Estates Dassenberg Road Franschhoek, Cape Town, 7690, South Africa" latitude=-33.9091014 longitude=19.1037483 name="Leeu Passant Winery"
10:52AM address="Vilhonkatu 2, 24100 Salo, Finland" latitude=60.3853006 longitude=23.1263324 name="Salo Main Library"
10:52AM address="118 Nguyễn Văn Thủ, Đa Kao, Quận 1, Hồ Chí Minh, Vietnam" latitude=10.7884646 longitude=106.6975431 name="GALERIE QUYNH CONTEMPORARY ART"
10:53AM address="Japan, 〒282-0004 Chiba, Narita, Furugome, 1−１ Terminal 1, Center 保安検査後(国際線 サテライト連絡通路エスカレーター先" latitude=35.774227 longitude=140.3898776 name="Soke Minamoto Kitchoan Narita Airport Terminal 2"
10:53AM address="Japan, 〒150-0033 Tokyo, Shibuya, Sarugakuchō, 17−５ 代官山Ｔ－ＳＩＴＥ １号館～３号館 １階～２階" latitude=35.6488866 longitude=139.6997888 name="Tsutaya Books Daikanyama"
10:53AM address="2-chōme-1-11 Nishiasakusa, Taito City, Tokyo 111-0035, Japan" latitude=35.7119914 longitude=139.7918882 name="Benitsuru Pancake"
10:53AM address="3-chōme-4-20 Hondamachi, Kanazawa, Ishikawa 920-0964, Japan" latitude=36.5576632 longitude=136.6609126 name="D.T. Suzuki Museum"
10:53AM address="Japan, 〒920-0849 Ishikawa, Kanazawa, Horikawashinmachi, 3−１ 6F" latitude=36.57934789999999 longitude=136.6498105 name="Morimori Sushi"
10:53AM address="11-1 Murasakino Higashifujinomorichō, Kita Ward, Kyoto, 603-8223, Japan" latitude=35.0375966 longitude=135.7467326 name="Teuichi Kanei Soba"
10:53AM ERR Giving up!!!!!!
10:53AM address="Japan, 〒553-0003 Osaka, Fukushima Ward, Fukushima, 2-chōme−8−２ 1階" latitude=34.6945818 longitude=135.4870692 name="Hanakujira honten"
10:53AM address="Japan, 〒553-0003 Osaka, Fukushima Ward, Fukushima, 1-chōme−5−２ 堀野ビル １階" latitude=34.6952289 longitude=135.4886857 name="La Pizza Napoletana Regalo"
10:53AM address="Japan, 〒530-0003 Osaka, Kita Ward, Dōjima, 1-chōme−5−１ エスパス北新地２３ １Ｆ" latitude=34.695967 longitude=135.497041 name="Yakitori Ichimatsu"
10:53AM address="Japan, 〒530-0004 Osaka, Kita Ward, Dōjimahama, 2-chōme−1−２ OFFICE 堂島" latitude=34.6948219 longitude=135.4954578 name="Pâtisserie Mon cher"
10:53AM address="R. Mal. Saldanha 1, 1249-069 Lisboa, Portugal" latitude=38.7101732 longitude=-9.1471679 name="Pharmacy Museum"
10:53AM address="SM023 M 54 L1 Av Chichén Itzá, Jabín Supermanzana 23, 77500 Cancún, Q.R., Mexico" latitude=21.168038 longitude=-86.82853539999999 name="PUNTO FRIO
```