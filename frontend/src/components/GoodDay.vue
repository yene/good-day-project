<template>
  <div class="hello">
    <h1>{{ today() }}</h1>
    <p>Hope you had a good day today! Tell us about it:</p>
    <div v-for="question in questions" :key="question.id">
      <div>
        {{ question.question }}
      </div>
      <div>
        <select v-model="question.answerID" placeholder="asdf">
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
    <button v-on:click="submitAnswers">Save my response</button>
  </div>
</template>

<script>
import { questions } from "./questions.js";

// initialize answers with empty
for (let i = 0; i < questions.length; i++) {
  questions[i].answerID = "";
}

export default {
  name: "GoodDay",
  data() {
    return {
      questions: questions,
    };
  },
  methods: {
    today() {
      let lang = (lang = navigator.language);
      let d = new Date();
      let options = {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
      };
      return d.toLocaleDateString(lang, options);
    },
    submitAnswers() {
      let answers = [];
      for (let q of this.questions) {
        if (q.answerID === "") {
          continue;
        }
        let choosenAnswer = q.answers.find((a) => {
          return a.id == q.answerID;
        });
        answers.push({
          id: q.id,
          question: q.question,
          answerID: q.answerID,
          answer: choosenAnswer.text,
        });
      }
      let url = "/api/answer";
      fetch(url, {
        method: "POST",
        mode: "cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(answers, null, 2),
      })
        .then((r) => {
          return r.json();
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
