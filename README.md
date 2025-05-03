# Go + HTMX CRUD Application 🛠️📷

This project is a modern CRUD (Create, Read, Update, Delete) web application built using **Go** for the backend and **HTMX** for frontend interactivity. It supports basic item management with image uploads and uses in-memory storage for simplicity.

## ✨ Features

- ⚙️ Server-side rendering with Go templates
- ⚡ Dynamic frontend with HTMX (no full-page reloads)
- 📤 File/image upload support
- 🌐 RESTful endpoints (GET, POST, DELETE)
- 🧭 Chi router for lightweight routing
- 🧠 In-memory data storage (for demonstration purposes)

---

## 📦 Tech Stack

- **Go** (Backend)
- **HTMX** (Frontend Interactivity)
- **Chi** (Router)
- **Go Templates** (HTML Rendering)
- **CORS** Middleware (Cross-Origin Support)

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/your-username/go-htmx-crud.git
cd go-htmx-crud
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Run the application
```bash
go run main.go
```

### 4. Visit in your browser
```
http://localhost:3000
```

---

## 📁 Project Structure

```
go-htmx-crud/
├── main.go                # Main application logic
├── templates/
│   ├── index.html         # Full page template
│   └── items-list.html    # Partial template for HTMX
└── uploads/               # Uploaded images are stored here
```

---

## 🧪 Example Use Cases

- Upload an item with name, description, and image
- View live-updating item list via HTMX
- Delete an item dynamically

---

## 🔐 Notes

- Uploaded files are saved in the `uploads/` directory.
- This app uses in-memory storage (no database). All data resets when the app restarts.
- In production, consider:
  - Switching to UUIDs for IDs
  - Adding persistent database storage (e.g., SQLite, Postgres)
  - Securing file uploads and handling duplicates

---

## 📸 Screenshot

> You can include a screenshot of the UI here for a better visual intro.

---

## 🧑‍💻 Author

Built with ❤️ by [Amir](https://github.com/your-username)  
BackBrace Channel — Tutorials | Tools | Security

---

## 📃 License

MIT License
