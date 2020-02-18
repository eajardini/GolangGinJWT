<template>
  <div class="inicio">
    <h1>Página Início</h1>
    <router-link to="hello">Hello</router-link>
    <br>
    <button v-on:click="digaOla">Diga olá</button>
    <br>
    <button v-on:click="sairDoSistema">Sair do Sistema</button>
  </div>
</template>

<script>
export default {  
  data() {
    return {
      token:"",
      bearer:"Bearer ",
    };
  },
  methods: {
    digaOla(){        
        this.token = this.$cookie.get('tokencliente');
        // Vue.prototype.$http.defaults.headers.common['Authorization'] = this.$cookie.get('tokencliente')
        this.bearer = this.bearer + this.token;
        this.$http.defaults.headers.common['Authorization'] = this.bearer;

        this.$http
          .get("/mens/olasemclains")
          .then(res => {
            console.log("[Inicio.vue digaOla] Retorno Token:" + res.data.mensagem)              
          })
          .catch(error => {
            console.log("[Inicio.vue Login] Retorno do Erro:" + error.response.data)
            alert("Erro de acesso:" + error.response.headers)
          })

        // this.$http
        //   .get("/mens/olasemclains",{
	      //             headers: {                            
	      //             'Authorization': 'Bearer ' + this.token
	      //   },})
        //   .then(res => {
        //     console.log("[Inicio.vue digaOla] Retorno Token:" + res.data.mensagem)              
        //   })
        //   .catch(error => {
        //     console.log("[Inicio.vue Login] Retorno do Erro:" + error.response.data)
        //     alert("Erro de acesso:" + error.response.headers)
        //   })
    },
    sairDoSistema(){
      //this.$cookie.set('borodin', true, '1d');
      this.$cookie.delete('autenticadocook');
      this.$cookie.delete('tokencliente');
      this.$router.push("/login" );
    }
  }
};
</script>
