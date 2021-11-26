<template>
  <v-container>
    <v-row class="text-center justify-center login">
      <v-col cols="6" sm="3" class="mt-3">
        <v-img src="@/assets/Imagen3.svg"  style="margin-bottom: -3vh;"/>
      </v-col>
      <v-col cols="12">
        <h1 class="display-4 font-weight-bold mb-3 noselect">Cinemax</h1>
      </v-col>
      <v-col cols="12">
        <v-btn color="primary" dark class="ma-2" @click="dialog = true">
          Iniciar Sesión
        </v-btn>
      </v-col>

      <v-dialog v-model="dialog" width="500">
        <v-card elevation="4" light>
          <v-card-title>Iniciar Sesión</v-card-title>
          <v-divider></v-divider>
          <v-alert
            v-model="alert"
            class="ma-4 py-2"
            close-text="Close Alert"
            color="red darken-4"
            dark
            text
            type="error"
            dismissible
          >
            {{message}}
          </v-alert>

          <v-card-text class="mt-3">
            <v-form class="mt-3">
              <v-text-field class="my-0"
                outline
                label="Email"
                type="text"
                v-model="user.email"
                :error-messages="emailErrors"
                @input="$v.user.email.$touch()"
                @blur="$v.user.email.$touch()"
                prepend-icon="mdi-account"
              ></v-text-field>
              <v-text-field class="my-0"
                outline
                label="Password"
                type="password"
                v-model="user.password"
                :error-messages="passwordErrors"
                @input="$v.user.password.$touch()"
                @blur="$v.user.password.$touch()"
                prepend-icon="mdi-lock"
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="grey darken-1" text @click="dialog=false">
            Cancelar
          </v-btn>
          <v-btn color="primary" :disabled="$v.user.$invalid" text @click="saludar">
            Iniciar Sesión
          </v-btn>
        </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </v-container>
</template>

<script>
import { mapActions } from "vuex";
import { required, email, minLength } from 'vuelidate/lib/validators'
import { loginApi, setTokenApi } from "../api/auth";

export default {
  name: "Login",
  data: () => ({
    dialog: false,
    alert: false,
    message: "",
    user:{
      email: "",
      password: "",
    }
  }),
  watch:{
    dialog: function (newV) {
      if (newV == false)
        this.resetForm()
    }
  },
   methods: {
     ...mapActions(["userLogin"]),
    saludar: function () {
      loginApi(this.user).then(res => {
        if (res.status == 200){
          setTokenApi(res.data.accessToken)
          this.userLogin()
        }
      }).catch(err => {
        const res = err.response
        if (res.status == 401){
          this.alert = true
          this.message = "El usuario o la contraseña no son válidos"
        } else{
          this.alert = true
          this.message = "Ocurrio un error, intentalo más tarde"
        }
        this.resetForm()
      })
    },
    resetForm: function(){
      this.user.email = this.user.password = ""
    }
  },
  validations: {
    user:{
      email: {
        required,
        email
      },
      password: {
        required,
        minLength: minLength(6)
      }
    }
  },
  computed:{
     emailErrors: function () {
      const errors = []
      if (!this.$v.user.email.$dirty || this.user.email === "") return errors
      !this.$v.user.email.email && errors.push('Introduce un email válido')
      !this.$v.user.email.required && errors.push('El email es requerido')
      return errors
    },
    passwordErrors: function () {
      const errors = []
      if (!this.$v.user.password.$dirty || this.user.password === "") return errors
      !this.$v.user.password.minLength && errors.push('Introduce una contraseña valida')
      return errors
    },
  }
};
</script>

<style>
#login {
  background: url('../assets/Imagen1.jpg');
  background-size: cover;
  height: 100%;
}
</style>
