package reflex

import (
	"fmt"
	"reflect"
)

// Some mock type
type Book struct {
	CurrentPage int
	maxPages    int
}

// Go to the next page, unless we are at the
// end of the book, reset the page.
func (b *Book) NextPage() int {
	if b.CurrentPage == 0 {
		// Initialize to 0
		b.Reset()
	}
	if b.CurrentPage >= b.maxPages {
		b.Reset()
	} else {
		b.CurrentPage++
	}
	return b.CurrentPage
}

// Set the beginning of the book back to page 1
func (b *Book) Reset() {
	b.CurrentPage = 1
}

func CheckExportedMethodsAndFields() {
	var err error
	SomeBook := Book{CurrentPage: 1, maxPages: 10}
	PtrToSomeBook := &SomeBook

	testMethods := []string{"NextPage", "String", "Reset"}
	testFields := []string{"SomeField", "CurrentPage", "maxPages"}

	// Test SomeBook
	fmt.Println("Test `SomeBook := Book{CurrentPage:1, maxPages:10}`")
	fmt.Println("\tMethods")
	for _, v := range testMethods {
		err = ReflectStructMethod(SomeBook, v)
		fmt.Print("\t\t")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Method `%s` found\n", v)
		}
	}
	fmt.Println("\tFields")
	for _, v := range testFields {
		err = ReflectStructField(SomeBook, v)
		fmt.Print("\t\t")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Method `%s` found\n", v)
		}
	}

	// Test PtrToSomeBook
	fmt.Println("Test `PtrToSomeBook := &SomeBook`")
	fmt.Println("\tMethods")
	for _, v := range testMethods {
		err = ReflectStructMethod(PtrToSomeBook, v)
		// Some tabs
		fmt.Print("\t\t")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Method `%s` found\n", v)
		}
	}
	fmt.Println("\tFields")
	for _, v := range testFields {
		err = ReflectStructField(PtrToSomeBook, v)
		// Some tabs
		fmt.Print("\t\t")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Field `%s` found\n", v)
		}
	}
}

// Reflect if an interface is either a struct or a pointer to a struct
// and has the defined member method. If error is nil, it means
// the MethodName is accessible with reflect.
func ReflectStructMethod(Iface interface{}, MethodName string) error {
	ValueIface := reflect.ValueOf(Iface)

	// Check if the passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		// Create a new type of Iface, so we have a pointer to work with
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// Get the method by name
	Method := ValueIface.MethodByName(MethodName)
	if !Method.IsValid() {

		return fmt.Errorf("Couldn't find method `%s` in interface `%s`, is it Exported?", MethodName, ValueIface.Type())
	}
	return nil
}

// Reflect if an interface is either a struct or a pointer to a struct
// and has the defined member field, if error is nil, the given
// FieldName exists and is accessible with reflect.
func ReflectStructField(Iface interface{}, FieldName string) error {
	ValueIface := reflect.ValueOf(Iface)

	// Check if the passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		// Create a new type of Iface's Type, so we have a pointer to work with
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// 'dereference' with Elem() and get the field by name
	Field := ValueIface.Elem().FieldByName(FieldName)
	if !Field.IsValid() {
		return fmt.Errorf("Interface `%s` does not have the field `%s`", ValueIface.Type(), FieldName)
	}
	return nil
}
