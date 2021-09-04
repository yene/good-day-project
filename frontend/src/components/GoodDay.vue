<template>
  <div class="hello">
    <h1>Saturday, August 2. 2021</h1>
    <p>Hope you had a good day today! Tell us about it:</p>
    <div v-for="question in questions" :key="question.id">
      <div>
        {{ question.question }}
      </div>
      <div>
        <select v-model="question.response" placeholder="asdf">
          <option value="" selected disabled hidden>
            {{ question.placeholder }}
          </option>
          <option
            v-for="answer in question.answers"
            :key="answer.id"
            :value="answer.id"
          >
            {{ answer.text }}
          </option>
        </select>
      </div>
    </div>
    <button v-on:click="submitAnswers">Submit</button>
  </div>
</template>

<script>
import { questions } from "./questions.js";

// initialize answers with empty
for (let i = 0; i < questions.length; i++) {
  questions[i].response = "";
}

export default {
  name: "GoodDay",
  data() {
    return {
      questions: questions,
    };
  },
  methods: {
    submitAnswers() {
      let data = this.questions;
      let url = "http://localhost:3000/api/answer";
      fetch(url, {
        method: "POST",
        mode: "cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      })
        .then((response) => {
          return response.json();
        })
        .then((data) => console.log("Response", data));
    },
  },
};
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
