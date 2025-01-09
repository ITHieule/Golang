<template>
  <div class="login-container">
    <h1>Login</h1>
    <div class="login-form">
      <input
        v-model="username"
        type="text"
        placeholder="Enter username"
        class="input-field"
      />
      <input
        v-model="password"
        type="password"
        placeholder="Enter password"
        class="input-field"
      />
      <button @click="login" class="login-btn">Login</button>
    </div>
    <div v-if="error" class="error">{{ error }}</div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      username: "",
      password: "",
      error: null,
    };
  },
  methods: {
    login() {
      axios
        .post("http://localhost:8081/api/login", {
          username: this.username,
          password: this.password,
        })
        .then((response) => {
          const token = response.data.token;
          localStorage.setItem("token", token);
          this.$router.push("/Todo");
        })
        .catch((err) => {
          this.error = "Login failed. Please check your credentials.";
          console.error(err);
        });
    },
  },
};
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 30px;
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
  text-align: center;
}
h1 {
  font-size: 28px;
  color: #333;
  margin-bottom: 20px;
}
.input-field {
  width: 100%;
  padding: 10px;
  margin-bottom: 15px;
  border-radius: 5px;
  border: 1px solid #ddd;
  font-size: 16px;
}
.login-btn {
  width: 100%;
  padding: 12px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 16px;
  cursor: pointer;
}
.login-btn:hover {
  background-color: #45a049;
}
.error {
  color: red;
  margin-top: 15px;
  font-size: 16px;
}
</style>
