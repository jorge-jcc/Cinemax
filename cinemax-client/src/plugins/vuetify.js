import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: colors.red.darken4, // #E53935
        secondary: colors.red.lighten4, // #FFCDD2
        accent: colors.red.base, // #3F51B5
      },
    },
  },
});
