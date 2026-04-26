package loadpuchases

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Purchase struct {
	Material     string    `json:"material"`
	Count        int       `json:"count"`
	UnitPrice    float64   `json:"unit_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}

// Метод для вычисления общей стоимости
func (p Purchase) TotalPrice() float64 {
	return float64(p.Count) * p.UnitPrice
}

func main() {
	// 1. Читаем JSON файл
	jsonFile, err := os.Open("purchases.json")
	if err != nil {
		log.Fatalf("Ошибка открытия JSON файла: %v", err)
	}
	defer jsonFile.Close()

	var purchases []Purchase
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&purchases)
	if err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	fmt.Printf("Загружено %d записей из JSON\n", len(purchases))

	// 2. Подключаемся к базе данных
	// Замените параметры на свои
	connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatalf("БД не отвечает: %v", err)
	}
	fmt.Println("Подключение к БД установлено")

	// 3. Создаём таблицу (если не существует)
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS purchases (
		id SERIAL PRIMARY KEY,
		material VARCHAR(255) NOT NULL,
		count INT NOT NULL,
		unit_price DECIMAL(10,2) NOT NULL,
		total_price DECIMAL(10,2) NOT NULL,
		purchase_date TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("Таблица purchases готова")

	// 4. Вставляем данные
	insertQuery := `
	INSERT INTO purchases (material, count, unit_price, total_price, purchase_date)
	VALUES ($1, $2, $3, $4, $5)
	`

	successCount := 0
	for _, p := range purchases {
		totalPrice := p.TotalPrice()
		_, err := db.Exec(insertQuery, p.Material, p.Count, p.UnitPrice, totalPrice, p.PurchaseDate)
		if err != nil {
			log.Printf("Ошибка вставки записи '%s': %v", p.Material, err)
		} else {
			successCount++
		}
	}

	fmt.Printf("Готово! Успешно загружено %d из %d записей\n", successCount, len(purchases))

	// 5. Показываем статистику
	var totalSum float64
	err = db.Get(&totalSum, "SELECT COALESCE(SUM(total_price), 0) FROM purchases")
	if err == nil {
		fmt.Printf("Общая сумма всех покупок: %.2f руб.\n", totalSum)
	}
}
