<template>
  <div id="app">
    <div id="nav">

      <!-- https://www.thepolyglotdeveloper.com/2018/04/simple-user-login-vuejs-web-application/ -->
      
    </div>
    <router-view @authenticated="setAuthenticated" />
  </div>
</template>

<script>
    export default {
        name: 'App',
        data() {
            return {
                authenticated: false,
                mockAccount: {
                    username: "teste",
                    password: "teste"
                }
            }
        },
        mounted() {
          console.log("[App.vue mounted] this.authenticated:" + this.authenticated)
            // if(!this.authenticated) {
            //    // this.$router.replace({ name: "login" });
            //    this.$router.push("/login");
            // }
            // this.authenticated = sessionStorage.getItem('autenticado')
            this.authenticated = this.$cookie.get('autenticadocook');
             console.log("[App.vue mounted] this.authenticated:" + this.authenticated)
            if(!this.authenticated) {
              this.$router.push("/login");
            } 
        },
        methods: {
            setAuthenticated(status) {
                this.authenticated = status;
            },
            getAuthenticated(){
              return this.authenticated;
            },
            logout() {
                this.authenticated = false;
            }
        }
    }
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
