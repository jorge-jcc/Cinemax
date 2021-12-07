import axios from "axios";
import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

// getInformacionSala obtiene la informacion de la sala en la que se proyectara 
// una función
export function getInformacionSala(funcionId){
  const url = `${API_HOST}/salas/by_funcion`;
  const params = { "funcion_id" : funcionId }
  return axios.get(url, {
    params,
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getSalasDisponibles(fechaInicio, peliculaId){
  const url = `${API_HOST}/salas`;
  const data = { fechaInicio, peliculaId }
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}