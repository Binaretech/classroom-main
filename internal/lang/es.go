package lang

func setUpEs() {
	es, _ := translator.GetTranslator("es")
	// ----- MESSAGES --------
	es.Add("updated user", "Usuario actualizado con éxito", true)
	// ----- ERRORS ----------
	es.Add("not found", "No encontrado", true)
	es.Add("internal error", "Ha ocurrido un error en el servidor.", true)
}
