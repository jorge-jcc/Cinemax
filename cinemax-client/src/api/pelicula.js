import axios from "axios";
import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getPeliculasEnCartelera(){
  const url = `${API_HOST}/pelicula/cartelera`;
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getIdiomas(){
  const url = `${API_HOST}/pelicula/idiomas`;
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getGeneros(){
  const url = `${API_HOST}/pelicula/generos`;
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getClasificaciones(){
  const url = `${API_HOST}/pelicula/clasificaciones`;
  return axios.get(url, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function CreatePeliculaAPI(pelicula) {
  const url = `${API_HOST}/pelicula/registrar`;
  const data = {...pelicula}
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function LoadImagePeliculaAPI(file, peliculaId) {
  const url = `${API_HOST}/pelicula/load_image`
  const formData = new FormData()
  formData.append("portada", file)
  formData.append("peliculaId", peliculaId)
  return axios.post(url, formData, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getPeliculasIdAPI(query = "", limit = 20, offset = 0){
  const url = `${API_HOST}/pelicula/search/id`
  const data = {query, limit, offset}
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getPeliculasAPI(query = "", limit = 20, offset = 0){
  const url = `${API_HOST}/pelicula/search`
  const data = {nombre: query, limit, offset}
  return axios.post(url, data, {
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}

export function getImagePeliculaByIdAPI(peliculaId){
  const url = `${API_HOST}/pelicula/image`
  const params = {"pelicula_id": peliculaId}
  return axios.get(url, {
    params,
    headers: {
      'Authorization': `Bearer ${getTokenApi()}`
    }
  })
}