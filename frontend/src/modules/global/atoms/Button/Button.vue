<template>
    <button @click="props.onClick">
        <slot />
    </button>
</template>
<script lang="ts" setup>

interface Props {
    onClick: Function
}
const props = defineProps<Props>()
</script>
<style src="./Button.scss"></style>