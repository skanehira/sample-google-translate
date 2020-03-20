const app = new Vue({
  el: '#app',
  data: {
    options: {
      en: "英語",
      ja: "日本語",
    },
    input: "",
    output: "",
    source: "en",
    target: "ja",
  },
  methods: {
    translate() {
      const form = {
        text: this.input,
        source: this.source,
        target: this.target,
      }

      axios.post("/translate", form).then(response => {
        this.output = response.data.output.split("\n")
      }).catch(error => {
        console.log(error);
      })
    }
  },
  mounted(){
  }
})

