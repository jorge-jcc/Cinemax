import axios from "axios";
import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getCarteleraAPI(){
  const url = `${API_HOST}/pelicula/cartelera`;
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}