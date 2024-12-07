<template>
  <h1>hello world</h1>
  <p>{{ value }}</p>
  <p>{{ greeting }}</p>
  <pre>{{ json }}</pre>
  <Button @click="() => console.log(value)">Value</Button>
  <Button @click="() => console.log(greeting)" class="small">Greet</Button>
</template>

<script lang="ts" setup>
import { GetJson, GetValue, Greet } from '@go/testpkg/Test';
import { useGetSimpleValue } from '@src/utils/composables';
import Button from '@src/modules/global/atoms/Button/button.vue';

const value = useGetSimpleValue(GetValue)
const greeting = useGetSimpleValue(Greet, " test")
const json = useGetSimpleValue(GetJson)
</script>


<style lang="sass" src="./global.scss"></style>
