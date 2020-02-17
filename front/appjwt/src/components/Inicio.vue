<template>
  <div class="inicio">
    <h1>Página Início</h1>
    <router-link to="hello">Hello</router-link>
    <br>
    <button v-on:click="sairDoSistema">Sair do Sistema</button>
    <br>
    <button v-on:click="sairDoSistema">Sair do Sistema</button>
  </div>
</template>

<script>
export default {  
  data() {
    return {
      
    };
  },
  methods: {
    sairDoSistema(){
      //this.$cookie.set('borodin', true, '1d');
      this.$cookie.delete('autenticadocook');
      this.$cookie.delete('tokencliente');
      this.$router.push("/login" );
    }
  }
};
</script>
