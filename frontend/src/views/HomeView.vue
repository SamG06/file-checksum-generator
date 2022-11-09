<template>
  <main>
    <h1>Hello Checksum Generator</h1>
    <p>{{ serverResponse }}</p>
    <div></div>
  </main>
</template>

<script setup lang="ts">
import { useServerStore } from "@/stores/server";
import { storeToRefs } from "pinia";
import { ref } from "vue";

const serverResponse = ref(
  "This is going to show what the server responds with"
);
const { serverAddress } = storeToRefs(useServerStore());

console.log(serverAddress);
const firstRequest = async () => {
  const response = await fetch(serverAddress.value);
  const data = await response.text();
  serverResponse.value = data;
};

firstRequest();
</script>

<style scoped>
main {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-flow: column;
  align-items: center;
  justify-content: center;
}
h1 {
  color: white;
}

div {
  background: rgb(14, 56, 97);
  width: 300px;
  height: 20px;
}
</style>
