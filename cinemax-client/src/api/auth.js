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
  const { exp } = jwtDecode(token);
  const expire = exp * 1000;
  const timeout = expire - Date.now();
  if (timeout < 0) 
    return true
  return false
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