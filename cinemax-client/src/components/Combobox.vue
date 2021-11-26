<template>
    <v-row>
      <v-col cols="12" class="pb-0"><h2>{{title}}</h2></v-col>
      <v-col class="mb-0 pb-0">
        <v-combobox 
          class="my-0"
          :disabled="disabled"
          v-model="select"
          :items="itemsValue"
          :label="label" 
          @input="$emit('select', itemsValue.indexOf(select))"
        />
      </v-col>
    </v-row>
</template>

<script>

export default {
  name: "ComboBox",
  props:{
    title: String,
    label: String,
    items: Array,
    name: String,
    disabled: Boolean
  },
  computed: {
    itemsValue: function(){
      return this.items.map(e => e[this.name])
    }
  },
  data: () => ({
    select: null,
  }),
}
</script>