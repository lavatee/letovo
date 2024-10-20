package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"

	backend "github.com/lavatee/children_backend"
	"github.com/lavatee/children_backend/internal/endpoint"
	"github.com/lavatee/children_backend/internal/repository"
	"github.com/lavatee/children_backend/internal/service"
	"github.com/spf13/viper"
)

func main() {
	children := map[int]backend.Child{
		1: {
			FirstName: "Татьтяна",
			LastName:  "Богреева",
			Id:        1,
			Gift:      "Все для маникюра: лампа, гель, полигель, обезжириватель, формы, топ, база, гель лаки, гелевые типсы, аппарат для маникюра, безворсовые салфетки",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/1%20%D0%91%D0%BE%D0%B3%D1%80%D0%B5%D0%B5%D0%B2%D0%B0%20%D0%A2%D0%B0%D1%82%D1%8C%D1%8F%D0%BD%D0%B0.jpeg",
		},
		2: {
			FirstName: "Лиля",
			LastName:  "",
			Id:        2,
			Gift:      "Большое лего",
			IsTaken:   false,
			Age:       7,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/2%20%D0%9B%D0%B8%D0%BB%D1%8F.jpeg",
		},
		3: {
			FirstName: "Слава",
			LastName:  "Богреев",
			Id:        3,
			Gift:      "Фотоаппарат",
			IsTaken:   false,
			Age:       9,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/3%20%D0%91%D0%BE%D0%B3%D1%80%D0%B5%D0%B5%D0%B2%20%D0%A1%D0%BB%D0%B0%D0%B2%D0%B0.jpeg",
		},
		4: {
			FirstName: "Артем",
			LastName:  "Черников",
			Id:        4,
			Gift:      "Большую машину с рулем и рычагом",
			IsTaken:   false,
			Age:       10,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/4%20%D0%A7%D0%B5%D1%80%D0%BD%D0%B8%D0%BA%D0%BE%D0%B2%20%D0%90%D1%80%D1%82%D0%B5%D0%BC.jpeg",
		},
		5: {
			FirstName: "Алексей",
			LastName:  "Шевченко",
			Id:        5,
			Gift:      "Аудиокнига с наушниками",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/5%20%D0%A8%D0%B5%D0%B2%D1%87%D0%B5%D0%BD%D0%BA%D0%BE%20%D0%90%D0%BB%D0%B5%D0%BA%D1%81%D0%B5%D0%B8%CC%86.jpeg",
		},
		6: {
			FirstName: "Миша",
			LastName:  "Батиенко",
			Id:        6,
			Gift:      "Спортивная форма и синие футзальные кроссовки",
			IsTaken:   false,
			Age:       15,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/6%20%D0%91%D0%B0%D1%82%D0%B8%D0%B5%D0%BD%D0%BA%D0%BE%20%D0%9C%D0%B8%D1%88%D0%B0.jpeg",
		},
		7: {
			FirstName: "Миша",
			LastName:  "Шевченко",
			Id:        7,
			Gift:      "Коньки",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/7%20%D0%A8%D0%B5%D0%B2%D1%87%D0%B5%D0%BD%D0%BA%D0%BE%20%D0%9C%D0%B8%D1%88%D0%B0.jpeg",
		},
		8: {
			FirstName: "Михаил",
			LastName:  "Кречетов",
			Id:        8,
			Gift:      "Коньки, клюшку и шайбу",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/8%20%D0%9A%D1%80%D0%B5%D1%87%D0%B5%D1%82%D0%BE%D0%B2%20%D0%9C%D0%B8%D1%88%D0%B0.jpeg",
		},
		9: {
			FirstName: "Максим",
			LastName:  "Богреев",
			Id:        9,
			Gift:      "Спортивный костюм и беспроводную колонку",
			IsTaken:   false,
			Age:       14,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/9%20%D0%91%D0%BE%D0%B3%D1%80%D0%B5%D0%B5%D0%B2%20%D0%9C%D0%B0%D0%BA%D1%81%D0%B8%D0%BC.jpeg",
		},
		10: {
			FirstName: "Анастасия",
			LastName:  "",
			Id:        10,
			Gift:      "Спортивную волейбольную форму и обувь",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/10%20%D0%90%D0%BD%D0%B0%D1%81%D1%82%D0%B0%D1%81%D0%B8%D1%8F.jpeg",
		},
		11: {
			FirstName: "Тамара",
			LastName:  "",
			Id:        11,
			Gift:      "Колонка с флешкой",
			IsTaken:   false,
			Age:       11,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/11%20%D0%A2%D0%B0%D0%BC%D0%B0%D1%80%D0%B0.jpeg",
		},
		12: {
			FirstName: "Никита",
			LastName:  "Гапеев",
			Id:        12,
			Gift:      "Конструктор-машина на радиоуправлении",
			IsTaken:   false,
			Age:       9,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/12%20%D0%93%D0%B0%D0%BF%D0%B5%D0%B5%D0%B2%20%D0%9D%D0%B8%D0%BA%D0%B8%D1%82%D0%B0.jpeg",
		},
		13: {
			FirstName: "Ренад",
			LastName:  "Гапеев",
			Id:        13,
			Gift:      "Машина на радиоуправлении Lego techtc",
			IsTaken:   false,
			Age:       12,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/13%20%D0%93%D0%B0%D0%BF%D0%B5%D0%B5%D0%B2%20%D0%A0%D0%B5%D0%BD%D0%B0%D0%B4.jpeg",
		},
		14: {
			FirstName: "Максим",
			LastName:  "Дмитриев",
			Id:        14,
			Gift:      "Беспроводную колонку и спортивный костюм",
			IsTaken:   false,
			Age:       16,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/14%20%D0%94%D0%BC%D0%B8%D1%82%D1%80%D0%B8%D0%B5%D0%B2%20%D0%9C%D0%B0%D0%BA%D1%81%D0%B8%D0%BC.jpeg",
		},
		15: {
			FirstName: "Светлана",
			LastName:  "",
			Id:        15,
			Gift:      "Беспроводной микрофон",
			IsTaken:   false,
			Age:       10,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/15%20%D0%A1%D0%B2%D0%B5%D1%82%D0%BB%D0%B0%D0%BD%D0%B0.jpeg",
		},
		16: {
			FirstName: "Даша",
			LastName:  "",
			Id:        16,
			Gift:      "Алмазная мозайка",
			IsTaken:   false,
			Age:       0,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/16%20%D0%94%D0%B0%D1%88%D0%B0.jpeg",
		},
		17: {
			FirstName: "Валентина",
			LastName:  "",
			Id:        17,
			Gift:      "Неоновая лента",
			IsTaken:   false,
			Age:       17,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/17%20%D0%92%D0%B0%D0%BB%D0%B5%D0%BD%D1%82%D0%B8%D0%BD%D0%B0.jpeg",
		},
		18: {
			FirstName: "Сергей",
			LastName:  "Белый",
			Id:        18,
			Gift:      "Внешний аккумулятор",
			IsTaken:   false,
			Age:       15,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/18%20%D0%91%D0%B5%D0%BB%D1%8B%D0%B8%CC%86%20%D0%A1%D0%B5%D1%80%D0%B3%D0%B5%D0%B8%CC%86.jpeg",
		},
		19: {
			FirstName: "Алена",
			LastName:  "Труфанова",
			Id:        19,
			Gift:      "Косметику (тушь, тональный крем, консилер, палитра теней, бальзам для губ, блестки, кисти, спонжи, косметичку)",
			IsTaken:   false,
			Age:       14,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/19%20%D0%A2%D1%80%D1%83%D1%84%D0%B0%D0%BD%D0%BE%D0%B2%D0%B0%20%D0%90%D0%BB%D0%B5%D0%BD%D0%B0.jpeg",
		},
		20: {
			FirstName: "Анжела",
			LastName:  "",
			Id:        20,
			Gift:      "Набор для рукоделия и создания украшений",
			IsTaken:   false,
			Age:       8,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/20%20%D0%90%D0%BD%D0%B6%D0%B5%D0%BB%D0%B0.jpeg",
		},
		21: {
			FirstName: "Оля",
			LastName:  "",
			Id:        21,
			Gift:      "Большую куклу",
			IsTaken:   false,
			Age:       5,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/21%20%D0%9E%D0%BB%D1%8F.jpeg",
		},
		22: {
			FirstName: "Люба",
			LastName:  "Ботова",
			Id:        22,
			Gift:      "Большой набор фломастеров и мягкую игрушку",
			IsTaken:   false,
			Age:       0,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/22%20%D0%91%D0%BE%D1%82%D0%BE%D0%B2%D0%B0%20%D0%9B%D1%8E%D0%B1%D0%B0.jpeg",
		},
		23: {
			FirstName: "Виктория",
			LastName:  "",
			Id:        23,
			Gift:      "Мульти стайлер 3 в 1",
			IsTaken:   false,
			Age:       14,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/23%20%D0%92%D0%B8%D0%BA%D1%82%D0%BE%D1%80%D0%B8%D1%8F.jpeg",
		},
		24: {
			FirstName: "Александра",
			LastName:  "",
			Id:        24,
			Gift:      "Фен, расческа, щипцы для укладки",
			IsTaken:   false,
			Age:       13,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/24%20%D0%90%D0%BB%D0%B5%D0%BA%D1%81%D0%B0%D0%BD%D0%B4%D1%80%D0%B0.jpeg",
		},
		25: {
			FirstName: "Вероника",
			LastName:  "",
			Id:        25,
			Gift:      "Набор кукол",
			IsTaken:   false,
			Age:       6,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/25%20%D0%92%D0%B5%D1%80%D0%BE%D0%BD%D0%B8%D0%BA%D0%B0.jpeg",
		},
		26: {
			FirstName: "Люба",
			LastName:  "Гапеева",
			Id:        26,
			Gift:      "Большая мягкая игрушка",
			IsTaken:   false,
			Age:       7,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/26%20%D0%9B%D1%8E%D0%B1%D0%B0.jpeg",
		},
		27: {
			FirstName: "Кристина",
			LastName:  "",
			Id:        27,
			Gift:      "Алмазная мозайка",
			IsTaken:   false,
			Age:       11,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/27%20%D0%9A%D1%80%D0%B8%D1%81%D1%82%D0%B8%D0%BD%D0%B0.jpeg",
		},
		28: {
			FirstName: "Полина",
			LastName:  "Черникова",
			Id:        28,
			Gift:      "Детскую коляску",
			IsTaken:   false,
			Age:       4,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/28%20%D0%A7%D0%B5%D1%80%D0%BD%D0%B8%D0%BA%D0%BE%D0%B2%D0%B0%20%D0%9F%D0%BE%D0%BB%D0%B8%D0%BD%D0%B0.jpeg",
		},
		29: {
			FirstName: "Снежана",
			LastName:  "",
			Id:        29,
			Gift:      "Алмазная мозайка",
			IsTaken:   false,
			Age:       9,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/29%20%D0%A1%D0%BD%D0%B5%D0%B6%D0%B0%D0%BD%D0%B0.jpeg",
		},
		30: {
			FirstName: "Николай",
			LastName:  "Яндуков",
			Id:        30,
			Gift:      "Большой сладкий сюрприз",
			IsTaken:   false,
			Age:       15,
			PhotoUrl:  "https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/children/30%20%D0%AF%D0%BD%D0%B4%D1%83%D0%BA%D0%BE%D0%B2%20%D0%9D%D0%B8%D0%BA%D0%BE%D0%BB%D0%B0%D0%B8%CC%86.jpeg",
		},
	}
	if err := InitConfig(); err != nil {
		log.Fatalf("config init error: %s", err.Error())
	}
	fmt.Println(viper.GetString("smtp.gmail"), viper.GetString("smtp.password"), viper.GetString("smtp.host"))
	auth := smtp.PlainAuth("", viper.GetString("smtp.gmail"), viper.GetString("smtp.password"), viper.GetString("smtp.host"))
	repo := repository.NewRepository(children)
	services := service.NewService(repo, auth, viper.GetString("smtp.gmail"), viper.GetString("smtp.host"), viper.GetString("smtp.port"))
	end := endpoint.NewEndpoint(services)
	handler := end.InitRoutes()
	server := &backend.Server{}
	go func() {
		if err := server.Run(viper.GetString("port"), handler); err != nil {
			log.Fatalf("server run error: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
