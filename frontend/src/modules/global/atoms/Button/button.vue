<template>
    <button @click="props.onClick">
        <slot />
    </button>
</template>
<script lang="ts" setup>


interface Props {
    onClick: ((payload: MouseEvent) => void) | undefined
}
const props = defineProps<Props>()
</script>
<style lang="scss" src="./button.scss"></style>