<template>
  <div class="container">
    <h1>Todo List</h1>
    <div class="input-container">
      <input v-model="newTodo" type="text" placeholder="Enter a new task" id="todo-input" />
      <button @click="createTodo" id="add-btn">Add Todo</button>
    </div>

    <!-- Edit Form -->
    <div v-if="showEditForm" id="edit-form">
      <input v-model="editForm.title" type="text" />
      <label for="completed-checkbox">Completed</label>
      <input v-model="editForm.completed" type="checkbox" id="completed-checkbox" />
      <button @click="saveEdit">Save</button>
      <button @click="cancelEdit">Cancel</button>
    </div>

    <!-- Todo Table -->
    <table id="todo-table">
      <thead>
        <tr>
          <th>Title</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" :key="todo.id">
          <td>{{ todo.title }}</td>
          <td>{{ todo.completed ? "Completed" : "Not Completed" }}</td>
          <td>
            <button @click="toggleComplete(todo.id, todo.completed)">
              Toggle
            </button>
            <button @click="deleteTodo(todo.id)">Delete</button>
            <button @click="openEditForm(todo.id, todo.title, todo.completed)">
              Edit
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      apiUrl: "http://localhost:8081/todos", // Backend API URL
      todos: [], // List of todos
      newTodo: "", // New todo title
      showEditForm: false, // Flag for edit form visibility
      editForm: {
        id: null,
        title: "",
        completed: false,
      },
    };
  },
  methods: {
    // Fetch todos from the API
    fetchTodos() {
      fetch(this.apiUrl)
        .then((response) => response.json())
        .then((data) => {
          this.todos = data;
        })
        .catch((error) => console.error("Error fetching todos:", error));
    },

    // Create a new todo
    createTodo() {
      if (!this.newTodo.trim()) return; // Prevent empty tasks
      const newTodo = { title: this.newTodo, completed: false };
      fetch(this.apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newTodo)
      })
        .then((response) => response.json())
        .then(() => {
          this.newTodo = ""; // Clear input field
          this.fetchTodos(); // Refresh todo list
        })
        .catch((error) => console.error("Error creating todo:", error));
    },

    // Open the edit form with pre-filled data
    openEditForm(id, title, completed) {
      this.editForm = { id, title, completed };
      this.showEditForm = true;
    },

    // Save edited todo
    saveEdit() {
      const { id, title, completed } = this.editForm;
      fetch(`${this.apiUrl}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, completed }),
      })
        .then(() => {
          this.showEditForm = false; // Hide edit form
          this.fetchTodos(); // Refresh todo list
        })
        .catch((error) => console.error("Error saving todo:", error));
    },

    // Cancel edit
    cancelEdit() {
      this.showEditForm = false; // Hide edit form
    },

    // Toggle completion status
    toggleComplete(id, currentCompleted) {
      fetch(`${this.apiUrl}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ completed: !currentCompleted }),
      })
        .then(() => this.fetchTodos())
        .catch((error) => console.error("Error toggling todo:", error));
    },

    // Delete a todo
    deleteTodo(id) {
      fetch(`${this.apiUrl}/${id}`, { method: "DELETE" })
        .then(() => this.fetchTodos())
        .catch((error) => console.error("Error deleting todo:", error));
    },
  },
  created() {
    const isLoggedIn = localStorage.getItem('isLoggedIn');
    if (!isLoggedIn) {
      this.$router.push('/login'); // Redirect to login if not logged in
    } else {
      this.fetchTodos(); // Fetch todo list after login
    }
  },
  mounted() {
    // Fetch todos when component mounts
    this.fetchTodos();
  },
};
</script>

<style scoped>
/* Styles similar to the original */
.container {
  max-width: 800px;
  margin: 50px auto;
  background-color: white;
  padding: 40px;
  border-radius: 10px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  color: #4a90e2;
  margin-bottom: 30px;
  font-size: 30px;
}

.input-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

input {
  width: 75%;
  padding: 12px;
  border-radius: 5px;
  border: 1px solid #ddd;
  font-size: 16px;
  background-color: #f9f9f9;
  transition: border-color 0.3s ease;
}

input:focus {
  border-color: #4caf50;
  outline: none;
}

button {
  padding: 12px 20px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #45a049;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 30px;
}

table th,
table td {
  padding: 12px;
  text-align: center;
  border-bottom: 1px solid #ddd;
}

table th {
  background-color: #4caf50;
  color: white;
  font-size: 18px;
}

table td {
  font-size: 16px;
}

table td button {
  padding: 8px 15px;
  margin: 2px;
  background-color: #f1f1f1;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

table td button:hover {
  background-color: #ddd;
}
</style>
