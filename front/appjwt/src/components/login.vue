<template>
  <form>
    <div id="FormularioLogin">
      <h1>Login</h1>
      <div>
        <input
          type="text"
          style="font-size: 24px"
          name="username"
          v-model="input.username"
          placeholder="Usuário"
        />
      </div>
      <div>
        <input
          type="password"
          style="font-size: 24px"
          name="password"
          v-model="input.password"
          placeholder="Password"
        />
      </div>
      <button type="button" v-on:click="login()">Login</button>
    </div>
  </form>
</template>

<script>
export default {  
  data() {
    return {
      input: {
        username: "",
        password: ""
      }
    };
  },
  methods: {
    login() {
      if (this.input.username != "" && this.input.password != "") {
        if (
          this.input.username == this.$parent.mockAccount.username &&
          this.input.password == this.$parent.mockAccount.password
        ) {
            console.log("[Login.vue mounted] this.authenticated:" + this.$parent.authenticated)
          this.$parent.authenticated = true;         
          this.$router.push("/home" );
          sessionStorage.setItem("autenticado", true);
          localStorage.autenticado = "sim";
          this.$cookie.set('autenticadocook', 'Hello world!', '1d');
        } else {
          console.log("The username and / or password is incorrect");
        }
      } else {
        console.log("A username and password must be present");
      }
    }
  }
};
</script>

<style scoped>
#login {
  width: 500px;
  border: 1px solid #cccccc;
  background-color: #ffffff;
  margin: auto;
  margin-top: 200px;
  padding: 20px;
}
</style>