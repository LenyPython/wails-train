<script lang="ts" setup>
import Button from './modules/global/atoms/Button/Button.vue';
import { GetValue, GetJson, Greet } from '../wailsjs/go/main/Test'
import { useGetSimpleValue } from './utils/composables';
const value = useGetSimpleValue(GetValue)
const greeting = useGetSimpleValue(Greet, " test")
const json = useGetSimpleValue(GetJson)
</script>

<template>
  <h1>hello world</h1>
  <p>{{ value }}</p>
  <p>{{ greeting }}</p>
  <pre>{{ json }}</pre>
  <Button @click="() => console.log('lol')">Click me</Button>
  <Button @click="() => console.log('lol2')" class="small">Click me2</Button>
</template>

<style lang="sass" src="./global.scss"></style>
