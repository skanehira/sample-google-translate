const app = new Vue({
  el: '#app',
  data: {
    input: "",
    output: "",
  },
  methods: {
    translate() {
      const form = {
        text: this.input,
      }

      axios.post("/translate", form).then(response => {
        this.output = response.data.output
      }).catch(error => {
        console.log(error);
      })
    }
  },
  mounted(){
  }
})

