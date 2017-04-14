http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	c, err := conn.UpgradeFromHTTP(ctx, conn.DefaultSettings(), w, r)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	for m := range c.Stream() {
		if m.IsBinaryMessage() {
			messageDataReceived <- m.Data
		}
	}
})
