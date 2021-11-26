import axios from "axios";
import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getAsientosByFuncion(funcionID){
  const url = `${API_HOST}/asiento`;
  const params = { "funcion_id" : funcionID }
  return axios.get(url, { 
    params,
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function seleccionarAsientoApi(asientoId, transaccionId){
  const url = `${API_HOST}/asiento`;
  const data = {
    asientoId, 
    transaccionId
  }
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function deseleccionarAsientoAPI(asientoId, transaccionId){
  const url = `${API_HOST}/asiento`;
  const data = {
    asientoId, 
    transaccionId
  }
  return axios.delete(url, {
    data,
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function deshacerTransaccionAPI(transaccionId){
  const url = `${API_HOST}/asiento/deshacer`;
  const data = { transaccionId }
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}