package core

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

// Server manages the embedded HTTP server for rendering markdown previews.
type Server struct {
	port    int
	mux     *http.ServeMux
	server  *http.Server
	html    string
	eventCh chan struct{}
	mu      sync.RWMutex
}

// NewServer creates and starts a new preview server on a random available port.
func NewServer() (*Server, error) {
	s := &Server{
		eventCh: make(chan struct{}, 1),
		mux:     http.NewServeMux(),
	}

	// Find an available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to find available port: %w", err)
	}
	s.port = listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	// Register handlers
	s.mux.HandleFunc("/", s.handleRoot)
	s.mux.HandleFunc("/preview", s.handlePreview)
	s.mux.HandleFunc("/events", s.handleEvents)

	s.server = &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", s.port),
		Handler: s.mux,
	}

	// Start server in background
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[md-viewer] server error: %v", err)
		}
	}()

	return s, nil
}

// Port returns the server's port number.
func (s *Server) Port() int {
	return s.port
}

// URL returns the base URL of the server.
func (s *Server) URL() string {
	return fmt.Sprintf("http://127.0.0.1:%d", s.port)
}

// PreviewURL returns the full URL to the preview page.
func (s *Server) PreviewURL() string {
	return fmt.Sprintf("http://127.0.0.1:%d/preview", s.port)
}

// SetHTML updates the rendered HTML content and triggers a reload event.
func (s *Server) SetHTML(html string) {
	s.mu.Lock()
	s.html = html
	s.mu.Unlock()

	// Send reload event (non-blocking)
	select {
	case s.eventCh <- struct{}{}:
	default:
	}
}

// Close shuts down the server.
func (s *Server) Close() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/preview", http.StatusFound)
}

func (s *Server) handlePreview(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	html := s.html
	s.mu.RUnlock()

	if html == "" {
		html = welcomePage()
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	fmt.Fprint(w, html)
}

func (s *Server) handleEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher.Flush()

	for {
		select {
		case <-s.eventCh:
			fmt.Fprintf(w, "data: reload\n\n")
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func welcomePage() string {
	return `<!DOCTYPE html>
<html lang="zh-TW">
<head>
  <meta charset="UTF-8">
  <title>md-viewer</title>
  <style>
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      max-width: 600px;
      margin: 80px auto;
      padding: 0 20px;
      color: #1f2328;
      background: #ffffff;
    }
    h1 { color: #0969da; }
    p { color: #656d76; line-height: 1.6; }
    code { background: #eff1f3; padding: 2px 6px; border-radius: 4px; font-size: 90%; }
    .tip {
      background: #fff8c5;
      padding: 12px 16px;
      border-radius: 6px;
      border-left: 4px solid #bf8700;
      margin: 20px 0;
    }
    .tip strong { color: #9a6700; }
    a { color: #0969da; }
    @media (prefers-color-scheme: dark) {
      body { background: #0d1117; color: #e6edf3; }
      h1 { color: #58a6ff; }
      p { color: #8b949e; }
      code { background: #30363d; }
      .tip { background: #3d2e00; border-left-color: #d29922; }
      .tip strong { color: #d29922; }
      a { color: #58a6ff; }
    }
  </style>
</head>
<body>
  <h1>📄 md-viewer</h1>
  <p>純 Markdown 預覽器 — 零編輯區，只給你看。</p>
  <p>請使用 <strong>⌘O</strong> 開啟檔案，或 <strong>⌘⇧O</strong> 開啟資料夾。</p>
  <div class="tip">💡 <strong>小技巧：</strong>支援 GFM 語法 — 表格、任務清單、刪除線、自動連結等。</div>
  <p style="color:#8b949e;font-size:13px;">Built with Go + goldmark + Fyne</p>
</body>
</html>`
}
