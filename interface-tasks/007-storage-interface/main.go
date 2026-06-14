package main

import "fmt"

type Storage interface {
	Save(data string)
	Load() string
}

type MemoryStorage struct {
	Data string
}

type FileStorage struct {
	FileName string
	Data     string
}

type DatabaseStorage struct {
	Table string
	Data  string
}

func (m *MemoryStorage) Save(data string) {
	m.Data = data
	fmt.Println("Data saved to memory:")
	
}

func (m *MemoryStorage) Load() string {
	return m.Data
} 

func (f *FileStorage) Save(data string) {
	f.Data = data
	fmt.Println("Data saved to file")

}

func (f *FileStorage) Load() string {
	return f.Data
}

func (d *DatabaseStorage) Save(data string) {
	d.Data = data
	fmt.Println("Data saved to database:")

}

func (d *DatabaseStorage) Load() string {
	return d.Data
}

func SaveData(s Storage, data string) {
	s.Save(data)
}

func LoadData(s Storage) {
	fmt.Println("Loaded data:", s.Load())
}


func main() {
	memory := MemoryStorage {}
	file :=FileStorage{FileName: "123.txt"}
	database := DatabaseStorage{Table: "salary"}
	SaveData(&memory,"balance 100")
	LoadData(&memory)
	SaveData(&file,"balance 120")
	LoadData(&file)
	SaveData(&database,"balance 152")
	LoadData(&database)
}