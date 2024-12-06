import { ref } from 'vue'

export const useGetSimpleValue = (fn: Function, ...args: any[]): any => {
	if (typeof fn !== 'function') {
		console.error(
			'Error: useGetSimpleValue - Argument should be a proper function, got: ',
			typeof fn
		)
		return `error when provided: ${fn}`
	}
	const value = ref()
	fn(...args).then((result: any) => {
		value.value = result
	})
	return value
}
