package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

/*Структура хранящая данные о товаре в базе*/
type Good struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Producer string `json:"producer"`
	Price    int    `json:"price"`
	Count    int    `json:"count"`
}


type Config struct {
	DefaultFile string `json:"default_file"`
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	args := os.Args
	argLength := len(os.Args)
	config := readConfigs()
	goods := readGoods(config.DefaultFile)

	if argLength > 1 {
		command := args[1]

		//Определение нового файла конфигурации товаров базы
		if strings.Compare(command, "-f") == 0 {
			config.DefaultFile = args[2]
			saveConfigs(config)
		}

		//Ввод нового товара в база
		if strings.Compare(command, "add") == 0 {
			var err error
			good := Good{}
			good.Id = nextID(goods)



			fmt.Print("Название товара: ")
			sc.Scan()
			good.Name = sc.Text()


			fmt.Print("Производитель: ")
			sc.Scan()
			good.Producer = sc.Text()

			for {
				fmt.Print("Кол-во товаров: ")
				sc.Scan()
				good.Count, err = strconv.Atoi(sc.Text())
				if err == nil{
					break
				}
			}

			for {
				fmt.Print("Цена товара: ")
				sc.Scan()
				good.Price, err = strconv.Atoi(sc.Text())
				if err == nil{
					break
				}
			}

			goods = append(goods, good)
			saveGoods(goods, config.DefaultFile)
			fmt.Println("appended good: ", good)
		}

		//Редактирование товаров в базе
		if strings.Compare(command, "edit") == 0 {
			index := findIndexByID(goods, args[2])
			fmt.Println("Редактировать поле:\n1. Название\n2. Производитель\n3. Цена\n4. Количество")
			var choose int
			_, _ = fmt.Scanf("%d\n", &choose)
			if choose == 1 {
				fmt.Print("Название товара: ")
				sc.Scan()
				goods[index].Name = sc.Text()
			} else if choose == 2 {
				fmt.Print("Производитель: ")
				sc.Scan()
				goods[index].Producer = sc.Text()
			} else if choose == 3 {
				for {
					fmt.Print("Цена: ")
					sc.Scan()
					num, err := strconv.Atoi(sc.Text())
					goods[index].Price = num
					if err == nil {
						break
					}
				}
			} else if choose == 4 {
				for {
					fmt.Print("Кол-во товаров: ")
					sc.Scan()
					num, err := strconv.Atoi(sc.Text())
					goods[index].Count = num
					if err == nil {
						break
					}
				}
			}
			saveGoods(goods, config.DefaultFile)
		}

		//Удаление товаров в базе
		if strings.Compare(command, "del") == 0 {
			switch len(args) {
			case 3:
				index := findIndexByID(goods, args[2])
				if index != -1 {
					goods = append(goods[:index], goods[index+1:]...)
					saveGoods(goods, config.DefaultFile)
				} else {
					fmt.Println("Неверно указанный ID")
				}
				break
			case 4:
				var rng [2]int
				rnga:= strings.Split(args[3], "-")
				rng[0], _ = strconv.Atoi(rnga[0])
				rng[1], _ = strconv.Atoi(rnga[1])
				if (rng[0] < rng[1])&&(rng[1] <= len(goods)){
					for i := rng[0] - 1; i < rng[1] - 1; i++ {
						goods = append(goods[:i], goods[i+1:]...)
					}
					saveGoods(goods, config.DefaultFile)
				} else {
					fmt.Println("Неверно указан диапазон")
				}
				break
			}
		}

		//Вывод товаров базы на экран
		if strings.Compare(command, "view") == 0 {
			switch len(args) {
			case 3:
				index := findIndexByID(goods, args[2])
				if index != -1 {
					good := goods[index]
					fmt.Println("-----------------------------------------------------------------")
					fmt.Printf("|%-20s|%-20s|%-8s|%-12s|\n", "Название", "   Производитель", "  Цена", " Количество ")
					fmt.Println("+--------------------+--------------------+--------+------------+")
					fmt.Printf("|%-20s|%-20s|%8d|%12d|\n", good.Name, good.Producer, good.Price, good.Count)
					fmt.Println("-----------------------------------------------------------------")
				} else {
					fmt.Println("Неверно указанный ID")
				}
				break
			case 4:
				var rng [2]int
				rngs := strings.Split(args[3], "-")
				rng[0], _ = strconv.Atoi(rngs[0])
				rng[1], _ = strconv.Atoi(rngs[1])
				fmt.Println("----------------------------------------------------------------------")
				fmt.Printf("|%-4s|%-20s|%-20s|%-8s|%-12s|\n", "Название", "   Производитель", "  Цена", " Количество ")
				if (rng[0] < rng[1])&&(rng[1] <= len(goods)){
					for i := rng[0] - 1; i < rng[1] - 1; i++ {
						fmt.Println("|--------------------------------------------------------------------|")
						fmt.Printf("|%4d|%-20s|%-20s|%8d|%12d|\n", goods[i].Id, goods[i].Name, goods[i].Producer, goods[i].Price, goods[i].Count)
					}
					fmt.Println("----------------------------------------------------------------------")
				} else {
					fmt.Println("Неверно указан диапазон")
				}
				break
			}

		}
	}

}


//Замена файла конфигурации на введенный пользователем
func saveConfigs(config *Config) {
	goodsJson, err := json.Marshal(config)
	if err != nil {
		log.Fatal("Cannot encode to JSON", err)
	}

	err = ioutil.WriteFile("config.json", goodsJson, 0777)
	if err != nil {
		log.Fatal("Cannot write data to file", err)
	}
}

//Считывание файла конфигурации в программу
func readConfigs() *Config {
	bytes, _ := ioutil.ReadFile("config.json")
	config := Config{}
	_ = json.Unmarshal(bytes, &config)
	return &config
}

//Сохранение базы в файл конфигурации
func saveGoods(goods []Good, filename string) {
	goodsJson, err := json.Marshal(goods)
	if err != nil {
		log.Fatal("Cannot encode to JSON", err)
	}

	err = ioutil.WriteFile(filename, goodsJson, 0777)
	if err != nil {
		log.Fatal("Cannot write data to file", err)
	}
}
//Считывание данных из файла конфигурации в программу
func readGoods(filename string) []Good {
	bytes, _ := ioutil.ReadFile(filename)
	var products []Good
	_ = json.Unmarshal(bytes, &products)
	return products
}

//Генерация нового ID
func nextID(goods []Good) int {
	maxID := -1
	for i := 0; i < len(goods); i++ {
		if maxID < goods[i].Id {
			maxID = goods[i].Id
		}
	}
	maxID++
	return maxID
}

//Нахождение индекса товара по ID/
func findIndexByID(goods []Good, id string) int {
	id_, _ := strconv.Atoi(id)
	for i := 0; i < len(goods); i++ {
		if goods[i].Id == id_ {
			return i
		}
	}
	return -1
}
