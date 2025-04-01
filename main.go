package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"password",
		"exercise2_db")

	a.Run(":8080")
}
