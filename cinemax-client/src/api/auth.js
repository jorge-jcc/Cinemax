import axios from "axios";
import { API_HOST, TOKEN } from "../utils/constant";
import  jwtDecode from "jwt-decode";

export function loginApi(user){
  const url = `${API_HOST}/empleado/login`
  const userTemp = {
    ...user,
    email: user.email.toLowerCase()
  }
  return axios.post(url, userTemp)
}

// setTokenApi guarda el token en el localStorage del navegador
export function setTokenApi(token){
  localStorage.setItem(TOKEN, token)
}

// setTokenApi obtiene el token del localStorage del navegador
export function getTokenApi(){
  return localStorage.getItem(TOKEN);
}

// logoutApi elimina el token del localStorage del navegador
export function logoutApi(){
  localStorage.removeItem(TOKEN);
}


// isExpired devuelve true si el token ya expiro, false si no
function isExpired(token){
  const { expired_at } = jwtDecode(token);
  return new Date(expired_at) < Date.now()
}

// isUserLogedApi Verifica el token se encuentre almacenado en el navegador
// y si aún es valido, devulve null o el contenido del token 
export function isUserLogedApi(){
  const token = localStorage.getItem(TOKEN)
  if (!token) {
    logoutApi()
    return
  }
  if(isExpired(token)){
    logoutApi()
    return
  }
  return jwtDecode(token)
}