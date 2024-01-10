<script setup>
import {ref} from 'vue'
let count = ref(0)
let inputValue = ref('')
let todos = ref([])

function addTodo() {
   todos.value.push(inputValue.value)
   inputValue.value=''
}

function deleteTodo(todo) {
    todos.value = todos.value.filter(t => t !== todo)
}

</script>

<template>
    <div>
        <h1>Count is {{ count }}</h1>
        <div>
          <input type ="text" v-model="inputValue"/>
          <button @click="addTodo">Add Todo</button>
        <div v-for="todo in todos" :key="todo">
            <h1>{{ todo }}</h1>
            <button @click="deleteTodo(todo)" >Delete</button>
            <img :src=" todo " alt="pic"/>
        </div> 
      </div>
    </div>
</template>