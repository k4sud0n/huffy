package template

import (
	"github.com/gofiber/fiber/v2"
	"huffy/database"
)

func CreateMenuItem(menus []database.Menu) string {
	var content string

	for _, menu := range menus {
		content = menu.Content
	}

	return content
}

func CreateMenuResponse(responseText string) fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": TemplateData{
			Outputs: []Output{
				{
					SimpleText: &SimpleText{
						Text: responseText,
					},
				},
			},
		},
	}

	return response
}

func CreateEomungwanResponse() fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": TemplateData{
			Outputs: []Output{
				{
					SimpleText: &SimpleText{
						Text: "<까르보네>\n\n크림양송이스파게티: ₩5500\n까르보나라스파게티: ₩6500\n매운까르보나라스파게티: ₩7100\n크림치킨스파게티: ₩6900\n크림고추새우스파게티: ₩7400\n매운로제스파게티: ₩8500\n미트소스스파게티: ₩5500\n토마토소스스파게티: ₩5900\n아마트리치아나스파게티: ₩6900\n알리오올리오스파게티: ₩5000\n봉골레스파게티: ₩6000\n중화짜장새우스파게티: ₩6000\n마라크림스파게티: ₩6500\n\n크림새우도리아: ₩7200\n매운낙지도리아: ₩7200\n\n베이컨포테이토피자: ₩8900\n고르곤졸라피자: ₩8900\n페페로니피자: ₩8900\n\n매콤누들떡볶이: ₩5900\n\n치즈추가 +₩1500",
					},
				},
				{
					SimpleText: &SimpleText{
						Text: "<헬로밀>\n\n오늘의 음료수 샌드위치: ₩3800\n오늘의 과일 샌드위치: ₩4100\n더블햄치즈에그 샌드위치: ₩4500\n닭가슴살 샌드위치: ₩4500\n치킨텐더 샌드위치: ₩4500\n\n에그단호박 샐러드: ₩4900\n닭가슴살 샐러드: ₩4900\n마늘우삼겹 샐러드: ₩5500\n치킨텐더 샐러드: ₩5500\n마늘새우 샐러드: ₩6000\n\n햄치즈 토스트: ₩3000\n햄치즈야채 토스트: ₩3200\n베이컨 토스트: ₩3500\n불갈비 토스트(매콤): ₩4200\n\n플레인 핫도그: ₩3600\n콘치즈 핫도그: ₩4000\n스위트어니언 핫도그: ₩4000\n스위트불고기 핫도그: ₩4200\n\n컵과일: ₩3000\n컵어묵: ₩1000\n그래과일: ₩3900\n타코야끼: ₩3000\n그래초코: ₩3900\n탄산음료: ₩1500\n\n해쉬브라운 + 탄산음료: +₩2000",
					},
				},
				{
					SimpleText: &SimpleText{
						Text: "<아리랑덮밥>\n\n김치찌개: ₩5900\n소고기오므라이스: ₩5900\n숯불소고기덮밥: ₩6900\n직화닭덮밥: ₩5900\n매운돼지갈비덮밥: ₩5900\n두부추덮밥: ₩5400\n묵은지참치덮밥: ₩5700\n참치두부덮밥: ₩5700\n된장두부삼겹덮밥: ₩5900",
					},
				},
			},
		},
	}

	return response
}

func CreateBabidundunResponse() fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": TemplateData{
			Outputs: []Output{
				{
					SimpleText: &SimpleText{
						Text: "<바비든든>\n\n싱글: ₩3300\n더블: ₩4300\n점보: ₩5300\n\n삼겹소금\n삼겹양념(매운맛)\n참치마요\n치킨마요\n스팸마요\n\n햄구이 +₩500\n체다치즈 +₩500\n계란후라이 +₩500\n고기추가 +₩1000\n참치마요추가 +₩1000\n피자치즈 +₩1000\n펩시 +₩1300\n사이다 +₩1300\n탐스제로 +₩1300",
					},
				},
			},
		},
	}

	return response
}
