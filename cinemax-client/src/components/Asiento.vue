<template>
  <div>
    <svg
      xmlns:dc="http://purl.org/dc/elements/1.1/"
      xmlns:cc="http://creativecommons.org/ns#"
      xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
      xmlns:svg="http://www.w3.org/2000/svg"
      xmlns="http://www.w3.org/2000/svg"
      id="svg979"
      version="1.1"
      viewBox="0 0 26.625523 22.821899"
      class="ma-1"
      :class="colorAsiento"
      @click="updateState">
    <defs
      id="defs973" />
    <metadata
      id="metadata976">
      <rdf:RDF>
        <cc:Work
          rdf:about="">
          <dc:format>image/svg+xml</dc:format>
          <dc:type
            rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
          <dc:title />
        </cc:Work>
      </rdf:RDF>
    </metadata>
    <g
      transform="translate(0.07898591,0.01653374)"
      id="layer1">
      <g
        transform="translate(-38.106926,-69.069162)"
        id="use951" />
      <g
        transform="translate(-57.859921,-112.94947)"
        id="g1574"
        clip-path="none">
        <g
          id="g1572">
          <g
            transform="matrix(0.26458333,0,0,0.26458333,15.37508,94.722647)"
            id="g1570"
            style="overflow:hidden">
            <path
              id="path1568"
              d="M 253.718,111.95429 H 246.53 V 90.390291 c 0,-11.889 -9.675,-21.564 -21.5638,-21.564 h -28.752 c -11.889,0 -21.5641,9.675 -21.5641,21.564 v 21.563999 h -7.188 c -3.975,0 -7.188,3.213 -7.188,7.188 v 28.752 c 0,3.975 3.213,7.188 7.188,7.188 h 86.2559 c 3.975,0 7.188,-3.213 7.188,-7.188 v -28.752 c 0,-3.975 -3.213,-7.188 -7.188,-7.188 z"
              />
          </g>
        </g>
      </g>
    </g>
    <text x="50%" y="50%" 
      dominant-baseline="middle" 
      class="txt-icon noselect" 
      text-anchor="middle">
      {{asiento.clave}}
    </text>
  </svg>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex"

export default {
  name: "Asiento",
  props: {
    asiento: Object,
  },
  data: () => ({
    dialog: false,
    select: true
  }),
  beforeUpdate: function (){
    this.select = true;
  },
  computed: {
    ...mapState(["asientos"]),
    colorAsiento: function(){
      return this.asiento.statusId === "DISPONIBLE" ? "disponible" :
      this.asiento.statusId === "EN PROCESO" ? "ocupado":
      this.asiento.statusId === "SELECCIONADO" ? "seleccionado": ""
    }
  },
  methods:{
    ...mapActions(["seleccionarAsiento", "deseleccionarAsiento"]),
    updateState: async function(){
      if (this.select){
        this.select = false
        switch (this.asiento.statusId){
          case "DISPONIBLE": {
            this.asiento.statusId = "SELECCIONADO"
            await this.seleccionarAsiento(this.asiento)
            break
          }
          case "SELECCIONADO":{
            await this.deseleccionarAsiento(this.asiento).then(() => {
              this.asiento.statusId = "DISPONIBLE"
            })
            break
          }
        }
        this.$emit('update')
      }
    }
  }
}
</script>

<style>

.seleccionado{
  fill:#FFCA28;
}
.seleccionado:hover{
  fill: #FFA000
}

.disponible{
  fill:#29B6F6;
}
.disponible:hover{
  fill: #0288D1; 
}

.ocupado{
  fill: #EF5350;
}
.ocupado:hover{
  fill: #D32F2F;
}

.txt-icon{
  font-family: 'Roboto';
  font-size:0.55rem !important; 
  fill:#ffffff;
}

</style>