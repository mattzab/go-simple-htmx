# Go Simple HTMX Demo

A simple, elegant demonstration of building modern web applications using HTMX, Go, and Bulma CSS. This project follows the KISS (Keep It Simple, Stupid) principle while showcasing powerful features of modern web development.

![Go Simple HTMX Demo](static/demo.png)

## 🌟 Features

- **HTMX Integration**: Dynamic content loading without writing JavaScript
- **Go Backend**: Fast, efficient server-side processing
- **Bulma CSS**: Clean, modern styling without JavaScript dependencies
- **Interactive Examples**: Real-world demonstrations of common web patterns
- **Zero JavaScript**: (Almost) No JavaScript required for core functionality

## 🚀 Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-simple-htmx.git
   cd go-simple-htmx
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and visit:
   ```
   http://localhost:8080
   ```

## 🛠️ Project Structure

```
go-simple-htmx/
├── main.go           # Go server implementation
├── static/           # Static assets
│   ├── htmx.min.js   # HTMX library
│   └── bulma.min.css # Bulma CSS framework
└── templates/        # HTML templates
    ├── index.html    # Main page template
    └── snippets.html # Reusable HTML snippets
```

## 📚 Examples

### 1. Dynamic Content Loading

```html
<button class="button is-info" 
        hx-get="/html/bulmainfo" 
        hx-target="#bulma-content" 
        hx-swap="innerHTML">
    What is Bulma CSS?
</button>
<div id="bulma-content"></div>
```

### 2. Server-Side Counter

```html
<div id="counter-container" 
     hx-get="/html/counter" 
     hx-trigger="load">
</div>
```

### 3. Form Submission

```html
<form hx-post="/html/submitform" 
      hx-target="#form-result" 
      hx-swap="innerHTML">
    <!-- form fields -->
</form>
```

## 🔧 How It Works

1. **HTMX**: Handles all dynamic interactions through HTML attributes
2. **Go**: Processes requests and renders HTML templates
3. **Bulma**: Provides beautiful, responsive styling

## 🎯 Key Concepts

### HTMX Attributes
- `hx-get`: Makes GET requests
- `hx-post`: Makes POST requests
- `hx-target`: Specifies where to put the response
- `hx-swap`: Defines how to insert the response
- `hx-trigger`: Specifies when to make the request

### Go Templates
- Uses Go's built-in template system
- Supports template inheritance
- Handles dynamic content rendering

### Bulma CSS
- Mobile-first responsive design
- Modern component library
- No JavaScript dependencies

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [HTMX](https://htmx.org/) - For making web development simpler
- [Bulma CSS](https://bulma.io/) - For beautiful, responsive styling
- [Go](https://golang.org/) - For fast, efficient server-side processing

## 📚 Resources

- [HTMX Documentation](https://htmx.org/docs/)
- [Bulma Documentation](https://bulma.io/documentation/)
- [Go Web Development](https://golang.org/doc/articles/wiki/) 