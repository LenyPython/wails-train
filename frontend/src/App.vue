<template>
  <h1>hello world</h1>
  <p>
    {{ value }}
  </p>
  <Button @click="handleClick">Change</Button>
</template>

<script lang="ts" setup>
import Button from '@js/modules/global/atoms/Button/button.vue';
import { onMounted, onUnmounted, ref } from 'vue';


const value = ref<string>('')
onMounted(() => {
  (window as any).runtime?.EventsOn('state_updated', (data: string) => {
    console.log(data)
    value.value = data
  })
})

onUnmounted(() => {
  (window as any).runtime?.EventsOff('state_updated')
})

const handleClick = () => {
  (window as any).runtime?.EventsEmit('change_state', 'changed')
  console.log('clicked')
}
</script>


<style lang="sass" src="./global.scss"></style>
