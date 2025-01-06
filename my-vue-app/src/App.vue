<template>
  <div class="container">
    <h1>Todo List</h1>

    <!-- Form Thêm Todo -->
    <div class="todo-input">
      <input v-model="newTodo" id="todo-input" placeholder="Enter a new task" />
      <button @click="createTodo" id="add-btn">Add Todo</button>
    </div>

    <!-- Form chỉnh sửa Todo -->
    <div v-if="editingTodoId" class="edit-form">
      <h2>Edit Todo</h2>
      <div class="form-group">
        <label for="edit-title">Title</label>
        <input v-model="editTitle" id="edit-title" type="text" placeholder="Edit todo title" />
      </div>
      <div class="form-group checkbox-group">
        
        <div class="checkbox-container">
          <input v-model="editCompleted" type="checkbox" id="completed-checkbox" class="checkbox-input" />
          <span class="status-text">{{ editCompleted ? 'Completed' : 'Not Completed' }}</span>
        </div>
      </div>
      <div class="button-group">
        <button @click="saveEdit" class="save-btn">Save</button>
        <button @click="cancelEdit" class="cancel-btn">Cancel</button>
      </div>
    </div>

    <!-- Danh sách Todos -->
    <table>
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
          <td>{{ todo.completed ? 'Completed' : 'Not Completed' }}</td>
          <td>
            <button @click="toggleComplete(todo.id, todo.completed)">Toggle</button>
            <button @click="deleteTodo(todo.id)">Delete</button>
            <button @click="showEditForm(todo)">Edit</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Thông báo lỗi -->
    <div v-if="error" class="error">
      <p>{{ error }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      todos: [],
      newTodo: '',
      editingTodoId: null,
      editTitle: '',
      editCompleted: false,
      error: null,
    };
  },
  methods: {
    getTodos() {
      axios
        .get('http://localhost:8081/todos')
        .then((response) => {
          this.todos = response.data;
          this.error = null;
        })
        .catch((error) => {
          this.error = 'Có lỗi khi lấy dữ liệu todos!';
          console.error('Có lỗi khi lấy dữ liệu todos:', error);
        });
    },
    createTodo() {
      if (this.newTodo.trim() === '') return;

      const todo = {
        title: this.newTodo,
        completed: false,
      };

      axios
        .post('http://localhost:8081/todos', todo)
        .then((response) => {
          this.todos.push(response.data);
          this.newTodo = '';
          this.error = null;
        })
        .catch((error) => {
          this.error = 'Có lỗi khi tạo todo!';
          console.error('Có lỗi khi tạo todo:', error);
        });
    },
    showEditForm(todo) {
      this.editingTodoId = todo.id;
      this.editTitle = todo.title;
      this.editCompleted = todo.completed;
    },
    saveEdit() {
      const updatedTodo = {
        title: this.editTitle,
        completed: this.editCompleted,
      };

      axios
        .put(`http://localhost:8081/todos/${this.editingTodoId}`, updatedTodo)
        .then((response) => {
          const index = this.todos.findIndex((todo) => todo.id === this.editingTodoId);
          this.todos[index] = response.data;
          this.editingTodoId = null;
          this.error = null;
        })
        .catch((error) => {
          this.error = 'Có lỗi khi cập nhật todo!';
          console.error('Có lỗi khi cập nhật todo:', error);
        });
    },
    cancelEdit() {
      this.editingTodoId = null;
    },
    toggleComplete(id, currentCompleted) {
      axios
        .put(`http://localhost:8081/todos/${id}`, { completed: !currentCompleted })
        .then(() => {
          this.getTodos();
        })
        .catch((error) => {
          this.error = 'Có lỗi khi thay đổi trạng thái todo!';
          console.error('Có lỗi khi thay đổi trạng thái todo:', error);
        });
    },
    deleteTodo(id) {
      axios
        .delete(`http://localhost:8081/todos/${id}`)
        .then(() => {
          this.getTodos();
        })
        .catch((error) => {
          this.error = 'Có lỗi khi xóa todo!';
          console.error('Có lỗi khi xóa todo:', error);
        });
    },
  },
  created() {
    this.getTodos();
  },
};
</script>

<style scoped>
/* Global Styles */
body {
  font-family: 'Arial', sans-serif;
  background-color: #f4f7fa;
  margin: 0;
  padding: 0;
}

.container {
  max-width: 600px;
  margin: 50px auto;
  background-color: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  color: #4a90e2;
  margin-bottom: 30px;
  font-size: 36px;
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 20px;
  font-size: 24px;
}

/* Input and Buttons */
.todo-input {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

#todo-input {
  width: 75%;
  padding: 12px;
  margin-right: 15px;
  border-radius: 5px;
  border: 1px solid #ddd;
  font-size: 16px;
  background-color: #f9f9f9;
  transition: border-color 0.3s ease;
}

#todo-input:focus {
  border-color: #4CAF50;
  outline: none;
}

#add-btn {
  padding: 12px 18px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

#add-btn:hover {
  background-color: #45a049;
}

/* Edit Form */
.edit-form {
  background-color: #f9f9f9;
  padding: 25px;
  margin-top: 20px;
  border-radius: 10px;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.1);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.form-group label {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.form-group input {
  padding: 12px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 5px;
  background-color: #f9f9f9;
  box-sizing: border-box;
}

.form-group input:focus {
  border-color: #4CAF50;
  outline: none;
}

/* Checkbox Styling */
.checkbox-group {
  display: grid;
  grid-template-columns: 1fr auto;
  align-items: center;
  gap: 10px;
}

.checkbox-label {
  font-size: 16px;
  color: #333;
}

.checkbox-container {
  display: flex;
  align-items: center;
}

.checkbox-input {
  width: 20px;
  height: 20px;
}

.status-text {
  margin-left: 10px;
  font-size: 16px;
  color: #333;
}

/* Button Group */
.button-group {
  display: flex;
  justify-content: space-between;
  gap: 20px;
}

.save-btn,
.cancel-btn {
  padding: 14px 22px;
  font-size: 16px;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  width: 48%;
}

.save-btn {
  background-color: #007bff;
  color: white;
}

.save-btn:hover {
  background-color: #0056b3;
}

.cancel-btn {
  background-color: #e53935;
  color: white;
}

.cancel-btn:hover {
  background-color: #c62828;
}

/* Table Styles */
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 30px;
}

table th,
table td {
  padding: 14px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

table th {
  background-color: #4CAF50;
  color: white;
}

table td button {
  padding: 10px 14px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 14px;
  cursor: pointer;
  margin-right: 10px;
  transition: background-color 0.3s ease;
}

table td button:hover {
  background-color: #0056b3;
}

/* Error Message */
.error {
  color: red;
  text-align: center;
  font-size: 18px;
  margin-top: 20px;
}
</style>
