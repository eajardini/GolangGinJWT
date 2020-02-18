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
    liberaAcesso(token){
          this.$cookie.set('autenticadocook', true, '1d');
          this.$cookie.set('tokencliente', token, '1d');
          // this.$http.defaults.headers.common['Authorization'] = token;
          this.$router.push("/inicio" );
    },

    login() {
      if (this.input.username != "" && this.input.password != "") {

        const formData = new FormData();
      
        formData.append("username", this.input.username);
        formData.append("password", this.input.password);
        this.$http
          .post("login", formData)
          .then(res => {
            console.log("[Login.vue Login] Retorno Token:" + res.data.token)      
            this.liberaAcesso(res.data.token);       
          })
          .catch(error => {
            console.log("[Login.vue Login] Retorno do Erro:" + error.response.data)
            alert("Usuário não cadastrado ou senha inválida:" + error.response.data)
          })


        // this.axios
        //   .post("/login", formData)
        //   .then(res => {
        //     console.log("[Login.vue Login] Retorno Token:" + res.data.token)      
        //     this.liberaAcesso(res.data.token);       
        //   })
        //   .catch(error => {
        //     console.log("[Login.vue Login] Retorno do Erro:" + error.response.data)
        //     alert("Usuário não cadastrado ou senha inválida:" + error.response.data)
        //   })



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