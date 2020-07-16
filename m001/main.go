package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var initial_len int


type File struct {
	Name string
	Path string
	Size int64
}

func AllFilesInFolder(path []string, all_files map[string]File) map[string]File {
	files, _ := ioutil.ReadDir(ArrayToPath(path))

	for _, file := range files{

		if file.IsDir() {
			path = append(path, file.Name())

			AllFilesInFolder(path, all_files)

			path = append(path[:len(path)- 1])

		}

		var new_file File
		new_file.Size = file.Size()
		new_file.Path = ArrayToPath(path[initial_len - 1:])
		new_file.Name = file.Name()
		all_files[file.Name()] = new_file

	}
	return all_files
}

//Перевод массива с путем в нужный формат
func ArrayToPath(array []string) string {

	var path string
	for i := 0; i < len(array); i++{
		path += "/" + array[i]
	}

	return path
}

func main() {

	//Словарь для хранения файлов
	all_files_map := make(map[string]File)

	var input string
	fmt.Scan(&input)

	path := make([]string, 0)

	//Делаем из пути массив строк
	//для того чтобы добавлять и убирать папки
	for i := 1; i < len(strings.Split(input, "/")); i++{
		path = append(path, string(strings.Split(input, "/")[i]))
	}
	initial_len = len(path)

	AllFilesInFolder(path, all_files_map)

	//Делаем из словаря слайс
	//Для того чтобы отсортировать по Size
	all_files_slice := make([]File, 0, len(all_files_map))

	for index, _ := range all_files_map{
		all_files_slice = append(all_files_slice, all_files_map[index])
	}

	//Сортировка слайса
	for i := 1; i < len(all_files_slice); i++{
		temp := all_files_slice[i]
		j := i
		for j > 0 && all_files_slice[j - 1].Size > temp.Size{
			all_files_slice[j] = all_files_slice[j - 1]
			j--
		}
		all_files_slice[j] = temp
	}

	//Вывод отсортированного слайса
	for i := 0; i < len(all_files_slice); i++{
		fmt.Println(all_files_slice[i].Path, all_files_slice[i].Name, all_files_slice[i].Size)
	}

}
