DOM (Document Object Model) is the data structure and content of an HTML document that looks like tree of objects:

- the entire page is the root (`<html>` element)
- inside are branches (`<head>`, `<body>`)
- those have leave (`<div>`, `<p>`, `<button>`) and so on

For example, this HTML

```html
<html>
    <body>
        <h1>Hello</h1>
        <button>Click me</button>
    </body>
</html>
```

has a DOM tree like this

```
Document
 └── html
      └── body
           ├── h1 ("Hello")
           └── button ("Click me")
```

Each element is a *node* in the tree.

Why the DOM matters

- browsers use DOM to display and update pages
- JavaScript uses DOM to
    - add/remove elements
    - change text or styles
    - respond to clicks and input

Example in JavaScript

```js
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>DOM Example</title>
</head>
<body>

  <h1>Hello</h1>
  <button>Click me</button>

  <script>
    // JavaScript interacting with the DOM
    document.querySelector("button").addEventListener("click", () => {
      document.querySelector("h1").textContent = "You clicked!";
    });
  </script>

</body>
</html>
```

This changes the DOM -> browser updates the view.