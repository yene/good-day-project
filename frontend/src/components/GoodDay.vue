<template>
  <div class="hello">
    <h1>Saturday, August 2. 2021</h1>
    <p>Hope you had a good day today! Tell us about it:</p>
    <div>
      <div>
        🤔 How was your workday?
      </div>
      <div>
        <select v-model="selectedAnswer" placeholder="asdf">
          <option value="" selected disabled hidden>I feel...</option>
          <option v-for="(answer) in answers" :key="answer.id" :value="answer.id">
            {{answer.text}}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GoodDay',
  data() {
    return {
      selectedAnswer: '',
      answers: [
        {text: 'Terrible', id: 'terrible'},
        {text: 'Bad', id: 'bad'},
        {text: 'OK', id: 'ok'},
        {text: 'Good', id: 'good'},
        {text: 'Awesome', id: 'awesome'}
      ]
    }
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}


</style>
