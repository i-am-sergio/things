package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAddJSONSerialization(t *testing.T) {
	// Crea una instancia de tu modelo Add
	originalAdd := Add{
		ProductID: 123,
		Price:     45.99,
		Time:      123456789,
		Date:      time.Now(),
		UserID:    456,
		View:      789,
	}

	// Serializa la instancia del modelo a JSON
	jsonData, err := json.Marshal(originalAdd)
	if err != nil {
		t.Errorf("Error al serializar el modelo a JSON: %v", err)
	}

	// Deserializa el JSON nuevamente en una instancia del modelo
	var deserializedAdd Add
	if err := json.Unmarshal(jsonData, &deserializedAdd); err != nil {
		t.Errorf("Error al deserializar JSON en modelo: %v", err)
	}

	// Comparar los campos relevantes de los modelos original y deserializado
	if originalAdd.ProductID != deserializedAdd.ProductID ||
		originalAdd.Price != deserializedAdd.Price ||
		originalAdd.Time != deserializedAdd.Time ||
		originalAdd.UserID != deserializedAdd.UserID ||
		originalAdd.View != deserializedAdd.View {
		t.Error("El modelo deserializado no coincide con el modelo original")
	}
}
func TestAddValidation(t *testing.T) {
	// Crear una instancia de Add con datos válidos
	validAdd := Add{
		ProductID: 123,
		Price:     45.99,
		Time:      123456789,
		Date:      time.Now(),
		UserID:    456,
		View:      789,
	}

	// Prueba de validación para datos válidos
	if err := validAdd.Validate(); err != nil {
		t.Errorf("Se esperaba que los datos válidos pasaran la validación, pero se encontró un error: %v", err)
	}

	// Crear una instancia de Add con datos inválidos
	invalidAdd := Add{
		ProductID: 0,                           // ProductID no puede ser cero
		Price:     -10,                         // Precio no puede ser negativo
		Time:      -100,                        // Tiempo no puede ser negativo
		Date:      time.Now().AddDate(0, 0, 1), // Fecha en el futuro no es válida
		UserID:    0,                           // UserID no puede ser cero
		View:      -5,                          // Vistas no pueden ser negativas
	}

	// Prueba de validación para datos inválidos
	if err := invalidAdd.Validate(); err == nil {
		t.Error("Se esperaba un error de validación para datos inválidos, pero no se encontró ninguno")
	}
}
