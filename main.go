package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// estructura de contactos

const fileName = "contacts.json"

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactManager struct {
	Contacts []Contact
}

// cargar contactos desde un archivo JSON
func (cm *ContactManager) loadContactsFromFile() error {

	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&cm.Contacts); err != nil {
		return err
	}

	return nil

}

// guardar contactos en un archivo JSON

func (cm *ContactManager) saveContactsToFile() error {
	reader := bufio.NewReader(os.Stdin)

	//trae el nombre del archivo desde la constante fileName declarada de forma global en el archivo main.go
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err = encoder.Encode(cm.Contacts); err != nil {
		return err
	}

	//inicio
	var c Contact

	fmt.Println("==== Agregar un nuevo contacto ====")

	fmt.Println("Nombre del contacto")
	c.Name, _ = reader.ReadString('\n')

	fmt.Println("Correo electrónico")
	c.Email, _ = reader.ReadString('\n')

	fmt.Println("Telefono")
	c.Phone, _ = reader.ReadString('\n')

	cm.Contacts = append(cm.Contacts, c)

	if err := cm.saveContactsToFile(); err != nil {
		fmt.Println("Error al guardar los contactos: ", err)
	}

	fmt.Println("---Contacto agregado correctamente")

	//fin

	return nil

}

func (cm *ContactManager) editContactsFromFile() {

	fmt.Println("===== Editar un contacto =====")

	//Primeramente se valida si el cm contine datos
	if len(cm.Contacts) == 0 {
		fmt.Println("No hay contactos para editar")
		return
	}

	// en caso que el tamaño del cm sea diferente de cero entonces vamos a mostrar la lista de contactos

	cm.showAllContactsFromFile()

	// ahora debemos seleccionar el contacto que se debe editar

	var op int

	fmt.Println("Digite el numero del contacto que desea editar (0 para cancelar)")

	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Println("Error al leer la opcion: ", err)
	}

	// si la opcion es 0 o la opcion es mayor del tamaño del cm se cancela
	if op == 0 || op > len(cm.Contacts) {
		fmt.Println("Edicion cancelada.")
		return
	}

	contact := &cm.Contacts[op-1]

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Nuevo nombre del contacto: ")
	newName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el nombre: ", err)
	}
	contact.Name = newName

	fmt.Println("Nuevo email del contacto: ")
	newEmail, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el email: ", err)
	}
	contact.Email = newEmail

	fmt.Println("Nuevo telefono del contacto: ")
	newPhone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el telefono: ", err)
	}
	contact.Phone = newPhone

	fmt.Println("---Contacto editado correctamente")

}

func (cm *ContactManager) showAllContactsFromFile() {
	fmt.Println("==== Mostrar todos los contactos ====")

	for index, contact := range cm.Contacts {

		fmt.Print(index+1, "Nombre de contacto: ", contact.Name, "\n")
		fmt.Print("Email de contacto:", contact.Email, "\n")
		fmt.Print("Telefono de contacto:", contact.Phone, "\n")

	}
}

func (cm *ContactManager) deleteContactsFromFile() {

	fmt.Println("==== Eliminar un contacto =====")

	cm.showAllContactsFromFile()

	fmt.Println("Digite el numero del contacto que desea eliminar (0 para cancelar)")

	var op int
	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Println("Error al leer la opcion: ", err)
	}

	// si la opcion es 0 o la opcion es mayor del tamaño del cm se cancela
	if op == 0 || op > len(cm.Contacts) {
		fmt.Println("Eliminacion cancelada.")
	}

	contact := &cm.Contacts[op-1]

	fmt.Println("Esta seguro que desea eliminar el contacto: ", contact.Name, "(s/n)")

	for {

		var confirmacion string

		if _, err := fmt.Scanln(&confirmacion); err != nil {
			fmt.Println("Error al leer la opcion: ", err)
		}

		if confirmacion != "s" {
			fmt.Println("Eliminacion cancelada")
		}

		if confirmacion == "s" {
			cm.Contacts = append(cm.Contacts[:op-1], cm.Contacts[op:]...)
			fmt.Println("---Contacto eliminado correctamente")
			break
		}
	}

}

func main() {

	var contacts []Contact

	contactManager := ContactManager{Contacts: contacts}

	if err := contactManager.loadContactsFromFile(); err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
		return
	}

	for {

		var op int

		fmt.Println("===== Gestion de contact =====")
		fmt.Println("1. Agregar un contacto")
		fmt.Println("2. Mostrar todos los contactos")
		fmt.Println("3. Editar un contacto")
		fmt.Println("4. Eliminar un contacto")
		fmt.Println("5. Salir")
		fmt.Println("Elige una opcion")

		if _, err := fmt.Scanln(&op); err != nil {
			fmt.Println("Error al leer la opcion: ", err)
		}

		switch op {
		case 1:
			contactManager.saveContactsToFile()

		case 2:
			contactManager.showAllContactsFromFile()

		case 3:
			contactManager.editContactsFromFile()

		case 4:
			contactManager.deleteContactsFromFile()

		case 5:

			fmt.Println("Fin del programa")

			return

		default:
			fmt.Println("Opcion incorrecta")
			continue

		}

	}
}
