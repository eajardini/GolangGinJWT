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
      requisicao: false,
      requisicaoTentativa: 5,
    };
  },
  mounted() {
       this.$http.defaults.headers.common['Authorization'] = this.setCabecalho();   
  },
  methods: {
    setCabecalho() {     
        this.token  = "";
       this.bearer = "Bearer ";
        this.token = this.$cookie.get('tokencliente');
        // Vue.prototype.$http.defaults.headers.common['Authorization'] = this.$cookie.get('tokencliente')
        this.bearer = this.bearer + this.token;
        return this.bearer;      
    },
    renovaAcesso(token){         
          this.$cookie.set('tokencliente', token, '1d');
          // this.$http.defaults.headers.common['Authorization'] = token;          
    },
    async digaOla(){        
        //this.$http.defaults.headers.common['Authorization'] = this.setCabecalho();
         this.requisicao = false;  
          this.requisicaoTentativa = 0;        
        do {
              this.requisicaoTentativa++;             
              console.log("[Inicio.vue digaOla n.1] Valor requiTentativa fora then:" + this.requisicaoTentativa);
              await this.$http //await necessário para o laço DO so continuar depois que o servidor respondeu
                .get("/mens/olasemclains")
                .then(res => {  
                  console.log("[Inicio.vue digaOla n.1] Valor requiTentativa dentro then:" + this.requisicaoTentativa);
                  alert("[Inicio.vue digaOla n.16] Retorno Token:" + res.data.mensagem);                  
                  this.requisicao = true;
                })
                .catch(error => {                  
                  const codigoDoErro = error.response.data.code;
                  const estaLogado = this.$cookie.get('autenticadocook');

                  if ((codigoDoErro == "401") && (estaLogado)) {
                    this.$http.defaults.headers.common['Authorization'] = "";
                    // alert("[Inicio.vue digaOla n.21] Valor do header:" + this.$http.defaults.headers.common['Authorization'])
                    this.$http.defaults.headers.common['Authorization'] = this.setCabecalho();                    
                    this.$http
                      .get("/auth/refresh_token")
                      .then(res => {
                            console.log("[Inicio.vue digaOla n.15] Retorno Token Renovado:" + res.data.message)              
                            this.renovaAcesso(res.data.token); 
                            this.$http.defaults.headers.common['Authorization'] = "";
                            this.$http.defaults.headers.common['Authorization'] = this.setCabecalho(); 
                            console.log("[Inicio.vue digaOla n.2] Valor requiTentativa dentro Refresh:" + this.requisicaoTentativa); 
                            
                      })
                      .catch(error => {
                        console.log("[Inicio.vue digaOla n.18] Erro ao renovar o Token:" + error.response.data.message)              
                      })
                  } else {
                    alert ("[Inicio.vue digaOla n. 12] Usuário não logado")
                    this.requisicao = true;
                  }
                })
        } while (this.requisicao != true) // && this.requisicaoTentativa < 5)
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
