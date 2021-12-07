import axios from "axios"
import { API_HOST } from "../utils/constant"
import { getTokenApi } from "./auth"

export function iniciarCompraAPI(payload){
  const url = `${API_HOST}/ticket/iniciar_compra`
  return axios.post(url, payload, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function crearTicketAPI(transaccionId, funcionId, monto, boletos){
  const url = `${API_HOST}/ticket/crear_ticket`
  const data = {
    transaccionId,
    funcionId,
    monto,
    boletos: boletos
  }
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

