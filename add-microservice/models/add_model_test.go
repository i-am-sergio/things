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
