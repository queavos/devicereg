package main

import (
	"fmt"
)

func main() {
	uuid, err := GetSystemUUID()
	if err != nil {
		fmt.Println("Error obteniendo UUID:", err)
		return
	}

	mac, err := GetMACAddress()
	if err != nil {
		fmt.Println("Error obteniendo MAC:", err)
		return
	}

	token := GenerateToken(uuid, mac)
	err = EncryptAndSaveToken(token, "config.dat")
	if err != nil {
		fmt.Println("Error guardando token:", err)
		return
	}

	//fmt.Println("Token generado y guardado exitosamente.")
	// Leer y descifrar el token
	recoveredToken, err := LoadAndDecryptToken("config.dat")
	if err != nil {
		fmt.Println("Error leyendo token:", err)
		return
	}

	fmt.Println("Token recuperado:", recoveredToken)

	// Comparar con el original
	if recoveredToken == token {
		fmt.Println("✅ El token recuperado coincide con el original.")
	} else {
		fmt.Println("❌ El token recuperado NO coincide.")
	}
	// Iniciar servidor local
	StartLocalServer(recoveredToken)

	//fmt.Println("Servidor iniciado correctamente.")
}
