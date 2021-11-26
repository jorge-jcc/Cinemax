import axios from "axios";
import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getFuncionesByPeliculaId(peliculaId){
  const url = `${API_HOST}/funcion/show`;
  const data = {
    peliculaId,
    fecha: new Date().toISOString()
  }
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}