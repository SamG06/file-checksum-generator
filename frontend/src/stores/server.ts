import { defineStore } from "pinia";
import { ref } from "vue";

export const useServerStore = defineStore("server", () => {
  const serverAddress = ref<string>("http://localhost:3000/");

  return { serverAddress };
});
