// // package main

// // import "fmt"

// // type Numbers struct {
// // 	Values []int
// // }

// // func (n *Numbers) Max() (int, error) {
// // 	if len(n.Values) == 0 {
// // 		return 0, fmt.Errorf("slice is empty ")
// // 	}
// // 	max := n.Values[0]
// // 	for _, v := range n.Values {
// // 		if v > max {
// // 			max = v
// // 		}
// // 	}
// // 	return max, nil
// // }

// // func main() {
// // 	nums := Numbers{Values: []int{3, 5, 3, 5}}
// // 	max, err := nums.Max()
// // 	if err != nil {
// // 		fmt.Println(err)

// // 	} else {
// // 		fmt.Println("Max: ", max)
// // 	}
// // }

// // //////////////////////////////////////////////////////////////////
// // package main

// // import (
// // 	"fmt"
// // )

// // type Numbers struct {
// // 	Values []int
// // }

// // func (n *Numbers) FilterEven() []int {
// // 	var evens []int
// // 	for _, v := range n.Values {
// // 		if v%2 == 0 {
// // 			evens = append(evens, v)
// // 		}
// // 	}
// // 	return evens
// // }

// // func main() {
// // 	nums := Numbers{Values: []int{1, 2, 3, 4, 5}}
// // 	evens := nums.FilterEven()
// // 	fmt.Println("Evens Numbers: ", evens)
// // }
// ////////////////////////////////////////////////////////////////////

// // package main

// // import "fmt"

// // type Product struct {
// // 	Name  string
// // 	Price float64
// // }

// // type Inventory struct {
// // 	Products []Product
// // }

// // // Increases price of all products by a given percentage
// // func (inv *Inventory) UpdatePrices(percent float64) {
// // 	for i := range inv.Products {
// // 		inv.Products[i].Price += inv.Products[i].Price * (percent / 100)
// // 	}
// // }

// // func main() {
// // 	inv := Inventory{
// // 		Products: []Product{
// // 			{"Laptop", 1000},
// // 			{"Phone", 500},
// // 		},
// // 	}

// // 	inv.UpdatePrices(10) // Increase all prices by 10%
// // 	fmt.Println(inv.Products)
// // 	fmt.Println(inv.Products)

// // 	// Output: [{Laptop 1100} {Phone 550}]
// // }
// //////////////////////////////////////////////////////////

// // package main

// // import "fmt"

// // type Product struct {
// // 	Name  string
// // 	Price float64
// // }

// // type Inventory struct {
// // 	Products []Product
// // }

// // // Searches for a product by name and returns a pointer to it
// // func (inv *Inventory) FindProduct(name string) *Product {
// // 	for i := range inv.Products {
// // 		if inv.Products[i].Name == name {
// // 			return &inv.Products[i] // Return a pointer to the found product
// // 		}
// // 	}
// // 	return nil // Return nil if not found
// // }

// // func main() {
// // 	inv := Inventory{
// // 		Products: []Product{
// // 			{"Laptop", 1000},
// // 			{"Phone", 500},
// // 		},
// // 	}

// //		product := inv.FindProduct("Phone")
// //		if product != nil {
// //			fmt.Println("Found:", product.Name, "Price:", product.Price) // Output: Found: Phone Price: 500
// //		} else {
// //			fmt.Println("Product not found")
// //		}
// //	}
// //
// // ///////////////////////////////////////////////////////////////

// // package main

// // import "fmt"

// // type Product struct {
// // 	Name  string
// // 	Price float64
// // }

// // type Inventory struct {
// // 	Products []Product
// // }

// // // Merges another inventory into the current one and returns a new Inventory
// // func (inv *Inventory) Merge(other Inventory) Inventory {
// // 	merged := Inventory{}
// // 	merged.Products = append(inv.Products, other.Products...)
// // 	return merged
// // }

// // func main() {
// // 	inv1 := Inventory{
// // 		Products: []Product{
// // 			{"Laptop", 1000},
// // 			{"Phone", 500},
// // 		},
// // 	}

// // 	inv2 := Inventory{
// // 		Products: []Product{
// // 			{"Tablet", 300},
// // 			{"Monitor", 200},
// // 		},
// // 	}

// // 	mergedInv := inv1.Merge(inv2)
// // 	fmt.Println(mergedInv.Products)
// // 	// Output: [{Laptop 1000} {Phone 500} {Tablet 300} {Monitor 200}]
// // }
// ////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// type Score struct {
// 	StudentScore map[string]int
// }

// func (s *Score) AboveThreshould(t int) []string {
// 	var topStudent []string

// 	for student, score := range s.StudentScore {
// 		if score > t {
// 			topStudent = append(topStudent, student)
// 		}
// 	}
// 	return topStudent
// }

// func main() {
// 	class := Score{StudentScore: map[string]int{
// 		"Alice": 653, "bob": 5353,
// 	}}

//		topscore := class.AboveThreshould(33)
//		fmt.Println("Student score above: ", topscore)
//	}

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := []byte("$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6GzGJSWG")
	err := bcrypt.CompareHashAndPassword(hash, []byte("my plain text password"))
	fmt.Println(err)
}
