package router

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn CustomHandler) {
	m.router.Handle("DELETE", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn CustomHandler) {
	m.router.Handle("GET", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn CustomHandler) {
	m.router.Handle("HEAD", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn CustomHandler) {
	m.router.Handle("OPTIONS", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn CustomHandler) {
	m.router.Handle("PATCH", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn CustomHandler) {
	m.router.Handle("POST", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn CustomHandler) {
	m.router.Handle("PUT", path, Handler{
		CustomHandler:   fn,
		CustomServeHTTP: m.CustomServeHTTP,
	})
}
