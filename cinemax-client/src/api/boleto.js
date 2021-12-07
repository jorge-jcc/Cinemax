import axios from "axios"
import { API_HOST } from "../utils/constant"
import { getTokenApi } from "./auth"

export function getPreciosBoletosAPI(){
  const url = `${API_HOST}/boleto`
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}


