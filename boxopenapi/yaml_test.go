package boxopenapi

import (
	"testing"
)

func TestConvertToYAML(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name: "Caso básico",
			input: map[string]interface{}{
				"nombre": "Juan",
				"edad":   30,
				"activo": true,
			},
			expected: "activo: true\nedad: 30\nnombre: Juan\n",
		},
		{
			name: "Lista simple",
			input: map[string]interface{}{
				"habilidades": []interface{}{"programación", "dibujo"},
			},
			expected: "habilidades:\n  - programación\n  - dibujo\n",
		},
		{
			name: "Mapa anidado",
			input: map[string]interface{}{
				"usuario": map[string]interface{}{
					"nombre": "Juan",
					"edad":   30,
				},
			},
			expected: "usuario:\n  edad: 30\n  nombre: Juan\n",
		},
		{
			name: "Valores nulos",
			input: map[string]interface{}{
				"nombre": nil,
				"edad":   nil,
			},
			expected: "edad: null\nnombre: null\n",
		},
		{
			name: "List of objects",
			input: map[string]interface{}{
				"people": []map[string]interface{}{
					{
						"name":   "Juan",
						"age":    30,
						"active": true,
					},
					{
						"name":   "Emma",
						"age":    65,
						"active": false,
					},
				},
			},
			expected: "people:\n- name: Juan\n  age: 30\n  active: true\n- name: Juan\n  age: 30\n  active: true\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToYAML(tt.input)
			if result != tt.expected {
				t.Errorf("convertToYAML() =\n %v, want\n %v", result, tt.expected)
			}
		})
	}
}
