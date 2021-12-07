import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import { isUserLogedApi, logoutApi } from "../api/auth"
import { deseleccionarAsientoAPI, deshacerTransaccionAPI, seleccionarAsientoApi } from '../api/asientos'
import { getPreciosBoletosAPI } from '../api/boleto'
import { crearTicketAPI } from '../api/ticket'

export default new Vuex.Store({
  state: {
    user: null,
    token: null,
    transactionId: null,
    dialogError: false,
    dialogErrorMessage: "",
    asientos: [],
    preciosBoletos: null,
    updateAsientos: false,
  },
  mutations: {
    setUser(state, payload ){
      state.user = payload;
    },
    setTransactionId(state, payload){
      state.transactionId = payload;
    },
    setAsiento(state, payload){
      state.asientos.push(payload);
    },
    deleteAsiento(state, payload){
      state.asientos = state.asientos.filter(a => a.id !== payload.id)
    },
    resetAsientos(state){
      state.asientos = []
    },
    setDialogError(state, payload){
      state.dialogError = payload
    },
    setDialogErrorMessage(state, payload){
      state.dialogErrorMessage = payload
    },
    setPreciosBoletos(state, payload){
      state.preciosBoletos = payload
    },
    setUpdateAsientos(state){
      state.updateAsientos ^= true;
    }
  },
  actions: {
    userLogin({commit}){
      const user = isUserLogedApi()
      commit('setUser', user)
    },
    userLogout({commit}){
      logoutApi()
      commit('setUser', null)
    },
    async seleccionarAsiento({commit, dispatch, state}, payload){
      await seleccionarAsientoApi(payload.id, state.transactionId).then((res)=>{
        if (res.status == 200){
          if (state.transactionId == null)
            commit('setTransactionId', res.data.transaccionId)
          commit('setAsiento', payload)
        }
      }).catch((err) => {
        if(err.response.status == 404)
          dispatch('resetTaquilla');
        commit('setDialogErrorMessage', err.response.data.error.message)
        commit('setDialogError', true)
      })
    },
    async deseleccionarAsiento({commit, state}, payload){
      await deseleccionarAsientoAPI(payload.id, state.transactionId).then(res =>{
        if (res.status == 204)
          commit('deleteAsiento', payload)
      }).catch(() => {})
    },
    resetTaquilla({commit}){
      commit('setTransactionId', null)
      commit('resetAsientos')
    },
    async deshacerTransaccionAPI({state, dispatch}){
      if (state.transactionId != null){
        await deshacerTransaccionAPI(state.transactionId).then(() =>{
          dispatch('resetTaquilla')
        })
      }
    },
    getPreciosBoletos({commit}){
      getPreciosBoletosAPI().then((res) => {
        if (res.status == 200){
          commit('setPreciosBoletos', res.data.precios)
        }
      })
    },
    async crearTicket({dispatch, commit, state}, {total, boletos}) {
      await crearTicketAPI(
        state.transactionId, state.asientos[0].funcionId, total, boletos
      ).then(()=>{}).catch(err => {
        commit('setDialogErrorMessage', err.response.data.error.message)
        commit('setDialogError', true)
      }).finally(()=>{
        dispatch('resetTaquilla')
        commit('setUpdateAsientos')
      })
    }
  },
  modules: {
  }
})
